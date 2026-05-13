/*
Copyright 2026 Grafana Labs
*/

package generateobserved

import (
	"embed"
	"log"
	"strings"
	"text/template"
)

//go:embed templates/*.tmpl
var templateFS embed.FS

var templates *template.Template

func init() {
	funcMap := template.FuncMap{
		// Used by parent templates
		"nestedStructs":       nestedStructsFn,
		"fieldDoc":            fieldDocFn,
		"toAttrsLegacy":       toAttrsLegacyFn,
		"fromResourceData":    fromResourceDataFn,
		"toTFTypesValue":      toTFTypesValueFn,
		"fromState":           fromStateFn,
		"connectionDetailsFn": connectionDetailsFnTmpl,
		"exampleValue":        exampleValueFn,

		// Used by helper snippets in _helpers.go.tmpl
		"hasPrefix":             strings.HasPrefix,
		"splitLines":            splitLines,
		"isScalar":              isScalarType,
		"hasScalarNestedFields": hasScalarNestedFields,
	}
	templates = template.Must(
		template.New("").Funcs(funcMap).ParseFS(templateFS, "templates/*.tmpl"),
	)
}

// templateData holds all possible template context fields. Each template uses
// only the subset it needs.
type templateData struct {
	Cfg        Config
	DS         *dsInfo
	CI         CategoryRule
	DSList     []*dsInfo
	GroupNames []string
	Group      string // pre-computed group string (varies by template)
	APIImport  string // pre-computed import path for the API types package
}

// connectionDetailsData is passed to the _connectionDetails template.
type connectionDetailsData struct {
	KindName string
	Scalars  []fieldInfo
}

// nestedStructData is passed to the _nestedStruct template.
type nestedStructData struct {
	Name   string
	Fields []fieldInfo
}

func executeTemplate(name string, data any) string {
	var b strings.Builder
	if err := templates.ExecuteTemplate(&b, name, data); err != nil {
		log.Fatalf("template %s: %v", name, err)
	}
	return b.String()
}

// =============================================================================
// Template helper functions (FuncMap)
// =============================================================================

// splitLines splits a string into lines (used by _fieldDoc template).
func splitLines(s string) []string {
	return strings.Split(s, "\n")
}

// nestedStructsFn emits Go struct definitions for all nested field types in
// a data source, deduplicating across forProvider and atProvider fields.
func nestedStructsFn(ds *dsInfo) string {
	var b strings.Builder
	emitted := make(map[string]bool)
	emitNestedStructs(&b, ds.ForProviderFields, emitted)
	emitNestedStructs(&b, ds.AtProviderFields, emitted)
	return b.String()
}

// emitNestedStructs recursively emits Go struct definitions for fields that
// have nested sub-fields (e.g. TypeList/TypeSet with Resource elements).
func emitNestedStructs(b *strings.Builder, fields []fieldInfo, emitted map[string]bool) {
	for _, f := range fields {
		if len(f.NestedFields) == 0 || emitted[f.NestedStructName] {
			continue
		}
		emitted[f.NestedStructName] = true
		// Recurse first so inner types are defined before outer types.
		emitNestedStructs(b, f.NestedFields, emitted)

		b.WriteString(executeTemplate("_nestedStruct", nestedStructData{
			Name:   f.NestedStructName,
			Fields: f.NestedFields,
		}))
	}
}

// fieldDocFn returns field documentation as a Go comment block.
func fieldDocFn(f fieldInfo) string {
	return executeTemplate("_fieldDoc", f)
}

// toAttrsLegacyFn emits the code to convert a ForProvider field to a
// map[string]string entry for legacy SDK data sources.
func toAttrsLegacyFn(f fieldInfo) string {
	return executeTemplate("_toAttrsLegacy", f)
}

// fromResourceDataFn emits the code to extract a single field from legacy
// SDK ResourceData into the AtProvider observation struct.
func fromResourceDataFn(f fieldInfo) string {
	if len(f.NestedFields) > 0 {
		return executeTemplate("_nestedFromResourceData", f)
	}
	return executeTemplate("_fromResourceData", f)
}

// toTFTypesValueFn emits the code to convert a ForProvider field to a
// tftypes.Value for Plugin Framework data sources.
func toTFTypesValueFn(f fieldInfo) string {
	return executeTemplate("_toTFTypesValue", f)
}

// fromStateFn emits the code to extract a single field from Plugin Framework
// tfsdk.State into the AtProvider observation struct.
func fromStateFn(f fieldInfo) string {
	return executeTemplate("_fromState", f)
}

// connectionDetailsFnTmpl emits the ConnectionDetailsFn field for a
// tfdatasource.Spec struct literal, or an empty string if there are no
// scalar atProvider fields.
func connectionDetailsFnTmpl(ds *dsInfo) string {
	var scalars []fieldInfo
	for _, f := range ds.AtProviderFields {
		if isScalarType(f.GoType) {
			scalars = append(scalars, f)
		}
	}
	if len(scalars) == 0 {
		return ""
	}
	return executeTemplate("_connectionDetails", connectionDetailsData{
		KindName: ds.KindName,
		Scalars:  scalars,
	})
}

// exampleValueFn returns a placeholder value for an example YAML manifest field.
func exampleValueFn(f fieldInfo) string {
	switch f.GoType {
	case goTypeString, goTypePtrString:
		return "\"example-" + f.TFName + "\""
	case goTypeInt64, goTypePtrInt64:
		return "0"
	case goTypeFloat64, goTypePtrFloat:
		return "0.0"
	case goTypeBool, goTypePtrBool:
		return "false"
	default:
		return "# TODO"
	}
}

// =============================================================================
// Predicate helpers (used by template functions and emit logic)
// =============================================================================

func isScalarType(goType string) bool {
	switch goType {
	case goTypeString, goTypePtrString, goTypeInt64, goTypePtrInt64,
		goTypeFloat64, goTypePtrFloat, goTypeBool, goTypePtrBool:
		return true
	default:
		return false
	}
}

// hasScalarNestedFields returns true if any of the nested fields are scalar
// (non-nested, non-complex). Used by _nestedFromResourceData template.
func hasScalarNestedFields(fields []fieldInfo) bool {
	for _, nf := range fields {
		if len(nf.NestedFields) == 0 && isScalarType(nf.GoType) {
			return true
		}
	}
	return false
}
