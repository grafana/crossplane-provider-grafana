/*
Copyright 2025 Grafana

generate-observed introspects the Terraform Grafana provider's data sources
and emits Crossplane observe-only resource types, controller specs, and
registration boilerplate.

Usage:

	go run ./cmd/generate-observed
*/
package main

import (
	grafanaProvider "github.com/grafana/terraform-provider-grafana/v4/pkg/provider"

	"github.com/grafana/crossplane-provider-grafana/v2/pkg/generateobserved"
)

func main() {
	cfg := generateobserved.Config{
		ModulePath:      "github.com/grafana/crossplane-provider-grafana/v2",
		APIVersion:      "v1alpha1",
		CopyrightHeader: "Copyright 2025 Grafana",
		GroupSuffix:     ".grafana.o.crossplane.io",

		TFProviderImport:        "github.com/grafana/terraform-provider-grafana/v4/pkg/provider",
		TFFrameworkProviderFunc: "grafanaProvider.FrameworkProvider",
		TFProviderImportAlias:   "grafanaProvider",

		CategoryRules: []generateobserved.CategoryRule{
			{GroupPrefix: "cloudprovider", DirName: "cloudprovider", TFPrefix: "grafana_cloud_provider_"},
			{GroupPrefix: "sm", DirName: "sm", TFPrefix: "grafana_synthetic_monitoring_"},
			{GroupPrefix: "fleetmanagement", DirName: "fleetmanagement", TFPrefix: "grafana_fleet_management_"},
			{GroupPrefix: "frontendobservability", DirName: "frontendobservability", TFPrefix: "grafana_frontend_o11y_"},
			{GroupPrefix: "connections", DirName: "connections", TFPrefix: "grafana_connections_"},
			{GroupPrefix: "oncall", DirName: "oncall", TFPrefix: "grafana_oncall_"},
			{GroupPrefix: "cloud", DirName: "cloud", TFPrefix: "grafana_cloud_"},
			{GroupPrefix: "k6", DirName: "k6", TFPrefix: "grafana_k6_"},
			{GroupPrefix: "slo", DirName: "slo", TFPrefix: "grafana_slo"},
			{GroupPrefix: "ml", DirName: "ml", TFPrefix: "grafana_machine_learning_"},
			// OSS: catch-all for remaining "grafana_*" data sources.
			{GroupPrefix: "oss", DirName: "oss", TFPrefix: "grafana_"},
		},

		CategoryOverrides: map[string]generateobserved.CategoryRule{
			"grafana_role": {GroupPrefix: "enterprise", DirName: "enterprise", TFPrefix: "grafana_"},
		},

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
