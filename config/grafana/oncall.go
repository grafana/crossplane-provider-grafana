/*
Copyright 2026 Grafana Labs
*/

package grafana

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

const externalNameExtractor = `github.com/crossplane/crossplane-runtime/v2/pkg/reference.ExternalName()`

// observedRef returns a ujconfig.Reference for an observed (datasource) oncall resource.
// It uses Type directly (not TerraformName) because observed resources are not in the
// Terraform resource map. The ExternalName extractor reads crossplane.io/external-name
// which the observed controller populates with the resource ID.
func observedRef(typePath, prefix string) ujconfig.Reference {
	return ujconfig.Reference{
		Type:              typePath,
		RefFieldName:      prefix + "Ref",
		SelectorFieldName: prefix + "Selector",
		Extractor:         externalNameExtractor,
	}
}

const (
	oncallTeamType         = "github.com/grafana/crossplane-provider-grafana/v2/apis/observed/oncall/v1alpha1.Team"
	oncallUserType         = "github.com/grafana/crossplane-provider-grafana/v2/apis/observed/oncall/v1alpha1.User"
	oncallSlackChannelType = "github.com/grafana/crossplane-provider-grafana/v2/apis/observed/oncall/v1alpha1.SlackChannel"
	oncallUserGroupType    = "github.com/grafana/crossplane-provider-grafana/v2/apis/observed/oncall/v1alpha1.UserGroup"
)

func configureOnCall(p *ujconfig.Provider) {
	p.AddResourceConfigurator("grafana_oncall_escalation", func(r *ujconfig.Resource) {
		r.References["escalation_chain_id"] = ujconfig.Reference{
			TerraformName:     "grafana_oncall_escalation_chain",
			RefFieldName:      "EscalationChainRef",
			SelectorFieldName: "EscalationChainSelector",
		}
		r.References["action_to_trigger"] = ujconfig.Reference{
			TerraformName:     "grafana_oncall_outgoing_webhook",
			RefFieldName:      "ActionToTriggerRef",
			SelectorFieldName: "ActionToTriggerSelector",
		}
		r.References["notify_on_call_from_schedule"] = ujconfig.Reference{
			TerraformName:     "grafana_oncall_schedule",
			RefFieldName:      "NotifyOnCallFromScheduleRef",
			SelectorFieldName: "NotifyOnCallFromScheduleSelector",
		}
		r.References["notify_to_team_members"] = observedRef(oncallTeamType, "Team")
		r.References["persons_to_notify"] = observedRef(oncallUserType, "PersonsToNotify")
		r.References["persons_to_notify_next_each_time"] = observedRef(oncallUserType, "PersonsToNotifyNextEachTime")
		r.References["group_to_notify"] = observedRef(oncallUserGroupType, "GroupToNotify")
	})

	p.AddResourceConfigurator("grafana_oncall_escalation_chain", func(r *ujconfig.Resource) {
		r.References["team_id"] = observedRef(oncallTeamType, "Team")
	})

	p.AddResourceConfigurator("grafana_oncall_integration", func(r *ujconfig.Resource) {
		r.References["default_route.escalation_chain_id"] = ujconfig.Reference{
			TerraformName:     "grafana_oncall_escalation_chain",
			RefFieldName:      "EscalationChainRef",
			SelectorFieldName: "EscalationChainSelector",
		}
		r.References["team_id"] = observedRef(oncallTeamType, "Team")
		r.References["default_route.slack.channel_id"] = observedRef(oncallSlackChannelType, "SlackChannel")
		// NOTE: labels and dynamic_labels are []map[string]*string because the upstream
		// Terraform schema uses TypeList+TypeMap instead of TypeList+schema.Resource.
		// The grafana_oncall_label data source is a no-op identity function (no API call),
		// so there is nothing to reference. Users should write labels inline:
		//   labels: [{key: "LabelKey", value: "LabelValue"}]
		// TODO: override ArgumentDocs and Description for labels/dynamic_labels to
		// replace the "using the grafana_oncall_label datasource" text. Also consider
		// doing the same for team_id and other fields that reference observed resources
		// (using the grafana_oncall_team/user/slack_channel/user_group datasource).
	})

	p.AddResourceConfigurator("grafana_oncall_on_call_shift", func(r *ujconfig.Resource) {
		r.References["team_id"] = observedRef(oncallTeamType, "Team")
		r.References["users"] = observedRef(oncallUserType, "Users")
		// NOTE: rolling_users is [][]*string (nested list) which native references cannot handle.
		// Use a KCL composition to resolve user names for rolling_users.
	})

	p.AddResourceConfigurator("grafana_oncall_outgoing_webhook", func(r *ujconfig.Resource) {
		r.References["team_id"] = observedRef(oncallTeamType, "Team")
		r.References["integration_filter"] = ujconfig.Reference{
			TerraformName:     "grafana_oncall_integration",
			RefFieldName:      "IntegrationFilterRef",
			SelectorFieldName: "IntegrationFilterSelector",
		}
	})

	p.AddResourceConfigurator("grafana_oncall_route", func(r *ujconfig.Resource) {
		r.References["escalation_chain_id"] = ujconfig.Reference{
			TerraformName:     "grafana_oncall_escalation_chain",
			RefFieldName:      "EscalationChainRef",
			SelectorFieldName: "EscalationChainSelector",
		}
		r.References["integration_id"] = ujconfig.Reference{
			TerraformName:     "grafana_oncall_integration",
			RefFieldName:      "IntegrationRef",
			SelectorFieldName: "IntegrationSelector",
		}
		r.References["slack.channel_id"] = observedRef(oncallSlackChannelType, "SlackChannel")
	})

	p.AddResourceConfigurator("grafana_oncall_schedule", func(r *ujconfig.Resource) {
		r.References["shifts"] = ujconfig.Reference{
			TerraformName:     "grafana_oncall_on_call_shift",
			RefFieldName:      "ShiftsRef",
			SelectorFieldName: "ShiftsSelector",
		}
		r.References["team_id"] = observedRef(oncallTeamType, "Team")
		r.References["slack.channel_id"] = observedRef(oncallSlackChannelType, "SlackChannel")
		r.References["slack.user_group_id"] = observedRef(oncallUserGroupType, "SlackUserGroup")
	})

	p.AddResourceConfigurator("grafana_oncall_user_notification_rule", func(r *ujconfig.Resource) {
		r.References["user_id"] = observedRef(oncallUserType, "OnCallUser")
	})
}
