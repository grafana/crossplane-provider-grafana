package teams

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	v1alpha1observe "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/observe/v1alpha1"
	"github.com/grafana/crossplane-provider-grafana/v2/internal/clients"

	oateams "github.com/grafana/grafana-openapi-client-go/client/teams"
)

// seedTeams is the canonical set of teams used by all unit tests. The mock
// server filters from this list based on name/query parameters, just like the
// real Grafana API does.
var seedTeams = []teamPayload{
	{ID: 1, UID: "uid-01", Name: "platform-core", Email: "platform-core@example.com", MemberCount: 8, OrgID: 1},
	{ID: 2, UID: "uid-02", Name: "platform-infra", Email: "platform-infra@example.com", MemberCount: 5, OrgID: 1},
	{ID: 3, UID: "uid-03", Name: "ops-east", Email: "ops-east@example.com", MemberCount: 4, OrgID: 1},
	{ID: 4, UID: "uid-04", Name: "ops-west", Email: "ops-west@example.com", MemberCount: 6, OrgID: 1},
	{ID: 5, UID: "uid-05", Name: "ops-apac", Email: "ops-apac@example.com", MemberCount: 3, OrgID: 1},
	{ID: 6, UID: "uid-06", Name: "frontend-web", Email: "frontend-web@example.com", MemberCount: 7, OrgID: 1},
	{ID: 7, UID: "uid-07", Name: "frontend-mobile", Email: "frontend-mobile@example.com", MemberCount: 4, OrgID: 1},
	{ID: 8, UID: "uid-08", Name: "backend-api", Email: "backend-api@example.com", MemberCount: 9, OrgID: 1},
	{ID: 9, UID: "uid-09", Name: "backend-workers", Email: "backend-workers@example.com", MemberCount: 3, OrgID: 1},
	{ID: 10, UID: "uid-10", Name: "data-engineering", Email: "data-eng@example.com", MemberCount: 6, OrgID: 1},
	{ID: 11, UID: "uid-11", Name: "data-science", Email: "data-sci@example.com", MemberCount: 4, OrgID: 1},
	{ID: 12, UID: "uid-12", Name: "security", Email: "security@example.com", MemberCount: 5, OrgID: 1},
	{ID: 13, UID: "uid-13", Name: "sre-observability", Email: "sre-obs@example.com", MemberCount: 3, OrgID: 2},
	{ID: 14, UID: "uid-14", Name: "sre-oncall", Email: "sre-oncall@example.com", MemberCount: 7, OrgID: 2},
	{ID: 15, UID: "uid-15", Name: "qa-automation", Email: "qa-auto@example.com", MemberCount: 4, OrgID: 1},
	{ID: 16, UID: "uid-16", Name: "qa-manual", Email: "qa-manual@example.com", MemberCount: 2, OrgID: 1},
	{ID: 17, UID: "uid-17", Name: "devrel", Email: "devrel@example.com", MemberCount: 3, OrgID: 1},
	{ID: 18, UID: "uid-18", Name: "design", Email: "design@example.com", MemberCount: 5, OrgID: 1},
	{ID: 19, UID: "uid-19", Name: "product-growth", Email: "product-growth@example.com", MemberCount: 4, OrgID: 1},
	{ID: 20, UID: "uid-20", Name: "product-enterprise", Email: "product-ent@example.com", MemberCount: 6, OrgID: 2},
}

// seedSummaries returns the TeamSummary representation of seedTeams.
func seedSummaries() []v1alpha1observe.TeamSummary {
	out := make([]v1alpha1observe.TeamSummary, len(seedTeams))
	for i, t := range seedTeams {
		out[i] = v1alpha1observe.TeamSummary{
			ID:          t.ID,
			UID:         t.UID,
			Name:        t.Name,
			Email:       t.Email,
			MemberCount: t.MemberCount,
			OrgID:       t.OrgID,
		}
	}
	return out
}

// filterSeedTeams returns the subset of seedTeams matching the given name
// (exact match) and query (substring match), mirroring Grafana API behavior.
func filterSeedTeams(name, query string) []teamPayload {
	var out []teamPayload
	for _, t := range seedTeams {
		if name != "" && t.Name != name {
			continue
		}
		if query != "" && !strings.Contains(t.Name, query) {
			continue
		}
		out = append(out, t)
	}
	return out
}

// searchResult is a helper for building mock Grafana search responses.
type searchResult struct {
	Teams      []teamPayload `json:"teams"`
	TotalCount int64         `json:"totalCount"`
	Page       int64         `json:"page"`
	PerPage    int64         `json:"perPage"`
}

type teamPayload struct {
	ID          int64  `json:"id"`
	UID         string `json:"uid"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	MemberCount int64  `json:"memberCount"`
	OrgID       int64  `json:"orgId"`
}

// newMockGrafana starts an httptest.Server that responds to GET /api/teams/search.
// The handler function receives (page, perPage, name, query) and returns a searchResult.
func newMockGrafana(t *testing.T, handler func(page, perPage int64, name, query string) searchResult) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/teams/search" {
			http.NotFound(w, r)
			return
		}

		q := r.URL.Query()
		page := int64(1)
		perPage := int64(1000)
		if v := q.Get("page"); v != "" {
			_ = json.Unmarshal([]byte(v), &page)
		}
		if v := q.Get("perpage"); v != "" {
			_ = json.Unmarshal([]byte(v), &perPage)
		}

		result := handler(page, perPage, q.Get("name"), q.Get("query"))
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(result)
	}))
	t.Cleanup(srv.Close)
	return srv
}

// newSeedMockGrafana starts a mock that serves from seedTeams, filtering by
// name (exact) and query (substring) and paginating the results.
func newSeedMockGrafana(t *testing.T) *httptest.Server {
	t.Helper()
	return newMockGrafana(t, func(page, perPage int64, name, query string) searchResult {
		filtered := filterSeedTeams(name, query)
		total := int64(len(filtered))

		start := (page - 1) * perPage
		if start > total {
			start = total
		}
		end := start + perPage
		if end > total {
			end = total
		}

		return searchResult{
			TotalCount: total,
			Page:       page,
			PerPage:    perPage,
			Teams:      filtered[start:end],
		}
	})
}

// newExternalFromServer builds an external client pointed at the given mock server.
func newExternalFromServer(t *testing.T, srv *httptest.Server) *external {
	t.Helper()
	c, err := clients.NewOAPIClient(srv.URL, "test-token")
	if err != nil {
		t.Fatalf("NewOAPIClient: %v", err)
	}
	return &external{teamsClient: c.Teams}
}

func ptr[T any](v T) *T { return &v }

// --- searchAllTeams unit tests ---

// TestSearchAllTeams_NoFilter verifies that an unfiltered search returns all 20
// seed teams with all fields correctly mapped to TeamSummary.
func TestSearchAllTeams_NoFilter(t *testing.T) {
	srv := newSeedMockGrafana(t)
	ext := newExternalFromServer(t, srv)

	got, err := ext.searchAllTeams(&v1alpha1observe.Teams{})
	if err != nil {
		t.Fatalf("searchAllTeams: %v", err)
	}
	want := seedSummaries()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("teams mismatch (-want +got):\n%s", diff)
	}
}

// TestSearchAllTeams_NameFilter verifies that setting Spec.ForProvider.Name
// passes the name parameter to the API and returns only exact matches.
func TestSearchAllTeams_NameFilter(t *testing.T) {
	srv := newSeedMockGrafana(t)
	ext := newExternalFromServer(t, srv)
	cr := &v1alpha1observe.Teams{
		Spec: v1alpha1observe.TeamsSpec{
			ForProvider: v1alpha1observe.TeamsParameters{Name: ptr("security")},
		},
	}

	got, err := ext.searchAllTeams(cr)
	if err != nil {
		t.Fatalf("searchAllTeams: %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("expected 1 team, got %d", len(got))
	}
	if got[0].Name != "security" {
		t.Errorf("expected team named security, got %q", got[0].Name)
	}
	if got[0].ID != 12 {
		t.Errorf("expected ID=12, got %d", got[0].ID)
	}
}

// TestSearchAllTeams_QueryFilter verifies that setting Spec.ForProvider.Query
// passes the query parameter to the API for substring matching.
func TestSearchAllTeams_QueryFilter(t *testing.T) {
	srv := newSeedMockGrafana(t)
	ext := newExternalFromServer(t, srv)
	cr := &v1alpha1observe.Teams{
		Spec: v1alpha1observe.TeamsSpec{
			ForProvider: v1alpha1observe.TeamsParameters{Query: ptr("ops")},
		},
	}

	got, err := ext.searchAllTeams(cr)
	if err != nil {
		t.Fatalf("searchAllTeams: %v", err)
	}
	wantNames := []string{"ops-east", "ops-west", "ops-apac"}
	if len(got) != len(wantNames) {
		t.Fatalf("expected %d teams, got %d", len(wantNames), len(got))
	}
	for i, name := range wantNames {
		if got[i].Name != name {
			t.Errorf("team[%d]: expected %q, got %q", i, name, got[i].Name)
		}
	}
}

// TestSearchAllTeams_Empty verifies that an empty API response (no matching
// teams) returns a zero-length slice without error.
func TestSearchAllTeams_Empty(t *testing.T) {
	srv := newSeedMockGrafana(t)
	ext := newExternalFromServer(t, srv)
	cr := &v1alpha1observe.Teams{
		Spec: v1alpha1observe.TeamsSpec{
			ForProvider: v1alpha1observe.TeamsParameters{Name: ptr("nonexistent-team")},
		},
	}

	got, err := ext.searchAllTeams(cr)
	if err != nil {
		t.Fatalf("searchAllTeams: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected 0 teams, got %d: %+v", len(got), got)
	}
}

// TestSearchAllTeams_Pagination verifies that searchAllTeams iterates through
// multiple pages when the API returns fewer teams than totalCount per page,
// collecting all 20 seed teams across pages.
func TestSearchAllTeams_Pagination(t *testing.T) {
	pageSize := int64(6)
	pagesRequested := 0

	// Mock that ignores the client's perPage and returns fixed-size pages
	// to force multiple round trips.
	srv := newMockGrafana(t, func(page, _ int64, name, query string) searchResult {
		pagesRequested++
		filtered := filterSeedTeams(name, query)
		total := int64(len(filtered))

		start := (page - 1) * pageSize
		if start > total {
			start = total
		}
		end := start + pageSize
		if end > total {
			end = total
		}

		return searchResult{
			TotalCount: total,
			Page:       page,
			PerPage:    pageSize,
			Teams:      filtered[start:end],
		}
	})

	ext := newExternalFromServer(t, srv)
	got, err := ext.searchAllTeams(&v1alpha1observe.Teams{})
	if err != nil {
		t.Fatalf("searchAllTeams paginated: %v", err)
	}
	if len(got) != len(seedTeams) {
		t.Errorf("expected %d teams after pagination, got %d", len(seedTeams), len(got))
	}
	// 20 teams at 6 per page = 4 pages (6+6+6+2).
	if pagesRequested != 4 {
		t.Errorf("expected 4 page requests, got %d", pagesRequested)
	}
}

// --- Observe unit tests ---

// TestObserve_UpToDate verifies that Observe reports ResourceUpToDate=true
// when the status already matches what the API returns.
func TestObserve_UpToDate(t *testing.T) {
	srv := newSeedMockGrafana(t)
	ext := newExternalFromServer(t, srv)
	cr := &v1alpha1observe.Teams{
		Status: v1alpha1observe.TeamsStatus{
			AtProvider: v1alpha1observe.TeamsObservation{Teams: seedSummaries()},
		},
	}

	obs, err := ext.Observe(context.Background(), cr)
	if err != nil {
		t.Fatalf("Observe: %v", err)
	}
	if !obs.ResourceExists {
		t.Error("expected ResourceExists=true")
	}
	if !obs.ResourceUpToDate {
		t.Error("expected ResourceUpToDate=true when status matches API")
	}
}

// TestObserve_NotUpToDate verifies that Observe reports ResourceUpToDate=false
// when the status is empty but the API returns teams, triggering an Update.
func TestObserve_NotUpToDate(t *testing.T) {
	srv := newSeedMockGrafana(t)
	ext := newExternalFromServer(t, srv)
	cr := &v1alpha1observe.Teams{}

	obs, err := ext.Observe(context.Background(), cr)
	if err != nil {
		t.Fatalf("Observe: %v", err)
	}
	if !obs.ResourceExists {
		t.Error("expected ResourceExists=true")
	}
	if obs.ResourceUpToDate {
		t.Error("expected ResourceUpToDate=false when status is empty but API has teams")
	}
}

// --- Update unit tests ---

// TestUpdate_PopulatesStatus verifies that Update writes the teams cached by
// Observe into cr.Status.AtProvider.Teams without making an additional API call.
func TestUpdate_PopulatesStatus(t *testing.T) {
	srv := newSeedMockGrafana(t)
	ext := newExternalFromServer(t, srv)
	cr := &v1alpha1observe.Teams{}

	// Observe first to populate lastObserved.
	_, err := ext.Observe(context.Background(), cr)
	if err != nil {
		t.Fatalf("Observe: %v", err)
	}

	_, err = ext.Update(context.Background(), cr)
	if err != nil {
		t.Fatalf("Update: %v", err)
	}

	want := seedSummaries()
	if diff := cmp.Diff(want, cr.Status.AtProvider.Teams); diff != "" {
		t.Errorf("status teams mismatch (-want +got):\n%s", diff)
	}
}

// --- Delete unit test ---

// TestDelete_IsNoop verifies that Delete succeeds without contacting the API,
// since Teams is an observe-only resource with nothing to delete in Grafana.
func TestDelete_IsNoop(t *testing.T) {
	ext := &external{teamsClient: oateams.New(nil, nil)}
	_, err := ext.Delete(context.Background(), &v1alpha1observe.Teams{})
	if err != nil {
		t.Errorf("Delete returned unexpected error: %v", err)
	}
}
