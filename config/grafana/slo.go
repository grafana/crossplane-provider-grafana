/*
Copyright 2026 Grafana Labs
*/

package grafana

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

func configureSLO(p *ujconfig.Provider) {
	p.AddResourceConfigurator("grafana_slo", func(r *ujconfig.Resource) {
		r.References["destination_datasource.uid"] = ujconfig.Reference{
			TerraformName:     "grafana_data_source",
			RefFieldName:      "Ref",
			SelectorFieldName: "Selector",
			Extractor:         optionalFieldExtractor("uid"),
		}
	})
}
