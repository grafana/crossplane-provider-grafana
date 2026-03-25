/*
Copyright 2021 Upbound Inc.
*/

package grafana

import (
	"encoding/json"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

func configureSyntheticMonitoring(p *ujconfig.Provider) {
	p.AddResourceConfigurator("grafana_synthetic_monitoring_installation", func(r *ujconfig.Resource) {
		r.References["stack_id"] = ujconfig.Reference{
			TerraformName:     "grafana_cloud_stack",
			RefFieldName:      "CloudStackRef",
			SelectorFieldName: "CloudStackSelector",
			Extractor:         computedFieldExtractor("id"),
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}

			providerConfig := map[string]string{}
			stackSmApiUrl, hasStackSmApiUrl := attr["stack_sm_api_url"].(string)
			smAccessToken, hasSmAccessToken := attr["sm_access_token"].(string)

			if hasStackSmApiUrl && hasSmAccessToken {
				providerConfig["sm_url"] = stackSmApiUrl
				providerConfig["sm_access_token"] = smAccessToken
				marshalled, err := json.Marshal(providerConfig)
				if err != nil {
					return nil, err
				}
				conn["smCredentials"] = marshalled
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("grafana_synthetic_monitoring_check_alerts",
		func(r *ujconfig.Resource) {
			r.References["check_id"] = ujconfig.Reference{
				TerraformName:     "grafana_synthetic_monitoring_check",
				RefFieldName:      "CheckRef",
				SelectorFieldName: "CheckSelector",
				Extractor:         optionalFieldExtractor("id"),
			}
		})
}
