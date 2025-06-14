/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AwsResourceMetadataScrapeJobInitParameters struct {

	// (String) The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this Resource Metadata Scrape Job. This can be provided by the resource_id attribute of the grafana_cloud_provider_aws_account resource.
	// The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this Resource Metadata Scrape Job. This can be provided by the `resource_id` attribute of the `grafana_cloud_provider_aws_account` resource.
	AwsAccountResourceID *string `json:"awsAccountResourceId,omitempty" tf:"aws_account_resource_id,omitempty"`

	// (Boolean) Whether the AWS Resource Metadata Scrape Job is enabled or not. Defaults to true.
	// Whether the AWS Resource Metadata Scrape Job is enabled or not. Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The name of the AWS Resource Metadata Scrape Job.
	// The name of the AWS Resource Metadata Scrape Job.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
	// A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
	// +listType=set
	RegionsSubsetOverride []*string `json:"regionsSubsetOverride,omitempty" tf:"regions_subset_override,omitempty"`

	// (Block List) One or more configuration blocks to configure AWS services for the Resource Metadata Scrape Job to scrape. Each block must have a distinct name attribute. When accessing this as an attribute reference, it is a list of objects. (see below for nested schema)
	// One or more configuration blocks to configure AWS services for the Resource Metadata Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
	Service []AwsResourceMetadataScrapeJobServiceInitParameters `json:"service,omitempty" tf:"service,omitempty"`

	// (String) The Stack ID of the Grafana Cloud instance.
	// The Stack ID of the Grafana Cloud instance.
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// (Map of String) A set of static labels to add to all metrics exported by this scrape job.
	// A set of static labels to add to all metrics exported by this scrape job.
	// +mapType=granular
	StaticLabels map[string]*string `json:"staticLabels,omitempty" tf:"static_labels,omitempty"`
}

type AwsResourceMetadataScrapeJobObservation struct {

	// (String) The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this Resource Metadata Scrape Job. This can be provided by the resource_id attribute of the grafana_cloud_provider_aws_account resource.
	// The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this Resource Metadata Scrape Job. This can be provided by the `resource_id` attribute of the `grafana_cloud_provider_aws_account` resource.
	AwsAccountResourceID *string `json:"awsAccountResourceId,omitempty" tf:"aws_account_resource_id,omitempty"`

	// (String) When the AWS Resource Metadata Scrape Job is disabled, this will show the reason that it is in that state.
	// When the AWS Resource Metadata Scrape Job is disabled, this will show the reason that it is in that state.
	DisabledReason *string `json:"disabledReason,omitempty" tf:"disabled_reason,omitempty"`

	// (Boolean) Whether the AWS Resource Metadata Scrape Job is enabled or not. Defaults to true.
	// Whether the AWS Resource Metadata Scrape Job is enabled or not. Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// This has the format "{{ stack_id }}:{{ name }}".
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the AWS Resource Metadata Scrape Job.
	// The name of the AWS Resource Metadata Scrape Job.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
	// A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
	// +listType=set
	RegionsSubsetOverride []*string `json:"regionsSubsetOverride,omitempty" tf:"regions_subset_override,omitempty"`

	// (Block List) One or more configuration blocks to configure AWS services for the Resource Metadata Scrape Job to scrape. Each block must have a distinct name attribute. When accessing this as an attribute reference, it is a list of objects. (see below for nested schema)
	// One or more configuration blocks to configure AWS services for the Resource Metadata Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
	Service []AwsResourceMetadataScrapeJobServiceObservation `json:"service,omitempty" tf:"service,omitempty"`

	// (String) The Stack ID of the Grafana Cloud instance.
	// The Stack ID of the Grafana Cloud instance.
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// (Map of String) A set of static labels to add to all metrics exported by this scrape job.
	// A set of static labels to add to all metrics exported by this scrape job.
	// +mapType=granular
	StaticLabels map[string]*string `json:"staticLabels,omitempty" tf:"static_labels,omitempty"`
}

type AwsResourceMetadataScrapeJobParameters struct {

	// (String) The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this Resource Metadata Scrape Job. This can be provided by the resource_id attribute of the grafana_cloud_provider_aws_account resource.
	// The ID assigned by the Grafana Cloud Provider API to an AWS Account resource that should be associated with this Resource Metadata Scrape Job. This can be provided by the `resource_id` attribute of the `grafana_cloud_provider_aws_account` resource.
	// +kubebuilder:validation:Optional
	AwsAccountResourceID *string `json:"awsAccountResourceId,omitempty" tf:"aws_account_resource_id,omitempty"`

	// (Boolean) Whether the AWS Resource Metadata Scrape Job is enabled or not. Defaults to true.
	// Whether the AWS Resource Metadata Scrape Job is enabled or not. Defaults to `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The name of the AWS Resource Metadata Scrape Job.
	// The name of the AWS Resource Metadata Scrape Job.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
	// A subset of the regions that are configured in the associated AWS Account resource to apply to this scrape job. If not set or empty, all of the Account resource's regions are scraped.
	// +kubebuilder:validation:Optional
	// +listType=set
	RegionsSubsetOverride []*string `json:"regionsSubsetOverride,omitempty" tf:"regions_subset_override,omitempty"`

	// (Block List) One or more configuration blocks to configure AWS services for the Resource Metadata Scrape Job to scrape. Each block must have a distinct name attribute. When accessing this as an attribute reference, it is a list of objects. (see below for nested schema)
	// One or more configuration blocks to configure AWS services for the Resource Metadata Scrape Job to scrape. Each block must have a distinct `name` attribute. When accessing this as an attribute reference, it is a list of objects.
	// +kubebuilder:validation:Optional
	Service []AwsResourceMetadataScrapeJobServiceParameters `json:"service,omitempty" tf:"service,omitempty"`

	// (String) The Stack ID of the Grafana Cloud instance.
	// The Stack ID of the Grafana Cloud instance.
	// +kubebuilder:validation:Optional
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// (Map of String) A set of static labels to add to all metrics exported by this scrape job.
	// A set of static labels to add to all metrics exported by this scrape job.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	StaticLabels map[string]*string `json:"staticLabels,omitempty" tf:"static_labels,omitempty"`
}

type AwsResourceMetadataScrapeJobServiceInitParameters struct {

	// (String) The name of the AWS Resource Metadata Scrape Job.
	// The name of the service to scrape. See https://grafana.com/docs/grafana-cloud/monitor-infrastructure/monitor-cloud-provider/aws/cloudwatch-metrics/services/ for supported services.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List) One or more configuration blocks to configure tag filters applied to discovery of resource entities in the associated AWS account. When accessing this as an attribute reference, it is a list of objects. (see below for nested schema)
	// One or more configuration blocks to configure tag filters applied to discovery of resource entities in the associated AWS account. When accessing this as an attribute reference, it is a list of objects.
	ResourceDiscoveryTagFilter []ServiceResourceDiscoveryTagFilterInitParameters `json:"resourceDiscoveryTagFilter,omitempty" tf:"resource_discovery_tag_filter,omitempty"`

	// cloud/monitor-infrastructure/monitor-cloud-provider/aws/cloudwatch-metrics/services/ for supported scrape intervals. Defaults to 300.
	// The interval in seconds to scrape the service. See https://grafana.com/docs/grafana-cloud/monitor-infrastructure/monitor-cloud-provider/aws/cloudwatch-metrics/services/ for supported scrape intervals. Defaults to `300`.
	ScrapeIntervalSeconds *float64 `json:"scrapeIntervalSeconds,omitempty" tf:"scrape_interval_seconds,omitempty"`
}

type AwsResourceMetadataScrapeJobServiceObservation struct {

	// (String) The name of the AWS Resource Metadata Scrape Job.
	// The name of the service to scrape. See https://grafana.com/docs/grafana-cloud/monitor-infrastructure/monitor-cloud-provider/aws/cloudwatch-metrics/services/ for supported services.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List) One or more configuration blocks to configure tag filters applied to discovery of resource entities in the associated AWS account. When accessing this as an attribute reference, it is a list of objects. (see below for nested schema)
	// One or more configuration blocks to configure tag filters applied to discovery of resource entities in the associated AWS account. When accessing this as an attribute reference, it is a list of objects.
	ResourceDiscoveryTagFilter []ServiceResourceDiscoveryTagFilterObservation `json:"resourceDiscoveryTagFilter,omitempty" tf:"resource_discovery_tag_filter,omitempty"`

	// cloud/monitor-infrastructure/monitor-cloud-provider/aws/cloudwatch-metrics/services/ for supported scrape intervals. Defaults to 300.
	// The interval in seconds to scrape the service. See https://grafana.com/docs/grafana-cloud/monitor-infrastructure/monitor-cloud-provider/aws/cloudwatch-metrics/services/ for supported scrape intervals. Defaults to `300`.
	ScrapeIntervalSeconds *float64 `json:"scrapeIntervalSeconds,omitempty" tf:"scrape_interval_seconds,omitempty"`
}

type AwsResourceMetadataScrapeJobServiceParameters struct {

	// (String) The name of the AWS Resource Metadata Scrape Job.
	// The name of the service to scrape. See https://grafana.com/docs/grafana-cloud/monitor-infrastructure/monitor-cloud-provider/aws/cloudwatch-metrics/services/ for supported services.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (Block List) One or more configuration blocks to configure tag filters applied to discovery of resource entities in the associated AWS account. When accessing this as an attribute reference, it is a list of objects. (see below for nested schema)
	// One or more configuration blocks to configure tag filters applied to discovery of resource entities in the associated AWS account. When accessing this as an attribute reference, it is a list of objects.
	// +kubebuilder:validation:Optional
	ResourceDiscoveryTagFilter []ServiceResourceDiscoveryTagFilterParameters `json:"resourceDiscoveryTagFilter,omitempty" tf:"resource_discovery_tag_filter,omitempty"`

	// cloud/monitor-infrastructure/monitor-cloud-provider/aws/cloudwatch-metrics/services/ for supported scrape intervals. Defaults to 300.
	// The interval in seconds to scrape the service. See https://grafana.com/docs/grafana-cloud/monitor-infrastructure/monitor-cloud-provider/aws/cloudwatch-metrics/services/ for supported scrape intervals. Defaults to `300`.
	// +kubebuilder:validation:Optional
	ScrapeIntervalSeconds *float64 `json:"scrapeIntervalSeconds,omitempty" tf:"scrape_interval_seconds,omitempty"`
}

type ServiceResourceDiscoveryTagFilterInitParameters struct {

	// (String) The key of the tag filter.
	// The key of the tag filter.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String) The value of the tag filter.
	// The value of the tag filter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ServiceResourceDiscoveryTagFilterObservation struct {

	// (String) The key of the tag filter.
	// The key of the tag filter.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String) The value of the tag filter.
	// The value of the tag filter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ServiceResourceDiscoveryTagFilterParameters struct {

	// (String) The key of the tag filter.
	// The key of the tag filter.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// (String) The value of the tag filter.
	// The value of the tag filter.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// AwsResourceMetadataScrapeJobSpec defines the desired state of AwsResourceMetadataScrapeJob
type AwsResourceMetadataScrapeJobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AwsResourceMetadataScrapeJobParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AwsResourceMetadataScrapeJobInitParameters `json:"initProvider,omitempty"`
}

// AwsResourceMetadataScrapeJobStatus defines the observed state of AwsResourceMetadataScrapeJob.
type AwsResourceMetadataScrapeJobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AwsResourceMetadataScrapeJobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AwsResourceMetadataScrapeJob is the Schema for the AwsResourceMetadataScrapeJobs API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type AwsResourceMetadataScrapeJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.awsAccountResourceId) || (has(self.initProvider) && has(self.initProvider.awsAccountResourceId))",message="spec.forProvider.awsAccountResourceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.stackId) || (has(self.initProvider) && has(self.initProvider.stackId))",message="spec.forProvider.stackId is a required parameter"
	Spec   AwsResourceMetadataScrapeJobSpec   `json:"spec"`
	Status AwsResourceMetadataScrapeJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AwsResourceMetadataScrapeJobList contains a list of AwsResourceMetadataScrapeJobs
type AwsResourceMetadataScrapeJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AwsResourceMetadataScrapeJob `json:"items"`
}

// Repository type metadata.
var (
	AwsResourceMetadataScrapeJob_Kind             = "AwsResourceMetadataScrapeJob"
	AwsResourceMetadataScrapeJob_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AwsResourceMetadataScrapeJob_Kind}.String()
	AwsResourceMetadataScrapeJob_KindAPIVersion   = AwsResourceMetadataScrapeJob_Kind + "." + CRDGroupVersion.String()
	AwsResourceMetadataScrapeJob_GroupVersionKind = CRDGroupVersion.WithKind(AwsResourceMetadataScrapeJob_Kind)
)

func init() {
	SchemeBuilder.Register(&AwsResourceMetadataScrapeJob{}, &AwsResourceMetadataScrapeJobList{})
}
