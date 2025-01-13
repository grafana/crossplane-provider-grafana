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

type AzureCredentialInitParameters struct {

	// (String) The client ID of the Azure Credential.
	// The client ID of the Azure Credential.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// (String, Sensitive) The client secret of the Azure Credential.
	// The client secret of the Azure Credential.
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// (String) The name of the Azure Credential.
	// The name of the Azure Credential.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List) The list of tag filters to apply to resources. (see below for nested schema)
	// The list of tag filters to apply to resources.
	ResourceDiscoveryTagFilter []AzureCredentialResourceDiscoveryTagFilterInitParameters `json:"resourceDiscoveryTagFilter,omitempty" tf:"resource_discovery_tag_filter,omitempty"`

	// (String) The StackID of the Grafana Cloud instance.
	// The StackID of the Grafana Cloud instance.
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// (String) The tenant ID of the Azure Credential.
	// The tenant ID of the Azure Credential.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AzureCredentialObservation struct {

	// (String) The client ID of the Azure Credential.
	// The client ID of the Azure Credential.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// This has the format "{{ stack_id }}:{{ resource_id }}".
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the Azure Credential.
	// The name of the Azure Credential.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List) The list of tag filters to apply to resources. (see below for nested schema)
	// The list of tag filters to apply to resources.
	ResourceDiscoveryTagFilter []AzureCredentialResourceDiscoveryTagFilterObservation `json:"resourceDiscoveryTagFilter,omitempty" tf:"resource_discovery_tag_filter,omitempty"`

	// (String) The ID given by the Grafana Cloud Provider API to this AWS Account resource.
	// The ID given by the Grafana Cloud Provider API to this AWS Account resource.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// (String) The StackID of the Grafana Cloud instance.
	// The StackID of the Grafana Cloud instance.
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// (String) The tenant ID of the Azure Credential.
	// The tenant ID of the Azure Credential.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AzureCredentialParameters struct {

	// (String) The client ID of the Azure Credential.
	// The client ID of the Azure Credential.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// (String, Sensitive) The client secret of the Azure Credential.
	// The client secret of the Azure Credential.
	// +kubebuilder:validation:Optional
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// (String) The name of the Azure Credential.
	// The name of the Azure Credential.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List) The list of tag filters to apply to resources. (see below for nested schema)
	// The list of tag filters to apply to resources.
	// +kubebuilder:validation:Optional
	ResourceDiscoveryTagFilter []AzureCredentialResourceDiscoveryTagFilterParameters `json:"resourceDiscoveryTagFilter,omitempty" tf:"resource_discovery_tag_filter,omitempty"`

	// (String) The StackID of the Grafana Cloud instance.
	// The StackID of the Grafana Cloud instance.
	// +kubebuilder:validation:Optional
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// (String) The tenant ID of the Azure Credential.
	// The tenant ID of the Azure Credential.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AzureCredentialResourceDiscoveryTagFilterInitParameters struct {

	// (String) The key of the tag filter.
	// The key of the tag filter.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String) The value of the tag filter.
	// The value of the tag filter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AzureCredentialResourceDiscoveryTagFilterObservation struct {

	// (String) The key of the tag filter.
	// The key of the tag filter.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String) The value of the tag filter.
	// The value of the tag filter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type AzureCredentialResourceDiscoveryTagFilterParameters struct {

	// (String) The key of the tag filter.
	// The key of the tag filter.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// (String) The value of the tag filter.
	// The value of the tag filter.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

// AzureCredentialSpec defines the desired state of AzureCredential
type AzureCredentialSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AzureCredentialParameters `json:"forProvider"`
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
	InitProvider AzureCredentialInitParameters `json:"initProvider,omitempty"`
}

// AzureCredentialStatus defines the observed state of AzureCredential.
type AzureCredentialStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AzureCredentialObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AzureCredential is the Schema for the AzureCredentials API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type AzureCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientId) || (has(self.initProvider) && has(self.initProvider.clientId))",message="spec.forProvider.clientId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientSecretSecretRef)",message="spec.forProvider.clientSecretSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.stackId) || (has(self.initProvider) && has(self.initProvider.stackId))",message="spec.forProvider.stackId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tenantId) || (has(self.initProvider) && has(self.initProvider.tenantId))",message="spec.forProvider.tenantId is a required parameter"
	Spec   AzureCredentialSpec   `json:"spec"`
	Status AzureCredentialStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureCredentialList contains a list of AzureCredentials
type AzureCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureCredential `json:"items"`
}

// Repository type metadata.
var (
	AzureCredential_Kind             = "AzureCredential"
	AzureCredential_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AzureCredential_Kind}.String()
	AzureCredential_KindAPIVersion   = AzureCredential_Kind + "." + CRDGroupVersion.String()
	AzureCredential_GroupVersionKind = CRDGroupVersion.WithKind(AzureCredential_Kind)
)

func init() {
	SchemeBuilder.Register(&AzureCredential{}, &AzureCredentialList{})
}
