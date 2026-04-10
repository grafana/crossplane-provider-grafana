/*
Copyright 2021-2026 Grafana Labs
*/

package grafana

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
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

// IgnoreOrgIDLateInit prevents upjet from late-initializing org_id on all
// resources that have the field. The Grafana API always returns org_id in
// responses, even when using API key auth. Without this, upjet's
// late-initialization copies the value back into spec.forProvider, causing
// subsequent reconciliations to fail with:
//
//	"org_id is only supported with basic auth. API keys are already org-scoped"
func IgnoreOrgIDLateInit(p *ujconfig.Provider) {
	for name, resource := range p.Resources {
		if resource.TerraformResource.Schema["org_id"] == nil {
			continue
		}
		p.AddResourceConfigurator(name, func(r *ujconfig.Resource) {
			r.LateInitializer.IgnoredFields = append(r.LateInitializer.IgnoredFields, "org_id")
		})
	}
}

// Configure configures the grafana group
func Configure(p *ujconfig.Provider) {
	// configures all resources to be synced without async callbacks, the Grafana API is synchronous
	for _, resource := range p.Resources {
		resource.UseAsync = false
	}

	configureAlerting(p)
	configureCloud(p)
	configureEnterprise(p)
	configureOSS(p)
	configureSyntheticMonitoring(p)
	configureML(p)
	configureSLO(p)
	configureK6(p)
	configureOnCall(p)
}
