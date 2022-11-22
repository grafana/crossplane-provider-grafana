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
		// "aws_route53_resolver_rule": "route53resolver" -> (route53resolver, Rule)
		words := strings.Split(strings.TrimPrefix(resource, "aws_"), "_")
		snakeKind := strings.Join(words[count:], "_")
		return group, name.NewFromSnake(snakeKind).Camel
	}
}

// GroupMap contains all overrides we'd like to make to the default group search.
var GroupMap = map[string]GroupKindCalculator{
	"grafana_dashboard": ReplaceGroupWords("oss", 0),
}

// KindMap contains kind string overrides.
var KindMap = map[string]string{}
