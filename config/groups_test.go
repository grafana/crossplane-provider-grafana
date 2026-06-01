/*
Copyright 2026 Grafana Labs
*/

package config

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func TestIsEmptyResourceIDDiagnostic(t *testing.T) {
	cases := map[string]struct {
		diags []*tfprotov6.Diagnostic
		want  bool
	}{
		"empty int id": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Failed to parse resource ID",
				Detail:   `expected int for field "id", got ""`,
			}},
			want: true,
		},
		"empty composite id": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Unable to parse resource ID",
				Detail:   `id "" does not match expected format. Should be in the format: folderUID:type (role, team, or user):identifier`,
			}},
			want: true,
		},
		"wrapped empty id": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Detail:   `Failed to parse resource ID: expected int for field "id", got ""`,
			}},
			want: true,
		},
		"non-empty invalid id": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Failed to parse resource ID",
				Detail:   `expected int for field "id", got "abc"`,
			}},
			want: false,
		},
		"warning": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityWarning,
				Summary:  "Failed to parse resource ID",
				Detail:   `expected int for field "id", got ""`,
			}},
			want: false,
		},
		"nil diagnostic": {
			diags: []*tfprotov6.Diagnostic{nil},
			want:  false,
		},
		"string id hits list endpoint - folder": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Failed to get folder",
				Detail:   `json: cannot unmarshal array into Go value of type models.Folder`,
			}},
			want: true,
		},
		"string id hits list endpoint - message template": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Failed to read message template",
				Detail:   `json: cannot unmarshal array into Go value of type models.NotificationTemplate`,
			}},
			want: true,
		},
		"unmarshal array warning ignored": {
			diags: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityWarning,
				Summary:  "Failed to get folder",
				Detail:   `json: cannot unmarshal array into Go value of type models.Folder`,
			}},
			want: false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if got := isEmptyResourceIDDiagnostic(tc.diags); got != tc.want {
				t.Fatalf("got %t, want %t", got, tc.want)
			}
		})
	}
}
