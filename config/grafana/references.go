/*
Copyright 2026 Grafana Labs
*/

package grafana

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

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
