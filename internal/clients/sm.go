package clients

import (
	"context"
	"fmt"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	smapi "github.com/grafana/synthetic-monitoring-api-go-client"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	clustersm "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster/sm/v1alpha1"
	namespacedsm "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/sm/v1alpha1"
)

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

	// Build a map of probe name to ID
	probeNameToID := make(map[string]int64)
	for _, p := range probes {
		probeNameToID[p.Name] = p.Id
	}

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
		return nil, fmt.Errorf("probe names not found: %v", missingNames)
	}

	return probeIDs, nil
}

// resolveSMCheckProbeNames checks if the managed resource is an SM Check with ProbeNames set,
// and if so, resolves the names to IDs and updates the Probes field.
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
		probeIDs, err := resolveProbeNames(ctx, smURL, smAccessToken, check.Spec.ForProvider.ProbeNames)
		if err != nil {
			return err
		}
		check.Spec.ForProvider.Probes = probeIDs
		if err := c.Update(ctx, check); err != nil {
			return errors.Wrap(err, "failed to update cluster SM Check with resolved probe IDs")
		}

	case *namespacedsm.Check:
		if len(check.Spec.ForProvider.ProbeNames) == 0 {
			return nil
		}
		probeIDs, err := resolveProbeNames(ctx, smURL, smAccessToken, check.Spec.ForProvider.ProbeNames)
		if err != nil {
			return err
		}
		check.Spec.ForProvider.Probes = probeIDs
		if err := c.Update(ctx, check); err != nil {
			return errors.Wrap(err, "failed to update namespaced SM Check with resolved probe IDs")
		}
	}

	return nil
}
