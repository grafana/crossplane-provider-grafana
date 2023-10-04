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

type APIKeyObservation struct {
	Expiration *string `json:"expiration,omitempty" tf:"expiration,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type APIKeyParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	SecondsToLive *float64 `json:"secondsToLive,omitempty" tf:"seconds_to_live,omitempty"`
}

// APIKeySpec defines the desired state of APIKey
type APIKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIKeyParameters `json:"forProvider"`
}

// APIKeyStatus defines the observed state of APIKey.
type APIKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APIKey is the Schema for the APIKeys API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type APIKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APIKeySpec   `json:"spec"`
	Status            APIKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIKeyList contains a list of APIKeys
type APIKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIKey `json:"items"`
}

// Repository type metadata.
var (
	APIKey_Kind             = "APIKey"
	APIKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APIKey_Kind}.String()
	APIKey_KindAPIVersion   = APIKey_Kind + "." + CRDGroupVersion.String()
	APIKey_GroupVersionKind = CRDGroupVersion.WithKind(APIKey_Kind)
)

func init() {
	SchemeBuilder.Register(&APIKey{}, &APIKeyList{})
}
