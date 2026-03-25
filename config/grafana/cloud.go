/*
Copyright 2026 Grafana Labs
*/

package grafana

import (
	"encoding/json"
	"errors"
	"fmt"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func configureCloud(p *ujconfig.Provider) {
	p.AddResourceConfigurator("grafana_cloud_access_policy", func(r *ujconfig.Resource) {
		r.References["realm.identifier"] = ujconfig.Reference{
			TerraformName:     "grafana_cloud_stack",
			RefFieldName:      "StackRef",
			SelectorFieldName: "StackSelector",
			Extractor:         computedFieldExtractor("id"),
		}
	})
	p.AddResourceConfigurator("grafana_cloud_access_policy_token", func(r *ujconfig.Resource) {
		r.References["access_policy_id"] = ujconfig.Reference{
			TerraformName:     "grafana_cloud_access_policy",
			RefFieldName:      "AccessPolicyRef",
			SelectorFieldName: "AccessPolicySelector",
			Extractor:         computedFieldExtractor("policyId"),
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}
			cloudConfig := map[string]string{}
			basicAuthConfig := map[string]string{}
			if a, ok := attr["token"].(string); ok {
				cloudConfig["cloud_access_policy_token"] = a
				basicAuthConfig["basicAuthPassword"] = a
				marshalledBasicAuthConfig, err := json.Marshal(basicAuthConfig)
				if err != nil {
					return nil, err
				}
				conn["basicAuthCredentials"] = marshalledBasicAuthConfig
				marshalledCloudConfig, err := json.Marshal(cloudConfig)
				if err != nil {
					return nil, err
				}
				conn["cloudCredentials"] = marshalledCloudConfig
			}
			return conn, nil
		}
	})
	p.AddResourceConfigurator("grafana_cloud_plugin_installation", func(r *ujconfig.Resource) {
		r.References["stack_slug"] = ujconfig.Reference{
			TerraformName:     "grafana_cloud_stack",
			RefFieldName:      "CloudStackRef",
			SelectorFieldName: "CloudStackSelector",
			Extractor:         fieldExtractor("slug"),
		}
	})
	p.AddResourceConfigurator("grafana_cloud_stack", func(r *ujconfig.Resource) {
		r.UseAsync = true

		// Cloud Stacks can either be imported by ID or by Slug
		// We'll default to slug instead of ID as the ID can't be known upfront
		// This'll allow us to import existing instances consistently
		// Also see: https://registry.terraform.io/providers/grafana/grafana/latest/docs/resources/cloud_stack#import
		r.ExternalName = ujconfig.ExternalName{
			SetIdentifierArgumentFn: ujconfig.NopSetIdentifierArgument,
			GetExternalNameFn: func(tfstate map[string]any) (string, error) {
				slug, ok := tfstate["slug"].(string)
				if !ok {
					return "", errors.New("cannot get slug attribute")
				}
				return slug, nil
			},
			GetIDFn:                ujconfig.ExternalNameAsID,
			DisableNameInitializer: true,
		}
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			// skip diff customization on create
			if state == nil || state.Empty() {
				return diff, nil
			}
			// skip no diff or destroy diffs
			if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
				return diff, nil
			}

			// ID is configured upon creation, don't try to update.
			// log: ResourceAttrDiff{"id":*terraform.ResourceAttrDiff{Old:"5527", New:"fedstartcrossplanetest"}}
			if diff.Attributes["id"] != nil {
				delete(diff.Attributes, "id")
			}

			return diff, nil
		}
	})

	p.AddResourceConfigurator("grafana_cloud_stack_service_account", func(r *ujconfig.Resource) {
		r.References["stack_slug"] = ujconfig.Reference{
			TerraformName:     "grafana_cloud_stack",
			RefFieldName:      "CloudStackRef",
			SelectorFieldName: "CloudStackSelector",
			Extractor:         fieldExtractor("slug"),
		}
	})
	p.AddResourceConfigurator("grafana_cloud_stack_service_account_token", func(r *ujconfig.Resource) {
		r.References["stack_slug"] = ujconfig.Reference{
			TerraformName:     "grafana_cloud_stack",
			RefFieldName:      "CloudStackRef",
			SelectorFieldName: "CloudStackSelector",
			Extractor:         fieldExtractor("slug"),
		}
		r.References["service_account_id"] = ujconfig.Reference{
			TerraformName:     "grafana_cloud_stack_service_account",
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}

			stackSlug, hasStackSlugAttribute := attr["stack_slug"].(string)
			key, hasKeyAttribute := attr["key"].(string)
			if hasStackSlugAttribute && hasKeyAttribute {
				instanceConfig := map[string]string{}
				instanceConfig["url"] = fmt.Sprintf("https://%s.grafana.net", stackSlug)
				instanceConfig["auth"] = key
				marshalled, err := json.Marshal(instanceConfig)
				if err != nil {
					return nil, err
				}
				conn["instanceCredentials"] = marshalled
			}

			return conn, nil
		}
	})
}
