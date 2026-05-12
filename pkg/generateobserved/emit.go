/*
Copyright 2026 Grafana Labs
*/

package generateobserved

import "fmt"

// Go type name constants used in code generation switch statements.
const (
	goTypeString    = "string"
	goTypePtrString = "*string"
	goTypeInt64     = "int64"
	goTypeSliceStr  = "[]string"
	goTypePtrInt64  = "*int64"
	goTypeFloat64   = "float64"
	goTypePtrFloat  = "*float64"
	goTypeBool      = "bool"
	goTypePtrBool   = "*bool"
)

// =============================================================================
// Code generation via templates
// =============================================================================

func generateGroupVersionInfo(cfg Config, ci CategoryRule) string {
	return executeTemplate("groupversion_info.go.tmpl", templateData{
		Cfg:   cfg,
		CI:    ci,
		Group: ci.GroupPrefix + cfg.GroupSuffix,
	})
}

func generateTypes(cfg Config, ds *dsInfo) string {
	return executeTemplate("types.go.tmpl", templateData{
		Cfg: cfg,
		DS:  ds,
	})
}

func generateSpec(cfg Config, ds *dsInfo, ci CategoryRule) string {
	apiImport := fmt.Sprintf("%s/apis/observed/%s/%s", cfg.ModulePath, ci.DirName, cfg.APIVersion)
	data := templateData{
		Cfg:       cfg,
		DS:        ds,
		CI:        ci,
		APIImport: apiImport,
	}
	if ds.IsLegacySDK {
		return executeTemplate("spec_legacy.go.tmpl", data)
	}
	return executeTemplate("spec_framework.go.tmpl", data)
}

func generateGroupSetup(cfg Config, dsList []*dsInfo, ci CategoryRule) string {
	return executeTemplate("group_setup.go.tmpl", templateData{
		Cfg:    cfg,
		DSList: dsList,
		CI:     ci,
	})
}

func generateRegister(cfg Config, groupNames []string) string {
	return executeTemplate("register.go.tmpl", templateData{
		Cfg:        cfg,
		GroupNames: groupNames,
	})
}

func generateTopSetup(cfg Config, groupNames []string) string {
	return executeTemplate("top_setup.go.tmpl", templateData{
		Cfg:        cfg,
		GroupNames: groupNames,
	})
}

func generateConnectFn(cfg Config) string {
	return executeTemplate("connect_fn.go.tmpl", templateData{
		Cfg: cfg,
	})
}

func generateLegacyFactories(cfg Config, dsList []*dsInfo, ci CategoryRule) string {
	return executeTemplate("legacy_factories.go.tmpl", templateData{
		Cfg:    cfg,
		DSList: dsList,
		CI:     ci,
	})
}

func generateFactories(cfg Config, dsList []*dsInfo, ci CategoryRule) string {
	return executeTemplate("factories.go.tmpl", templateData{
		Cfg:    cfg,
		DSList: dsList,
		CI:     ci,
	})
}

func generateExample(cfg Config, ds *dsInfo, ci CategoryRule) string {
	group := ci.GroupPrefix + cfg.GroupSuffix
	return executeTemplate("example.yaml.tmpl", templateData{
		Cfg:   cfg,
		DS:    ds,
		CI:    ci,
		Group: group + "/" + cfg.APIVersion,
	})
}

// =============================================================================
// Predicates used by emitFiles in generate.go
// =============================================================================

func hasFrameworkDS(dsList []*dsInfo) bool {
	for _, ds := range dsList {
		if !ds.IsLegacySDK {
			return true
		}
	}
	return false
}

func hasLegacyDS(dsList []*dsInfo) bool {
	for _, ds := range dsList {
		if ds.IsLegacySDK {
			return true
		}
	}
	return false
}
