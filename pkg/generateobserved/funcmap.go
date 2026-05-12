/*
Copyright 2026 Grafana Labs
*/

package generateobserved

import (
	"embed"
	"fmt"
	"log"
	"strings"
	"text/template"
)

//go:embed templates/*.tmpl
var templateFS embed.FS

var funcMap = template.FuncMap{
	"nestedStructs":     nestedStructsFn,
	"fieldDoc":          fieldDocFn,
	"toAttrsLegacy":     toAttrsLegacyFn,
	"fromResourceData":  fromResourceDataFn,
	"toTFTypesValue":    toTFTypesValueFn,
	"fromState":         fromStateFn,
	"connectionDetailsFn": connectionDetailsFnTmpl,
	"exampleValue":      exampleValueFn,
}

var templates = template.Must(
	template.New("").Funcs(funcMap).ParseFS(templateFS, "templates/*.tmpl"),
)

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

		fmt.Fprintf(b, "type %s struct {\n", f.NestedStructName)
		for _, nf := range f.NestedFields {
			writeFieldDoc(b, nf)
			fmt.Fprintf(b, "\t%s %s `json:\"%s,omitempty\"`\n\n", nf.GoName, nf.GoType, nf.JSONName)
		}
		b.WriteString("}\n\n")
	}
}

func writeFieldDoc(b *strings.Builder, f fieldInfo) {
	if f.Description != "" {
		for _, line := range strings.Split(f.Description, "\n") {
			fmt.Fprintf(b, "\t// %s\n", line)
		}
	}
}

// fieldDocFn returns field documentation as a Go comment block.
func fieldDocFn(f fieldInfo) string {
	var b strings.Builder
	writeFieldDoc(&b, f)
	return b.String()
}

// toAttrsLegacyFn emits the code to convert a ForProvider field to a
// map[string]string entry for legacy SDK data sources.
func toAttrsLegacyFn(f fieldInfo) string {
	switch {
	case strings.HasPrefix(f.GoType, "*"):
		return fmt.Sprintf("\t\t\tif cr.Spec.ForProvider.%s != nil {\n\t\t\t\tattrs[%q] = fmt.Sprintf(\"%%v\", *cr.Spec.ForProvider.%s)\n\t\t\t}\n", f.GoName, f.TFName, f.GoName)
	case f.GoType == goTypeString:
		return fmt.Sprintf("\t\t\tattrs[%q] = cr.Spec.ForProvider.%s\n", f.TFName, f.GoName)
	default:
		return fmt.Sprintf("\t\t\tattrs[%q] = fmt.Sprintf(\"%%v\", cr.Spec.ForProvider.%s)\n", f.TFName, f.GoName)
	}
}

// fromResourceDataFn emits the code to extract a single field from legacy
// SDK ResourceData into the AtProvider observation struct.
func fromResourceDataFn(f fieldInfo) string {
	if len(f.NestedFields) > 0 {
		return nestedFromResourceDataFn(f)
	}
	switch f.GoType {
	case goTypePtrString:
		return fmt.Sprintf("\t\t\tif v, ok := d.GetOk(%q); ok {\n\t\t\t\tif s, ok := v.(string); ok {\n\t\t\t\t\tcr.Status.AtProvider.%s = &s\n\t\t\t\t}\n\t\t\t}\n", f.TFName, f.GoName)
	case goTypeString:
		return fmt.Sprintf("\t\t\tif v, ok := d.GetOk(%q); ok {\n\t\t\t\tif s, ok := v.(string); ok {\n\t\t\t\t\tcr.Status.AtProvider.%s = s\n\t\t\t\t}\n\t\t\t}\n", f.TFName, f.GoName)
	case goTypePtrInt64:
		return fmt.Sprintf("\t\t\tif v, ok := d.GetOk(%q); ok {\n\t\t\t\tif i, ok := v.(int); ok {\n\t\t\t\t\tv := int64(i)\n\t\t\t\t\tcr.Status.AtProvider.%s = &v\n\t\t\t\t}\n\t\t\t}\n", f.TFName, f.GoName)
	case goTypeInt64:
		return fmt.Sprintf("\t\t\tif v, ok := d.GetOk(%q); ok {\n\t\t\t\tif i, ok := v.(int); ok {\n\t\t\t\t\tcr.Status.AtProvider.%s = int64(i)\n\t\t\t\t}\n\t\t\t}\n", f.TFName, f.GoName)
	case goTypePtrFloat:
		return fmt.Sprintf("\t\t\tif v, ok := d.GetOk(%q); ok {\n\t\t\t\tif f, ok := v.(float64); ok {\n\t\t\t\t\tcr.Status.AtProvider.%s = &f\n\t\t\t\t}\n\t\t\t}\n", f.TFName, f.GoName)
	case goTypeFloat64:
		return fmt.Sprintf("\t\t\tif v, ok := d.GetOk(%q); ok {\n\t\t\t\tif f, ok := v.(float64); ok {\n\t\t\t\t\tcr.Status.AtProvider.%s = f\n\t\t\t\t}\n\t\t\t}\n", f.TFName, f.GoName)
	case goTypePtrBool:
		return fmt.Sprintf("\t\t\tif v, ok := d.GetOk(%q); ok {\n\t\t\t\tif b, ok := v.(bool); ok {\n\t\t\t\t\tcr.Status.AtProvider.%s = &b\n\t\t\t\t}\n\t\t\t}\n", f.TFName, f.GoName)
	case goTypeBool:
		return fmt.Sprintf("\t\t\tif v, ok := d.GetOk(%q); ok {\n\t\t\t\tif b, ok := v.(bool); ok {\n\t\t\t\t\tcr.Status.AtProvider.%s = b\n\t\t\t\t}\n\t\t\t}\n", f.TFName, f.GoName)
	default:
		return fmt.Sprintf("\t\t\t// TODO: complex type %s for %s\n", f.GoType, f.TFName)
	}
}

// nestedFromResourceDataFn emits the code to extract a nested (List/Set of
// objects) field from legacy SDK ResourceData.
func nestedFromResourceDataFn(f fieldInfo) string {
	hasScalarFields := false
	for _, nf := range f.NestedFields {
		if len(nf.NestedFields) == 0 && isScalarType(nf.GoType) {
			hasScalarFields = true
			break
		}
	}

	var b strings.Builder
	fmt.Fprintf(&b, "\t\t\tif v, ok := d.GetOk(%q); ok {\n", f.TFName)
	fmt.Fprintf(&b, "\t\t\t\tvar items []v1alpha1.%s\n", f.NestedStructName)
	b.WriteString("\t\t\t\tvar list []interface{}\n")
	if f.IsSet {
		b.WriteString("\t\t\t\tif s, ok := v.(*sdkschema.Set); ok {\n\t\t\t\t\tlist = s.List()\n\t\t\t\t}\n")
	} else {
		b.WriteString("\t\t\t\tlist, _ = v.([]interface{})\n")
	}
	b.WriteString("\t\t\t\tfor _, raw := range list {\n")
	fmt.Fprintf(&b, "\t\t\t\t\titem := v1alpha1.%s{}\n", f.NestedStructName)
	if hasScalarFields {
		b.WriteString("\t\t\t\t\tm, _ := raw.(map[string]interface{})\n")
		for _, nf := range f.NestedFields {
			if len(nf.NestedFields) > 0 || !isScalarType(nf.GoType) {
				continue
			}
			switch nf.GoType {
			case goTypeString:
				fmt.Fprintf(&b, "\t\t\t\t\tif val, ok := m[%q].(string); ok {\n\t\t\t\t\t\titem.%s = val\n\t\t\t\t\t}\n", nf.TFName, nf.GoName)
			case goTypePtrString:
				fmt.Fprintf(&b, "\t\t\t\t\tif val, ok := m[%q].(string); ok {\n\t\t\t\t\t\titem.%s = &val\n\t\t\t\t\t}\n", nf.TFName, nf.GoName)
			case goTypeInt64:
				fmt.Fprintf(&b, "\t\t\t\t\tif val, ok := m[%q].(int); ok {\n\t\t\t\t\t\tv := int64(val)\n\t\t\t\t\t\titem.%s = v\n\t\t\t\t\t}\n", nf.TFName, nf.GoName)
			case goTypePtrInt64:
				fmt.Fprintf(&b, "\t\t\t\t\tif val, ok := m[%q].(int); ok {\n\t\t\t\t\t\tv := int64(val)\n\t\t\t\t\t\titem.%s = &v\n\t\t\t\t\t}\n", nf.TFName, nf.GoName)
			case goTypeBool:
				fmt.Fprintf(&b, "\t\t\t\t\tif val, ok := m[%q].(bool); ok {\n\t\t\t\t\t\titem.%s = val\n\t\t\t\t\t}\n", nf.TFName, nf.GoName)
			case goTypePtrBool:
				fmt.Fprintf(&b, "\t\t\t\t\tif val, ok := m[%q].(bool); ok {\n\t\t\t\t\t\titem.%s = &val\n\t\t\t\t\t}\n", nf.TFName, nf.GoName)
			case goTypeFloat64:
				fmt.Fprintf(&b, "\t\t\t\t\tif val, ok := m[%q].(float64); ok {\n\t\t\t\t\t\titem.%s = val\n\t\t\t\t\t}\n", nf.TFName, nf.GoName)
			case goTypePtrFloat:
				fmt.Fprintf(&b, "\t\t\t\t\tif val, ok := m[%q].(float64); ok {\n\t\t\t\t\t\titem.%s = &val\n\t\t\t\t\t}\n", nf.TFName, nf.GoName)
			}
		}
	} else {
		b.WriteString("\t\t\t\t\t_ = raw\n")
	}
	b.WriteString("\t\t\t\t\titems = append(items, item)\n")
	b.WriteString("\t\t\t\t}\n")
	fmt.Fprintf(&b, "\t\t\t\tcr.Status.AtProvider.%s = items\n", f.GoName)
	b.WriteString("\t\t\t}\n")
	return b.String()
}

// toTFTypesValueFn emits the code to convert a ForProvider field to a
// tftypes.Value for Plugin Framework data sources.
func toTFTypesValueFn(f fieldInfo) string {
	switch f.GoType {
	case goTypeString:
		return fmt.Sprintf("\t\t\tattrs[%q] = tftypes.NewValue(tftypes.String, cr.Spec.ForProvider.%s)\n", f.TFName, f.GoName)
	case goTypePtrString:
		return fmt.Sprintf("\t\t\tif cr.Spec.ForProvider.%s != nil {\n\t\t\t\tattrs[%q] = tftypes.NewValue(tftypes.String, *cr.Spec.ForProvider.%s)\n\t\t\t}\n", f.GoName, f.TFName, f.GoName)
	case goTypeInt64:
		return fmt.Sprintf("\t\t\tattrs[%q] = tftypes.NewValue(tftypes.Number, cr.Spec.ForProvider.%s)\n", f.TFName, f.GoName)
	case goTypePtrInt64:
		return fmt.Sprintf("\t\t\tif cr.Spec.ForProvider.%s != nil {\n\t\t\t\tattrs[%q] = tftypes.NewValue(tftypes.Number, *cr.Spec.ForProvider.%s)\n\t\t\t}\n", f.GoName, f.TFName, f.GoName)
	case goTypeBool:
		return fmt.Sprintf("\t\t\tattrs[%q] = tftypes.NewValue(tftypes.Bool, cr.Spec.ForProvider.%s)\n", f.TFName, f.GoName)
	case goTypePtrBool:
		return fmt.Sprintf("\t\t\tif cr.Spec.ForProvider.%s != nil {\n\t\t\t\tattrs[%q] = tftypes.NewValue(tftypes.Bool, *cr.Spec.ForProvider.%s)\n\t\t\t}\n", f.GoName, f.TFName, f.GoName)
	case goTypeFloat64:
		return fmt.Sprintf("\t\t\tattrs[%q] = tftypes.NewValue(tftypes.Number, cr.Spec.ForProvider.%s)\n", f.TFName, f.GoName)
	case goTypePtrFloat:
		return fmt.Sprintf("\t\t\tif cr.Spec.ForProvider.%s != nil {\n\t\t\t\tattrs[%q] = tftypes.NewValue(tftypes.Number, *cr.Spec.ForProvider.%s)\n\t\t\t}\n", f.GoName, f.TFName, f.GoName)
	default:
		return fmt.Sprintf("\t\t\t// TODO: complex type %s for %s\n", f.GoType, f.TFName)
	}
}

// fromStateFn emits the code to extract a single field from Plugin Framework
// tfsdk.State into the AtProvider observation struct.
func fromStateFn(f fieldInfo) string {
	switch f.GoType {
	case goTypePtrString:
		return fmt.Sprintf("\t\t\t{\n\t\t\t\tvar v *string\n\t\t\t\tif diags := state.GetAttribute(ctx, path.Root(%q), &v); !diags.HasError() && v != nil {\n\t\t\t\t\tcr.Status.AtProvider.%s = v\n\t\t\t\t}\n\t\t\t}\n", f.TFName, f.GoName)
	case goTypeString:
		return fmt.Sprintf("\t\t\t{\n\t\t\t\tvar v string\n\t\t\t\tif diags := state.GetAttribute(ctx, path.Root(%q), &v); !diags.HasError() {\n\t\t\t\t\tcr.Status.AtProvider.%s = v\n\t\t\t\t}\n\t\t\t}\n", f.TFName, f.GoName)
	case goTypePtrInt64:
		return fmt.Sprintf("\t\t\t{\n\t\t\t\tvar v *int64\n\t\t\t\tif diags := state.GetAttribute(ctx, path.Root(%q), &v); !diags.HasError() && v != nil {\n\t\t\t\t\tcr.Status.AtProvider.%s = v\n\t\t\t\t}\n\t\t\t}\n", f.TFName, f.GoName)
	case goTypeInt64:
		return fmt.Sprintf("\t\t\t{\n\t\t\t\tvar v int64\n\t\t\t\tif diags := state.GetAttribute(ctx, path.Root(%q), &v); !diags.HasError() {\n\t\t\t\t\tcr.Status.AtProvider.%s = v\n\t\t\t\t}\n\t\t\t}\n", f.TFName, f.GoName)
	case goTypePtrBool:
		return fmt.Sprintf("\t\t\t{\n\t\t\t\tvar v *bool\n\t\t\t\tif diags := state.GetAttribute(ctx, path.Root(%q), &v); !diags.HasError() && v != nil {\n\t\t\t\t\tcr.Status.AtProvider.%s = v\n\t\t\t\t}\n\t\t\t}\n", f.TFName, f.GoName)
	case goTypeBool:
		return fmt.Sprintf("\t\t\t{\n\t\t\t\tvar v bool\n\t\t\t\tif diags := state.GetAttribute(ctx, path.Root(%q), &v); !diags.HasError() {\n\t\t\t\t\tcr.Status.AtProvider.%s = v\n\t\t\t\t}\n\t\t\t}\n", f.TFName, f.GoName)
	case goTypeSliceStr:
		return fmt.Sprintf("\t\t\t{\n\t\t\t\tvar v []string\n\t\t\t\tif diags := state.GetAttribute(ctx, path.Root(%q), &v); !diags.HasError() && len(v) > 0 {\n\t\t\t\t\tcr.Status.AtProvider.%s = v\n\t\t\t\t}\n\t\t\t}\n", f.TFName, f.GoName)
	default:
		return fmt.Sprintf("\t\t\t// TODO: complex type %s for %s\n", f.GoType, f.TFName)
	}
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

	var b strings.Builder
	b.WriteString("\tConnectionDetailsFn: func(mg resource.Managed) managed.ConnectionDetails {\n")
	fmt.Fprintf(&b, "\t\tcr := mg.(*v1alpha1.%s)\n", ds.KindName)
	b.WriteString("\t\tcd := managed.ConnectionDetails{}\n")
	b.WriteString("\t\tif id := meta.GetExternalName(cr); id != \"\" {\n")
	b.WriteString("\t\t\tcd[\"id\"] = []byte(id)\n")
	b.WriteString("\t\t}\n")
	for _, f := range scalars {
		switch f.GoType {
		case goTypePtrString:
			fmt.Fprintf(&b, "\t\tif cr.Status.AtProvider.%s != nil {\n", f.GoName)
			fmt.Fprintf(&b, "\t\t\tcd[%q] = []byte(*cr.Status.AtProvider.%s)\n", f.TFName, f.GoName)
			b.WriteString("\t\t}\n")
		case goTypeString:
			fmt.Fprintf(&b, "\t\tif cr.Status.AtProvider.%s != \"\" {\n", f.GoName)
			fmt.Fprintf(&b, "\t\t\tcd[%q] = []byte(cr.Status.AtProvider.%s)\n", f.TFName, f.GoName)
			b.WriteString("\t\t}\n")
		case goTypePtrInt64:
			fmt.Fprintf(&b, "\t\tif cr.Status.AtProvider.%s != nil {\n", f.GoName)
			fmt.Fprintf(&b, "\t\t\tcd[%q] = []byte(strconv.FormatInt(*cr.Status.AtProvider.%s, 10))\n", f.TFName, f.GoName)
			b.WriteString("\t\t}\n")
		case goTypeInt64:
			fmt.Fprintf(&b, "\t\tcd[%q] = []byte(strconv.FormatInt(cr.Status.AtProvider.%s, 10))\n", f.TFName, f.GoName)
		case goTypePtrFloat:
			fmt.Fprintf(&b, "\t\tif cr.Status.AtProvider.%s != nil {\n", f.GoName)
			fmt.Fprintf(&b, "\t\t\tcd[%q] = []byte(strconv.FormatFloat(*cr.Status.AtProvider.%s, 'f', -1, 64))\n", f.TFName, f.GoName)
			b.WriteString("\t\t}\n")
		case goTypeFloat64:
			fmt.Fprintf(&b, "\t\tcd[%q] = []byte(strconv.FormatFloat(cr.Status.AtProvider.%s, 'f', -1, 64))\n", f.TFName, f.GoName)
		case goTypePtrBool:
			fmt.Fprintf(&b, "\t\tif cr.Status.AtProvider.%s != nil {\n", f.GoName)
			fmt.Fprintf(&b, "\t\t\tcd[%q] = []byte(strconv.FormatBool(*cr.Status.AtProvider.%s))\n", f.TFName, f.GoName)
			b.WriteString("\t\t}\n")
		case goTypeBool:
			fmt.Fprintf(&b, "\t\tcd[%q] = []byte(strconv.FormatBool(cr.Status.AtProvider.%s))\n", f.TFName, f.GoName)
		}
	}
	b.WriteString("\t\treturn cd\n\t},\n")
	return b.String()
}

// exampleValueFn returns a placeholder value for an example YAML manifest field.
func exampleValueFn(f fieldInfo) string {
	switch f.GoType {
	case goTypeString, goTypePtrString:
		return fmt.Sprintf("\"example-%s\"", f.TFName)
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
