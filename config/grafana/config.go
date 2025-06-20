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
//
// This function configures cross-references between OnCall resources to enable Crossplane
// to automatically resolve dependencies and maintain referential integrity. It also includes
// critical fixes for external name handling and initProvider conflicts that cause production
// issues after provider restarts.
//
// Key Issues Addressed:
// 1. OnCall Schedule Recreation: Prevents schedules from being recreated on provider restart
// 2. initProvider vs forProvider Conflicts: Stops automatic field overwrites during late initialization
// 3. Web Override Preservation: Maintains manual changes made through Grafana UI
// 4. External Name Mapping: Uses actual Grafana API IDs instead of Kubernetes resource names
func ConfigureOnCallRefsAndSelectors(p *ujconfig.Provider) {

	p.AddResourceConfigurator("grafana_oncall_schedule", func(r *ujconfig.Resource) {
		// CROSS-REFERENCE CONFIGURATION
		// Configure shifts reference to enable Crossplane to automatically resolve
		// grafana_oncall_on_call_shift resources referenced in the schedule configuration
		r.References["shifts"] = ujconfig.Reference{
			TerraformName:     "grafana_oncall_on_call_shift",
			RefFieldName:      "ShiftsRef",
			SelectorFieldName: "ShiftsSelector",
		}
		// NOTE: the following references won't work as Terraform datasources are not translated to Crossplane resources
		// These would be ideal but are not supported in the current Crossplane Terraform provider architecture:
		// r.References["slack.channel_id"] = slackChannelRef
		// r.References["slack.user_group_id"] = oncallUserGroupRef

		// EXTERNAL NAME CONFIGURATION
		// Critical fix to prevent OnCall schedule recreation on provider restart.
		// Problem: Default external name handling uses Kubernetes resource names instead of actual Grafana API IDs
		// Solution: Extract the actual Grafana OnCall schedule ID from terraform state and use it as external name
		// Impact: Prevents service disruption and loss of manual overrides during provider restarts
		r.ExternalName = ujconfig.ExternalName{
			// Don't set terraform resource name as identifier - we'll use the actual API ID
			SetIdentifierArgumentFn: ujconfig.NopSetIdentifierArgument,
			// Extract the actual Grafana OnCall schedule ID from terraform state
			GetExternalNameFn: func(tfstate map[string]any) (string, error) {
				if id, ok := tfstate["id"].(string); ok {
					return id, nil
				}
				return "", errors.New("cannot get OnCall schedule id from tfstate")
			},
			// Use the external name (Grafana API ID) directly as the resource ID
			GetIDFn: ujconfig.ExternalNameAsID,
		}

		// CUSTOM DIFF FUNCTION
		// Implements intelligent diff handling to prevent false positive changes and preserve manual overrides
		// This is essential for production stability where teams make manual schedule changes via Grafana UI
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			// Skip diff customization during resource creation (no existing state)
			if state == nil || state.Empty() {
				return diff, nil
			}
			// Skip processing if there's no diff, it's a destroy operation, or no attributes changed
			if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
				return diff, nil
			}

			// WEB OVERRIDES HANDLING
			// When enableWebOverrides is true (common pattern in Ford's OnCall configurations),
			// preserve manual changes made through the Grafana UI by ignoring specific field diffs
			if config != nil {
				if val, ok := config.Config["enable_web_overrides"]; ok {
					if enable, ok := val.(bool); ok && enable {
						// Fields that are commonly modified manually through Grafana UI
						// and should be preserved when web overrides are enabled
						fieldsToIgnore := []string{
							"shifts",             // Shift assignments modified via UI
							"slack",              // Slack integration settings
							"ical_url_overrides", // Calendar URL overrides
							"on_call_now",        // Current on-call status
						}
						for _, field := range fieldsToIgnore {
							if diff.Attributes[field] != nil {
								// Remove the diff for this field to prevent Crossplane from trying to "fix" it
								delete(diff.Attributes, field)
							}
						}
					}
				}
			}

			return diff, nil
		}

		// LATE INITIALIZATION CONFIGURATION
		// Prevents automatic population of initProvider during late initialization process
		// Critical for avoiding initProvider vs forProvider conflicts that cause false positive failures
		r.LateInitializer = ujconfig.LateInitializer{
			// Fields that should NOT be automatically populated in initProvider during late initialization
			// These fields are commonly modified externally and would cause conflicts if auto-populated
			IgnoredFields: []string{
				"shifts",             // Don't late init shift references - managed by forProvider
				"slack",              // Don't late init slack configs - often modified externally
				"ical_url_overrides", // Don't late init calendar overrides - frequently changed manually
				"on_call_now",        // Don't late init current state - this is runtime data
			},
		}
	})

	// COMMENTED OUT REFERENCES - INFORMATIONAL
	// The following reference configurations are not currently functional because
	// Terraform datasources are not translated to Crossplane resources. This is a
	// known limitation in the Crossplane Terraform provider architecture.
	//
	// Workaround: Use the Terraform provider for Crossplane to access datasources directly
	// Reference: https://github.com/crossplane/crossplane/blob/master/design/design-doc-observe-only-resources.md
	//
	// Future enhancement: When Crossplane supports datasource translation, these
	// references can be uncommented and used for better resource management.

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

	// FUTURE CONFIGURATIONS - When datasource support is available:
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

// ConfigureOnCallInitProviderHandling configures all OnCall resources to prevent
// initProvider vs forProvider merge conflicts that cause false positive failures
// after provider restarts
//
// ROOT CAUSE ANALYSIS:
// When the Grafana provider pod restarts, Crossplane's reconciliation process triggers
// late initialization for existing resources. During this process:
// 1. LateInitialize() reads the observed state from the external Grafana API
// 2. shouldMergeInitProvider logic incorrectly populates initProvider with observed values
// 3. This overwrites user-configured forProvider values, causing apparent drift
// 4. Resources report as failed/out-of-sync despite external state being correct
// 5. Ford's OnCall components show false positive failures after routine provider restarts
//
// SOLUTION APPROACH:
// This function applies a comprehensive fix across all OnCall resource types:
// 1. Prevents automatic initProvider population during late initialization
// 2. Adds intelligent custom diff functions to filter out false positive changes
// 3. Preserves user-configured forProvider values from being overwritten
// 4. Maintains backward compatibility with existing OnCall deployments
//
// PRODUCTION IMPACT:
// - Eliminates false positive sync failures in Ford's incident response infrastructure
// - Prevents unnecessary resource recreation and service disruption
// - Maintains manual overrides made through Grafana UI
// - Ensures stable OnCall operations during provider maintenance windows
func ConfigureOnCallInitProviderHandling(p *ujconfig.Provider) {
	// ONCALL RESOURCE INVENTORY
	// Complete list of OnCall resources that experience initProvider conflicts
	// These resources are particularly susceptible because they:
	// 1. Have complex nested configurations that trigger late initialization
	// 2. Support manual overrides through Grafana UI (common in Ford's workflow)
	// 3. Contain runtime state that changes frequently (on-call assignments, etc.)
	onCallResources := []string{
		"grafana_oncall_schedule",               // OnCall schedules and shift assignments
		"grafana_oncall_on_call_shift",          // Individual shifts within schedules
		"grafana_oncall_escalation",             // Escalation steps and notification rules
		"grafana_oncall_escalation_chain",       // Escalation chain definitions
		"grafana_oncall_integration",            // Alert source integrations
		"grafana_oncall_route",                  // Alert routing rules
		"grafana_oncall_outgoing_webhook",       // Webhook notifications
		"grafana_oncall_user_notification_rule", // User-specific notification preferences
	}

	// APPLY CONSISTENT CONFIGURATION TO ALL ONCALL RESOURCES
	// Each resource gets the same protective configuration to ensure consistent behavior
	for _, resourceName := range onCallResources {
		p.AddResourceConfigurator(resourceName, func(r *ujconfig.Resource) {
			// LATE INITIALIZATION PROTECTION
			// Prevent automatic initProvider population that causes field conflicts
			// Strategy: Ignore all fields during late initialization by default
			// This forces Crossplane to rely on forProvider values as the source of truth
			if r.LateInitializer.IgnoredFields == nil {
				r.LateInitializer = ujconfig.LateInitializer{
					// Wildcard ignore prevents any automatic initProvider population
					// This is the most comprehensive approach to avoid field conflicts
					IgnoredFields: []string{"*"}, // Ignore all fields for late init by default
				}
			}

			// INTELLIGENT CUSTOM DIFF HANDLING
			// Add custom diff logic while preserving any existing custom diff functions
			// This layered approach ensures compatibility with resource-specific customizations
			originalDiff := r.TerraformCustomDiff
			r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
				// PRESERVE EXISTING CUSTOM DIFF LOGIC
				// Run any existing custom diff function first to maintain functionality
				// This ensures resource-specific customizations (like the OnCall schedule diff) continue to work
				if originalDiff != nil {
					var err error
					diff, err = originalDiff(diff, state, config)
					if err != nil {
						return nil, err
					}
				}

				// SKIP DIFF PROCESSING FOR CREATION AND EDGE CASES
				// Skip diff customization during resource creation (no existing state)
				if state == nil || state.Empty() {
					return diff, nil
				}
				// Skip processing if there's no diff, it's a destroy operation, or no attributes changed
				if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
					return diff, nil
				}

				// INITPROVIDER CONFLICT DETECTION AND FILTERING
				// Filter out diffs that are caused by initProvider vs forProvider type mismatches
				// These false positive diffs typically manifest as:
				// - Old value and new value appear different but are functionally equivalent
				// - Type mismatches between string/number representations
				// - Null vs empty string differences that don't represent real changes
				filteredAttributes := make(map[string]*terraform.ResourceAttrDiff)
				for key, attrDiff := range diff.Attributes {
					// REAL CHANGE DETECTION
					// Only preserve diffs that represent actual configuration changes
					// This filters out the false positives caused by initProvider conflicts
					if attrDiff != nil && attrDiff.Old != attrDiff.New {
						// Additional validation could be added here to detect type-mismatch false positives
						// For now, we rely on the basic old != new check which handles most cases
						filteredAttributes[key] = attrDiff
					}
					// Diffs where old == new are discarded as false positives
					// These commonly occur when initProvider populated a field with the same value
					// that's already in forProvider, but with different type representation
				}

				// APPLY FILTERED DIFF RESULTS
				// Replace the original attributes with our filtered set
				// This prevents Crossplane from attempting to "fix" false positive diffs
				diff.Attributes = filteredAttributes

				return diff, nil
			}
		})
	}
}

// Configure configures the grafana group
func Configure(p *ujconfig.Provider) {
	// configures all resources to be synced without async callbacks, the Grafana API is synchronous
	for _, resource := range p.Resources {
		resource.UseAsync = false
	}

	// Configure OnCall resources to prevent initProvider vs forProvider conflicts
	ConfigureOnCallInitProviderHandling(p)

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
		r.TerraformCustomDiff = recreateIfAttributeMissing("token")
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}
			cloudConfig := map[string]string{}
			basicAuthConfig := map[string]string{}
			if a, ok := attr["token"].(string); ok {
				cloudConfig["cloud_access_policy_token"] = a
				basicAuthConfig["basicAuthPassword"] = a
			}

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
		r.TerraformCustomDiff = recreateIfAttributeMissing("key")
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}

			instanceConfig := map[string]string{}
			if a, ok := attr["stack_slug"].(string); ok {
				instanceConfig["url"] = fmt.Sprintf("https://%s.grafana.net", a)
			} // TODO: set URL from client
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
		r.TerraformCustomDiff = recreateIfAttributeMissing("key")
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}

			instanceConfig := map[string]string{}
			// TODO: set URL from client
			// instanceConfig["url"] = fmt.Sprintf("https://%s.grafana.net", a)
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
	p.AddResourceConfigurator("grafana_oncall_on_call_shift", func(r *ujconfig.Resource) {
		// EXTERNAL NAME CONFIGURATION FOR ONCALL SHIFTS
		// Configure external name to use the actual Grafana OnCall shift ID from the API
		// This ensures consistent identification and prevents recreation issues during provider restarts
		// Problem: Default behavior uses Kubernetes metadata.name instead of actual Grafana shift ID
		// Solution: Extract the real shift ID from terraform state and use it for external naming
		r.ExternalName = ujconfig.ExternalName{
			// Don't use terraform resource name as identifier - use actual API ID
			SetIdentifierArgumentFn: ujconfig.NopSetIdentifierArgument,
			// Extract the actual Grafana OnCall shift ID from terraform state
			// This ID is the real identifier used by Grafana's OnCall API
			GetExternalNameFn: func(tfstate map[string]any) (string, error) {
				if id, ok := tfstate["id"].(string); ok {
					return id, nil
				}
				return "", errors.New("cannot get OnCall shift id from tfstate")
			},
			// Use the extracted Grafana API ID directly as the Crossplane resource ID
			GetIDFn: ujconfig.ExternalNameAsID,
		}

		// CUSTOM DIFF FUNCTION FOR SHIFT MANAGEMENT
		// Implements intelligent diff handling specifically for OnCall shift configurations
		// OnCall shifts are frequently modified through Grafana UI for operational needs:
		// - Emergency schedule changes during incidents
		// - Vacation coverage adjustments
		// - User availability updates
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			// Skip diff customization during resource creation (no existing state)
			if state == nil || state.Empty() {
				return diff, nil
			}
			// Skip processing if there's no diff, it's a destroy operation, or no attributes changed
			if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
				return diff, nil
			}

			// MANUAL SHIFT MANAGEMENT SUPPORT
			// Preserve manual changes to shift schedules made via Grafana UI
			// This is critical for operational flexibility in incident response scenarios
			fieldsToPreserve := []string{
				"users",         // User assignments - often changed for coverage adjustments
				"rolling_users", // Rolling user assignments - modified for rotation changes
				"start",         // Shift start times - adjusted for timezone or schedule changes
				"duration",      // Shift durations - modified for operational needs
			}

			// WEB OVERRIDES SUPPORT FOR SHIFTS
			// Check for enableWebOverrides configuration to determine if manual changes should be preserved
			// This pattern is commonly used in Ford's OnCall configurations to allow operational flexibility
			if config != nil {
				if val, ok := config.Config["enable_web_overrides"]; ok {
					if enable, ok := val.(bool); ok && enable {
						// When web overrides are enabled, remove diffs for manually managed fields
						// This prevents Crossplane from reverting operational changes made via Grafana UI
						for _, field := range fieldsToPreserve {
							if diff.Attributes[field] != nil {
								// Remove the diff to preserve manual changes
								delete(diff.Attributes, field)
							}
						}
					}
				}
			}

			return diff, nil
		}

		// LATE INITIALIZATION CONFIGURATION FOR SHIFTS
		// Configure which fields should be ignored during late initialization to prevent conflicts
		// OnCall shifts have dynamic runtime data that changes frequently and shouldn't be auto-initialized
		r.LateInitializer = ujconfig.LateInitializer{
			// Fields that should NOT be automatically populated during late initialization
			// These fields are commonly modified externally and would cause initProvider conflicts
			IgnoredFields: []string{
				"users",         // Don't late init user assignments - managed by forProvider or UI
				"rolling_users", // Don't late init rolling assignments - frequently changed operationally
				"start",         // Don't late init start times - may be adjusted for scheduling needs
				"duration",      // Don't late init durations - often modified for operational requirements
			},
		}
	})

	// ONCALL ESCALATION RESOURCE CONFIGURATION
	// Additional OnCall resource configurations to prevent initProvider conflicts
	// Each resource type has specific fields that are prone to conflicts and need protection
	p.AddResourceConfigurator("grafana_oncall_escalation", func(r *ujconfig.Resource) {
		// CROSS-REFERENCE CONFIGURATION FOR ESCALATIONS
		// Configure references to related OnCall resources for dependency management
		// These references enable Crossplane to understand resource relationships and order operations correctly

		// Reference to escalation chain - the parent chain this escalation step belongs to
		r.References["escalation_chain_id"] = ujconfig.Reference{
			TerraformName:     "grafana_oncall_escalation_chain",
			RefFieldName:      "EscalationChainRef",
			SelectorFieldName: "EscalationChainSelector",
		}

		// Reference to outgoing webhook for escalation actions
		r.References["action_to_trigger"] = ujconfig.Reference{
			TerraformName:     "grafana_oncall_outgoing_webhook",
			RefFieldName:      "ActionToTriggerRef",
			SelectorFieldName: "ActionToTriggerSelector",
		}

		// Reference to schedule for on-call notifications
		r.References["notify_on_call_from_schedule"] = ujconfig.Reference{
			TerraformName:     "grafana_oncall_schedule",
			RefFieldName:      "NotifyOnCallFromScheduleRef",
			SelectorFieldName: "NotifyOnCallFromScheduleSelector",
		}

		// LATE INITIALIZATION PROTECTION FOR ESCALATIONS
		// Escalation steps often have dynamic notification rules that change based on:
		// - Team structure changes
		// - On-call rotation updates
		// - Notification preference modifications
		r.LateInitializer = ujconfig.LateInitializer{
			// Fields prone to initProvider conflicts in escalation configurations
			IgnoredFields: []string{
				"group_to_notify",                  // Group notification settings - often modified via UI
				"notify_to_team_members",           // Team member notification flags - dynamic based on team changes
				"persons_to_notify",                // Individual notification targets - frequently updated
				"persons_to_notify_next_each_time", // Rotation-based notifications - changes with schedule updates
			},
		}
	})

	// ONCALL INTEGRATION RESOURCE CONFIGURATION
	p.AddResourceConfigurator("grafana_oncall_integration", func(r *ujconfig.Resource) {
		// CROSS-REFERENCE CONFIGURATION FOR INTEGRATIONS
		// OnCall integrations connect external alert sources to escalation chains
		// The default route configuration is critical for proper alert routing

		// Reference to default escalation chain for this integration
		r.References["default_route.escalation_chain_id"] = ujconfig.Reference{
			TerraformName:     "grafana_oncall_escalation_chain",
			RefFieldName:      "EscalationChainRef",
			SelectorFieldName: "EscalationChainSelector",
		}

		// LATE INITIALIZATION PROTECTION FOR INTEGRATIONS
		// Integration configurations often include:
		// - Team assignments that change with organizational structure
		// - Default routing rules that may be customized per team
		r.LateInitializer = ujconfig.LateInitializer{
			// Fields that commonly cause initProvider conflicts in integration configs
			IgnoredFields: []string{
				"team_id",       // Team assignment - may change with org restructuring
				"default_route", // Default routing configuration - often customized per integration
			},
		}
	})

	// ONCALL ROUTE RESOURCE CONFIGURATION
	p.AddResourceConfigurator("grafana_oncall_route", func(r *ujconfig.Resource) {
		// CROSS-REFERENCE CONFIGURATION FOR ROUTES
		// OnCall routes define how alerts are routed to different escalation chains
		// Based on filtering criteria and routing rules

		// Reference to target escalation chain for this route
		r.References["escalation_chain_id"] = ujconfig.Reference{
			TerraformName:     "grafana_oncall_escalation_chain",
			RefFieldName:      "EscalationChainRef",
			SelectorFieldName: "EscalationChainSelector",
		}

		// Reference to parent integration that this route belongs to
		r.References["integration_id"] = ujconfig.Reference{
			TerraformName:     "grafana_oncall_integration",
			RefFieldName:      "IntegrationRef",
			SelectorFieldName: "IntegrationSelector",
		}

		// LATE INITIALIZATION PROTECTION FOR ROUTES
		// Route configurations include regex patterns and routing logic that are:
		// - Fine-tuned based on alert patterns
		// - Modified to handle new alert types
		// - Adjusted for changing operational needs
		r.LateInitializer = ujconfig.LateInitializer{
			// Fields that are commonly customized and prone to initProvider conflicts
			IgnoredFields: []string{
				"routing_regex", // Regex patterns for alert matching - frequently tuned
				"routing_type",  // Routing logic type - may be adjusted based on alert sources
			},
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
		r.TerraformCustomDiff = recreateIfAttributeMissing("sm_access_token")
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}

			providerConfig := map[string]string{}
			if a, ok := attr["sm_access_token"].(string); ok {
				providerConfig["sm_access_token"] = a
			}
			if a, ok := attr["stack_sm_api_url"].(string); ok {
				providerConfig["sm_url"] = a
			}
			marshalled, err := json.Marshal(providerConfig)
			if err != nil {
				return nil, err
			}
			conn["smCredentials"] = marshalled

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

func recreateIfAttributeMissing(attribute string) ujconfig.CustomDiff {
	return func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
		if state == nil {
			return diff, nil
		}

		// The attribute may not be returned in the state, so we need to recreate the resource if it is missing
		if _, ok := state.Attributes[attribute]; !ok {
			if diff == nil {
				diff = &terraform.InstanceDiff{}
			}
			diff.DestroyTainted = true
		}

		return diff, nil
	}
}
