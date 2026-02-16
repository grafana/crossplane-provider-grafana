/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"fmt"

	v1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	grafanaProvider "github.com/grafana/terraform-provider-grafana/v4/pkg/provider"
	terraformSDK "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	clusterapis "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster/v1beta1"
	namespacedapis "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errUpdateStatus         = "cannot update ProviderConfig status"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal grafana credentials as JSON"
)

type Credentials struct {
	Source                       v1.CredentialsSource
	v1.CommonCredentialSelectors `json:",inline"`
}

type Config struct {
	URL                string
	CloudAPIURL        string
	CloudProviderURL   string
	ConnectionsAPIURL  string
	FleetManagementURL string
	OnCallURL          string
	SMURL              string
	OrgID              *int
	StackID            *int
	Credentials        Credentials
}

func useLegacyProviderConfig(ctx context.Context, c client.Client, mg resource.LegacyManaged) (*Config, error) {
	ref := mg.GetProviderConfigReference()
	if ref == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	pc := &clusterapis.ProviderConfig{}
	if err := c.Get(ctx, types.NamespacedName{Name: ref.Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	tracker := resource.NewLegacyProviderConfigUsageTracker(c, &clusterapis.ProviderConfigUsage{})
	if err := tracker.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	if len(pc.Status.Conditions) == 0 {
		pc.Status.SetConditions(v1.Available())
		if err := c.Status().Update(ctx, pc); err != nil {
			return nil, errors.Wrap(err, errUpdateStatus)
		}
	}

	config := &Config{
		URL:                pc.Spec.URL,
		CloudAPIURL:        pc.Spec.CloudAPIURL,
		CloudProviderURL:   pc.Spec.CloudProviderURL,
		ConnectionsAPIURL:  pc.Spec.ConnectionsAPIURL,
		FleetManagementURL: pc.Spec.FleetManagementURL,
		OnCallURL:          pc.Spec.OnCallURL,
		SMURL:              pc.Spec.SMURL,
		OrgID:              pc.Spec.OrgID,
		StackID:            pc.Spec.StackID,
	}

	// Convert v1 to v2 types explicitly
	// Best-effort simple field copies (types differ between v1 and v2 runtime packages); we only use SecretRef today.
	config.Credentials.Source = v1.CredentialsSource(string(pc.Spec.Credentials.Source))
	if secret := pc.Spec.Credentials.SecretRef; secret != nil {
		config.Credentials.SecretRef = &v1.SecretKeySelector{
			SecretReference: v1.SecretReference{
				Name:      secret.Name,
				Namespace: secret.Namespace,
			},
			Key: secret.Key,
		}
	}

	return config, nil
}

func useModernProviderConfig(ctx context.Context, c client.Client, mg resource.ModernManaged) (*Config, error) {
	ref := mg.GetProviderConfigReference()
	if ref == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	var spec *namespacedapis.ProviderConfigSpec
	switch ref.Kind {
	case "ProviderConfig":
		pc := &namespacedapis.ProviderConfig{}
		if err := c.Get(ctx, types.NamespacedName{Name: ref.Name, Namespace: mg.GetNamespace()}, pc); err != nil {
			return nil, errors.Wrap(err, errGetProviderConfig)
		}
		spec = &pc.Spec

		if len(pc.Status.Conditions) == 0 {
			pc.Status.SetConditions(v1.Available())
			if err := c.Status().Update(ctx, pc); err != nil {
				return nil, errors.Wrap(err, errUpdateStatus)
			}
		}
	case "ClusterProviderConfig":
		cpc := &namespacedapis.ClusterProviderConfig{}
		if err := c.Get(ctx, types.NamespacedName{Name: ref.Name, Namespace: mg.GetNamespace()}, cpc); err != nil {
			return nil, errors.Wrap(err, errGetProviderConfig)
		}
		spec = &cpc.Spec

		if len(cpc.Status.Conditions) == 0 {
			cpc.Status.SetConditions(v1.Available())
			if err := c.Status().Update(ctx, cpc); err != nil {
				return nil, errors.Wrap(err, errUpdateStatus)
			}
		}
	default:
		return nil, fmt.Errorf("unsupported ProviderConfig kind: %s", ref.Kind)
	}

	tracker := resource.NewProviderConfigUsageTracker(c, &namespacedapis.ProviderConfigUsage{})
	if err := tracker.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return &Config{
		URL:                spec.URL,
		CloudAPIURL:        spec.CloudAPIURL,
		CloudProviderURL:   spec.CloudProviderURL,
		ConnectionsAPIURL:  spec.ConnectionsAPIURL,
		FleetManagementURL: spec.FleetManagementURL,
		OnCallURL:          spec.OnCallURL,
		SMURL:              spec.SMURL,
		OrgID:              spec.OrgID,
		StackID:            spec.StackID,
		Credentials: Credentials{
			Source:                    spec.Credentials.Source,
			CommonCredentialSelectors: spec.Credentials.CommonCredentialSelectors,
		},
	}, nil
}

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder() terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{}

		var credSpec Config

		switch mr := mg.(type) {
		case resource.LegacyManaged:
			config, err := useLegacyProviderConfig(ctx, client, mr)
			if err != nil {
				return ps, errors.Wrapf(err, "cannot use legacy provider config")
			}
			credSpec = *config
		case resource.ModernManaged:
			config, err := useModernProviderConfig(ctx, client, mr)
			if err != nil {
				return ps, errors.Wrapf(err, "cannot use modern provider config")
			}
			credSpec = *config
		default:
			return ps, fmt.Errorf("unknown managed resource type %T", mg)
		}

		data, err := resource.CommonCredentialExtractor(ctx, credSpec.Credentials.Source, client, credSpec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Resolve ProbeNames to Probes for SM Check resources
		if err := resolveSMCheckProbeNames(ctx, client, mg, creds, credSpec); err != nil {
			return ps, errors.Wrap(err, "failed to resolve probe names")
		}

		// Set credentials in Terraform provider configuration.
		// https://registry.terraform.io/providers/grafana/grafana/latest/docs
		ps.Configuration = map[string]any{}
		for _, k := range []string{
			"auth",
			"url",

			"cloud_access_policy_token",
			"cloud_api_url",

			"cloud_provider_access_token",
			"cloud_provider_url",

			"connections_api_access_token",
			"connections_api_url",

			"fleet_management_auth",
			"fleet_management_url",

			"frontend_o11y_api_access_token",

			"oncall_access_token",
			"oncall_url",

			"sm_access_token",
			"sm_url",

			"cloud_api_key", // don't see it in the TF config
			"org_id",        // don't see it in the TF config

			// required for k6 resources
			"stack_id",
			"k6_access_token",
		} {
			if v, ok := creds[k]; ok {
				ps.Configuration[k] = v
			}
		}

		if credSpec.URL != "" {
			ps.Configuration["url"] = credSpec.URL
		}
		if credSpec.CloudAPIURL != "" {
			ps.Configuration["cloud_api_url"] = credSpec.CloudAPIURL
		}
		if credSpec.CloudProviderURL != "" {
			ps.Configuration["cloud_provider_url"] = credSpec.CloudProviderURL
		}
		if credSpec.ConnectionsAPIURL != "" {
			ps.Configuration["connections_api_url"] = credSpec.ConnectionsAPIURL
		}
		if credSpec.FleetManagementURL != "" {
			ps.Configuration["fleet_management_url"] = credSpec.FleetManagementURL
		}
		if credSpec.OnCallURL != "" {
			ps.Configuration["oncall_url"] = credSpec.OnCallURL
		}
		if credSpec.SMURL != "" {
			ps.Configuration["sm_url"] = credSpec.SMURL
		}
		if credSpec.OrgID != nil {
			ps.Configuration["org_id"] = *credSpec.OrgID
		}
		if credSpec.StackID != nil {
			ps.Configuration["stack_id"] = *credSpec.StackID
		}

		if err := configureNoForkGrafanaClient(ctx, &ps); err != nil {
			return ps, errors.Wrap(err, "failed to configure the no-fork Grafana client")
		}

		return ps, nil
	}
}

func configureNoForkGrafanaClient(ctx context.Context, ps *terraform.Setup) error {
	ps.FrameworkProvider = grafanaProvider.FrameworkProvider("crossplane")

	cb := grafanaProvider.Provider("crossplane")
	diags := cb.Configure(ctx, terraformSDK.NewResourceConfigRaw(ps.Configuration))
	if diags.HasError() {
		return fmt.Errorf("failed to configure the Grafana provider: %v", diags)
	}

	ps.Meta = cb.Meta()
	return nil
}
