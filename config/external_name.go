/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"grafana_api_key":                           config.IdentifierFromProvider,
	"grafana_service_account":                   config.IdentifierFromProvider,
	"grafana_service_account_permission":        config.IdentifierFromProvider,
	"grafana_service_account_token":             config.IdentifierFromProvider,
	"grafana_cloud_access_policy":               config.IdentifierFromProvider,
	"grafana_cloud_access_policy_token":         config.IdentifierFromProvider,
	"grafana_cloud_api_key":                     config.IdentifierFromProvider,
	"grafana_cloud_plugin_installation":         config.IdentifierFromProvider,
	"grafana_cloud_stack":                       config.IdentifierFromProvider,
	"grafana_cloud_stack_service_account":       config.IdentifierFromProvider,
	"grafana_cloud_stack_service_account_token": config.IdentifierFromProvider,
	"grafana_contact_point":                     config.IdentifierFromProvider,
	"grafana_dashboard_permission":              config.IdentifierFromProvider,
	"grafana_dashboard":                         config.IdentifierFromProvider,
	"grafana_data_source":                       config.IdentifierFromProvider,
	"grafana_folder_permission":                 config.IdentifierFromProvider,
	"grafana_folder":                            config.IdentifierFromProvider,
	"grafana_message_template":                  config.IdentifierFromProvider,
	"grafana_mute_timing":                       config.IdentifierFromProvider,
	"grafana_notification_policy":               config.IdentifierFromProvider,
	"grafana_oncall_escalation_chain":           config.IdentifierFromProvider,
	"grafana_oncall_escalation":                 config.IdentifierFromProvider,
	"grafana_oncall_integration":                config.IdentifierFromProvider,
	"grafana_oncall_on_call_shift":              config.IdentifierFromProvider,
	"grafana_oncall_outgoing_webhook":           config.IdentifierFromProvider,
	"grafana_oncall_route":                      config.IdentifierFromProvider,
	"grafana_oncall_schedule":                   config.IdentifierFromProvider,
	"grafana_report":                            config.IdentifierFromProvider,
	"grafana_rule_group":                        config.IdentifierFromProvider,
	"grafana_synthetic_monitoring_check":        config.IdentifierFromProvider,
	"grafana_synthetic_monitoring_installation": config.IdentifierFromProvider,
	"grafana_synthetic_monitoring_probe":        config.IdentifierFromProvider,
	"grafana_team":                              config.IdentifierFromProvider,
	"grafana_user":                              config.IdentifierFromProvider,
	"grafana_organization":                      config.IdentifierFromProvider,
	"grafana_organization_preferences":          config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
