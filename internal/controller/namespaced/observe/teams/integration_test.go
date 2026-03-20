package teams

/*
Integration tests for the Teams observe controller against a real Grafana instance.

Environment variables:

  GRAFANA_URL    Grafana base URL, e.g. https://my-instance.grafana.net
                 (default: http://localhost:3000)
  GRAFANA_AUTH   API token or "user:password" basic-auth string (required)

Run with:

  GRAFANA_AUTH=admin:admin go test -v ./internal/controller/namespaced/observe/teams/... -run Integration
  GRAFANA_AUTH=glsa_mytoken go test -v ./internal/controller/namespaced/observe/teams/... -run Integration

If GRAFANA_AUTH is not set the integration tests are skipped automatically.

To run against a local Docker container (started automatically, local development only):

  GRAFANA_DOCKER=1 go test -v -run Integration ./internal/controller/namespaced/observe/teams/...
  make go.test.integration.docker

In CI, set GRAFANA_AUTH (and optionally GRAFANA_URL) and use go.test.integration instead.
*/

import (
	"context"
	"os"
	"sort"
	"testing"

	v1alpha1observe "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/observe/v1alpha1"
	"github.com/grafana/crossplane-provider-grafana/v2/internal/clients"
)

const defaultGrafanaURL = "http://localhost:3000"

func grafanaClientFromEnv(t *testing.T) *external {
	t.Helper()

	auth := os.Getenv("GRAFANA_AUTH")
	if auth == "" {
		t.Skip("skipping integration test: GRAFANA_AUTH environment variable not set")
	}

	url := os.Getenv("GRAFANA_URL")
	if url == "" {
		url = defaultGrafanaURL
	}

	c, err := clients.NewOAPIClient(url, auth)
	if err != nil {
		t.Fatalf("NewOAPIClient(%q): %v", url, err)
	}
	return &external{teamsClient: c.Teams}
}

// TestIntegration_SearchAllTeams verifies that SearchTeams returns a valid
// (possibly empty) list without error.
func TestIntegration_SearchAllTeams(t *testing.T) {
	ext := grafanaClientFromEnv(t)

	teams, err := ext.searchAllTeams(&v1alpha1observe.Teams{})
	if err != nil {
		t.Fatalf("searchAllTeams: %v", err)
	}
	t.Logf("found %d team(s)", len(teams))
	for _, tm := range teams {
		t.Logf("  id=%d uid=%q name=%q members=%d orgId=%d", tm.ID, tm.UID, tm.Name, tm.MemberCount, tm.OrgID)
	}

	if dockerMode && len(teams) < len(seededTeams.All) {
		t.Errorf("dockerMode: expected at least %d teams (seeded), got %d", len(seededTeams.All), len(teams))
	}
}

// TestIntegration_Observe verifies the full Observe cycle.
// First call should report ResourceUpToDate=false (empty status).
// A subsequent call after Update should report ResourceUpToDate=true.
func TestIntegration_Observe(t *testing.T) {
	ext := grafanaClientFromEnv(t)
	ctx := context.Background()

	cr := &v1alpha1observe.Teams{}

	obs, err := ext.Observe(ctx, cr)
	if err != nil {
		t.Fatalf("Observe (first): %v", err)
	}
	if !obs.ResourceExists {
		t.Error("ResourceExists should always be true for Teams")
	}

	// Populate status via Update (simulating what the reconciler does).
	if _, err := ext.Update(ctx, cr); err != nil {
		t.Fatalf("Update: %v", err)
	}
	t.Logf("status populated with %d team(s)", len(cr.Status.AtProvider.Teams))

	// Now Observe again — status should match the API.
	obs, err = ext.Observe(ctx, cr)
	if err != nil {
		t.Fatalf("Observe (second): %v", err)
	}
	if !obs.ResourceUpToDate {
		t.Error("ResourceUpToDate should be true after Update")
	}
}

// TestIntegration_NameFilter verifies that the name filter reaches the API.
// In dockerMode the expected team name is deterministic (seededTeams.ExactName).
// Otherwise the test uses the first team returned by an unfiltered search.
func TestIntegration_NameFilter(t *testing.T) {
	ext := grafanaClientFromEnv(t)

	var target string
	if dockerMode {
		target = seededTeams.ExactName
	} else {
		all, err := ext.searchAllTeams(&v1alpha1observe.Teams{})
		if err != nil {
			t.Fatalf("searchAllTeams (all): %v", err)
		}
		if len(all) == 0 {
			t.Skip("no teams in this Grafana instance; skipping filter sub-test")
		}
		target = all[0].Name
	}

	cr := &v1alpha1observe.Teams{
		Spec: v1alpha1observe.TeamsSpec{
			ForProvider: v1alpha1observe.TeamsParameters{Name: &target},
		},
	}
	filtered, err := ext.searchAllTeams(cr)
	if err != nil {
		t.Fatalf("searchAllTeams (name=%q): %v", target, err)
	}
	if len(filtered) == 0 {
		t.Errorf("expected at least 1 result for name=%q, got 0", target)
	}
	for _, tm := range filtered {
		if tm.Name != target {
			t.Errorf("name filter returned unexpected team %q (wanted %q)", tm.Name, target)
		}
	}
	t.Logf("name=%q → %d result(s)", target, len(filtered))
}

// TestIntegration_QueryFilter verifies that the query filter returns teams
// whose name contains the query string.
// In dockerMode it asserts exactly the two ops-* teams were seeded and returned.
func TestIntegration_QueryFilter(t *testing.T) {
	ext := grafanaClientFromEnv(t)

	query := "ops"
	cr := &v1alpha1observe.Teams{
		Spec: v1alpha1observe.TeamsSpec{
			ForProvider: v1alpha1observe.TeamsParameters{Query: &query},
		},
	}

	got, err := ext.searchAllTeams(cr)
	if err != nil {
		t.Fatalf("searchAllTeams (query=%q): %v", query, err)
	}
	t.Logf("query=%q → %d result(s)", query, len(got))
	for _, tm := range got {
		t.Logf("  id=%d name=%q", tm.ID, tm.Name)
	}

	if dockerMode {
		if len(got) != len(seededTeams.OpsQuery) {
			t.Errorf("dockerMode: expected %d teams for query=%q, got %d", len(seededTeams.OpsQuery), query, len(got))
		}
		gotNames := make([]string, len(got))
		for i, tm := range got {
			gotNames[i] = tm.Name
		}
		sort.Strings(gotNames)
		want := make([]string, len(seededTeams.OpsQuery))
		copy(want, seededTeams.OpsQuery)
		sort.Strings(want)
		for i := range want {
			if i >= len(gotNames) || gotNames[i] != want[i] {
				t.Errorf("dockerMode: team names mismatch: want %v, got %v", want, gotNames)
				break
			}
		}
	}
}
