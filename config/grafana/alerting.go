/*
Copyright 2026 Grafana Labs
*/

package grafana

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

func configureAlerting(p *ujconfig.Provider) {
	p.AddResourceConfigurator("grafana_notification_policy", func(r *ujconfig.Resource) {
		contactPointRef := ujconfig.Reference{
			TerraformName:     "grafana_contact_point",
			RefFieldName:      "ContactPointRef",
			SelectorFieldName: "ContactPointSelector",
			Extractor:         fieldExtractor("name"),
		}
		r.References["contact_point"] = contactPointRef
		r.References["policy.contact_point"] = contactPointRef
		r.References["policy.policy.contact_point"] = contactPointRef
		r.References["policy.policy.policy.contact_point"] = contactPointRef
		r.References["policy.policy.policy.policy.contact_point"] = contactPointRef

		muteTimingRef := ujconfig.Reference{
			TerraformName:     "grafana_mute_timing",
			RefFieldName:      "MuteTimingRef",
			SelectorFieldName: "MuteTimingSelector",
			Extractor:         fieldExtractor("name"),
		}
		r.References["policy.mute_timings"] = muteTimingRef
		r.References["policy.policy.mute_timings"] = muteTimingRef
		r.References["policy.policy.policy.mute_timings"] = muteTimingRef
		r.References["policy.policy.policy.policy.mute_timings"] = muteTimingRef
	})
	p.AddResourceConfigurator("grafana_contact_point", func(r *ujconfig.Resource) {
		r.References["oncall.url"] = ujconfig.Reference{
			TerraformName:     "grafana_oncall_integration",
			RefFieldName:      "OncallIntegrationRef",
			SelectorFieldName: "OncallIntegrationSelector",
			Extractor:         computedFieldExtractor("link"),
		}
	})
	p.AddResourceConfigurator("grafana_rule_group", func(r *ujconfig.Resource) {
		r.References["folder_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_folder",
			RefFieldName:      "FolderRef",
			SelectorFieldName: "FolderSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
		r.References["rule.notification_settings.contact_point"] = ujconfig.Reference{
			TerraformName:     "grafana_contact_point",
			RefFieldName:      "ContactPointRef",
			SelectorFieldName: "ContactPointSelector",
			Extractor:         fieldExtractor("name"),
		}
	})
}
