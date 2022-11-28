/*
Copyright 2021 Upbound Inc.
*/

package grafana

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/grafana/crossplane-provider-grafana/config/grafana"
)

// Configure configures the grafana group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("grafana_contact_point", func(r *ujconfig.Resource) {
		// TODO: Make maps work!
		delete(r.TerraformResource.Schema["alertmanager"].Elem.(*schema.Resource).Schema, "settings")
		delete(r.TerraformResource.Schema["dingding"].Elem.(*schema.Resource).Schema, "settings")
		delete(r.TerraformResource.Schema["discord"].Elem.(*schema.Resource).Schema, "settings")
		delete(r.TerraformResource.Schema["email"].Elem.(*schema.Resource).Schema, "settings")
		delete(r.TerraformResource.Schema["googlechat"].Elem.(*schema.Resource).Schema, "settings")
		delete(r.TerraformResource.Schema["kafka"].Elem.(*schema.Resource).Schema, "settings")
		delete(r.TerraformResource.Schema["opsgenie"].Elem.(*schema.Resource).Schema, "settings")
		delete(r.TerraformResource.Schema["pagerduty"].Elem.(*schema.Resource).Schema, "settings")
		delete(r.TerraformResource.Schema["pushover"].Elem.(*schema.Resource).Schema, "settings")
		delete(r.TerraformResource.Schema["sensugo"].Elem.(*schema.Resource).Schema, "settings")
		delete(r.TerraformResource.Schema["slack"].Elem.(*schema.Resource).Schema, "settings")
		delete(r.TerraformResource.Schema["teams"].Elem.(*schema.Resource).Schema, "settings")
		delete(r.TerraformResource.Schema["telegram"].Elem.(*schema.Resource).Schema, "settings")
		delete(r.TerraformResource.Schema["threema"].Elem.(*schema.Resource).Schema, "settings")
		delete(r.TerraformResource.Schema["victorops"].Elem.(*schema.Resource).Schema, "settings")
		delete(r.TerraformResource.Schema["webhook"].Elem.(*schema.Resource).Schema, "settings")
		delete(r.TerraformResource.Schema["wecom"].Elem.(*schema.Resource).Schema, "settings")
	})
	p.AddResourceConfigurator("grafana_notification_policy", func(r *ujconfig.Resource) {
		r.References["contact_point"] = ujconfig.Reference{
			TerraformName:     "grafana_contact_point",
			RefFieldName:      "ContactPointRef",
			SelectorFieldName: "ContactPointSelector",
			Extractor:         SelfPackagePath + ".NameExtractor()",
		}
	})
	p.AddResourceConfigurator("grafana_rule_group", func(r *ujconfig.Resource) {
		r.References["folder_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_folder",
			RefFieldName:      "FolderRef",
			SelectorFieldName: "FolderSelector",
			Extractor:         SelfPackagePath + ".UIDExtractor()",
		}
	})
	p.AddResourceConfigurator("grafana_api_key", func(r *ujconfig.Resource) {
		r.References["cloud_stack_slug"] = ujconfig.Reference{
			TerraformName:     "grafana_cloud_stack",
			RefFieldName:      "CloudStackRef",
			SelectorFieldName: "CloudStackSelector",
			Extractor:         SelfPackagePath + ".CloudStackSlugExtractor()",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}

			instanceConfig := map[string]string{}
			if a, ok := attr["cloud_stack_slug"].(string); ok {
				instanceConfig["url"] = fmt.Sprintf("https://%s.grafana.net", a)
			}
			if a, ok := attr["key"].(string); ok {
				instanceConfig["auth"] = a
			}
			marshalled, err := json.Marshal(instanceConfig)
			if err != nil {
				return nil, err
			}
			conn["instanceCredentials"] = marshalled

			return conn, nil
		}
	})
	p.AddResourceConfigurator("grafana_dashboard", func(r *ujconfig.Resource) {
		r.References["folder"] = ujconfig.Reference{
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
		r.References["members"] = ujconfig.Reference{
			TerraformName:     "grafana_user",
			RefFieldName:      "MemberRefs",
			SelectorFieldName: "MemberSelector",
			Extractor:         SelfPackagePath + ".UserEmailExtractor()",
		}
	})
}
