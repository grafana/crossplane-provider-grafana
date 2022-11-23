/*
Copyright 2021 Upbound Inc.
*/

package grafana

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the grafana group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("grafana_data_source", func(r *ujconfig.Resource) {
		delete(r.TerraformResource.Schema, "basic_auth_password") // Deprecated
		delete(r.TerraformResource.Schema, "password")            // Deprecated
		delete(r.TerraformResource.Schema, "json_data")           // Deprecated
		delete(r.TerraformResource.Schema, "secure_json_data")    // Deprecated
		delete(r.TerraformResource.Schema, "http_headers")        // TODO: Make this work!
	})
}
