/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"fmt"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	grafanaProvider "github.com/grafana/terraform-provider-grafana/v3/pkg/provider"
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

// resourcesByFramework returns resources by framework (legacy SDK or plugin framework)
func resourcesByFramework() ([]string, []string) {
	resourcesMap := map[string]bool{}
	for _, r := range grafanaProvider.Resources() {
		resourcesMap[r.Name] = r.PluginFrameworkSchema != nil
		if _, ok := GroupMap[r.Name]; !ok {
			// list resource not yet included in config/groups.go
			fmt.Printf("Not in groupMap: %s\n", r.Name)
		}

	}

	var legacySDKResources, pluginFrameworkResources []string
	for crossplaneResource := range GroupMap {
		isPluginFrameworkResource := resourcesMap[crossplaneResource]
		regexResource := "^" + crossplaneResource + "$"
		if isPluginFrameworkResource {
			pluginFrameworkResources = append(pluginFrameworkResources, regexResource)
		} else {
			legacySDKResources = append(legacySDKResources, regexResource)
		}
	}

	return legacySDKResources, pluginFrameworkResources
}

// GetProvider returns provider configuration
func GetProvider(generationProvider bool) (*ujconfig.Provider, error) {
	var p *schema.Provider
	var err error
	if generationProvider {
		p, err = getProviderSchema(providerSchema)
	} else {
		p = grafanaProvider.Provider("crossplane")
	}
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get the Terraform provider schema with generation mode set to %t", generationProvider)
	}

	legacySDKResources, pluginFrameworkResources := resourcesByFramework()
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithShortName("grafana"),
		ujconfig.WithRootGroup("grafana.crossplane.io"),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithIncludeList([]string{}),
		ujconfig.WithTerraformPluginSDKIncludeList(legacySDKResources),
		ujconfig.WithTerraformPluginFrameworkIncludeList(pluginFrameworkResources),
		ujconfig.WithTerraformProvider(p),
		ujconfig.WithTerraformPluginFrameworkProvider(grafanaProvider.FrameworkProvider("crossplane")),
		ujconfig.WithDefaultResourceOptions(
			GroupKindOverrides(),
			KindOverrides(),
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		grafana.ConfigureOrgIDRefs,
		grafana.Configure,
		grafana.ConfigureOnCallRefsAndSelectors,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}
