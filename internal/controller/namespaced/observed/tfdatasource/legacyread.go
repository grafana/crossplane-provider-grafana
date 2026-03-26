/*
Copyright 2025 Grafana
*/

package tfdatasource

import (
	"context"
	"fmt"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	grafanaProvider "github.com/grafana/terraform-provider-grafana/v4/pkg/provider"
	sdkschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	terraformSDK "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// NewLegacyReadFn builds a ReadFn from a TF SDK v2 data source name and two
// typed callbacks:
//   - toAttrs maps the Crossplane resource's ForProvider fields to the flat
//     attribute map that the TF data source expects as input (Required/Optional fields).
//   - fromData extracts computed output fields from the ResourceData after a
//     successful read and populates the Crossplane resource's AtProvider + external name.
func NewLegacyReadFn(
	dsName string,
	toAttrs func(resource.Managed) map[string]string,
	fromData func(resource.Managed, *sdkschema.ResourceData),
) ReadFn {
	// Resolve the data source schema once at setup time.
	p := grafanaProvider.Provider("crossplane")
	ds, ok := p.DataSourcesMap[dsName]
	if !ok {
		return func(_ context.Context, _ resource.Managed, _ any) error {
			return fmt.Errorf("data source %q not found in provider", dsName)
		}
	}

	return func(ctx context.Context, mg resource.Managed, providerMeta any) error {
		attrs := toAttrs(mg)

		// Apply schema defaults for attributes not provided by the caller.
		// Without this, the TF SDK uses zero values (e.g. 0 for TypeInt)
		// instead of the schema-defined defaults (e.g. -1 for user_id),
		// which can cause incorrect API calls.
		for name, field := range ds.Schema {
			if _, ok := attrs[name]; !ok && field.Default != nil {
				attrs[name] = fmt.Sprintf("%v", field.Default)
			}
		}

		state := &terraformSDK.InstanceState{Attributes: attrs}
		d := ds.Data(state)

		// providerMeta is the configured provider's Meta() — pass it as the
		// second argument to ReadContext, which expects the provider's client.
		if ds.ReadContext != nil {
			if diags := ds.ReadContext(ctx, d, providerMeta); diags.HasError() {
				return fmt.Errorf("data source %q read failed: %v", dsName, diags)
			}
		} else if ds.Read != nil {
			if err := ds.Read(d, providerMeta); err != nil {
				return fmt.Errorf("data source %q read failed: %w", dsName, err)
			}
		} else {
			return fmt.Errorf("data source %q has no Read or ReadContext function", dsName)
		}

		fromData(mg, d)
		return nil
	}
}
