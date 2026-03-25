/*
Copyright 2021 Upbound Inc.
*/

package grafana

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

func configureK6(p *ujconfig.Provider) {
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
}
