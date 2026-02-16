package clients

import (
	"context"
	"os"
	"strings"
	"testing"
)

/*
Integration tests for SM Check probe name resolution.

To run these tests locally, you need a Grafana Cloud stack with Synthetic Monitoring enabled.

Environment variables:
  SM_URL          - SM API URL (default: https://synthetic-monitoring-api-dev.grafana.net)
  SM_ACCESS_TOKEN - SM API access token (required)
  SM_PROBE_NAMES  - Comma-separated probe names to test (default: Paris,Oregon,Spain)

Run with:
  source .env && go test -v ./internal/clients/... -run Integration

To get SM credentials:
1. Go to your Grafana Cloud portal
2. Navigate to Synthetic Monitoring > Config
3. Generate or copy the API access token
*/

const (
	defaultSMURL      = "https://synthetic-monitoring-api-dev.grafana.net"
	defaultProbeNames = "Paris,Oregon,Spain"
)

// TestIntegration_ResolveProbeNames tests probe name resolution against a real SM API.
func TestIntegration_ResolveProbeNames(t *testing.T) {
	smURL := os.Getenv("SM_URL")
	if smURL == "" {
		smURL = defaultSMURL
	}

	smToken := os.Getenv("SM_ACCESS_TOKEN")
	if smToken == "" {
		t.Skip("Skipping integration test: SM_ACCESS_TOKEN environment variable not set")
	}

	probeNamesEnv := os.Getenv("SM_PROBE_NAMES")
	if probeNamesEnv == "" {
		probeNamesEnv = defaultProbeNames
	}
	probeNames := strings.Split(probeNamesEnv, ",")

	ctx := context.Background()

	// Test resolving a single probe name
	probeIDs, err := resolveProbeNames(ctx, smURL, smToken, []string{probeNames[0]})
	if err != nil {
		t.Fatalf("failed to resolve probe name %q: %v", probeNames[0], err)
	}
	if len(probeIDs) != 1 {
		t.Errorf("expected 1 probe ID, got %d", len(probeIDs))
	}
	t.Logf("Successfully resolved %q to ID %v", probeNames[0], *probeIDs[0])

	// Test resolving multiple probe names
	probeIDs, err = resolveProbeNames(ctx, smURL, smToken, probeNames)
	if err != nil {
		t.Fatalf("failed to resolve probe names %v: %v", probeNames, err)
	}
	if len(probeIDs) != len(probeNames) {
		t.Errorf("expected %d probe IDs, got %d", len(probeNames), len(probeIDs))
	}
	t.Logf("Successfully resolved %v to IDs: %v", probeNames, probeIDsToString(probeIDs))

	// Test error case: non-existent probe name
	_, err = resolveProbeNames(ctx, smURL, smToken, []string{"NonExistentProbeName12345"})
	if err == nil {
		t.Error("expected error for non-existent probe name, got nil")
	} else {
		t.Logf("Correctly got error for non-existent probe: %v", err)
	}
}

// probeIDsToString converts probe IDs to a string for logging.
func probeIDsToString(ids []*float64) []float64 {
	result := make([]float64, len(ids))
	for i, id := range ids {
		if id != nil {
			result[i] = *id
		}
	}
	return result
}
