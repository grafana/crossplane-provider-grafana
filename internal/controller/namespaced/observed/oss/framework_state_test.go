/*
Copyright 2026 Grafana Labs
*/

package oss

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	v1alpha1 "github.com/grafana/crossplane-provider-grafana/v2/apis/observed/oss/v1alpha1"
)

func TestTeamSetTeamsStateDecode(t *testing.T) {
	t.Parallel()

	teamAttributeTypes := map[string]attr.Type{
		"email":        types.StringType,
		"id":           types.Int64Type,
		"member_count": types.Int64Type,
		"name":         types.StringType,
		"org_id":       types.Int64Type,
		"uid":          types.StringType,
	}
	teamTerraformType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{
		"email":        tftypes.String,
		"id":           tftypes.Number,
		"member_count": tftypes.Number,
		"name":         tftypes.String,
		"org_id":       tftypes.Number,
		"uid":          tftypes.String,
	}}
	teamsTerraformType := tftypes.List{ElementType: teamTerraformType}

	tests := map[string]struct {
		values []tftypes.Value
		want   []v1alpha1.TeamSetTeams
	}{
		"populated": {
			values: []tftypes.Value{
				tftypes.NewValue(teamTerraformType, map[string]tftypes.Value{
					"email":        tftypes.NewValue(tftypes.String, "platform@example.com"),
					"id":           tftypes.NewValue(tftypes.Number, int64(42)),
					"member_count": tftypes.NewValue(tftypes.Number, int64(7)),
					"name":         tftypes.NewValue(tftypes.String, "Platform"),
					"org_id":       tftypes.NewValue(tftypes.Number, int64(3)),
					"uid":          tftypes.NewValue(tftypes.String, "team-platform"),
				}),
			},
			want: []v1alpha1.TeamSetTeams{
				{
					Email:       pointerTo("platform@example.com"),
					ID:          pointerTo(int64(42)),
					MemberCount: pointerTo(int64(7)),
					Name:        pointerTo("Platform"),
					OrgID:       pointerTo(int64(3)),
					UID:         pointerTo("team-platform"),
				},
			},
		},
		"empty": {
			values: []tftypes.Value{},
			want:   []v1alpha1.TeamSetTeams{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			state := tfsdk.State{
				Raw: tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{"teams": teamsTerraformType},
				}, map[string]tftypes.Value{
					"teams": tftypes.NewValue(teamsTerraformType, tt.values),
				}),
				Schema: schema.Schema{Attributes: map[string]schema.Attribute{
					"teams": schema.ListAttribute{
						Computed:    true,
						ElementType: types.ObjectType{AttrTypes: teamAttributeTypes},
					},
				}},
			}

			var got []v1alpha1.TeamSetTeams
			if diags := state.GetAttribute(context.Background(), path.Root("teams"), &got); diags.HasError() {
				t.Fatalf("state.GetAttribute() diagnostics = %v", diags)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("state.GetAttribute() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func pointerTo[T any](value T) *T {
	return &value
}
