/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

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

// workaround for the TF Azure v3.57.0-based no-fork release: We would like to
// keep the types in the generated CRDs intact
// (prevent number->int type replacements).
func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns provider configuration
func GetProvider() (*ujconfig.Provider, error) {
	p, err := getProviderSchema(providerSchema)
	if err != nil {
		return nil, err
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithShortName("grafana"),
		ujconfig.WithRootGroup("grafana.crossplane.io"),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithIncludeList([]string{}),
		ujconfig.WithTerraformPluginSDKIncludeList(ExternalNameConfigured()),
		ujconfig.WithTerraformPluginFrameworkIncludeList([]string{}), // For future resources
		ujconfig.WithTerraformProvider(p),
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
	return pc, nil
}
