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

type MetricsEndpointScrapeJobInitParameters struct {

	// (String, Sensitive) Password for basic authentication, use if scrape job is using basic authentication method
	// Password for basic authentication, use if scrape job is using basic authentication method
	AuthenticationBasicPasswordSecretRef *v1.SecretKeySelector `json:"authenticationBasicPasswordSecretRef,omitempty" tf:"-"`

	// (String) Username for basic authentication, use if scrape job is using basic authentication method
	// Username for basic authentication, use if scrape job is using basic authentication method
	AuthenticationBasicUsername *string `json:"authenticationBasicUsername,omitempty" tf:"authentication_basic_username,omitempty"`

	// (String, Sensitive) Bearer token used for authentication, use if scrape job is using bearer authentication method
	// Bearer token used for authentication, use if scrape job is using bearer authentication method
	AuthenticationBearerTokenSecretRef *v1.SecretKeySelector `json:"authenticationBearerTokenSecretRef,omitempty" tf:"-"`

	// (String) Method to pass authentication credentials: basic or bearer.
	// Method to pass authentication credentials: basic or bearer.
	AuthenticationMethod *string `json:"authenticationMethod,omitempty" tf:"authentication_method,omitempty"`

	// (Boolean) Whether the metrics endpoint scrape job is enabled or not.
	// Whether the metrics endpoint scrape job is enabled or not.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The name of the metrics endpoint scrape job.
	// The name of the metrics endpoint scrape job.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number) Frequency for scraping the metrics endpoint: 30, 60, or 120 seconds.
	// Frequency for scraping the metrics endpoint: 30, 60, or 120 seconds.
	ScrapeIntervalSeconds *float64 `json:"scrapeIntervalSeconds,omitempty" tf:"scrape_interval_seconds,omitempty"`

	// (String) The Stack ID of the Grafana Cloud instance.
	// The Stack ID of the Grafana Cloud instance.
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// (String) The url to scrape metrics from; a valid HTTPs URL is required.
	// The url to scrape metrics from; a valid HTTPs URL is required.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type MetricsEndpointScrapeJobObservation struct {

	// (String) Username for basic authentication, use if scrape job is using basic authentication method
	// Username for basic authentication, use if scrape job is using basic authentication method
	AuthenticationBasicUsername *string `json:"authenticationBasicUsername,omitempty" tf:"authentication_basic_username,omitempty"`

	// (String) Method to pass authentication credentials: basic or bearer.
	// Method to pass authentication credentials: basic or bearer.
	AuthenticationMethod *string `json:"authenticationMethod,omitempty" tf:"authentication_method,omitempty"`

	// (Boolean) Whether the metrics endpoint scrape job is enabled or not.
	// Whether the metrics endpoint scrape job is enabled or not.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// This has the format "{{ stack_id }}:{{ name }}".
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the metrics endpoint scrape job.
	// The name of the metrics endpoint scrape job.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number) Frequency for scraping the metrics endpoint: 30, 60, or 120 seconds.
	// Frequency for scraping the metrics endpoint: 30, 60, or 120 seconds.
	ScrapeIntervalSeconds *float64 `json:"scrapeIntervalSeconds,omitempty" tf:"scrape_interval_seconds,omitempty"`

	// (String) The Stack ID of the Grafana Cloud instance.
	// The Stack ID of the Grafana Cloud instance.
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// (String) The url to scrape metrics from; a valid HTTPs URL is required.
	// The url to scrape metrics from; a valid HTTPs URL is required.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type MetricsEndpointScrapeJobParameters struct {

	// (String, Sensitive) Password for basic authentication, use if scrape job is using basic authentication method
	// Password for basic authentication, use if scrape job is using basic authentication method
	// +kubebuilder:validation:Optional
	AuthenticationBasicPasswordSecretRef *v1.SecretKeySelector `json:"authenticationBasicPasswordSecretRef,omitempty" tf:"-"`

	// (String) Username for basic authentication, use if scrape job is using basic authentication method
	// Username for basic authentication, use if scrape job is using basic authentication method
	// +kubebuilder:validation:Optional
	AuthenticationBasicUsername *string `json:"authenticationBasicUsername,omitempty" tf:"authentication_basic_username,omitempty"`

	// (String, Sensitive) Bearer token used for authentication, use if scrape job is using bearer authentication method
	// Bearer token used for authentication, use if scrape job is using bearer authentication method
	// +kubebuilder:validation:Optional
	AuthenticationBearerTokenSecretRef *v1.SecretKeySelector `json:"authenticationBearerTokenSecretRef,omitempty" tf:"-"`

	// (String) Method to pass authentication credentials: basic or bearer.
	// Method to pass authentication credentials: basic or bearer.
	// +kubebuilder:validation:Optional
	AuthenticationMethod *string `json:"authenticationMethod,omitempty" tf:"authentication_method,omitempty"`

	// (Boolean) Whether the metrics endpoint scrape job is enabled or not.
	// Whether the metrics endpoint scrape job is enabled or not.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The name of the metrics endpoint scrape job.
	// The name of the metrics endpoint scrape job.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number) Frequency for scraping the metrics endpoint: 30, 60, or 120 seconds.
	// Frequency for scraping the metrics endpoint: 30, 60, or 120 seconds.
	// +kubebuilder:validation:Optional
	ScrapeIntervalSeconds *float64 `json:"scrapeIntervalSeconds,omitempty" tf:"scrape_interval_seconds,omitempty"`

	// (String) The Stack ID of the Grafana Cloud instance.
	// The Stack ID of the Grafana Cloud instance.
	// +kubebuilder:validation:Optional
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// (String) The url to scrape metrics from; a valid HTTPs URL is required.
	// The url to scrape metrics from; a valid HTTPs URL is required.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

// MetricsEndpointScrapeJobSpec defines the desired state of MetricsEndpointScrapeJob
type MetricsEndpointScrapeJobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MetricsEndpointScrapeJobParameters `json:"forProvider"`
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
	InitProvider MetricsEndpointScrapeJobInitParameters `json:"initProvider,omitempty"`
}

// MetricsEndpointScrapeJobStatus defines the observed state of MetricsEndpointScrapeJob.
type MetricsEndpointScrapeJobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MetricsEndpointScrapeJobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MetricsEndpointScrapeJob is the Schema for the MetricsEndpointScrapeJobs API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type MetricsEndpointScrapeJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authenticationMethod) || (has(self.initProvider) && has(self.initProvider.authenticationMethod))",message="spec.forProvider.authenticationMethod is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.stackId) || (has(self.initProvider) && has(self.initProvider.stackId))",message="spec.forProvider.stackId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.url) || (has(self.initProvider) && has(self.initProvider.url))",message="spec.forProvider.url is a required parameter"
	Spec   MetricsEndpointScrapeJobSpec   `json:"spec"`
	Status MetricsEndpointScrapeJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MetricsEndpointScrapeJobList contains a list of MetricsEndpointScrapeJobs
type MetricsEndpointScrapeJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MetricsEndpointScrapeJob `json:"items"`
}

// Repository type metadata.
var (
	MetricsEndpointScrapeJob_Kind             = "MetricsEndpointScrapeJob"
	MetricsEndpointScrapeJob_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MetricsEndpointScrapeJob_Kind}.String()
	MetricsEndpointScrapeJob_KindAPIVersion   = MetricsEndpointScrapeJob_Kind + "." + CRDGroupVersion.String()
	MetricsEndpointScrapeJob_GroupVersionKind = CRDGroupVersion.WithKind(MetricsEndpointScrapeJob_Kind)
)

func init() {
	SchemeBuilder.Register(&MetricsEndpointScrapeJob{}, &MetricsEndpointScrapeJobList{})
}
