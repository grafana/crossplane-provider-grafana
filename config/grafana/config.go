/*
Copyright 2021 Upbound Inc.
*/

package grafana

import (
	"encoding/json"
	"fmt"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/grafana/crossplane-provider-grafana/config/grafana"
)

// ConfigureOrgIDRefs adds an organization reference to the org_id field for all resources that have the field.
func ConfigureOrgIDRefs(p *ujconfig.Provider) {
	for name, resource := range p.Resources {
		if resource.TerraformResource.Schema["org_id"] == nil {
			continue
		}
		p.AddResourceConfigurator(name, func(r *ujconfig.Resource) {
			r.References["org_id"] = ujconfig.Reference{
				TerraformName:     "grafana_organization",
				RefFieldName:      "OrganizationRef",
				SelectorFieldName: "OrganizationSelector",
			}
		})
	}
}

// Configure configures the grafana group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("grafana_api_key", func(r *ujconfig.Resource) {
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}

			instanceConfig := map[string]string{}
			// TODO: set URL from client
			if a, ok := attr["key"].(string); ok {
				instanceConfig["auth"] = a
			}
			marshalled, err := json.Marshal(instanceConfig)
			if err != nil {
				return nil, err
			}
			conn["instanceCredentials"] = marshalled

			return conn, nil
		}
	})
	p.AddResourceConfigurator("grafana_cloud_stack_service_account", func(r *ujconfig.Resource) {
		r.References["stack_slug"] = ujconfig.Reference{
			TerraformName:     "grafana_cloud_stack",
			RefFieldName:      "CloudStackRef",
			SelectorFieldName: "CloudStackSelector",
			Extractor:         SelfPackagePath + ".CloudStackSlugExtractor()",
		}
	})
	p.AddResourceConfigurator("grafana_cloud_stack_service_account_token", func(r *ujconfig.Resource) {
		r.References["stack_slug"] = ujconfig.Reference{
			TerraformName:     "grafana_cloud_stack",
			RefFieldName:      "CloudStackRef",
			SelectorFieldName: "CloudStackSelector",
			Extractor:         SelfPackagePath + ".CloudStackSlugExtractor()",
		}
		r.References["service_account_id"] = ujconfig.Reference{
			TerraformName:     "grafana_cloud_stack_service_account",
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}

			instanceConfig := map[string]string{}
			if a, ok := attr["stack_slug"].(string); ok {
				instanceConfig["url"] = fmt.Sprintf("https://%s.grafana.net", a)
			} // TODO: set URL from client
			if a, ok := attr["key"].(string); ok {
				instanceConfig["auth"] = a
			}
			marshalled, err := json.Marshal(instanceConfig)
			if err != nil {
				return nil, err
			}
			conn["instanceCredentials"] = marshalled

			return conn, nil
		}
	})
	p.AddResourceConfigurator("grafana_service_account_permission", func(r *ujconfig.Resource) {
		r.References["service_account_id"] = ujconfig.Reference{
			TerraformName:     "grafana_service_account",
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
	})
	p.AddResourceConfigurator("grafana_service_account_token", func(r *ujconfig.Resource) {
		r.References["service_account_id"] = ujconfig.Reference{
			TerraformName:     "grafana_service_account",
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}

			instanceConfig := map[string]string{}
			// TODO: set URL from client
			// instanceConfig["url"] = fmt.Sprintf("https://%s.grafana.net", a)
			if a, ok := attr["key"].(string); ok {
				instanceConfig["auth"] = a
			}
			marshalled, err := json.Marshal(instanceConfig)
			if err != nil {
				return nil, err
			}
			conn["instanceCredentials"] = marshalled

			return conn, nil
		}
	})
	p.AddResourceConfigurator("grafana_dashboard", func(r *ujconfig.Resource) {
		r.References["folder"] = ujconfig.Reference{
			TerraformName:     "grafana_folder",
			RefFieldName:      "FolderRef",
			SelectorFieldName: "FolderSelector",
			Extractor:         SelfPackagePath + ".UIDExtractor()",
		}
	})
	p.AddResourceConfigurator("grafana_dashboard_permission", func(r *ujconfig.Resource) {
		delete(r.TerraformResource.Schema, "dashboard_id") // Deprecated
		r.References["dashboard_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_dashboard",
			RefFieldName:      "DashboardRef",
			SelectorFieldName: "DashboardSelector",
			Extractor:         SelfPackagePath + ".UIDExtractor()",
		}
	})
	p.AddResourceConfigurator("grafana_folder_permission", func(r *ujconfig.Resource) {
		r.References["folder_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_folder",
			RefFieldName:      "FolderRef",
			SelectorFieldName: "FolderSelector",
			Extractor:         SelfPackagePath + ".UIDExtractor()",
		}
	})
	p.AddResourceConfigurator("grafana_notification_policy", func(r *ujconfig.Resource) {
		r.References["contact_point"] = ujconfig.Reference{
			TerraformName:     "grafana_contact_point",
			RefFieldName:      "ContactPointRef",
			SelectorFieldName: "ContactPointSelector",
			Extractor:         SelfPackagePath + ".NameExtractor()",
		}
	})
	p.AddResourceConfigurator("grafana_report", func(r *ujconfig.Resource) {
		delete(r.TerraformResource.Schema, "dashboard_id") // Deprecated
		r.References["dashboard_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_dashboard",
			RefFieldName:      "DashboardRef",
			SelectorFieldName: "DashboardSelector",
			Extractor:         SelfPackagePath + ".UIDExtractor()",
		}
	})
	p.AddResourceConfigurator("grafana_rule_group", func(r *ujconfig.Resource) {
		r.References["folder_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_folder",
			RefFieldName:      "FolderRef",
			SelectorFieldName: "FolderSelector",
			Extractor:         SelfPackagePath + ".UIDExtractor()",
		}
	})
	p.AddResourceConfigurator("grafana_team", func(r *ujconfig.Resource) {
		r.References["members"] = ujconfig.Reference{
			TerraformName:     "grafana_user",
			RefFieldName:      "MemberRefs",
			SelectorFieldName: "MemberSelector",
			Extractor:         SelfPackagePath + ".UserEmailExtractor()",
		}
	})
	p.AddResourceConfigurator("grafana_synthetic_monitoring_installation", func(r *ujconfig.Resource) {
		r.References["stack_id"] = ujconfig.Reference{
			TerraformName:     "grafana_cloud_stack",
			RefFieldName:      "CloudStackRef",
			SelectorFieldName: "CloudStackSelector",
		}

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}

			providerConfig := map[string]string{}
			if a, ok := attr["sm_access_token"].(string); ok {
				providerConfig["sm_access_token"] = a
			}
			if a, ok := attr["stack_sm_api_url"].(string); ok {
				providerConfig["sm_url"] = a
			}
			marshalled, err := json.Marshal(providerConfig)
			if err != nil {
				return nil, err
			}
			conn["smCredentials"] = marshalled

			return conn, nil
		}
	})
}
