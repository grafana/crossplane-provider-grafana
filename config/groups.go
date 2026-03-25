/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"fmt"
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/types/name"
	grafanaProvider "github.com/grafana/terraform-provider-grafana/v4/pkg/provider"
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

// CategoryInfo holds the short group name and prefix word count for a resource category.
type CategoryInfo struct {
	ShortGroup string
	WordCount  int
}

// CategoryConfig maps each upstream ResourceCategory string to the short group
// name and the number of prefix words to strip when computing the Kind.
// Category string values match common.ResourceCategory constants from the
// upstream Terraform provider.
var CategoryConfig = map[string]CategoryInfo{
	"Alerting":               {"alerting", 0},
	"Cloud":                  {"cloud", 1},
	"Grafana Enterprise":     {"enterprise", 0},
	"Grafana OSS":            {"oss", 0},
	"Grafana Apps":           {"oss", 0},
	"Machine Learning":       {"ml", 2},
	"OnCall":                 {"oncall", 1},
	"SLO":                    {"slo", 0},
	"Synthetic Monitoring":   {"sm", 2},
	"Cloud Provider":         {"cloudprovider", 2},
	"Connections":            {"connections", 1},
	"Fleet Management":       {"fleetmanagement", 2},
	"Frontend Observability": {"frontendobservability", 2},
	"k6":                     {"k6", 1},
	"Knowledge Graph":        {"asserts", 1},
}

// GroupMap contains all overrides we'd like to make to the default group search.
// It is populated dynamically from the upstream provider's resource categories.
var GroupMap = map[string]GroupKindCalculator{}

func init() {
	for _, r := range grafanaProvider.Resources() {
		cat := string(r.Category)
		cfg, ok := CategoryConfig[cat]
		if !ok {
			panic(fmt.Sprintf(
				"unknown category %q for resource %s\n"+
					"Add an entry to categoryConfig in config/groups.go",
				cat, r.Name))
		}
		GroupMap[r.Name] = ReplaceGroupWords(cfg.ShortGroup, cfg.WordCount)
	}
	for _, r := range grafanaProvider.AppPlatformResources() {
		cat := string(r.Category)
		cfg, ok := CategoryConfig[cat]
		if !ok {
			panic(fmt.Sprintf(
				"unknown category %q for resource %s\n"+
					"Add an entry to categoryConfig in config/groups.go",
				cat, r.Name))
		}
		// App platform resources always strip 2 prefix words (apps_{appname}).
		GroupMap[r.Name] = ReplaceGroupWords(cfg.ShortGroup, 2)
	}
}

// KindMap contains kind string overrides.
var KindMap = map[string]string{}
