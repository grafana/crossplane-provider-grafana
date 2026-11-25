/*
Copyright 2026 Grafana Labs
*/

package grafana

import (
	"encoding/json"
	"fmt"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

func configureOSS(p *ujconfig.Provider) {
	p.AddResourceConfigurator("grafana_annotation", func(r *ujconfig.Resource) {
		r.References["dashboard_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_dashboard",
			RefFieldName:      "DashboardRef",
			SelectorFieldName: "DashboardSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
	})
	p.AddResourceConfigurator("grafana_dashboard", func(r *ujconfig.Resource) {
		r.References["folder"] = ujconfig.Reference{
			TerraformName:     "grafana_folder",
			RefFieldName:      "FolderRef",
			SelectorFieldName: "FolderSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
	})
	p.AddResourceConfigurator("grafana_dashboard_public", func(r *ujconfig.Resource) {
		r.References["dashboard_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_dashboard",
			RefFieldName:      "DashboardRef",
			SelectorFieldName: "DashboardSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
	})
	p.AddResourceConfigurator("grafana_dashboard_permission_item", func(r *ujconfig.Resource) {
		r.References["dashboard_uid"] = dashboardReference("Dashboard")
		addObservedReference(r, "team", ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRef",
			SelectorFieldName: "TeamSelector",
		}, "observed_team", observedRef(ossTeamType, "ObservedTeam"))
		addUserReferences(r, "user", "User", false, "", externalNameExtractor)
	})
	p.AddResourceConfigurator("grafana_data_source", func(r *ujconfig.Resource) {
		r.References["private_data_source_connect_network_id"] = ujconfig.Reference{
			TerraformName:     "grafana_cloud_private_data_source_connect_network",
			RefFieldName:      "PrivateDataSourceConnectNetworkRef",
			SelectorFieldName: "PrivateDataSourceConnectNetworkSelector",
			Extractor:         computedFieldExtractor("pdcNetworkId"),
		}
	})
	p.AddResourceConfigurator("grafana_dashboard_permission", func(r *ujconfig.Resource) {
		r.References["dashboard_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_dashboard",
			RefFieldName:      "DashboardRef",
			SelectorFieldName: "DashboardSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
		addObservedReference(r, "permissions.team_id", ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRef",
			SelectorFieldName: "TeamSelector",
		}, "observed_team_id", observedRef(ossTeamType, "ObservedTeam"))
		addUserReferences(r, "permissions.user_id", "User", false, "", externalNameExtractor)
	})
	p.AddResourceConfigurator("grafana_folder", func(r *ujconfig.Resource) {
		r.References["parent_folder_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_folder",
			RefFieldName:      "FolderRef",
			SelectorFieldName: "FolderSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
	})
	p.AddResourceConfigurator("grafana_folder_permission", func(r *ujconfig.Resource) {
		r.References["folder_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_folder",
			RefFieldName:      "FolderRef",
			SelectorFieldName: "FolderSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
		addObservedReference(r, "permissions.team_id", ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRef",
			SelectorFieldName: "TeamSelector",
		}, "observed_team_id", observedRef(ossTeamType, "ObservedTeam"))
		addUserReferences(r, "permissions.user_id", "User", false, "", externalNameExtractor)
	})
	p.AddResourceConfigurator("grafana_folder_permission_item", func(r *ujconfig.Resource) {
		r.References["folder_uid"] = folderReference()
		addObservedReference(r, "team", ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRef",
			SelectorFieldName: "TeamSelector",
		}, "observed_team", observedRef(ossTeamType, "ObservedTeam"))
		addUserReferences(r, "user", "User", false, "", externalNameExtractor)
	})
	p.AddResourceConfigurator("grafana_library_panel", func(r *ujconfig.Resource) {
		r.References["folder_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_folder",
			RefFieldName:      "FolderRef",
			SelectorFieldName: "FolderSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
	})
	p.AddResourceConfigurator("grafana_service_account_permission", func(r *ujconfig.Resource) {
		r.References["service_account_id"] = ujconfig.Reference{
			TerraformName:     "grafana_service_account",
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
		addObservedReference(r, "permissions.team_id", ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRef",
			SelectorFieldName: "TeamSelector",
		}, "observed_team_id", observedRef(ossTeamType, "ObservedTeam"))
		addUserReferences(r, "permissions.user_id", "User", false, "", externalNameExtractor)
	})
	p.AddResourceConfigurator("grafana_service_account_permission_item", func(r *ujconfig.Resource) {
		r.References["service_account_id"] = ujconfig.Reference{
			TerraformName:     "grafana_service_account",
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
		addObservedReference(r, "team", ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRef",
			SelectorFieldName: "TeamSelector",
		}, "observed_team", observedRef(ossTeamType, "ObservedTeam"))
		addUserReferences(r, "user", "User", false, "", externalNameExtractor)
	})
	p.AddResourceConfigurator("grafana_service_account_rotating_token", func(r *ujconfig.Resource) {
		r.References["service_account_id"] = ujconfig.Reference{
			TerraformName:     "grafana_service_account",
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
	})
	p.AddResourceConfigurator("grafana_organization_preferences", func(r *ujconfig.Resource) {
		r.References["home_dashboard_uid"] = dashboardReference("HomeDashboard")
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
				marshalled, err := json.Marshal(instanceConfig)
				if err != nil {
					return nil, err
				}
				conn["instanceCredentials"] = marshalled
			}

			return conn, nil
		}
	})
	p.AddResourceConfigurator("grafana_team", func(r *ujconfig.Resource) {
		addUserReferences(r, "members", "Member", true, fieldExtractor("email"), optionalFieldExtractor("email"))
		r.References["preferences.home_dashboard_uid"] = dashboardReference("HomeDashboard")
	})
	p.AddResourceConfigurator("grafana_organization", func(r *ujconfig.Resource) {
		addUserReferences(r, "admins", "Admin", true, fieldExtractor("email"), optionalFieldExtractor("email"))
		addUserReferences(r, "editors", "Editor", true, fieldExtractor("email"), optionalFieldExtractor("email"))
		addUserReferences(r, "viewers", "Viewer", true, fieldExtractor("email"), optionalFieldExtractor("email"))
		addUserReferences(r, "users_without_access", "UserWithoutAccess", true, fieldExtractor("email"), optionalFieldExtractor("email"))
		addUserReferences(r, "admin_user", "AdminUser", false, fieldExtractor("login"), optionalFieldExtractor("login"))
	})
	p.AddResourceConfigurator(
		"grafana_apps_alertenrichment_alertenrichment_v1beta1",
		func(r *ujconfig.Resource) {
			if err := ujconfig.TraverseSchemas(r.Name, r, &ujconfig.SingletonListEmbedder{}); err != nil {
				panic(fmt.Errorf("failed to configure singleton blocks for %s: %w", r.Name, err))
			}
		},
	)
}
