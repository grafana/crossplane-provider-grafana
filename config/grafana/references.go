/*
Copyright 2026 Grafana Labs
*/

package grafana

import (
	"strings"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/grafana/crossplane-provider-grafana/v2/config/multirefs"
)

const externalNameExtractor = `github.com/crossplane/crossplane-runtime/v2/pkg/reference.ExternalName()`

const ossUserType = "github.com/grafana/crossplane-provider-grafana/v2/apis/observed/oss/v1alpha1.User"

const ossOrganizationUserType = "github.com/grafana/crossplane-provider-grafana/v2/apis/observed/oss/v1alpha1.OrganizationUser"

const ossTeamType = "github.com/grafana/crossplane-provider-grafana/v2/apis/observed/oss/v1alpha1.Team"

// observedRef returns a ujconfig.Reference for an observed (data source) resource.
// It uses Type directly (not TerraformName) because observed resources are not in the
// Terraform resource map. The ExternalName extractor reads crossplane.io/external-name
// which the observed controller populates with the resource ID.
func observedRef(typePath, prefix string) ujconfig.Reference {
	return ujconfig.Reference{
		Type:              typePath,
		RefFieldName:      prefix + "Ref",
		SelectorFieldName: prefix + "Selector",
		Extractor:         externalNameExtractor,
	}
}

func referenceWithFieldNames(r ujconfig.Reference, prefix string, plural bool) ujconfig.Reference {
	suffix := "Ref"
	if plural {
		suffix = "Refs"
	}
	r.RefFieldName = prefix + suffix
	r.SelectorFieldName = prefix + "Selector"
	return r
}

func addObservedReference(
	r *ujconfig.Resource,
	fieldPath string,
	managed ujconfig.Reference,
	syntheticName string,
	observed ujconfig.Reference,
) {
	r.References[fieldPath] = managed
	multirefs.Add(r, fieldPath, multirefs.Alternative{Name: syntheticName, Reference: observed})
}

// addUserReferences adds multiple references to lookup users (grafana_user, observed grafana_user, observed OrganizatioUser)
// Note that OSS Grafana supports grafana_user while Grafana Cloud does not support this API endpoint
// For Grafana Cloud, organization user can be used instead to observe an existing user
func addUserReferences(
	r *ujconfig.Resource,
	fieldPath string,
	prefix string,
	plural bool,
	managedExtractor string,
	observedExtractor string,
) {
	managed := referenceWithFieldNames(ujconfig.Reference{
		TerraformName: "grafana_user",
		Extractor:     managedExtractor,
	}, prefix, plural)
	observedUser := referenceWithFieldNames(ujconfig.Reference{
		Type:      ossUserType,
		Extractor: observedExtractor,
	}, "Observed"+prefix, plural)

	organizationUserPrefix := "ObservedOrganizationUser"
	if prefix != "User" {
		organizationUserPrefix += prefix
	}
	observedOrganizationUser := referenceWithFieldNames(ujconfig.Reference{
		Type:      ossOrganizationUserType,
		Extractor: observedExtractor,
	}, organizationUserPrefix, plural)

	leaf := fieldPath
	if i := strings.LastIndexByte(fieldPath, '.'); i >= 0 {
		leaf = fieldPath[i+1:]
	}
	observedUserName := "observed_" + leaf
	if !strings.HasPrefix(leaf, "user") {
		observedUserName = "observed_user_" + leaf
	}
	observedOrganizationUserName := "observed_organization_" + leaf
	if !strings.HasPrefix(leaf, "user") {
		observedOrganizationUserName = "observed_organization_user_" + leaf
	}

	r.References[fieldPath] = managed
	multirefs.Add(r, fieldPath,
		multirefs.Alternative{Name: observedUserName, Reference: observedUser},
		multirefs.Alternative{Name: observedOrganizationUserName, Reference: observedOrganizationUser},
	)
}

func dashboardReference(prefix string) ujconfig.Reference {
	return ujconfig.Reference{
		TerraformName:     "grafana_dashboard",
		RefFieldName:      prefix + "Ref",
		SelectorFieldName: prefix + "Selector",
		Extractor:         optionalFieldExtractor("uid"),
	}
}

func dataSourceReference(prefix string) ujconfig.Reference {
	return ujconfig.Reference{
		TerraformName:     "grafana_data_source",
		RefFieldName:      prefix + "Ref",
		SelectorFieldName: prefix + "Selector",
		Extractor:         optionalFieldExtractor("uid"),
	}
}

func folderReference() ujconfig.Reference {
	return ujconfig.Reference{
		TerraformName:     "grafana_folder",
		RefFieldName:      "FolderRef",
		SelectorFieldName: "FolderSelector",
		Extractor:         optionalFieldExtractor("uid"),
	}
}

func cloudStackIDReference() ujconfig.Reference {
	return ujconfig.Reference{
		TerraformName:     "grafana_cloud_stack",
		RefFieldName:      "CloudStackRef",
		SelectorFieldName: "CloudStackSelector",
		Extractor:         computedFieldExtractor("id"),
	}
}
