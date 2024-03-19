/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/terraform-svchost/auth"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

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
			"url",
			"auth",
			"cloud_api_key",
			"cloud_api_url",
			"oncall_access_token",
			"oncall_url",
			"sm_access_token",
			"sm_url",
			"org_id",
		} {
			if v, ok := creds[k]; ok {
				ps.Configuration[k] = v
			}
		}
		return ps, errors.Wrap(configureNoForkGrafanaClient(pc, &ps), "failed to configure the no-fork Azure client")
	}
}

func configureNoForkGrafanaClient(pc *v1beta1.ProviderConfig, ps *terraform.Setup) error {
	cb := grafanaProvider.AzureClientBuilder{}
	switch pc.Spec.Credentials.Source { //nolint:exhaustive
	case xpv1.CredentialsSourceSecret:
		cb.SubscriptionID = ps.Configuration[keySubscriptionID].(string)
		cb.AuthConfig = &auth.Credentials{
			ClientID:                              ps.Configuration[keyClientID].(string),
			TenantID:                              ps.Configuration[keyTenantID].(string),
			ClientSecret:                          ps.Configuration[keyClientSecret].(string),
			EnableAuthenticatingUsingClientSecret: true,
		}
	}
	// TODO: we need to check how to prepare environment's
	// authorization context, this may differ especially if
	// we are preparing an environment inside an EKS cluster, etc.
	cb.AuthConfig.Environment = *environments.AzurePublic()
	c, err := cb.GetClient(context.WithoutCancel(ctx))
	if err != nil {
		return errors.Wrap(err, "failed to build the Terraform Azure provider client")
	}
	ps.Meta = c
	return nil
}
