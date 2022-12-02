/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"strings"

	"github.com/upbound/upjet/pkg/config"
	"github.com/upbound/upjet/pkg/types/name"
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
	"grafana_cloud_api_key": ReplaceGroupWords("cloud", 1),
	"grafana_cloud_stack":   ReplaceGroupWords("cloud", 1),

	// Enterprise
	"grafana_report": ReplaceGroupWords("enterprise", 0),

	// OnCall
	"grafana_oncall_escalation":       ReplaceGroupWords("oncall", 1),
	"grafana_oncall_escalation_chain": ReplaceGroupWords("oncall", 1),
	"grafana_oncall_integration":      ReplaceGroupWords("oncall", 1),
	"grafana_oncall_on_call_shift":    ReplaceGroupWords("oncall", 1),
	"grafana_oncall_outgoing_webhook": ReplaceGroupWords("oncall", 1),
	"grafana_oncall_route":            ReplaceGroupWords("oncall", 1),
	"grafana_oncall_schedule":         ReplaceGroupWords("oncall", 1),

	// OSS
	"grafana_api_key":              ReplaceGroupWords("oss", 0),
	"grafana_data_source":          ReplaceGroupWords("oss", 0),
	"grafana_dashboard":            ReplaceGroupWords("oss", 0),
	"grafana_dashboard_permission": ReplaceGroupWords("oss", 0),
	"grafana_folder":               ReplaceGroupWords("oss", 0),
	"grafana_folder_permission":    ReplaceGroupWords("oss", 0),
	"grafana_team":                 ReplaceGroupWords("oss", 0),
	"grafana_user":                 ReplaceGroupWords("oss", 0),

	// Synthetic Monitoring
	"grafana_synthetic_monitoring_check":        ReplaceGroupWords("sm", 2),
	"grafana_synthetic_monitoring_installation": ReplaceGroupWords("sm", 2),
	"grafana_synthetic_monitoring_probe":        ReplaceGroupWords("sm", 2),
}

// KindMap contains kind string overrides.
var KindMap = map[string]string{}
