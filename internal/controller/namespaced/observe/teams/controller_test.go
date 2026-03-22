package teams

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"

	v1alpha1observe "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/observe/v1alpha1"
	"github.com/grafana/crossplane-provider-grafana/v2/internal/clients"

	oateams "github.com/grafana/grafana-openapi-client-go/client/teams"
)

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

// TestSearchAllTeams_NoFilter verifies that an unfiltered search returns all
// teams from the API with all fields correctly mapped to TeamSummary.
func TestSearchAllTeams_NoFilter(t *testing.T) {
	want := []v1alpha1observe.TeamSummary{
		{ID: 1, UID: "aaa", Name: "Alpha", Email: "alpha@example.com", MemberCount: 3, OrgID: 1},
		{ID: 2, UID: "bbb", Name: "Beta", Email: "beta@example.com", MemberCount: 5, OrgID: 1},
	}

	srv := newMockGrafana(t, func(_, _ int64, _, _ string) searchResult {
		return searchResult{
			TotalCount: 2,
			Teams: []teamPayload{
				{ID: 1, UID: "aaa", Name: "Alpha", Email: "alpha@example.com", MemberCount: 3, OrgID: 1},
				{ID: 2, UID: "bbb", Name: "Beta", Email: "beta@example.com", MemberCount: 5, OrgID: 1},
			},
		}
	})

	ext := newExternalFromServer(t, srv)
	cr := &v1alpha1observe.Teams{}

	got, err := ext.searchAllTeams(cr)
	if err != nil {
		t.Fatalf("searchAllTeams: %v", err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("teams mismatch (-want +got):\n%s", diff)
	}
}

// TestSearchAllTeams_NameFilter verifies that setting Spec.ForProvider.Name
// passes the name parameter to the API and returns only exact matches.
func TestSearchAllTeams_NameFilter(t *testing.T) {
	srv := newMockGrafana(t, func(_, _ int64, name, _ string) searchResult {
		if name != "Alpha" {
			return searchResult{TotalCount: 0}
		}
		return searchResult{
			TotalCount: 1,
			Teams:      []teamPayload{{ID: 1, UID: "aaa", Name: "Alpha", OrgID: 1}},
		}
	})

	ext := newExternalFromServer(t, srv)
	cr := &v1alpha1observe.Teams{
		Spec: v1alpha1observe.TeamsSpec{
			ForProvider: v1alpha1observe.TeamsParameters{Name: ptr("Alpha")},
		},
	}

	got, err := ext.searchAllTeams(cr)
	if err != nil {
		t.Fatalf("searchAllTeams: %v", err)
	}
	if len(got) != 1 || got[0].Name != "Alpha" {
		t.Errorf("expected 1 team named Alpha, got %+v", got)
	}
}

// TestSearchAllTeams_QueryFilter verifies that setting Spec.ForProvider.Query
// passes the query parameter to the API for substring matching.
func TestSearchAllTeams_QueryFilter(t *testing.T) {
	srv := newMockGrafana(t, func(_, _ int64, _, query string) searchResult {
		if query != "ops" {
			return searchResult{TotalCount: 0}
		}
		return searchResult{
			TotalCount: 2,
			Teams: []teamPayload{
				{ID: 10, Name: "ops-eu", OrgID: 1},
				{ID: 11, Name: "ops-us", OrgID: 1},
			},
		}
	})

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
	if len(got) != 2 {
		t.Errorf("expected 2 teams, got %d", len(got))
	}
}

// TestSearchAllTeams_Empty verifies that an empty API response (no teams)
// returns a zero-length slice without error.
func TestSearchAllTeams_Empty(t *testing.T) {
	srv := newMockGrafana(t, func(_, _ int64, _, _ string) searchResult {
		return searchResult{TotalCount: 0, Teams: nil}
	})

	ext := newExternalFromServer(t, srv)
	got, err := ext.searchAllTeams(&v1alpha1observe.Teams{})
	if err != nil {
		t.Fatalf("searchAllTeams: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected 0 teams, got %d: %+v", len(got), got)
	}
}

// TestSearchAllTeams_Pagination verifies that searchAllTeams iterates through
// multiple pages when the API returns fewer teams than totalCount per page,
// collecting all teams across pages.
func TestSearchAllTeams_Pagination(t *testing.T) {
	const total = 5
	allTeams := []teamPayload{
		{ID: 1, Name: "T1", OrgID: 1},
		{ID: 2, Name: "T2", OrgID: 1},
		{ID: 3, Name: "T3", OrgID: 1},
		{ID: 4, Name: "T4", OrgID: 1},
		{ID: 5, Name: "T5", OrgID: 1},
	}

	pagesRequested := 0
	srv := newMockGrafana(t, func(page, perPage int64, _, _ string) searchResult {
		pagesRequested++
		start := (page - 1) * perPage
		end := start + perPage
		if end > int64(len(allTeams)) {
			end = int64(len(allTeams))
		}
		return searchResult{
			TotalCount: total,
			Teams:      allTeams[start:end],
		}
	})

	// Build client with a small page size (2) to force pagination.
	grafanaClient, err := clients.NewOAPIClient(srv.URL, "test-token")
	if err != nil {
		t.Fatalf("NewOAPIClient: %v", err)
	}
	ext := &external{teamsClient: grafanaClient.Teams}

	// Override page size by driving the loop manually through SearchTeams.
	// Instead, test via searchAllTeams but mock the server to return 2 per page.
	srv2 := newMockGrafana(t, func(page, perPage int64, _, _ string) searchResult {
		pagesRequested++
		pageSize := int64(2)
		start := (page - 1) * pageSize
		end := start + pageSize
		if end > int64(len(allTeams)) {
			end = int64(len(allTeams))
		}
		return searchResult{
			TotalCount: total,
			Page:       page,
			PerPage:    pageSize,
			Teams:      allTeams[start:end],
		}
	})
	pagesRequested = 0
	grafanaClient2, err := clients.NewOAPIClient(srv2.URL, "test-token")
	if err != nil {
		t.Fatalf("NewOAPIClient: %v", err)
	}
	ext2 := &external{teamsClient: grafanaClient2.Teams}

	// searchAllTeams requests pages of 1000; the mock ignores perPage and
	// returns 2 per page with totalCount=5, so 3 pages are needed.
	got, err := ext2.searchAllTeams(&v1alpha1observe.Teams{})
	if err != nil {
		t.Fatalf("searchAllTeams paginated: %v", err)
	}
	if len(got) != total {
		t.Errorf("expected %d teams after pagination, got %d", total, len(got))
	}
	if pagesRequested < 3 {
		t.Errorf("expected at least 3 page requests, got %d", pagesRequested)
	}

	// silence the unused variable warning for ext
	_ = ext
}

// --- Observe unit tests ---

// TestObserve_UpToDate verifies that Observe reports ResourceUpToDate=true
// when the status already matches what the API returns.
func TestObserve_UpToDate(t *testing.T) {
	existing := []v1alpha1observe.TeamSummary{
		{ID: 1, Name: "Alpha", OrgID: 1},
	}
	srv := newMockGrafana(t, func(_, _ int64, _, _ string) searchResult {
		return searchResult{
			TotalCount: 1,
			Teams:      []teamPayload{{ID: 1, Name: "Alpha", OrgID: 1}},
		}
	})

	ext := newExternalFromServer(t, srv)
	cr := &v1alpha1observe.Teams{
		Status: v1alpha1observe.TeamsStatus{
			AtProvider: v1alpha1observe.TeamsObservation{Teams: existing},
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
	srv := newMockGrafana(t, func(_, _ int64, _, _ string) searchResult {
		return searchResult{
			TotalCount: 1,
			Teams:      []teamPayload{{ID: 1, Name: "Alpha", OrgID: 1}},
		}
	})

	ext := newExternalFromServer(t, srv)
	// Status is empty → does not match → not up to date.
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

// TestUpdate_PopulatesStatus verifies that Update fetches teams from the API
// and writes them into cr.Status.AtProvider.Teams.
func TestUpdate_PopulatesStatus(t *testing.T) {
	srv := newMockGrafana(t, func(_, _ int64, _, _ string) searchResult {
		return searchResult{
			TotalCount: 2,
			Teams: []teamPayload{
				{ID: 1, UID: "aaa", Name: "Alpha", Email: "alpha@example.com", MemberCount: 2, OrgID: 1},
				{ID: 2, UID: "bbb", Name: "Beta", Email: "beta@example.com", MemberCount: 4, OrgID: 1},
			},
		}
	})

	ext := newExternalFromServer(t, srv)
	cr := &v1alpha1observe.Teams{}

	_, err := ext.Update(context.Background(), cr)
	if err != nil {
		t.Fatalf("Update: %v", err)
	}

	if len(cr.Status.AtProvider.Teams) != 2 {
		t.Fatalf("expected 2 teams in status, got %d", len(cr.Status.AtProvider.Teams))
	}
	if cr.Status.AtProvider.Teams[0].Name != "Alpha" {
		t.Errorf("expected first team Alpha, got %q", cr.Status.AtProvider.Teams[0].Name)
	}
	if cr.Status.AtProvider.Teams[1].Name != "Beta" {
		t.Errorf("expected second team Beta, got %q", cr.Status.AtProvider.Teams[1].Name)
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
