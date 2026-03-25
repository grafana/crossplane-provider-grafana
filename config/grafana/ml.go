/*
Copyright 2026 Grafana Labs
*/

package grafana

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

func configureML(p *ujconfig.Provider) {
	p.AddResourceConfigurator("grafana_machine_learning_job", func(r *ujconfig.Resource) {
		r.References["datasource_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_data_source",
			RefFieldName:      "DataSourceRef",
			SelectorFieldName: "DataSourceSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
	})

	p.AddResourceConfigurator("grafana_machine_learning_outlier_detector", func(r *ujconfig.Resource) {
		r.References["datasource_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_data_source",
			RefFieldName:      "DataSourceRef",
			SelectorFieldName: "DataSourceSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
	})
}
