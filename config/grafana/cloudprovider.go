/*
Copyright 2026 Grafana Labs
*/

package grafana

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

func configureCloudProvider(p *ujconfig.Provider) {
	for _, name := range []string{
		"grafana_cloud_provider_aws_account",
		"grafana_cloud_provider_aws_cloudwatch_scrape_job",
		"grafana_cloud_provider_aws_resource_metadata_scrape_job",
		"grafana_cloud_provider_azure_credential",
		"grafana_connections_metrics_endpoint_scrape_job",
		"grafana_frontend_o11y_app",
	} {
		p.AddResourceConfigurator(name, func(r *ujconfig.Resource) {
			r.References["stack_id"] = cloudStackIDReference()
		})
	}

	awsAccountRef := ujconfig.Reference{
		TerraformName:     "grafana_cloud_provider_aws_account",
		RefFieldName:      "AwsAccountRef",
		SelectorFieldName: "AwsAccountSelector",
		Extractor:         computedFieldExtractor("resourceId"),
	}
	for _, name := range []string{
		"grafana_cloud_provider_aws_cloudwatch_scrape_job",
		"grafana_cloud_provider_aws_resource_metadata_scrape_job",
	} {
		p.AddResourceConfigurator(name, func(r *ujconfig.Resource) {
			r.References["aws_account_resource_id"] = awsAccountRef
		})
	}
}
