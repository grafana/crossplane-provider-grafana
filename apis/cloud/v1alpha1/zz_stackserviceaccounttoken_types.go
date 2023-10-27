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

type StackServiceAccountTokenInitParameters struct {

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number)
	SecondsToLive *float64 `json:"secondsToLive,omitempty" tf:"seconds_to_live,omitempty"`
}

type StackServiceAccountTokenObservation struct {

	// (String)
	Expiration *string `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// (Boolean)
	HasExpired *bool `json:"hasExpired,omitempty" tf:"has_expired,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number)
	SecondsToLive *float64 `json:"secondsToLive,omitempty" tf:"seconds_to_live,omitempty"`

	// (String)
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

	// (String)
	StackSlug *string `json:"stackSlug,omitempty" tf:"stack_slug,omitempty"`
}

type StackServiceAccountTokenParameters struct {

	// Reference to a Stack in cloud to populate stackSlug.
	// +kubebuilder:validation:Optional
	CloudStackRef *v1.Reference `json:"cloudStackRef,omitempty" tf:"-"`

	// Selector for a Stack in cloud to populate stackSlug.
	// +kubebuilder:validation:Optional
	CloudStackSelector *v1.Selector `json:"cloudStackSelector,omitempty" tf:"-"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	SecondsToLive *float64 `json:"secondsToLive,omitempty" tf:"seconds_to_live,omitempty"`

	// (String)
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/cloud/v1alpha1.StackServiceAccount
	// +crossplane:generate:reference:refFieldName=ServiceAccountRef
	// +crossplane:generate:reference:selectorFieldName=ServiceAccountSelector
	// +kubebuilder:validation:Optional
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

	// Reference to a StackServiceAccount in cloud to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountRef *v1.Reference `json:"serviceAccountRef,omitempty" tf:"-"`

	// Selector for a StackServiceAccount in cloud to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountSelector *v1.Selector `json:"serviceAccountSelector,omitempty" tf:"-"`

	// (String)
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/cloud/v1alpha1.Stack
	// +crossplane:generate:reference:extractor=github.com/grafana/crossplane-provider-grafana/config/grafana.CloudStackSlugExtractor()
	// +crossplane:generate:reference:refFieldName=CloudStackRef
	// +crossplane:generate:reference:selectorFieldName=CloudStackSelector
	// +kubebuilder:validation:Optional
	StackSlug *string `json:"stackSlug,omitempty" tf:"stack_slug,omitempty"`
}

// StackServiceAccountTokenSpec defines the desired state of StackServiceAccountToken
type StackServiceAccountTokenSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StackServiceAccountTokenParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider StackServiceAccountTokenInitParameters `json:"initProvider,omitempty"`
}

// StackServiceAccountTokenStatus defines the observed state of StackServiceAccountToken.
type StackServiceAccountTokenStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StackServiceAccountTokenObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StackServiceAccountToken is the Schema for the StackServiceAccountTokens API. Note: This resource is available only with Grafana 9.1+. Manages service account tokens of a Grafana Cloud stack using the Cloud API This can be used to bootstrap a management service account token for a new stack Official documentation https://grafana.com/docs/grafana/latest/administration/service-accounts/HTTP API https://grafana.com/docs/grafana/latest/developers/http_api/serviceaccount/#service-account-api
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type StackServiceAccountToken struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   StackServiceAccountTokenSpec   `json:"spec"`
	Status StackServiceAccountTokenStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StackServiceAccountTokenList contains a list of StackServiceAccountTokens
type StackServiceAccountTokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StackServiceAccountToken `json:"items"`
}

// Repository type metadata.
var (
	StackServiceAccountToken_Kind             = "StackServiceAccountToken"
	StackServiceAccountToken_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StackServiceAccountToken_Kind}.String()
	StackServiceAccountToken_KindAPIVersion   = StackServiceAccountToken_Kind + "." + CRDGroupVersion.String()
	StackServiceAccountToken_GroupVersionKind = CRDGroupVersion.WithKind(StackServiceAccountToken_Kind)
)

func init() {
	SchemeBuilder.Register(&StackServiceAccountToken{}, &StackServiceAccountTokenList{})
}
