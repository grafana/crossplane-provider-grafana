/*
Copyright 2025 Grafana
*/

// Package generateobserved provides a provider-agnostic code generator that
// introspects a Terraform provider's data sources (both legacy SDK and Plugin
// Framework) and emits Crossplane observe-only resource types, controller
// specs, factory initializers, and scheme/controller registration boilerplate.
//
// The generator produces the following file tree from TF provider introspection:
//
//	apis/observed/<category>/v1alpha1/
//	    zz_groupversion_info.go       — CRD group constants and scheme builder
//	    zz_<name>_types.go            — ForProvider/AtProvider structs, CRD type, GVK
//	internal/controller/namespaced/observed/<category>/
//	    zz_<name>_spec.go             — tfdatasource.Spec with read callbacks
//	    zz_factories.go               — Plugin Framework data source factories (if needed)
//	    zz_setup.go                   — per-group controller registration
//	apis/observed/zz_register.go      — aggregated AddToScheme across all groups
//	internal/controller/namespaced/observed/zz_setup.go — aggregated Setup/SetupGated
//
// All provider-specific configuration is supplied via [Config]. The
// generation engine itself has no knowledge of any particular Terraform
// provider.
//
// # Usage
//
//	cfg := generateobserved.Config{
//	    ModulePath:  "github.com/example/crossplane-provider-foo/v2",
//	    GroupSuffix: ".foo.o.crossplane.io",
//	    // ...
//	}
//	p := fooProvider.Provider("crossplane")
//	fwp := fooProvider.FrameworkProvider("crossplane")
//	generateobserved.Generate(cfg, p, fwp)
//
// # Data source classification
//
// Each TF data source is assigned to a category (CRD group) by matching its
// name against [Config.CategoryRules] (longest prefix wins).
// [Config.CategoryOverrides] takes precedence for individual names
// that don't fit the prefix convention.
//
// Within each category, field classification follows Terraform schema
// conventions: Required fields become ForProvider (spec inputs),
// Computed-only fields become AtProvider (status outputs), and Optional
// fields appear in both (since data sources populate all fields on read).
package generateobserved

// Config parameterizes the generator for a specific Terraform provider. All
// provider-specific settings are collected here; the generation engine itself
// is provider-agnostic.
type Config struct {
	// ModulePath is the Go module path of the Crossplane provider.
	// e.g. "github.com/grafana/crossplane-provider-grafana/v2"
	ModulePath string

	// APIVersion is the CRD API version for generated resources.
	// e.g. "v1alpha1"
	APIVersion string

	// CopyrightHeader is the copyright line for generated files.
	// e.g. "Copyright 2025 Grafana"
	CopyrightHeader string

	// GroupSuffix is appended to the category prefix to form the CRD group.
	// e.g. ".grafana.o.crossplane.io" → "oncall.grafana.o.crossplane.io"
	GroupSuffix string

	// TFProviderImport is the Go import path for the TF provider package,
	// used in generated factory files.
	// e.g. "github.com/grafana/terraform-provider-grafana/v4/pkg/provider"
	TFProviderImport string

	// TFFrameworkProviderFunc is the expression to call FrameworkProvider in
	// generated factory code. Must accept a version string argument.
	// e.g. "grafanaProvider.FrameworkProvider"
	TFFrameworkProviderFunc string

	// TFProviderImportAlias is the import alias for TFProviderImport in
	// generated factory files.
	// e.g. "grafanaProvider"
	TFProviderImportAlias string

	// CategoryRules map TF data source name prefixes to CRD groups.
	// Ordered longest-prefix-first so matching is unambiguous.
	CategoryRules []CategoryRule

	// CategoryOverrides maps specific TF data source names to a category
	// that differs from what the prefix rules would assign.
	CategoryOverrides map[string]CategoryRule

	// FallbackTFPrefix is stripped as a last resort when deriving Kind names
	// from TF data source names that don't match any category prefix.
	// e.g. "grafana_"
	FallbackTFPrefix string

	// TFDataSourcePkg is the import path for the public tfdatasource runtime.
	// e.g. "github.com/grafana/crossplane-provider-grafana/v2/pkg/tfdatasource"
	TFDataSourcePkg string

	// TFLegacyProviderFunc is the Go expression that returns a
	// *sdkschema.Provider for legacy data source factories and the connect
	// function. e.g. `grafanaProvider.Provider("crossplane")`
	TFLegacyProviderFunc string

	// ConfigExtractorImport is the import path for the package containing
	// the config extraction function used in the generated connect function.
	// e.g. "github.com/grafana/crossplane-provider-grafana/v2/internal/clients"
	ConfigExtractorImport string

	// ConfigExtractorImportAlias is the import alias for ConfigExtractorImport.
	// e.g. "clients"
	ConfigExtractorImportAlias string

	// ConfigExtractorFunc is the qualified function call to extract provider
	// config from a ModernManaged resource.
	// e.g. "clients.ExtractModernConfig"
	ConfigExtractorFunc string

	// TFConfigBuilderFunc is the qualified function call to build a TF config
	// map from the extracted config and credentials.
	// e.g. "clients.BuildTFConfig"
	TFConfigBuilderFunc string

	// Acronyms to keep uppercased in Go identifiers (e.g. "ID", "URL", "API").
	Acronyms []string
}

// CategoryRule maps a TF data source name prefix to a CRD group.
type CategoryRule struct {
	GroupPrefix string // e.g. "oncall"
	DirName     string // e.g. "oncall"
	TFPrefix    string // e.g. "grafana_oncall_"
}
