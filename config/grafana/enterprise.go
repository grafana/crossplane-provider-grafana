/*
Copyright 2021 Upbound Inc.
*/

package grafana

import (
	"context"
	"errors"
	"fmt"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func configureEnterprise(p *ujconfig.Provider) {
	p.AddResourceConfigurator("grafana_data_source_permission", func(r *ujconfig.Resource) {
		r.References["datasource_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_data_source",
			RefFieldName:      "DataSourceRef",
			SelectorFieldName: "DataSourceSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
		r.References["permissions.team_id"] = ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRef",
			SelectorFieldName: "TeamSelector",
		}
		r.References["permissions.user_id"] = ujconfig.Reference{
			TerraformName:     "grafana_user",
			RefFieldName:      "UserRef",
			SelectorFieldName: "UserSelector",
		}
	})
	p.AddResourceConfigurator("grafana_report", func(r *ujconfig.Resource) {
		r.References["dashboard_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_dashboard",
			RefFieldName:      "DashboardRef",
			SelectorFieldName: "DashboardSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
	})
	p.AddResourceConfigurator("grafana_role", func(r *ujconfig.Resource) {
		r.InitializerFns = append(r.InitializerFns, createroleInitializer)
	})
	p.AddResourceConfigurator("grafana_role_assignment", func(r *ujconfig.Resource) {
		r.References["role_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_role",
			RefFieldName:      "RoleRef",
			SelectorFieldName: "RoleSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
		r.References["service_accounts"] = ujconfig.Reference{
			TerraformName:     "grafana_service_account",
			RefFieldName:      "ServiceAccountRefs",
			SelectorFieldName: "ServiceAccountSelector",
		}
		r.References["teams"] = ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRefs",
			SelectorFieldName: "TeamSelector",
		}
		r.References["users"] = ujconfig.Reference{
			TerraformName:     "grafana_user",
			RefFieldName:      "UserRefs",
			SelectorFieldName: "UserSelector",
		}
		r.ExternalName = ujconfig.ExternalName{
			SetIdentifierArgumentFn: ujconfig.NopSetIdentifierArgument,
			GetExternalNameFn: func(tfstate map[string]any) (string, error) {
				roleUID, ok := tfstate["role_uid"].(string)
				if !ok {
					return "", errors.New("cannot get role_uid attribute")
				}
				return roleUID, nil
			},
			GetIDFn:                ujconfig.ExternalNameAsID,
			DisableNameInitializer: true,
		}
	})
	p.AddResourceConfigurator("grafana_role_assignment_item", func(r *ujconfig.Resource) {
		r.References["role_uid"] = ujconfig.Reference{
			TerraformName:     "grafana_role",
			RefFieldName:      "RoleRef",
			SelectorFieldName: "RoleSelector",
			Extractor:         optionalFieldExtractor("uid"),
		}
		r.References["service_account_id"] = ujconfig.Reference{
			TerraformName:     "grafana_service_account",
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
		r.References["team_id"] = ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRef",
			SelectorFieldName: "TeamSelector",
		}
		r.References["user_id"] = ujconfig.Reference{
			TerraformName:     "grafana_user",
			RefFieldName:      "UserRef",
			SelectorFieldName: "UserSelector",
		}
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			// skip diff customization on create
			if state == nil || state.Empty() {
				return diff, nil
			}
			// skip no diff or destroy diffs
			if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
				return diff, nil
			}

			// ID is configured upon creation, don't try to update.
			if diff.Attributes["id"] != nil {
				delete(diff.Attributes, "id")
			}

			return diff, nil
		}
		r.ExternalName = ujconfig.ExternalName{
			SetIdentifierArgumentFn: ujconfig.NopSetIdentifierArgument,
			GetExternalNameFn: func(tfstate map[string]any) (string, error) {
				roleUID, ok := tfstate["role_uid"].(string)
				if !ok {
					return "", errors.New("cannot get role_uid attribute")
				}
				if serviceAccountID, ok := tfstate["service_account_id"].(string); ok {
					return fmt.Sprintf("%s:service_account:%s", roleUID, serviceAccountID), nil
				}
				if teamID, ok := tfstate["team_id"].(string); ok {
					return fmt.Sprintf("%s:team:%s", roleUID, teamID), nil
				}
				if userID, ok := tfstate["user_id"].(string); ok {
					return fmt.Sprintf("%s:user:%s", roleUID, userID), nil
				}
				return "", errors.New("cannot get either serviceAccountId, teamId or userId attribute")
			},
			GetIDFn: func(_ context.Context, externalName string, parameters map[string]any, _ map[string]any) (string, error) {
				roleUID, ok := parameters["role_uid"].(string)
				if !ok {
					return "", errors.New("cannot get role_uid attribute")
				}
				if serviceAccountID, ok := parameters["service_account_id"].(string); ok {
					return fmt.Sprintf("%s:service_account:%s", roleUID, serviceAccountID), nil
				}
				if teamID, ok := parameters["team_id"].(string); ok {
					return fmt.Sprintf("%s:team:%s", roleUID, teamID), nil
				}
				if userID, ok := parameters["user_id"].(string); ok {
					return fmt.Sprintf("%s:user:%s", roleUID, userID), nil
				}
				return "", errors.New("cannot get either serviceAccountId, teamId or userId attribute")
			},
			DisableNameInitializer: true,
		}
	})
	p.AddResourceConfigurator("grafana_team_external_group", func(r *ujconfig.Resource) {
		r.References["team_id"] = ujconfig.Reference{
			TerraformName:     "grafana_team",
			RefFieldName:      "TeamRef",
			SelectorFieldName: "TeamSelector",
		}
	})
}
