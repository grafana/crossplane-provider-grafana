package multirefs

import (
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAdd(t *testing.T) {
	tests := map[string]struct {
		resource  *ujconfig.Resource
		fieldPath string
		jsonMap   map[string]any
		tfMap     map[string]any
		want      map[string]any
		wantErr   bool
	}{
		"TopLevelAlternative": {
			resource:  resourceWithTeamField(),
			fieldPath: "team",
			jsonMap:   map[string]any{"observedTeam": "42"},
			tfMap:     map[string]any{"observed_team": "42"},
			want:      map[string]any{"team": "42"},
		},
		"SecondAlternative": {
			resource:  resourceWithTeamField(),
			fieldPath: "team",
			jsonMap:   map[string]any{"importedTeam": "42"},
			tfMap:     map[string]any{"imported_team": "42"},
			want:      map[string]any{"team": "42"},
		},
		"OriginalField": {
			resource:  resourceWithTeamField(),
			fieldPath: "team",
			jsonMap:   map[string]any{"team": "42"},
			tfMap:     map[string]any{"team": "42"},
			want:      map[string]any{"team": "42"},
		},
		"ConflictingFields": {
			resource:  resourceWithTeamField(),
			fieldPath: "team",
			jsonMap:   map[string]any{"team": "1", "observedTeam": "2"},
			tfMap:     map[string]any{"team": "1", "observed_team": "2"},
			wantErr:   true,
		},
		"ConflictingAlternatives": {
			resource:  resourceWithTeamField(),
			fieldPath: "team",
			jsonMap:   map[string]any{"observedTeam": "1", "importedTeam": "2"},
			tfMap:     map[string]any{"observed_team": "1", "imported_team": "2"},
			wantErr:   true,
		},
		"NestedAlternative": {
			resource:  resourceWithNestedTeamField(),
			fieldPath: "permissions.team_id",
			jsonMap: map[string]any{
				"permissions": []any{map[string]any{"observedTeamId": "42"}},
			},
			tfMap: map[string]any{
				"permissions": []any{map[string]any{"observed_team_id": "42"}},
			},
			want: map[string]any{
				"permissions": []any{map[string]any{"team_id": "42"}},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			Add(tc.resource, tc.fieldPath,
				Alternative{
					Name:      "observed_" + lastPathSegment(tc.fieldPath),
					Reference: ujconfig.Reference{Type: "example.org/observed.Team"},
				},
				Alternative{
					Name:      "imported_" + lastPathSegment(tc.fieldPath),
					Reference: ujconfig.Reference{Type: "example.org/imported.Team"},
				},
			)

			if tc.resource.TerraformConfigurationInjector == nil {
				t.Fatal("TerraformConfigurationInjector was not configured")
			}
			err := tc.resource.TerraformConfigurationInjector(tc.jsonMap, tc.tfMap)
			if (err != nil) != tc.wantErr {
				t.Fatalf("TerraformConfigurationInjector() error = %v, wantErr %v", err, tc.wantErr)
			}
			if tc.wantErr {
				return
			}
			if diff := cmp.Diff(tc.want, tc.tfMap); diff != "" {
				t.Errorf("Terraform map mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func resourceWithTeamField() *ujconfig.Resource {
	return &ujconfig.Resource{
		TerraformResource: &schema.Resource{Schema: map[string]*schema.Schema{
			"team": {Type: schema.TypeString, Optional: true},
		}},
		References: ujconfig.References{
			"team": {TerraformName: "grafana_team"},
		},
	}
}

func resourceWithNestedTeamField() *ujconfig.Resource {
	return &ujconfig.Resource{
		TerraformResource: &schema.Resource{Schema: map[string]*schema.Schema{
			"permissions": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"team_id": {Type: schema.TypeString, Optional: true},
				}},
			},
		}},
		References: ujconfig.References{
			"permissions.team_id": {TerraformName: "grafana_team"},
		},
	}
}

func lastPathSegment(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '.' {
			return path[i+1:]
		}
	}
	return path
}
