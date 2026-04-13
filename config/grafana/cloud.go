/*
Copyright 2026 Grafana Labs
*/

package grafana

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

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
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}
			for _, key := range []string{
				"alertmanager_ip_allow_list_cname",
				"alertmanager_name",
				"alertmanager_status",
				"alertmanager_url",
				"cluster_name",
				"cluster_slug",
				"description",
				"fleet_management_name",
				"fleet_management_private_connectivity_info_private_dns",
				"fleet_management_private_connectivity_info_service_name",
				"fleet_management_status",
				"fleet_management_url",
				"grafanas_ip_allow_list_cname",
				"graphite_ip_allow_list_cname",
				"graphite_name",
				"graphite_private_connectivity_info_private_dns",
				"graphite_private_connectivity_info_service_name",
				"graphite_status",
				"graphite_url",
				"influx_url",
				"logs_ip_allow_list_cname",
				"logs_name",
				"logs_private_connectivity_info_private_dns",
				"logs_private_connectivity_info_service_name",
				"logs_status",
				"logs_url",
				"name",
				"oncall_api_url",
				"org_name",
				"org_slug",
				"otlp_private_connectivity_info_private_dns",
				"otlp_private_connectivity_info_service_name",
				"otlp_url",
				"pdc_api_private_connectivity_info_private_dns",
				"pdc_api_private_connectivity_info_service_name",
				"pdc_gateway_private_connectivity_info_private_dns",
				"pdc_gateway_private_connectivity_info_service_name",
				"profiles_ip_allow_list_cname",
				"profiles_name",
				"profiles_private_connectivity_info_private_dns",
				"profiles_private_connectivity_info_service_name",
				"profiles_status",
				"profiles_url",
				"prometheus_ip_allow_list_cname",
				"prometheus_name",
				"prometheus_private_connectivity_info_private_dns",
				"prometheus_private_connectivity_info_service_name",
				"prometheus_remote_endpoint",
				"prometheus_remote_write_endpoint",
				"prometheus_status",
				"prometheus_url",
				"region_slug",
				"slug",
				"status",
				"traces_ip_allow_list_cname",
				"traces_name",
				"traces_private_connectivity_info_private_dns",
				"traces_private_connectivity_info_service_name",
				"traces_status",
				"traces_url",
				"url",
				"wait_for_readiness_timeout",
			} {
				if v, ok := attr[key].(string); ok && v != "" {
					conn[key] = []byte(v)
				}
			}
			for _, key := range []string{
				"alertmanager_user_id",
				"fleet_management_user_id",
				"graphite_user_id",
				"logs_user_id",
				"org_id",
				"profiles_user_id",
				"prometheus_user_id",
				"traces_user_id",
			} {
				switch v := attr[key].(type) {
				case float64:
					conn[key] = []byte(strconv.FormatFloat(v, 'f', -1, 64))
				case int:
					conn[key] = []byte(strconv.Itoa(v))
				}
			}
			// Export the stack ID (TF uses "id" but ProviderConfig expects "stack_id")
			if v, ok := attr["id"].(string); ok && v != "" {
				conn["id"] = []byte(v)
			}
			return conn, nil
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
