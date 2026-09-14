/*
Copyright 2026 Grafana Labs
*/

package generateobserved

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestBuildFrameworkObjectCollectionFieldInfo(t *testing.T) {
	t.Parallel()

	objectType := types.ObjectType{AttrTypes: map[string]attr.Type{
		"count":   types.Int32Type,
		"labels":  types.MapType{ElemType: types.StringType},
		"aliases": types.SetType{ElemType: types.StringType},
		"schedule": types.ObjectType{AttrTypes: map[string]attr.Type{
			"active": types.BoolType,
		}},
		"history": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
			"id": types.Int64Type,
		}}},
	}}

	tests := map[string]struct {
		attribute schema.Attribute
		wantGo    string
		wantState string
		wantSet   bool
	}{
		"list": {
			attribute: schema.ListAttribute{Computed: true, ElementType: objectType},
			wantGo:    "[]ExampleItems",
			wantState: "[]v1alpha1.ExampleItems",
		},
		"set": {
			attribute: schema.SetAttribute{Computed: true, ElementType: objectType},
			wantGo:    "[]ExampleItems",
			wantState: "[]v1alpha1.ExampleItems",
			wantSet:   true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := buildFrameworkFieldInfo(nil, "Example", "items", tt.attribute)
			if err != nil {
				t.Fatalf("buildFrameworkFieldInfo() error = %v", err)
			}
			if got.GoType != tt.wantGo || got.StateGoType != tt.wantState || got.IsSet != tt.wantSet {
				t.Fatalf("buildFrameworkFieldInfo() type = (%q, %q, set %t), want (%q, %q, set %t)", got.GoType, got.StateGoType, got.IsSet, tt.wantGo, tt.wantState, tt.wantSet)
			}
			if !got.HasObject || !got.UseTypedStateDecode || got.NestedStructName != "ExampleItems" {
				t.Fatalf("buildFrameworkFieldInfo() object metadata = (object %t, typed decode %t, name %q)", got.HasObject, got.UseTypedStateDecode, got.NestedStructName)
			}

			wantFields := []fieldInfo{
				{TFName: "aliases", GoName: "Aliases", JSONName: "aliases", GoType: "[]string", StateGoType: "[]string", IsSet: true},
				{TFName: "count", GoName: "Count", JSONName: "count", GoType: "*int32", StateGoType: "*int32"},
				{TFName: "history", GoName: "History", JSONName: "history", GoType: "[]ExampleItemsHistory", StateGoType: "[]v1alpha1.ExampleItemsHistory", NestedStructName: "ExampleItemsHistory", HasObject: true, NestedFields: []fieldInfo{
					{TFName: "id", GoName: "Id", JSONName: "id", GoType: "*int64", StateGoType: "*int64"},
				}},
				{TFName: "labels", GoName: "Labels", JSONName: "labels", GoType: "map[string]string", StateGoType: "map[string]string"},
				{TFName: "schedule", GoName: "Schedule", JSONName: "schedule", GoType: "*ExampleItemsSchedule", StateGoType: "*v1alpha1.ExampleItemsSchedule", NestedStructName: "ExampleItemsSchedule", HasObject: true, NestedFields: []fieldInfo{
					{TFName: "active", GoName: "Active", JSONName: "active", GoType: "*bool", StateGoType: "*bool"},
				}},
			}
			if diff := cmp.Diff(wantFields, got.NestedFields); diff != "" {
				t.Errorf("nested fields mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestBuildFrameworkScalarCollectionFieldInfo(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		attribute schema.Attribute
		wantGo    string
	}{
		"list string": {
			attribute: schema.ListAttribute{Computed: true, ElementType: types.StringType},
			wantGo:    "[]string",
		},
		"set int32": {
			attribute: schema.SetAttribute{Computed: true, ElementType: types.Int32Type},
			wantGo:    "[]string",
		},
		"map int64": {
			attribute: schema.MapAttribute{Computed: true, ElementType: types.Int64Type},
			wantGo:    "map[string]string",
		},
		"nested map list": {
			attribute: schema.MapAttribute{Computed: true, ElementType: types.ListType{ElemType: types.BoolType}},
			wantGo:    "map[string]string",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := buildFrameworkFieldInfo(nil, "Example", "values", tt.attribute)
			if err != nil {
				t.Fatalf("buildFrameworkFieldInfo() error = %v", err)
			}
			if got.GoType != tt.wantGo || got.StateGoType != tt.wantGo {
				t.Errorf("buildFrameworkFieldInfo() = (%q, %q), want (%q, %q)", got.GoType, got.StateGoType, tt.wantGo, tt.wantGo)
			}
			if got.UseTypedStateDecode {
				t.Error("scalar-only framework collection unexpectedly enabled typed state decoding")
			}
		})
	}
}

func TestFrameworkObjectCollectionGeneration(t *testing.T) {
	t.Parallel()

	field, err := buildFrameworkFieldInfo(nil, "Example", "items", schema.ListAttribute{
		Computed: true,
		ElementType: types.ObjectType{AttrTypes: map[string]attr.Type{
			"count": types.Int32Type,
			"name":  types.StringType,
		}},
	})
	if err != nil {
		t.Fatalf("buildFrameworkFieldInfo() error = %v", err)
	}
	ds := &dsInfo{KindName: "Example", TFName: "example", AtProviderFields: []fieldInfo{field}}

	typesOutput := generateTypes(Config{APIVersion: "v1alpha1"}, ds)
	for _, want := range []string{
		"type ExampleItems struct {",
		"Count *int32 `json:\"count,omitempty\" tfsdk:\"count\"`",
		"Name *string `json:\"name,omitempty\" tfsdk:\"name\"`",
		"Items []ExampleItems `json:\"items,omitempty\"`",
	} {
		if !strings.Contains(typesOutput, want) {
			t.Errorf("generated types missing %q:\n%s", want, typesOutput)
		}
	}

	specOutput := generateSpec(Config{ModulePath: "example.test/provider", APIVersion: "v1alpha1", TFDataSourcePkg: "example.test/tfdatasource"}, ds, CategoryRule{DirName: "test"})
	for _, want := range []string{
		"func(ctx context.Context, mg resource.Managed, state tfsdk.State)",
		"var v []v1alpha1.ExampleItems",
		"state.GetAttribute(ctx, path.Root(\"items\"), &v)",
		"cr.Status.AtProvider.Items = v",
	} {
		if !strings.Contains(specOutput, want) {
			t.Errorf("generated spec missing %q:\n%s", want, specOutput)
		}
	}
	if strings.Contains(specOutput, "len(v) > 0") {
		t.Errorf("generated decode does not clear an empty collection:\n%s", specOutput)
	}
	if !strings.Contains(specOutput, "cr := mg.(*v1alpha1.Example)\n\t\t\t{") {
		t.Errorf("generated typed decode introduced a blank line after the resource assertion:\n%s", specOutput)
	}
}

func TestFrameworkScalarStateGenerationIsUnchanged(t *testing.T) {
	t.Parallel()

	field, err := buildFrameworkFieldInfo(nil, "Example", "name", schema.StringAttribute{Computed: true})
	if err != nil {
		t.Fatalf("buildFrameworkFieldInfo() error = %v", err)
	}
	if field.UseTypedStateDecode {
		t.Fatal("scalar field unexpectedly enabled typed state decoding")
	}

	output := fromStateFn(field)
	for _, want := range []string{
		"var v *string",
		"state.GetAttribute(ctx, path.Root(\"name\"), &v)",
		"!diags.HasError() && v != nil",
	} {
		if !strings.Contains(output, want) {
			t.Errorf("generated scalar state decoding missing %q:\n%s", want, output)
		}
	}
}

func TestFrameworkScalarCollectionStateGenerationIsUnchanged(t *testing.T) {
	t.Parallel()

	field, err := buildFrameworkFieldInfo(nil, "Example", "values", schema.SetAttribute{Computed: true, ElementType: types.Int32Type})
	if err != nil {
		t.Fatalf("buildFrameworkFieldInfo() error = %v", err)
	}
	output := fromStateFn(field)
	for _, want := range []string{
		"var v []string",
		"!diags.HasError() && len(v) > 0",
	} {
		if !strings.Contains(output, want) {
			t.Errorf("generated scalar collection state decoding missing %q:\n%s", want, output)
		}
	}
}

func TestFrameworkTopLevelInt32FallbackIsUnchanged(t *testing.T) {
	t.Parallel()

	field, err := buildFrameworkFieldInfo(nil, "Example", "count", schema.Int32Attribute{Computed: true})
	if err != nil {
		t.Fatalf("buildFrameworkFieldInfo() error = %v", err)
	}
	if field.GoType != "*string" || field.StateGoType != "*string" || field.UseTypedStateDecode {
		t.Errorf("top-level Int32 field = (%q, %q, typed %t), want (*string, *string, typed false)", field.GoType, field.StateGoType, field.UseTypedStateDecode)
	}
}

func TestNestedStructTFSDKTags(t *testing.T) {
	t.Parallel()

	field := fieldInfo{
		TFName:           "items",
		GoName:           "Items",
		JSONName:         "items",
		GoType:           "[]ExampleItems",
		NestedStructName: "ExampleItems",
		NestedFields: []fieldInfo{
			{TFName: "name", GoName: "Name", JSONName: "name", GoType: "string"},
		},
	}

	tests := map[string]struct {
		ds      *dsInfo
		wantTag bool
	}{
		"framework": {ds: &dsInfo{KindName: "Example", AtProviderFields: []fieldInfo{field}}, wantTag: true},
		"legacy":    {ds: &dsInfo{KindName: "Example", IsLegacySDK: true, AtProviderFields: []fieldInfo{field}}},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			output := generateTypes(Config{APIVersion: "v1alpha1"}, tt.ds)
			hasTag := strings.Contains(output, `tfsdk:"name"`)
			if hasTag != tt.wantTag {
				t.Errorf("generated nested struct tfsdk tag = %t, want %t:\n%s", hasTag, tt.wantTag, output)
			}
		})
	}
}

func TestFrameworkScalarCollectionInputGenerationIsUnchanged(t *testing.T) {
	t.Parallel()

	field, err := buildFrameworkFieldInfo(nil, "Example", "values", schema.ListAttribute{
		Optional:    true,
		ElementType: types.Int32Type,
	})
	if err != nil {
		t.Fatalf("buildFrameworkFieldInfo() error = %v", err)
	}
	ds := &dsInfo{KindName: "Example", TFName: "example", ForProviderFields: []fieldInfo{field}}
	output := generateSpec(Config{ModulePath: "example.test/provider", APIVersion: "v1alpha1", TFDataSourcePkg: "example.test/tfdatasource"}, ds, CategoryRule{DirName: "test"})

	if want := "// TODO: complex type []string for values"; !strings.Contains(output, want) {
		t.Errorf("generated input serialization missing %q:\n%s", want, output)
	}
}

func TestValidateFrameworkInput(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		field   fieldInfo
		wantErr bool
	}{
		"scalar collection": {field: fieldInfo{GoType: "[]string"}},
		"object collection": {field: fieldInfo{GoType: "[]ExampleItems", HasObject: true}, wantErr: true},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			err := validateFrameworkInput(tt.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateFrameworkInput() error = %v, wantErr %t", err, tt.wantErr)
			}
		})
	}
}
