/*
Copyright 2021 Upbound Inc.
*/

package grafana

import (
	"github.com/upbound/upjet/pkg/config"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/grafana/crossplane-provider-grafana/config/grafana"
)

// Configure configures the grafana group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("grafana_api_key", func(r *ujconfig.Resource) {
		r.References["cloud_stack_slug"] = config.Reference{
			TerraformName:     "grafana_cloud_stack",
			RefFieldName:      "CloudStackRef",
			SelectorFieldName: "CloudStackSelector",
		}
	})
	p.AddResourceConfigurator("grafana_dashboard", func(r *ujconfig.Resource) {
		r.References["folder"] = config.Reference{
			TerraformName:     "grafana_folder",
			RefFieldName:      "FolderRef",
			SelectorFieldName: "FolderSelector",
		}
	})
	p.AddResourceConfigurator("grafana_data_source", func(r *ujconfig.Resource) {
		delete(r.TerraformResource.Schema, "basic_auth_password") // Deprecated
		delete(r.TerraformResource.Schema, "password")            // Deprecated
		delete(r.TerraformResource.Schema, "json_data")           // Deprecated
		delete(r.TerraformResource.Schema, "secure_json_data")    // Deprecated
		delete(r.TerraformResource.Schema, "http_headers")        // TODO: Make this work!
	})
	p.AddResourceConfigurator("grafana_team", func(r *ujconfig.Resource) {
		r.References["members"] = config.Reference{
			TerraformName:     "grafana_user",
			RefFieldName:      "MemberRefs",
			SelectorFieldName: "MemberSelector",
			Extractor:         SelfPackagePath + ".UserEmailExtractor()",
		}
	})
}
