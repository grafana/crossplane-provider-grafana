/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	grafana "github.com/grafana/crossplane-provider-grafana/config/grafana"
)

const (
	resourcePrefix = "grafana"
	modulePath     = "github.com/grafana/crossplane-provider-grafana"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithShortName("grafana"),
		ujconfig.WithRootGroup("grafana.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			GroupKindOverrides(),
			KindOverrides(),
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		grafana.ConfigureOrgIDRefs,
		grafana.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
