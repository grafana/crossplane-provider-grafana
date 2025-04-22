/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"strings"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/types/name"
)

// GroupKindOverrides overrides the group and kind of the resource if it matches
// any entry in the GroupMap.
func GroupKindOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		if f, ok := GroupMap[r.Name]; ok {
			r.ShortGroup, r.Kind = f(r.Name)
		}
	}
}

// KindOverrides overrides the kind of the resources given in KindMap.
func KindOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		if k, ok := KindMap[r.Name]; ok {
			r.Kind = k
		}
	}
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if _, ok := GroupMap[r.Name]; ok {
			r.ExternalName = config.IdentifierFromProvider
		}
	}
}

// GroupKindCalculator returns the correct group and kind name for given TF
// resource.
type GroupKindCalculator func(resource string) (string, string)

// ReplaceGroupWords uses given group as the group of the resource and removes
// a number of words in resource name before calculating the kind of the resource.
func ReplaceGroupWords(group string, count int) GroupKindCalculator {
	return func(resource string) (string, string) {
		words := strings.Split(strings.TrimPrefix(resource, "grafana_"), "_")
		snakeKind := strings.Join(words[count:], "_")
		return group, name.NewFromSnake(snakeKind).Camel
	}
}

// GroupMap contains all overrides we'd like to make to the default group search.
// Keep the same structure as in the Terraform docs: https://registry.terraform.io/providers/grafana/grafana/latest/docs
var GroupMap = map[string]GroupKindCalculator{
	// Alerting
	"grafana_contact_point":       ReplaceGroupWords("alerting", 0),
	"grafana_message_template":    ReplaceGroupWords("alerting", 0),
	"grafana_mute_timing":         ReplaceGroupWords("alerting", 0),
	"grafana_notification_policy": ReplaceGroupWords("alerting", 0),
	"grafana_rule_group":          ReplaceGroupWords("alerting", 0),

	// Cloud
	"grafana_cloud_access_policy":                             ReplaceGroupWords("cloud", 1),
	"grafana_cloud_access_policy_token":                       ReplaceGroupWords("cloud", 1),
	"grafana_cloud_org_member":                                ReplaceGroupWords("cloud", 1),
	"grafana_cloud_plugin_installation":                       ReplaceGroupWords("cloud", 1),
	"grafana_cloud_private_data_source_connect_network":       ReplaceGroupWords("cloud", 1),
	"grafana_cloud_private_data_source_connect_network_token": ReplaceGroupWords("cloud", 1),
	"grafana_cloud_stack":                                     ReplaceGroupWords("cloud", 1),
	"grafana_cloud_stack_service_account":                     ReplaceGroupWords("cloud", 1),
	"grafana_cloud_stack_service_account_token":               ReplaceGroupWords("cloud", 1),

	// Cloud Provider
	"grafana_cloud_provider_aws_account":               ReplaceGroupWords("cloudprovider", 2),
	"grafana_cloud_provider_aws_cloudwatch_scrape_job": ReplaceGroupWords("cloudprovider", 2),
	"grafana_cloud_provider_azure_credential":          ReplaceGroupWords("cloudprovider", 2),

	// Connections
	"grafana_connections_metrics_endpoint_scrape_job": ReplaceGroupWords("connections", 1),

	// Enterprise
	"grafana_data_source_config_lbac_rules": ReplaceGroupWords("enterprise", 0),
	"grafana_data_source_permission":        ReplaceGroupWords("enterprise", 0),
	"grafana_data_source_permission_item":   ReplaceGroupWords("enterprise", 0),
	"grafana_report":                        ReplaceGroupWords("enterprise", 0),
	"grafana_role":                          ReplaceGroupWords("enterprise", 0),
	"grafana_role_assignment":               ReplaceGroupWords("enterprise", 0),
	"grafana_role_assignment_item":          ReplaceGroupWords("enterprise", 0),
	"grafana_team_external_group":           ReplaceGroupWords("enterprise", 0),

	// Machine Learning
	"grafana_machine_learning_alert":            ReplaceGroupWords("ml", 2),
	"grafana_machine_learning_holiday":          ReplaceGroupWords("ml", 2),
	"grafana_machine_learning_job":              ReplaceGroupWords("ml", 2),
	"grafana_machine_learning_outlier_detector": ReplaceGroupWords("ml", 2),

	// OnCall
	"grafana_oncall_escalation":             ReplaceGroupWords("oncall", 1),
	"grafana_oncall_escalation_chain":       ReplaceGroupWords("oncall", 1),
	"grafana_oncall_integration":            ReplaceGroupWords("oncall", 1),
	"grafana_oncall_on_call_shift":          ReplaceGroupWords("oncall", 1),
	"grafana_oncall_outgoing_webhook":       ReplaceGroupWords("oncall", 1),
	"grafana_oncall_route":                  ReplaceGroupWords("oncall", 1),
	"grafana_oncall_schedule":               ReplaceGroupWords("oncall", 1),
	"grafana_oncall_user_notification_rule": ReplaceGroupWords("oncall", 1),

	// OSS
	"grafana_annotation":                      ReplaceGroupWords("oss", 0),
	"grafana_dashboard":                       ReplaceGroupWords("oss", 0),
	"grafana_dashboard_permission":            ReplaceGroupWords("oss", 0),
	"grafana_dashboard_permission_item":       ReplaceGroupWords("oss", 0),
	"grafana_dashboard_public":                ReplaceGroupWords("oss", 0),
	"grafana_data_source":                     ReplaceGroupWords("oss", 0),
	"grafana_data_source_config":              ReplaceGroupWords("oss", 0),
	"grafana_folder":                          ReplaceGroupWords("oss", 0),
	"grafana_folder_permission":               ReplaceGroupWords("oss", 0),
	"grafana_folder_permission_item":          ReplaceGroupWords("oss", 0),
	"grafana_library_panel":                   ReplaceGroupWords("oss", 0),
	"grafana_organization":                    ReplaceGroupWords("oss", 0),
	"grafana_organization_preferences":        ReplaceGroupWords("oss", 0),
	"grafana_playlist":                        ReplaceGroupWords("oss", 0),
	"grafana_service_account":                 ReplaceGroupWords("oss", 0),
	"grafana_service_account_permission":      ReplaceGroupWords("oss", 0),
	"grafana_service_account_permission_item": ReplaceGroupWords("oss", 0),
	"grafana_service_account_token":           ReplaceGroupWords("oss", 0),
	"grafana_sso_settings":                    ReplaceGroupWords("oss", 0),
	"grafana_team":                            ReplaceGroupWords("oss", 0),
	"grafana_user":                            ReplaceGroupWords("oss", 0),

	// SLO
	"grafana_slo": ReplaceGroupWords("slo", 0),

	// Synthetic Monitoring
	"grafana_synthetic_monitoring_check":        ReplaceGroupWords("sm", 2),
	"grafana_synthetic_monitoring_installation": ReplaceGroupWords("sm", 2),
	"grafana_synthetic_monitoring_probe":        ReplaceGroupWords("sm", 2),

	// Fleet Management
	"grafana_fleet_management_collector": ReplaceGroupWords("fleetmanagement", 2),
	"grafana_fleet_management_pipeline":  ReplaceGroupWords("fleetmanagement", 2),

	// Frontend Observability
	"grafana_frontend_o11y_app": ReplaceGroupWords("frontendobservability", 2),
}

// KindMap contains kind string overrides.
var KindMap = map[string]string{}
