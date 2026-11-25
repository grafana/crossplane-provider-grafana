/*
Copyright 2026 Grafana Labs
*/

package grafana

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

func configureApps(p *ujconfig.Provider) {
	resourcesWithFolder := []string{
		"grafana_apps_alertenrichment_alertenrichment_v1beta1",
		"grafana_apps_dashboard_dashboard_v1beta1",
		"grafana_apps_dashboard_dashboard_v2",
		"grafana_apps_dashboard_dashboard_v2beta1",
		"grafana_apps_notifications_inhibitionrule_v1beta1",
		"grafana_apps_playlist_playlist_v0alpha1",
		"grafana_apps_playlist_playlist_v1",
		"grafana_apps_provisioning_connection_v0alpha1",
		"grafana_apps_provisioning_repository_v0alpha1",
		"grafana_apps_rules_alertrule_v0alpha1",
		"grafana_apps_rules_recordingrule_v0alpha1",
		"grafana_apps_secret_keeper_activation_v1beta1",
		"grafana_apps_secret_keeper_v1beta1",
		"grafana_apps_secret_securevalue_v1beta1",
	}
	for _, name := range resourcesWithFolder {
		p.AddResourceConfigurator(name, func(r *ujconfig.Resource) {
			r.References["metadata.folder_uid"] = folderReference()
		})
	}

	p.AddResourceConfigurator("grafana_apps_alertenrichment_alertenrichment_v1beta1", func(r *ujconfig.Resource) {
		ref := dataSourceReference("DataSource")
		r.References["spec.step.data_source.logs_query.data_source_uid"] = ref
		r.References["spec.step.conditional.then.step.data_source.logs_query.data_source_uid"] = ref
		r.References["spec.step.conditional.else.step.data_source.logs_query.data_source_uid"] = ref
	})

	p.AddResourceConfigurator("grafana_apps_provisioning_repository_v0alpha1", func(r *ujconfig.Resource) {
		r.References["spec.connection.name"] = ujconfig.Reference{
			TerraformName:     "grafana_apps_provisioning_connection_v0alpha1",
			RefFieldName:      "ConnectionRef",
			SelectorFieldName: "ConnectionSelector",
			Extractor:         fieldExtractor("metadata.uid"),
		}
	})

	p.AddResourceConfigurator("grafana_apps_rules_recordingrule_v0alpha1", func(r *ujconfig.Resource) {
		r.References["spec.target_datasource_uid"] = dataSourceReference("TargetDataSource")
	})
}
