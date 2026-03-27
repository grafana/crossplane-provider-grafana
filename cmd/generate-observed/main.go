/*
Copyright 2026 Grafana Labs

generate-observed introspects the Terraform Grafana provider's data sources
and emits Crossplane observe-only resource types, controller specs, and
registration boilerplate.

Usage:

	go run ./cmd/generate-observed
*/
package main

import (
	"log"
	"strings"

	grafanaProvider "github.com/grafana/terraform-provider-grafana/v4/pkg/provider"

	"github.com/grafana/crossplane-provider-grafana/v2/config"
	"github.com/grafana/crossplane-provider-grafana/v2/pkg/generateobserved"
)

func main() {
	overrides := map[string]generateobserved.CategoryRule{}
	for _, ds := range grafanaProvider.DataSources() {
		ci, ok := config.CategoryConfig[string(ds.Category)]
		if !ok {
			log.Fatalf("unknown category %q for data source %s", ds.Category, ds.Name)
		}
		words := strings.Split(strings.TrimPrefix(ds.Name, "grafana_"), "_")
		tfPrefix := "grafana_"
		if ci.WordCount > 0 {
			tfPrefix += strings.Join(words[:ci.WordCount], "_") + "_"
		}
		overrides[ds.Name] = generateobserved.CategoryRule{
			GroupPrefix: ci.ShortGroup,
			DirName:     ci.ShortGroup,
			TFPrefix:    tfPrefix,
		}
	}

	cfg := generateobserved.Config{
		ModulePath:      "github.com/grafana/crossplane-provider-grafana/v2",
		APIVersion:      "v1alpha1",
		CopyrightHeader: "Copyright 2026 Grafana Labs",
		GroupSuffix:     ".grafana.o.crossplane.io",

		TFProviderImport:        "github.com/grafana/terraform-provider-grafana/v4/pkg/provider",
		TFFrameworkProviderFunc: "grafanaProvider.FrameworkProvider",
		TFProviderImportAlias:   "grafanaProvider",

		CategoryOverrides: overrides,

		TFDataSourcePkg:            "github.com/grafana/crossplane-provider-grafana/v2/pkg/tfdatasource",
		TFLegacyProviderFunc:       `grafanaProvider.Provider("crossplane")`,
		ConfigExtractorImport:      "github.com/grafana/crossplane-provider-grafana/v2/internal/clients",
		ConfigExtractorImportAlias: "clients",
		ConfigExtractorFunc:        "clients.ExtractModernConfig",
		TFConfigBuilderFunc:        "clients.BuildTFConfig",

		FallbackTFPrefix: "grafana_",

		Acronyms: []string{
			"ID", "URL", "API", "IP", "IPS", "HTTP", "HTTPS", "SSH", "SSL", "TLS",
			"DNS", "TCP", "UDP", "HTML", "CSS", "JSON", "XML", "YAML", "SQL",
			"AWS", "GCP", "SMTP", "LDAP", "SAML", "RBAC", "UID", "UUID", "URI",
			"SLO", "SM", "O11Y", "K6", "LBAC",
		},
	}

	p := grafanaProvider.Provider("crossplane")
	fwp := grafanaProvider.FrameworkProvider("crossplane")

	generateobserved.Generate(cfg, p, fwp)
}
