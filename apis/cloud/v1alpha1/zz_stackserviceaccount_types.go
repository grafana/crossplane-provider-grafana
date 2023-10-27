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

type StackServiceAccountInitParameters struct {

	// (Boolean) The disabled status for the service account. Defaults to false.
	// The disabled status for the service account. Defaults to `false`.
	IsDisabled *bool `json:"isDisabled,omitempty" tf:"is_disabled,omitempty"`

	// (String) The name of the service account.
	// The name of the service account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The basic role of the service account in the organization.
	// The basic role of the service account in the organization.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type StackServiceAccountObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) The disabled status for the service account. Defaults to false.
	// The disabled status for the service account. Defaults to `false`.
	IsDisabled *bool `json:"isDisabled,omitempty" tf:"is_disabled,omitempty"`

	// (String) The name of the service account.
	// The name of the service account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The basic role of the service account in the organization.
	// The basic role of the service account in the organization.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// (String)
	StackSlug *string `json:"stackSlug,omitempty" tf:"stack_slug,omitempty"`
}

type StackServiceAccountParameters struct {

	// Reference to a Stack in cloud to populate stackSlug.
	// +kubebuilder:validation:Optional
	CloudStackRef *v1.Reference `json:"cloudStackRef,omitempty" tf:"-"`

	// Selector for a Stack in cloud to populate stackSlug.
	// +kubebuilder:validation:Optional
	CloudStackSelector *v1.Selector `json:"cloudStackSelector,omitempty" tf:"-"`

	// (Boolean) The disabled status for the service account. Defaults to false.
	// The disabled status for the service account. Defaults to `false`.
	// +kubebuilder:validation:Optional
	IsDisabled *bool `json:"isDisabled,omitempty" tf:"is_disabled,omitempty"`

	// (String) The name of the service account.
	// The name of the service account.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The basic role of the service account in the organization.
	// The basic role of the service account in the organization.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// (String)
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/cloud/v1alpha1.Stack
	// +crossplane:generate:reference:extractor=github.com/grafana/crossplane-provider-grafana/config/grafana.CloudStackSlugExtractor()
	// +crossplane:generate:reference:refFieldName=CloudStackRef
	// +crossplane:generate:reference:selectorFieldName=CloudStackSelector
	// +kubebuilder:validation:Optional
	StackSlug *string `json:"stackSlug,omitempty" tf:"stack_slug,omitempty"`
}

// StackServiceAccountSpec defines the desired state of StackServiceAccount
type StackServiceAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StackServiceAccountParameters `json:"forProvider"`
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
	InitProvider StackServiceAccountInitParameters `json:"initProvider,omitempty"`
}

// StackServiceAccountStatus defines the observed state of StackServiceAccount.
type StackServiceAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StackServiceAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StackServiceAccount is the Schema for the StackServiceAccounts API. Note: This resource is available only with Grafana 9.1+. Manages service accounts of a Grafana Cloud stack using the Cloud API This can be used to bootstrap a management service account for a new stack Official documentation https://grafana.com/docs/grafana/latest/administration/service-accounts/HTTP API https://grafana.com/docs/grafana/latest/developers/http_api/serviceaccount/#service-account-api
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type StackServiceAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   StackServiceAccountSpec   `json:"spec"`
	Status StackServiceAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StackServiceAccountList contains a list of StackServiceAccounts
type StackServiceAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StackServiceAccount `json:"items"`
}

// Repository type metadata.
var (
	StackServiceAccount_Kind             = "StackServiceAccount"
	StackServiceAccount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StackServiceAccount_Kind}.String()
	StackServiceAccount_KindAPIVersion   = StackServiceAccount_Kind + "." + CRDGroupVersion.String()
	StackServiceAccount_GroupVersionKind = CRDGroupVersion.WithKind(StackServiceAccount_Kind)
)

func init() {
	SchemeBuilder.Register(&StackServiceAccount{}, &StackServiceAccountList{})
}
