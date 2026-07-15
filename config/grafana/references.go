/*
Copyright 2026 Grafana Labs
*/

package grafana

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

const externalNameExtractor = `github.com/crossplane/crossplane-runtime/v2/pkg/reference.ExternalName()`

const ossUserType = "github.com/grafana/crossplane-provider-grafana/v2/apis/observed/oss/v1alpha1.User"

// observedRef returns a ujconfig.Reference for an observed (datasource) oncall resource.
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

// observedFieldRef returns a reference to an observed resource that extracts a
// specific field instead of its external name. OptionalFieldExtractor first
// reads the field from spec.forProvider and falls back to status.atProvider.
// This is useful when the target Terraform argument expects a value such as an
// email address or login rather than the data source resource ID.
func observedFieldRef(typePath, prefix, field string) ujconfig.Reference {
	r := observedRef(typePath, prefix)
	r.Extractor = optionalFieldExtractor(field)
	return r
}

// observedFieldRefs is the list-reference variant of observedFieldRef. It uses
// a plural RefFieldName while retaining the same field extraction behavior.
func observedFieldRefs(typePath, prefix, field string) ujconfig.Reference {
	r := observedFieldRef(typePath, prefix, field)
	r.RefFieldName = prefix + "Refs"
	return r
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
