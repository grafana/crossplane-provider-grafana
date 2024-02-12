// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type InstallationInitParameters struct {

	// Reference to a Stack in cloud to populate stackId.
	// +kubebuilder:validation:Optional
	CloudStackRef *v1.Reference `json:"cloudStackRef,omitempty" tf:"-"`

	// Selector for a Stack in cloud to populate stackId.
	// +kubebuilder:validation:Optional
	CloudStackSelector *v1.Selector `json:"cloudStackSelector,omitempty" tf:"-"`

	// (String) The ID or slug of the stack to install SM on.
	// The ID or slug of the stack to install SM on.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/cloud/v1alpha1.Stack
	// +crossplane:generate:reference:refFieldName=CloudStackRef
	// +crossplane:generate:reference:selectorFieldName=CloudStackSelector
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// cloud/monitor-public-endpoints/private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
	// The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/monitor-public-endpoints/private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
	StackSmAPIURL *string `json:"stackSmApiUrl,omitempty" tf:"stack_sm_api_url,omitempty"`
}

type InstallationObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Generated token to access the SM API.
	// Generated token to access the SM API.
	SmAccessToken *string `json:"smAccessToken,omitempty" tf:"sm_access_token,omitempty"`

	// (String) The ID or slug of the stack to install SM on.
	// The ID or slug of the stack to install SM on.
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// cloud/monitor-public-endpoints/private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
	// The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/monitor-public-endpoints/private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
	StackSmAPIURL *string `json:"stackSmApiUrl,omitempty" tf:"stack_sm_api_url,omitempty"`
}

type InstallationParameters struct {

	// Reference to a Stack in cloud to populate stackId.
	// +kubebuilder:validation:Optional
	CloudStackRef *v1.Reference `json:"cloudStackRef,omitempty" tf:"-"`

	// Selector for a Stack in cloud to populate stackId.
	// +kubebuilder:validation:Optional
	CloudStackSelector *v1.Selector `json:"cloudStackSelector,omitempty" tf:"-"`

	// (String, Sensitive) The Cloud API Key with the MetricsPublisher role used to publish metrics to the SM API
	// The Cloud API Key with the `MetricsPublisher` role used to publish metrics to the SM API
	// +kubebuilder:validation:Optional
	MetricsPublisherKeySecretRef v1.SecretKeySelector `json:"metricsPublisherKeySecretRef" tf:"-"`

	// (String) The ID or slug of the stack to install SM on.
	// The ID or slug of the stack to install SM on.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/cloud/v1alpha1.Stack
	// +crossplane:generate:reference:refFieldName=CloudStackRef
	// +crossplane:generate:reference:selectorFieldName=CloudStackSelector
	// +kubebuilder:validation:Optional
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// cloud/monitor-public-endpoints/private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
	// The URL of the SM API to install SM on. This depends on the stack region, find the list of API URLs here: https://grafana.com/docs/grafana-cloud/monitor-public-endpoints/private-probes/#probe-api-server-url. A static mapping exists in the provider but it may not contain all the regions. If it does contain the stack's region, this field is computed automatically and readable.
	// +kubebuilder:validation:Optional
	StackSmAPIURL *string `json:"stackSmApiUrl,omitempty" tf:"stack_sm_api_url,omitempty"`
}

// InstallationSpec defines the desired state of Installation
type InstallationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstallationParameters `json:"forProvider"`
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
	InitProvider InstallationInitParameters `json:"initProvider,omitempty"`
}

// InstallationStatus defines the observed state of Installation.
type InstallationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstallationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Installation is the Schema for the Installations API. Sets up Synthetic Monitoring on a Grafana cloud stack and generates a token. Once a Grafana Cloud stack is created, a user can either use this resource or go into the UI to install synthetic monitoring. This resource cannot be imported but it can be used on an existing Synthetic Monitoring installation without issues. Note that this resource must be used on a provider configured with Grafana Cloud credentials. Official documentation https://grafana.com/docs/grafana-cloud/monitor-public-endpoints/installation/API documentation https://github.com/grafana/synthetic-monitoring-api-go-client/blob/main/docs/API.md#apiv1registerinstall
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type Installation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.metricsPublisherKeySecretRef)",message="spec.forProvider.metricsPublisherKeySecretRef is a required parameter"
	Spec   InstallationSpec   `json:"spec"`
	Status InstallationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstallationList contains a list of Installations
type InstallationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Installation `json:"items"`
}

// Repository type metadata.
var (
	Installation_Kind             = "Installation"
	Installation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Installation_Kind}.String()
	Installation_KindAPIVersion   = Installation_Kind + "." + CRDGroupVersion.String()
	Installation_GroupVersionKind = CRDGroupVersion.WithKind(Installation_Kind)
)

func init() {
	SchemeBuilder.Register(&Installation{}, &InstallationList{})
}
