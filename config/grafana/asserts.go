/*
Copyright 2026 Grafana Labs
*/

package grafana

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

func configureAsserts(p *ujconfig.Provider) {
	for _, name := range []string{
		"grafana_asserts_log_config",
		"grafana_asserts_profile_config",
		"grafana_asserts_trace_config",
	} {
		p.AddResourceConfigurator(name, func(r *ujconfig.Resource) {
			r.References["data_source_uid"] = dataSourceReference("DataSource")
		})
	}
}
