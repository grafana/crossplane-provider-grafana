/*
Copyright 2021 Upbound Inc.
*/

package grafana

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
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
		// NOTE: the following references won't work as Terraform datasources are not translated to Crossplane resources
		// r.References["group_to_notify"] = oncallUserGroupRef
		// r.References["notify_to_team_members"] = oncallTeamRef
		// r.References["persons_to_notify"] = oncallUserRef
		// r.References["persons_to_notify_next_each_time"] = oncallUserRef
	})

	p.AddResourceConfigurator("grafana_oncall_integration", func(r *ujconfig.Resource) {
		r.References["default_route.escalation_chain_id"] = ujconfig.Reference{
			TerraformName:     "grafana_oncall_escalation_chain",
			RefFieldName:      "EscalationChainRef",
			SelectorFieldName: "EscalationChainSelector",
		}
		// NOTE: the following references won't work as Terraform datasources are not translated to Crossplane resources
		// r.References["team_id"] = oncallTeamRef
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
	})

	p.AddResourceConfigurator("grafana_oncall_schedule", func(r *ujconfig.Resource) {
		r.References["shifts"] = ujconfig.Reference{
			TerraformName:     "grafana_oncall_on_call_shift",
			RefFieldName:      "ShiftsRef",
			SelectorFieldName: "ShiftsSelector",
		}
		// NOTE: the following references won't work as Terraform datasources are not translated to Crossplane resources
		// r.References["slack.channel_id"] = slackChannelRef
		// r.References["slack.user_group_id"] = oncallUserGroupRef
	})

	// NOTE: the following refs will not work as Terraform datasources are not translated to Crossplane resources
	// the workaround is to use the Terraform provider for Crossplane to use the datasources directly
	// https://github.com/crossplane/crossplane/blob/master/design/design-doc-observe-only-resources.md
	// oncallTeamRef := ujconfig.Reference{
	// 	TerraformName:     "grafana_oncall_team",
	// 	RefFieldName:      "TeamRef",
	// 	SelectorFieldName: "TeamSelector",
	// }
	// oncallUserGroupRef := ujconfig.Reference{
	// 	TerraformName:     "grafana_oncall_user_group",
	// 	RefFieldName:      "OnCallUserGroupRef",
	// 	SelectorFieldName: "OnCallUserGroupSelector",
	// }
	// oncallUserRef := ujconfig.Reference{
	// 	TerraformName:     "grafana_oncall_user",
	// 	RefFieldName:      "OnCallUserRef",
	// 	SelectorFieldName: "OnCallUserSelector",
	// }
	// slackChannelRef := ujconfig.Reference{
	// 	TerraformName:     "grafana_oncall_slack_channel",
	// 	RefFieldName:      "SlackChannelRef",
	// 	SelectorFieldName: "SlackChannelSelector",
	// }

	// p.AddResourceConfigurator("grafana_oncall_escalation_chain", func(r *ujconfig.Resource) {
	// 	r.References["team_id"] = oncallTeamRef
	// })

	// p.AddResourceConfigurator("grafana_on_call_shift", func(r *ujconfig.Resource) {
	// 	r.References["users"] = oncallUserRef
	// 	r.References["rolling_users"] = oncallUserRef
	// 	r.References["team_id"] = oncallTeamRef
	// })

	// p.AddResourceConfigurator("grafana_oncall_outgoing_webhook", func(r *ujconfig.Resource) {
	// 	r.References["team_id"] = oncallTeamRef
	// })
}
