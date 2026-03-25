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
	p.AddResourceConfigurator("grafana_dashboard_permission", func(r *ujconfig.Resource) {
		r.References["dashboard_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_dashboard",
			RefFieldName:      "DashboardRef",
			SelectorFieldName: "DashboardSelector",
			Extractor:         optionalFieldExtractor("uid"),
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
		r.References["members"] = ujconfig.Reference{
			TerraformName:     "grafana_user",
			RefFieldName:      "MemberRefs",
			SelectorFieldName: "MemberSelector",
			Extractor:         fieldExtractor("email"),
		}
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
