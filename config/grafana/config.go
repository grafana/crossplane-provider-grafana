/*
Copyright 2021 Upbound Inc.
*/

package grafana

import (
	"encoding/json"
	"errors"
	"fmt"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// ConfigureOrgIDRefs adds an organization reference to the org_id field for all resources that have the field.
func ConfigureOrgIDRefs(p *ujconfig.Provider) {
	for name, resource := range p.Resources {
		if resource.TerraformResource.Schema["org_id"] == nil {
			continue
		}
		p.AddResourceConfigurator(name, func(r *ujconfig.Resource) {
			r.References["org_id"] = ujconfig.Reference{
				TerraformName:     "grafana_organization",
				RefFieldName:      "OrganizationRef",
				SelectorFieldName: "OrganizationSelector",
			}
		})
	}
}

// ConfigureOnCallRefsAndSelectors add reference and selector fields for Grafana OnCall resources
func ConfigureOnCallRefsAndSelectors(p *ujconfig.Provider) {
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

	p.AddResourceConfigurator(
		"grafana_apps_alertenrichment_alertenrichment_v1beta1",
		func(r *ujconfig.Resource) {
			if err := ujconfig.TraverseSchemas(r.Name, r, &ujconfig.SingletonListEmbedder{}); err != nil {
				panic(fmt.Errorf("failed to configure singleton blocks for %s: %w", r.Name, err))
			}
		},
	)

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

// Configure configures the grafana group
func Configure(p *ujconfig.Provider) {
	// configures all resources to be synced without async callbacks, the Grafana API is synchronous
	for _, resource := range p.Resources {
		resource.UseAsync = false
	}

	p.AddResourceConfigurator("grafana_annotation", func(r *ujconfig.Resource) {
		r.References["dashboard_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_dashboard",
			RefFieldName:      "DashboardRef",
			SelectorFieldName: "DashboardSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
	})
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
	p.AddResourceConfigurator("grafana_cloud_access_policy", func(r *ujconfig.Resource) {
		r.References["realm.identifier"] = ujconfig.Reference{
			TerraformName:     "grafana_cloud_stack",
			RefFieldName:      "StackRef",
			SelectorFieldName: "StackSelector",
			Extractor:         computedFieldExtractor("id"),
		}
	})
	p.AddResourceConfigurator("grafana_cloud_access_policy_token", func(r *ujconfig.Resource) {
		r.References["access_policy_id"] = ujconfig.Reference{
			TerraformName:     "grafana_cloud_access_policy",
			RefFieldName:      "AccessPolicyRef",
			SelectorFieldName: "AccessPolicySelector",
			Extractor:         computedFieldExtractor("policyId"),
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}
			cloudConfig := map[string]string{}
			basicAuthConfig := map[string]string{}
			if a, ok := attr["token"].(string); ok {
				cloudConfig["cloud_access_policy_token"] = a
				basicAuthConfig["basicAuthPassword"] = a
				marshalledBasicAuthConfig, err := json.Marshal(basicAuthConfig)
				if err != nil {
					return nil, err
				}
				conn["basicAuthCredentials"] = marshalledBasicAuthConfig
				marshalledCloudConfig, err := json.Marshal(cloudConfig)
				if err != nil {
					return nil, err
				}
				conn["cloudCredentials"] = marshalledCloudConfig
			}
			return conn, nil
		}
	})
	p.AddResourceConfigurator("grafana_cloud_plugin_installation", func(r *ujconfig.Resource) {
		r.References["stack_slug"] = ujconfig.Reference{
			TerraformName:     "grafana_cloud_stack",
			RefFieldName:      "CloudStackRef",
			SelectorFieldName: "CloudStackSelector",
			Extractor:         fieldExtractor("slug"),
		}
	})
	p.AddResourceConfigurator("grafana_cloud_stack", func(r *ujconfig.Resource) {
		r.UseAsync = true

		// Cloud Stacks can either be imported by ID or by Slug
		// We'll default to slug instead of ID as the ID can't be known upfront
		// This'll allow us to import existing instances consistently
		// Also see: https://registry.terraform.io/providers/grafana/grafana/latest/docs/resources/cloud_stack#import
		r.ExternalName = ujconfig.ExternalName{
			SetIdentifierArgumentFn: ujconfig.NopSetIdentifierArgument,
			GetExternalNameFn: func(tfstate map[string]any) (string, error) {
				slug, ok := tfstate["slug"].(string)
				if !ok {
					return "", errors.New("cannot get slug attribute")
				}
				return slug, nil
			},
			GetIDFn:                ujconfig.ExternalNameAsID,
			DisableNameInitializer: true,
		}
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			// skip diff customization on create
			if state == nil || state.Empty() {
				return diff, nil
			}
			// skip no diff or destroy diffs
			if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
				return diff, nil
			}

			// ID is configured upon creation, don't try to update.
			// log: ResourceAttrDiff{"id":*terraform.ResourceAttrDiff{Old:"5527", New:"fedstartcrossplanetest"}}
			if diff.Attributes["id"] != nil {
				delete(diff.Attributes, "id")
			}

			return diff, nil
		}
	})

	p.AddResourceConfigurator("grafana_cloud_stack_service_account", func(r *ujconfig.Resource) {
		r.References["stack_slug"] = ujconfig.Reference{
			TerraformName:     "grafana_cloud_stack",
			RefFieldName:      "CloudStackRef",
			SelectorFieldName: "CloudStackSelector",
			Extractor:         fieldExtractor("slug"),
		}
	})
	p.AddResourceConfigurator("grafana_cloud_stack_service_account_token", func(r *ujconfig.Resource) {
		r.References["stack_slug"] = ujconfig.Reference{
			TerraformName:     "grafana_cloud_stack",
			RefFieldName:      "CloudStackRef",
			SelectorFieldName: "CloudStackSelector",
			Extractor:         fieldExtractor("slug"),
		}
		r.References["service_account_id"] = ujconfig.Reference{
			TerraformName:     "grafana_cloud_stack_service_account",
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}

			stackSlug, hasStackSlugAttribute := attr["stack_slug"].(string)
			key, hasKeyAttribute := attr["key"].(string)
			if hasStackSlugAttribute && hasKeyAttribute {
				instanceConfig := map[string]string{}
				instanceConfig["url"] = fmt.Sprintf("https://%s.grafana.net", stackSlug)
				instanceConfig["auth"] = key
				marshalled, err := json.Marshal(instanceConfig)
				if err != nil {
					return nil, err
				}
				conn["instanceCredentials"] = marshalled
			}

			return conn, nil
		}
	})
	p.AddResourceConfigurator("grafana_service_account_permission", func(r *ujconfig.Resource) {
		r.References["service_account_id"] = ujconfig.Reference{
			TerraformName:     "grafana_service_account",
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
		r.References["permissions.team_id"] = ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRef",
			SelectorFieldName: "TeamSelector",
		}
		r.References["permissions.user_id"] = ujconfig.Reference{
			TerraformName:     "grafana_user",
			RefFieldName:      "UserRef",
			SelectorFieldName: "UserSelector",
		}
	})
	p.AddResourceConfigurator("grafana_service_account_token", func(r *ujconfig.Resource) {
		r.References["service_account_id"] = ujconfig.Reference{
			TerraformName:     "grafana_service_account",
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}

			instanceConfig := map[string]string{}
			// TODO: set URL from client
			// instanceConfig["url"] = fmt.Sprintf("https://%s.grafana.net", a)
			if a, ok := attr["key"].(string); ok {
				instanceConfig["auth"] = a
				marshalled, err := json.Marshal(instanceConfig)
				if err != nil {
					return nil, err
				}
				conn["instanceCredentials"] = marshalled
			}

			return conn, nil
		}
	})
	p.AddResourceConfigurator("grafana_dashboard", func(r *ujconfig.Resource) {
		r.References["folder"] = ujconfig.Reference{
			TerraformName:     "grafana_folder",
			RefFieldName:      "FolderRef",
			SelectorFieldName: "FolderSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
	})
	p.AddResourceConfigurator("grafana_dashboard_public", func(r *ujconfig.Resource) {
		r.References["dashboard_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_dashboard",
			RefFieldName:      "DashboardRef",
			SelectorFieldName: "DashboardSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
	})
	p.AddResourceConfigurator("grafana_dashboard_permission", func(r *ujconfig.Resource) {
		r.References["dashboard_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_dashboard",
			RefFieldName:      "DashboardRef",
			SelectorFieldName: "DashboardSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
		r.References["permissions.team_id"] = ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRef",
			SelectorFieldName: "TeamSelector",
		}
		r.References["permissions.user_id"] = ujconfig.Reference{
			TerraformName:     "grafana_user",
			RefFieldName:      "UserRef",
			SelectorFieldName: "UserSelector",
		}
	})
	p.AddResourceConfigurator("grafana_data_source_permission", func(r *ujconfig.Resource) {
		r.References["datasource_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_data_source",
			RefFieldName:      "DataSourceRef",
			SelectorFieldName: "DataSourceSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
		r.References["permissions.team_id"] = ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRef",
			SelectorFieldName: "TeamSelector",
		}
		r.References["permissions.user_id"] = ujconfig.Reference{
			TerraformName:     "grafana_user",
			RefFieldName:      "UserRef",
			SelectorFieldName: "UserSelector",
		}
	})
	p.AddResourceConfigurator("grafana_folder", func(r *ujconfig.Resource) {
		r.References["parent_folder_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_folder",
			RefFieldName:      "FolderRef",
			SelectorFieldName: "FolderSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
	})
	p.AddResourceConfigurator("grafana_folder_permission", func(r *ujconfig.Resource) {
		r.References["folder_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_folder",
			RefFieldName:      "FolderRef",
			SelectorFieldName: "FolderSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
		r.References["permissions.team_id"] = ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRef",
			SelectorFieldName: "TeamSelector",
		}
		r.References["permissions.user_id"] = ujconfig.Reference{
			TerraformName:     "grafana_user",
			RefFieldName:      "UserRef",
			SelectorFieldName: "UserSelector",
		}
	})
	p.AddResourceConfigurator("grafana_library_panel", func(r *ujconfig.Resource) {
		r.References["folder_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_folder",
			RefFieldName:      "FolderRef",
			SelectorFieldName: "FolderSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
	})
	p.AddResourceConfigurator("grafana_notification_policy", func(r *ujconfig.Resource) {
		r.References["contact_point"] = ujconfig.Reference{
			TerraformName:     "grafana_contact_point",
			RefFieldName:      "ContactPointRef",
			SelectorFieldName: "ContactPointSelector",
			Extractor:         fieldExtractor("name"),
		}
	})
	p.AddResourceConfigurator("grafana_report", func(r *ujconfig.Resource) {
		r.References["dashboard_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_dashboard",
			RefFieldName:      "DashboardRef",
			SelectorFieldName: "DashboardSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
	})
	p.AddResourceConfigurator("grafana_role", func(r *ujconfig.Resource) {
		r.InitializerFns = append(r.InitializerFns, createroleInitializer)
	})
	p.AddResourceConfigurator("grafana_role_assignment", func(r *ujconfig.Resource) {
		r.References["role_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_role",
			RefFieldName:      "RoleRef",
			SelectorFieldName: "RoleSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
		r.References["service_accounts"] = ujconfig.Reference{
			TerraformName:     "grafana_service_account",
			RefFieldName:      "ServiceAccountRefs",
			SelectorFieldName: "ServiceAccountSelector",
		}
		r.References["teams"] = ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRefs",
			SelectorFieldName: "TeamSelector",
		}
		r.References["users"] = ujconfig.Reference{
			TerraformName:     "grafana_user",
			RefFieldName:      "UserRefs",
			SelectorFieldName: "UserSelector",
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
	p.AddResourceConfigurator("grafana_team", func(r *ujconfig.Resource) {
		r.References["members"] = ujconfig.Reference{
			TerraformName:     "grafana_user",
			RefFieldName:      "MemberRefs",
			SelectorFieldName: "MemberSelector",
			Extractor:         fieldExtractor("email"),
		}
	})
	p.AddResourceConfigurator("grafana_team_external_group", func(r *ujconfig.Resource) {
		r.References["team_id"] = ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRef",
			SelectorFieldName: "TeamSelector",
		}
	})
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

	p.AddResourceConfigurator("grafana_slo", func(r *ujconfig.Resource) {
		r.References["destination_datasource.uid"] = ujconfig.Reference{
			TerraformName:     "grafana_data_source",
			RefFieldName:      "Ref",
			SelectorFieldName: "Selector",
			Extractor:         optionalFieldExtractor("uid"),
		}
	})

	// Configuration for k6 resources
	p.AddResourceConfigurator("grafana_k6_project_limits", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{
			TerraformName:     "grafana_k6_project",
			RefFieldName:      "ProjectRef",
			SelectorFieldName: "ProjectSelector",
		}
	})
	p.AddResourceConfigurator("grafana_k6_load_test", func(r *ujconfig.Resource) {
		r.References["project_id"] = ujconfig.Reference{
			TerraformName:     "grafana_k6_project",
			RefFieldName:      "ProjectRef",
			SelectorFieldName: "ProjectSelector",
		}
	})
}
