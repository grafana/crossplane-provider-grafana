package clients

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	smapi "github.com/grafana/synthetic-monitoring-api-go-client"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	clustersm "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster/sm/v1alpha1"
	namespacedsm "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/sm/v1alpha1"
)

const (
	// probeNamesHashAnnotation stores a hash of probeNames to detect changes
	// and avoid unnecessary API calls on every reconciliation.
	probeNamesHashAnnotation = "sm.grafana.crossplane.io/probe-names-hash"
)

// hashProbeNames computes a hash of the probe names for change detection.
func hashProbeNames(names []string) string {
	sorted := make([]string, len(names))
	copy(sorted, names)
	sort.Strings(sorted)
	h := sha256.Sum256([]byte(strings.Join(sorted, ",")))
	return hex.EncodeToString(h[:8]) // First 8 bytes is enough
}

// resolveProbeNames resolves probe names to probe IDs using the SM API.
// It returns the list of probe IDs corresponding to the given names.
func resolveProbeNames(ctx context.Context, smURL, smAccessToken string, probeNames []string) ([]*float64, error) {
	if smURL == "" || smAccessToken == "" {
		return nil, errors.New("sm_url and sm_access_token are required to resolve probe names")
	}

	smClient := smapi.NewClient(smURL, smAccessToken, nil)
	probes, err := smClient.ListProbes(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list probes from SM API")
	}

	// Build a map of probe name to ID and collect available names
	probeNameToID := make(map[string]int64)
	availableNames := make([]string, 0, len(probes))
	for _, p := range probes {
		probeNameToID[p.Name] = p.Id
		availableNames = append(availableNames, p.Name)
	}
	sort.Strings(availableNames)

	// Resolve each name to an ID
	probeIDs := make([]*float64, 0, len(probeNames))
	var missingNames []string
	for _, name := range probeNames {
		if id, ok := probeNameToID[name]; ok {
			idFloat := float64(id)
			probeIDs = append(probeIDs, &idFloat)
		} else {
			missingNames = append(missingNames, name)
		}
	}

	if len(missingNames) > 0 {
		return nil, fmt.Errorf("failed to resolve probe names to IDs. "+
			"Requested probe names: %v. "+
			"Probes not found: %v. "+
			"Available probes: %v",
			probeNames, missingNames, availableNames)
	}

	return probeIDs, nil
}

// resolveSMCheckProbeNames checks if the managed resource is an SM Check with ProbeNames set,
// and if so, resolves the names to IDs and updates the Probes field.
// It uses an annotation hash to avoid API calls when probeNames hasn't changed.
func resolveSMCheckProbeNames(ctx context.Context, c client.Client, mg resource.Managed, creds map[string]string, credSpec Config) error {
	// Get SM credentials
	smURL := creds["sm_url"]
	if credSpec.SMURL != "" {
		smURL = credSpec.SMURL
	}
	smAccessToken := creds["sm_access_token"]

	// Check if this is an SM Check resource with ProbeNames set
	switch check := mg.(type) {
	case *clustersm.Check:
		if len(check.Spec.ForProvider.ProbeNames) == 0 {
			return nil
		}

		// Check if hash matches to skip API call
		currentHash := hashProbeNames(check.Spec.ForProvider.ProbeNames)
		annotations := check.GetAnnotations()
		if annotations != nil && annotations[probeNamesHashAnnotation] == currentHash {
			// Hash matches, probeNames unchanged - skip API call
			return nil
		}

		probeIDs, err := resolveProbeNames(ctx, smURL, smAccessToken, check.Spec.ForProvider.ProbeNames)
		if err != nil {
			return err
		}

		// Update probes and set hash annotation
		check.Spec.ForProvider.Probes = probeIDs
		if annotations == nil {
			annotations = make(map[string]string)
		}
		annotations[probeNamesHashAnnotation] = currentHash
		check.SetAnnotations(annotations)

		if err := c.Update(ctx, check); err != nil {
			return errors.Wrap(err, "failed to update cluster SM Check with resolved probe IDs")
		}

	case *namespacedsm.Check:
		if len(check.Spec.ForProvider.ProbeNames) == 0 {
			return nil
		}

		// Check if hash matches to skip API call
		currentHash := hashProbeNames(check.Spec.ForProvider.ProbeNames)
		annotations := check.GetAnnotations()
		if annotations != nil && annotations[probeNamesHashAnnotation] == currentHash {
			// Hash matches, probeNames unchanged - skip API call
			return nil
		}

		probeIDs, err := resolveProbeNames(ctx, smURL, smAccessToken, check.Spec.ForProvider.ProbeNames)
		if err != nil {
			return err
		}

		// Update probes and set hash annotation
		check.Spec.ForProvider.Probes = probeIDs
		if annotations == nil {
			annotations = make(map[string]string)
		}
		annotations[probeNamesHashAnnotation] = currentHash
		check.SetAnnotations(annotations)

		if err := c.Update(ctx, check); err != nil {
			return errors.Wrap(err, "failed to update namespaced SM Check with resolved probe IDs")
		}
	}

	return nil
}
