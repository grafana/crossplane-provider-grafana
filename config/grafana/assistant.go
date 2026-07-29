/*
Copyright 2026 Grafana Labs
*/

package grafana

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

func configureAssistant(p *ujconfig.Provider) {
	p.AddResourceConfigurator("grafana_assistant_skill", func(r *ujconfig.Resource) {
		r.References["allowed_tools.integration_id"] = ujconfig.Reference{
			TerraformName:     "grafana_assistant_mcp_server",
			RefFieldName:      "IntegrationRef",
			SelectorFieldName: "IntegrationSelector",
			Extractor:         computedFieldExtractor("id"),
		}
	})
}
