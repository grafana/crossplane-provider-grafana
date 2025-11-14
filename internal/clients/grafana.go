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

	clusterapis "github.com/grafana/crossplane-provider-grafana/apis/cluster/v1beta1"
	namespacedapis "github.com/grafana/crossplane-provider-grafana/apis/namespaced/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal grafana credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder() terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{}

		// Attempt to resolve modern (namespaced) ProviderConfig first; fall back to legacy (cluster) ProviderConfig.
		// ModernManaged exposes ProviderConfigReference (typed) for namespaced resources.
		var pcURLName string
		if mm, ok := mg.(resource.ModernManaged); ok && mm.GetProviderConfigReference() != nil {
			pcURLName = mm.GetProviderConfigReference().Name
		} else if legacy, ok := mg.(interface {
			GetProviderConfigReference() *v1.ProviderConfigReference
		}); ok && legacy.GetProviderConfigReference() != nil {
			pcURLName = legacy.GetProviderConfigReference().Name
		}
		if pcURLName == "" {
			return ps, errors.New(errNoProviderConfig)
		}

		// Prefer namespaced ProviderConfig if it exists in the namespace of the MR.
		// We don't know the namespace of cluster-scoped MRs (they are cluster-scoped), so only attempt namespaced lookup when mg has a namespace.
		var credSpec struct {
			URL                string
			CloudAPIURL        string
			CloudProviderURL   string
			ConnectionsAPIURL  string
			FleetManagementURL string
			OnCallURL          string
			SMURL              string
			OrgID              *int
			StackID            *int
			Credentials        struct {
				Source                       v1.CredentialsSource
				v1.CommonCredentialSelectors `json:",inline"`
			}
		}

		// pcStatus conditions update will differ by type; capture minimal interface.
		type statusSetter interface {
			SetConditions(...v1.Condition)
			GetCondition(v1.ConditionType) v1.Condition
		}

		// Prefer namespaced ProviderConfig when managed resource itself is namespaced.
		if nsGetter, ok := mg.(interface{ GetNamespace() string }); ok && nsGetter.GetNamespace() != "" {
			npc := &namespacedapis.ProviderConfig{}
			if err := client.Get(ctx, types.NamespacedName{Name: pcURLName, Namespace: nsGetter.GetNamespace()}, npc); err == nil {
				credSpec.URL = npc.Spec.URL
				credSpec.CloudAPIURL = npc.Spec.CloudAPIURL
				credSpec.CloudProviderURL = npc.Spec.CloudProviderURL
				credSpec.ConnectionsAPIURL = npc.Spec.ConnectionsAPIURL
				credSpec.FleetManagementURL = npc.Spec.FleetManagementURL
				credSpec.OnCallURL = npc.Spec.OnCallURL
				credSpec.SMURL = npc.Spec.SMURL
				credSpec.OrgID = npc.Spec.OrgID
				credSpec.StackID = npc.Spec.StackID
				credSpec.Credentials.Source = npc.Spec.Credentials.Source
				credSpec.Credentials.CommonCredentialSelectors = npc.Spec.Credentials.CommonCredentialSelectors
			}
		}
		// Fall back to cluster-scoped ProviderConfig if URL not set yet.
		if credSpec.URL == "" {
			cpc := &clusterapis.ProviderConfig{}
			if err := client.Get(ctx, types.NamespacedName{Name: pcURLName}, cpc); err != nil {
				return ps, errors.Wrap(err, errGetProviderConfig)
			}
			credSpec.URL = cpc.Spec.URL
			credSpec.CloudAPIURL = cpc.Spec.CloudAPIURL
			credSpec.CloudProviderURL = cpc.Spec.CloudProviderURL
			credSpec.ConnectionsAPIURL = cpc.Spec.ConnectionsAPIURL
			credSpec.FleetManagementURL = cpc.Spec.FleetManagementURL
			credSpec.OnCallURL = cpc.Spec.OnCallURL
			credSpec.SMURL = cpc.Spec.SMURL
			credSpec.OrgID = cpc.Spec.OrgID
			credSpec.StackID = cpc.Spec.StackID
			// convert v1 to v2 types explicitly
			// Best-effort simple field copies (types differ between v1 and v2 runtime packages); we only use SecretRef today.
			credSpec.Credentials.Source = v1.CredentialsSource(string(cpc.Spec.Credentials.Source))
			if cpc.Spec.Credentials.SecretRef != nil {
				credSpec.Credentials.SecretRef = &v1.SecretKeySelector{SecretReference: v1.SecretReference{Name: cpc.Spec.Credentials.SecretRef.Name, Namespace: cpc.Spec.Credentials.SecretRef.Namespace}, Key: cpc.Spec.Credentials.SecretRef.Key}
			}
		}

		data, err := resource.CommonCredentialExtractor(ctx, credSpec.Credentials.Source, client, credSpec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
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
