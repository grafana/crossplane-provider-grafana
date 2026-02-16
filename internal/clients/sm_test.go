package clients

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	synthetic_monitoring "github.com/grafana/synthetic-monitoring-agent/pkg/pb/synthetic_monitoring"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	clustersm "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster/sm/v1alpha1"
	namespacedsm "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/sm/v1alpha1"
)

func createSMScheme(t *testing.T) *runtime.Scheme {
	t.Helper()
	scheme := runtime.NewScheme()
	_ = clustersm.SchemeBuilder.AddToScheme(scheme)
	_ = namespacedsm.SchemeBuilder.AddToScheme(scheme)
	return scheme
}

func createMockSMServer(t *testing.T, probes []synthetic_monitoring.Probe) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/probe/list" {
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(probes); err != nil {
				t.Fatalf("failed to encode probes: %v", err)
			}
			return
		}
		http.NotFound(w, r)
	}))
}

// TestResolveProbeNames verifies that probe names are correctly resolved to IDs via the SM API.
func TestResolveProbeNames(t *testing.T) {
	cases := []struct {
		name          string
		probes        []synthetic_monitoring.Probe
		probeNames    []string
		wantIDs       []*float64
		wantErr       bool
		wantErrSubstr string
	}{
		{
			name: "resolves single probe name",
			probes: []synthetic_monitoring.Probe{
				{Id: 1, Name: "Amsterdam"},
				{Id: 2, Name: "New York"},
			},
			probeNames: []string{"Amsterdam"},
			wantIDs:    []*float64{float64Ptr(1.0)},
		},
		{
			name: "resolves multiple probe names",
			probes: []synthetic_monitoring.Probe{
				{Id: 1, Name: "Amsterdam"},
				{Id: 2, Name: "New York"},
				{Id: 3, Name: "Tokyo"},
			},
			probeNames: []string{"Amsterdam", "Tokyo"},
			wantIDs:    []*float64{float64Ptr(1.0), float64Ptr(3.0)},
		},
		{
			name: "returns error for missing probe name",
			probes: []synthetic_monitoring.Probe{
				{Id: 1, Name: "Amsterdam"},
			},
			probeNames:    []string{"Amsterdam", "NonExistent"},
			wantErr:       true,
			wantErrSubstr: "probe names not found: [NonExistent]",
		},
		{
			name: "returns error for all missing probe names",
			probes: []synthetic_monitoring.Probe{
				{Id: 1, Name: "Amsterdam"},
			},
			probeNames:    []string{"NonExistent1", "NonExistent2"},
			wantErr:       true,
			wantErrSubstr: "probe names not found:",
		},
		{
			name:       "handles empty probe names list",
			probes:     []synthetic_monitoring.Probe{{Id: 1, Name: "Amsterdam"}},
			probeNames: []string{},
			wantIDs:    []*float64{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			server := createMockSMServer(t, tc.probes)
			defer server.Close()

			ids, err := resolveProbeNames(context.Background(), server.URL, "test-token", tc.probeNames)

			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tc.wantErrSubstr)
				}
				if tc.wantErrSubstr != "" && !containsSubstring(err.Error(), tc.wantErrSubstr) {
					t.Errorf("expected error containing %q, got %q", tc.wantErrSubstr, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if diff := cmp.Diff(tc.wantIDs, ids); diff != "" {
				t.Errorf("probe IDs mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// TestResolveProbeNamesMissingCredentials verifies that an error is returned when SM credentials are missing.
func TestResolveProbeNamesMissingCredentials(t *testing.T) {
	cases := []struct {
		name       string
		smURL      string
		smToken    string
		wantErrMsg string
	}{
		{
			name:       "missing SM URL",
			smURL:      "",
			smToken:    "token",
			wantErrMsg: "sm_url and sm_access_token are required",
		},
		{
			name:       "missing SM token",
			smURL:      "http://example.com",
			smToken:    "",
			wantErrMsg: "sm_url and sm_access_token are required",
		},
		{
			name:       "missing both",
			smURL:      "",
			smToken:    "",
			wantErrMsg: "sm_url and sm_access_token are required",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := resolveProbeNames(context.Background(), tc.smURL, tc.smToken, []string{"test"})
			if err == nil {
				t.Fatal("expected error, got nil")
			}
			if !containsSubstring(err.Error(), tc.wantErrMsg) {
				t.Errorf("expected error containing %q, got %q", tc.wantErrMsg, err.Error())
			}
		})
	}
}

// TestResolveSMCheckProbeNames_ClusterCheck verifies probe name resolution for cluster-scoped SM Check resources.
func TestResolveSMCheckProbeNames_ClusterCheck(t *testing.T) {
	probes := []synthetic_monitoring.Probe{
		{Id: 1, Name: "Amsterdam"},
		{Id: 2, Name: "New York"},
	}
	server := createMockSMServer(t, probes)
	defer server.Close()

	check := &clustersm.Check{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-check",
		},
		Spec: clustersm.CheckSpec{
			ForProvider: clustersm.CheckParameters{
				ProbeNames: []string{"Amsterdam", "New York"},
			},
		},
	}

	scheme := createSMScheme(t)
	client := fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(check).
		Build()

	creds := map[string]string{
		"sm_url":          server.URL,
		"sm_access_token": "test-token",
	}

	err := resolveSMCheckProbeNames(context.Background(), client, check, creds, Config{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify the check was updated with probe IDs
	updatedCheck := &clustersm.Check{}
	if err := client.Get(context.Background(), types.NamespacedName{Name: "test-check"}, updatedCheck); err != nil {
		t.Fatalf("failed to get updated check: %v", err)
	}

	wantProbes := []*float64{float64Ptr(1.0), float64Ptr(2.0)}
	if diff := cmp.Diff(wantProbes, updatedCheck.Spec.ForProvider.Probes); diff != "" {
		t.Errorf("probes mismatch (-want +got):\n%s", diff)
	}
}

// TestResolveSMCheckProbeNames_NamespacedCheck verifies probe name resolution for namespaced SM Check resources.
func TestResolveSMCheckProbeNames_NamespacedCheck(t *testing.T) {
	probes := []synthetic_monitoring.Probe{
		{Id: 10, Name: "Tokyo"},
		{Id: 20, Name: "Sydney"},
	}
	server := createMockSMServer(t, probes)
	defer server.Close()

	check := &namespacedsm.Check{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-check",
			Namespace: "default",
		},
		Spec: namespacedsm.CheckSpec{
			ForProvider: namespacedsm.CheckParameters{
				ProbeNames: []string{"Tokyo"},
			},
		},
	}

	scheme := createSMScheme(t)
	client := fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(check).
		Build()

	creds := map[string]string{
		"sm_url":          server.URL,
		"sm_access_token": "test-token",
	}

	err := resolveSMCheckProbeNames(context.Background(), client, check, creds, Config{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify the check was updated with probe IDs
	updatedCheck := &namespacedsm.Check{}
	if err := client.Get(context.Background(), types.NamespacedName{Name: "test-check", Namespace: "default"}, updatedCheck); err != nil {
		t.Fatalf("failed to get updated check: %v", err)
	}

	wantProbes := []*float64{float64Ptr(10.0)}
	if diff := cmp.Diff(wantProbes, updatedCheck.Spec.ForProvider.Probes); diff != "" {
		t.Errorf("probes mismatch (-want +got):\n%s", diff)
	}
}

// TestResolveSMCheckProbeNames_NoProbeNames verifies that no API call is made when ProbeNames is empty.
func TestResolveSMCheckProbeNames_NoProbeNames(t *testing.T) {
	check := &clustersm.Check{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-check",
		},
		Spec: clustersm.CheckSpec{
			ForProvider: clustersm.CheckParameters{
				// ProbeNames is nil/empty
			},
		},
	}

	scheme := createSMScheme(t)
	client := fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(check).
		Build()

	// No SM server - should not be called
	creds := map[string]string{}

	err := resolveSMCheckProbeNames(context.Background(), client, check, creds, Config{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// TestResolveSMCheckProbeNames_SMURLFromCredSpec verifies that SM URL from Config overrides the credentials map.
func TestResolveSMCheckProbeNames_SMURLFromCredSpec(t *testing.T) {
	probes := []synthetic_monitoring.Probe{
		{Id: 1, Name: "Amsterdam"},
	}
	server := createMockSMServer(t, probes)
	defer server.Close()

	check := &clustersm.Check{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-check",
		},
		Spec: clustersm.CheckSpec{
			ForProvider: clustersm.CheckParameters{
				ProbeNames: []string{"Amsterdam"},
			},
		},
	}

	scheme := createSMScheme(t)
	client := fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(check).
		Build()

	// SM URL from credSpec should override creds
	creds := map[string]string{
		"sm_url":          "http://wrong-url",
		"sm_access_token": "test-token",
	}
	credSpec := Config{
		SMURL: server.URL, // This should be used instead
	}

	err := resolveSMCheckProbeNames(context.Background(), client, check, creds, credSpec)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// Helper functions

func float64Ptr(f float64) *float64 {
	return &f
}

func containsSubstring(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > 0 && len(substr) > 0 && searchSubstring(s, substr)))
}

func searchSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
