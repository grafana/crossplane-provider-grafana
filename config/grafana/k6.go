/*
Copyright 2026 Grafana Labs
*/

package grafana

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

func configureK6(p *ujconfig.Provider) {
	p.AddResourceConfigurator("grafana_k6_installation", func(r *ujconfig.Resource) {
		r.References["stack_id"] = cloudStackIDReference()
	})
	p.AddResourceConfigurator("grafana_k6_project_allowed_load_zones", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{
			TerraformName:     "grafana_k6_project",
			RefFieldName:      "ProjectRef",
			SelectorFieldName: "ProjectSelector",
		}
	})
	p.AddResourceConfigurator("grafana_k6_project_limits", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{
			TerraformName:     "grafana_k6_project",
			RefFieldName:      "ProjectRef",
			SelectorFieldName: "ProjectSelector",
		}
	})
	p.AddResourceConfigurator("grafana_k6_load_test", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{
			TerraformName:     "grafana_k6_project",
			RefFieldName:      "ProjectRef",
			SelectorFieldName: "ProjectSelector",
		}
	})
	p.AddResourceConfigurator("grafana_k6_schedule", func(r *ujconfig.Resource) {
		r.References["load_test_id"] = ujconfig.Reference{
			TerraformName:     "grafana_k6_load_test",
			RefFieldName:      "LoadTestRef",
			SelectorFieldName: "LoadTestSelector",
		}
	})
}
