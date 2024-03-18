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

type ItemInitParameters struct {

	// (Number)
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// (String)
	Title *string `json:"title,omitempty" tf:"title,omitempty"`

	// (String)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String)
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ItemObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Number)
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// (String)
	Title *string `json:"title,omitempty" tf:"title,omitempty"`

	// (String)
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String)
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ItemParameters struct {

	// (Number)
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order" tf:"order,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Title *string `json:"title" tf:"title,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PlaylistInitParameters struct {

	// (String)
	Interval *string `json:"interval,omitempty" tf:"interval,omitempty"`

	// (Block Set, Min: 1) (see below for nested schema)
	Item []ItemInitParameters `json:"item,omitempty" tf:"item,omitempty"`

	// (String) The name of the playlist.
	// The name of the playlist.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The Organization ID. If not set, the Org ID defined in the provider block will be used.
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oss/v1alpha1.Organization
	// +crossplane:generate:reference:refFieldName=OrganizationRef
	// +crossplane:generate:reference:selectorFieldName=OrganizationSelector
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Reference to a Organization in oss to populate orgId.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`

	// Selector for a Organization in oss to populate orgId.
	// +kubebuilder:validation:Optional
	OrganizationSelector *v1.Selector `json:"organizationSelector,omitempty" tf:"-"`
}

type PlaylistObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Interval *string `json:"interval,omitempty" tf:"interval,omitempty"`

	// (Block Set, Min: 1) (see below for nested schema)
	Item []ItemObservation `json:"item,omitempty" tf:"item,omitempty"`

	// (String) The name of the playlist.
	// The name of the playlist.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The Organization ID. If not set, the Org ID defined in the provider block will be used.
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`
}

type PlaylistParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Interval *string `json:"interval,omitempty" tf:"interval,omitempty"`

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Item []ItemParameters `json:"item,omitempty" tf:"item,omitempty"`

	// (String) The name of the playlist.
	// The name of the playlist.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The Organization ID. If not set, the Org ID defined in the provider block will be used.
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oss/v1alpha1.Organization
	// +crossplane:generate:reference:refFieldName=OrganizationRef
	// +crossplane:generate:reference:selectorFieldName=OrganizationSelector
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Reference to a Organization in oss to populate orgId.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`

	// Selector for a Organization in oss to populate orgId.
	// +kubebuilder:validation:Optional
	OrganizationSelector *v1.Selector `json:"organizationSelector,omitempty" tf:"-"`
}

// PlaylistSpec defines the desired state of Playlist
type PlaylistSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PlaylistParameters `json:"forProvider"`
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
	InitProvider PlaylistInitParameters `json:"initProvider,omitempty"`
}

// PlaylistStatus defines the observed state of Playlist.
type PlaylistStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PlaylistObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Playlist is the Schema for the Playlists API. Official documentation https://grafana.com/docs/grafana/latest/dashboards/create-manage-playlists/HTTP API https://grafana.com/docs/grafana/latest/developers/http_api/playlist/
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type Playlist struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.interval) || (has(self.initProvider) && has(self.initProvider.interval))",message="spec.forProvider.interval is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.item) || (has(self.initProvider) && has(self.initProvider.item))",message="spec.forProvider.item is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   PlaylistSpec   `json:"spec"`
	Status PlaylistStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PlaylistList contains a list of Playlists
type PlaylistList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Playlist `json:"items"`
}

// Repository type metadata.
var (
	Playlist_Kind             = "Playlist"
	Playlist_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Playlist_Kind}.String()
	Playlist_KindAPIVersion   = Playlist_Kind + "." + CRDGroupVersion.String()
	Playlist_GroupVersionKind = CRDGroupVersion.WithKind(Playlist_Kind)
)

func init() {
	SchemeBuilder.Register(&Playlist{}, &PlaylistList{})
}
