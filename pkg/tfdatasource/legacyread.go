/*
Copyright 2026 Grafana Labs
*/

package tfdatasource

import (
	"context"
	"fmt"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	sdkschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	terraformSDK "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// NewLegacyReadFn builds a ReadFn from a TF SDK v2 data source name and two
// typed callbacks:
//   - toAttrs maps the Crossplane resource's ForProvider fields to the flat
//     attribute map that the TF data source expects as input (Required/Optional fields).
//   - fromData extracts computed output fields from the ResourceData after a
//     successful read and populates the Crossplane resource's AtProvider + external name.
//
// The caller must resolve the *sdkschema.Resource from the provider at init
// time and pass it directly.
func NewLegacyReadFn(
	dsName string,
	ds *sdkschema.Resource,
	toAttrs func(resource.Managed) map[string]string,
	fromData func(resource.Managed, *sdkschema.ResourceData),
) ReadFn {
	if ds == nil {
		return func(_ context.Context, _ resource.Managed, _ any) error {
			return fmt.Errorf("data source %q not found in provider (nil schema)", dsName)
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
		switch {
		case ds.ReadContext != nil:
			if diags := ds.ReadContext(ctx, d, providerMeta); diags.HasError() {
				return fmt.Errorf("data source %q read failed: %v", dsName, diags)
			}
		case ds.Read != nil: //nolint:staticcheck // ds.Read is the fallback for older TF providers that haven't migrated to ReadContext.
			if err := ds.Read(d, providerMeta); err != nil { //nolint:staticcheck
				return fmt.Errorf("data source %q read failed: %w", dsName, err)
			}
		default:
			return fmt.Errorf("data source %q has no Read or ReadContext function", dsName)
		}

		fromData(mg, d)
		return nil
	}
}
