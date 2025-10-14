/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"fmt"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/terraform"
	grafanaProvider "github.com/grafana/terraform-provider-grafana/v4/pkg/provider"
	terraformSDK "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/grafana/crossplane-provider-grafana/apis/v1beta1"
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

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
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

		if pc.Spec.URL != "" {
			ps.Configuration["url"] = pc.Spec.URL
		}
		if pc.Spec.CloudAPIURL != "" {
			ps.Configuration["cloud_api_url"] = pc.Spec.CloudAPIURL
		}
		if pc.Spec.CloudProviderURL != "" {
			ps.Configuration["cloud_provider_url"] = pc.Spec.CloudProviderURL
		}
		if pc.Spec.ConnectionsAPIURL != "" {
			ps.Configuration["connections_api_url"] = pc.Spec.ConnectionsAPIURL
		}
		if pc.Spec.FleetManagementURL != "" {
			ps.Configuration["fleet_management_url"] = pc.Spec.FleetManagementURL
		}
		if pc.Spec.OnCallURL != "" {
			ps.Configuration["oncall_url"] = pc.Spec.OnCallURL
		}
		if pc.Spec.SMURL != "" {
			ps.Configuration["sm_url"] = pc.Spec.SMURL
		}
		if pc.Spec.OrgID != nil {
			ps.Configuration["org_id"] = *pc.Spec.OrgID
		}
		if pc.Spec.StackID != nil {
			ps.Configuration["stack_id"] = *pc.Spec.StackID
		}

		if err := configureNoForkGrafanaClient(ctx, &ps); err != nil {
			return ps, errors.Wrap(err, "failed to configure the no-fork Grafana client")
		}

		// Set Ready condition to true and write back the status.
		if len(pc.Status.Conditions) == 0 {
			pc.Status.SetConditions(v1.Available())
			if err := client.Status().Update(ctx, pc); err != nil {
				return ps, errors.Wrap(err, "cannot update ProviderConfig status")
			}
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
