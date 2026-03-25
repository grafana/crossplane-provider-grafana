/*
Copyright 2025 Grafana
*/

package tfdatasource

import (
	"context"
	"fmt"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	fwschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// NewFrameworkReadFn builds a ReadFn from a Plugin Framework data source.
//
//   - newDS returns a fresh DataSource instance (must implement DataSourceWithConfigure).
//   - toAttrs maps the Crossplane resource's ForProvider fields to a map of
//     typed tftypes.Value entries keyed by attribute name.
//   - fromState extracts computed output fields from the tfsdk.State after a
//     successful read and populates the Crossplane resource's AtProvider + external name.
func NewFrameworkReadFn(
	newDS func() datasource.DataSourceWithConfigure,
	toAttrs func(resource.Managed) map[string]tftypes.Value,
	fromState func(resource.Managed, tfsdk.State),
) ReadFn {
	return func(ctx context.Context, mg resource.Managed, providerMeta any) error {
		ds := newDS()

		// Configure the data source with the provider client.
		var cfgResp datasource.ConfigureResponse
		ds.Configure(ctx, datasource.ConfigureRequest{ProviderData: providerMeta}, &cfgResp)
		if cfgResp.Diagnostics.HasError() {
			return fmt.Errorf("data source configure failed: %s", cfgResp.Diagnostics.Errors())
		}

		// Get the schema.
		var schemaResp datasource.SchemaResponse
		ds.Schema(ctx, datasource.SchemaRequest{}, &schemaResp)
		if schemaResp.Diagnostics.HasError() {
			return fmt.Errorf("data source schema failed: %s", schemaResp.Diagnostics.Errors())
		}

		// Build the config from typed input attrs.
		inputAttrs := toAttrs(mg)
		config, err := buildConfig(ctx, schemaResp.Schema, inputAttrs)
		if err != nil {
			return fmt.Errorf("cannot build tfsdk.Config: %w", err)
		}

		// Read.
		readResp := datasource.ReadResponse{
			State: tfsdk.State{Schema: schemaResp.Schema},
		}
		ds.Read(ctx, datasource.ReadRequest{Config: *config}, &readResp)
		if readResp.Diagnostics.HasError() {
			return fmt.Errorf("data source read failed: %s", readResp.Diagnostics.Errors())
		}

		fromState(mg, readResp.State)
		return nil
	}
}

// buildConfig constructs a tfsdk.Config from a schema and a set of input attribute values.
func buildConfig(ctx context.Context, s fwschema.Schema, inputAttrs map[string]tftypes.Value) (*tfsdk.Config, error) {
	// Build the tftypes.Object type from the schema attributes.
	attrTypes := make(map[string]tftypes.Type)
	for name, attr := range s.Attributes {
		attrTypes[name] = attr.GetType().TerraformType(ctx)
	}
	for name, block := range s.Blocks {
		attrTypes[name] = block.Type().TerraformType(ctx)
	}

	// Build the full value map, filling in nulls for attributes not in inputAttrs.
	vals := make(map[string]tftypes.Value)
	for name, t := range attrTypes {
		if v, ok := inputAttrs[name]; ok {
			vals[name] = v
		} else {
			vals[name] = tftypes.NewValue(t, nil) // null
		}
	}

	objType := tftypes.Object{AttributeTypes: attrTypes}
	objVal := tftypes.NewValue(objType, vals)

	return &tfsdk.Config{
		Schema: s,
		Raw:    objVal,
	}, nil
}
