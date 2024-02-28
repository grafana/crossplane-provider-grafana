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
	// configures all resources to be synced without async callbacks, the Grafana API is synchronous
	for _, resource := range p.Resources {
		resource.UseAsync = false
	}

	p.AddResourceConfigurator("grafana_annotation", func(r *ujconfig.Resource) {
		delete(r.TerraformResource.Schema, "dashboard_id") // Deprecated
		r.References["dashboard_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_dashboard",
			RefFieldName:      "DashboardRef",
			SelectorFieldName: "DashboardSelector",
			Extractor:         SelfPackagePath + ".UIDExtractor()",
		}
	})
	p.AddResourceConfigurator("grafana_notification_policy", func(r *ujconfig.Resource) {
		contactPointRef := ujconfig.Reference{
			TerraformName:     "grafana_contact_point",
			RefFieldName:      "ContactPointRef",
			SelectorFieldName: "ContactPointSelector",
			Extractor:         SelfPackagePath + ".NameExtractor()",
		}
		r.References["contact_point"] = contactPointRef
		r.References["policy.contact_point"] = contactPointRef
		r.References["policy.policy.contact_point"] = contactPointRef
		r.References["policy.policy.policy.contact_point"] = contactPointRef
		r.References["policy.policy.policy.policy.contact_point"] = contactPointRef

		muteTimingRef := ujconfig.Reference{
			TerraformName:     "grafana_mute_timing",
			RefFieldName:      "MuteTimingRef",
			SelectorFieldName: "MuteTimingSelector",
			Extractor:         SelfPackagePath + ".NameExtractor()",
		}
		r.References["policy.mute_timings"] = muteTimingRef
		r.References["policy.policy.mute_timings"] = muteTimingRef
		r.References["policy.policy.policy.mute_timings"] = muteTimingRef
		r.References["policy.policy.policy.policy.mute_timings"] = muteTimingRef

	})
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
	p.AddResourceConfigurator("grafana_cloud_access_policy_token", func(r *ujconfig.Resource) {
		r.References["access_policy_id"] = ujconfig.Reference{
			TerraformName:     "grafana_cloud_access_policy",
			RefFieldName:      "AccessPolicyRef",
			SelectorFieldName: "AccessPolicySelector",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}
			cloudConfig := map[string]string{}
			if a, ok := attr["token"].(string); ok {
				cloudConfig["cloud_access_policy_token"] = a
			}
			marshalled, err := json.Marshal(cloudConfig)
			if err != nil {
				return nil, err
			}
			conn["cloudCredentials"] = marshalled
			return conn, nil
		}
	})
	p.AddResourceConfigurator("grafana_cloud_plugin_installation", func(r *ujconfig.Resource) {
		r.References["stack_slug"] = ujconfig.Reference{
			TerraformName:     "grafana_cloud_stack",
			RefFieldName:      "CloudStackRef",
			SelectorFieldName: "CloudStackSelector",
			Extractor:         SelfPackagePath + ".CloudStackSlugExtractor()",
		}
	})
	p.AddResourceConfigurator("grafana_cloud_stack", func(r *ujconfig.Resource) {
		r.UseAsync = true
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
		r.References["permissions.team_id"] = ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRef",
			SelectorFieldName: "TeamSelector",
		}
		r.References["permissions.user_id"] = ujconfig.Reference{
			TerraformName:     "grafana_user",
			RefFieldName:      "UserRef",
			SelectorFieldName: "UserSelector",
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
		r.InitializerFns = append(r.InitializerFns, createDashboardConfigInitializer)
	})
	p.AddResourceConfigurator("grafana_dashboard_public", func(r *ujconfig.Resource) {
		r.References["dashboard_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_dashboard",
			RefFieldName:      "DashboardRef",
			SelectorFieldName: "DashboardSelector",
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
		r.References["permissions.team_id"] = ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRef",
			SelectorFieldName: "TeamSelector",
		}
		r.References["permissions.user_id"] = ujconfig.Reference{
			TerraformName:     "grafana_user",
			RefFieldName:      "UserRef",
			SelectorFieldName: "UserSelector",
		}
	})
	p.AddResourceConfigurator("grafana_data_source_permission", func(r *ujconfig.Resource) {
		r.References["datasource_id"] = ujconfig.Reference{
			TerraformName:     "grafana_data_source",
			RefFieldName:      "DataSourceRef",
			SelectorFieldName: "DataSourceSelector",
		}
		r.References["permissions.team_id"] = ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRef",
			SelectorFieldName: "TeamSelector",
		}
		r.References["permissions.user_id"] = ujconfig.Reference{
			TerraformName:     "grafana_user",
			RefFieldName:      "UserRef",
			SelectorFieldName: "UserSelector",
		}
	})
	p.AddResourceConfigurator("grafana_folder", func(r *ujconfig.Resource) {
		r.References["parent_folder_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_folder",
			RefFieldName:      "FolderRef",
			SelectorFieldName: "FolderSelector",
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
		r.References["permissions.team_id"] = ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRef",
			SelectorFieldName: "TeamSelector",
		}
		r.References["permissions.user_id"] = ujconfig.Reference{
			TerraformName:     "grafana_user",
			RefFieldName:      "UserRef",
			SelectorFieldName: "UserSelector",
		}
	})
	p.AddResourceConfigurator("grafana_library_panel", func(r *ujconfig.Resource) {
		r.References["folder_id"] = ujconfig.Reference{
			TerraformName:     "grafana_folder",
			RefFieldName:      "FolderRef",
			SelectorFieldName: "FolderSelector",
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
	p.AddResourceConfigurator("grafana_role_assignment", func(r *ujconfig.Resource) {
		r.References["role_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_role",
			RefFieldName:      "RoleRef",
			SelectorFieldName: "RoleSelector",
			Extractor:         SelfPackagePath + ".UIDExtractor()",
		}
		r.References["service_accounts"] = ujconfig.Reference{
			TerraformName:     "grafana_service_account",
			RefFieldName:      "ServiceAccountRefs",
			SelectorFieldName: "ServiceAccountSelector",
		}
		r.References["teams"] = ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRefs",
			SelectorFieldName: "TeamSelector",
		}
		r.References["users"] = ujconfig.Reference{
			TerraformName:     "grafana_user",
			RefFieldName:      "UserRefs",
			SelectorFieldName: "UserSelector",
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
	p.AddResourceConfigurator("grafana_team_external_group", func(r *ujconfig.Resource) {
		r.References["team_id"] = ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRef",
			SelectorFieldName: "TeamSelector",
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
