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

type TeamExternalGroupInitParameters struct {

	// (Set of String) The team external groups list
	// The team external groups list
	// +listType=set
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// (String) The Team ID
	// The Team ID
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oss/v1alpha1.Team
	// +crossplane:generate:reference:refFieldName=TeamRef
	// +crossplane:generate:reference:selectorFieldName=TeamSelector
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// Reference to a Team in oss to populate teamId.
	// +kubebuilder:validation:Optional
	TeamRef *v1.Reference `json:"teamRef,omitempty" tf:"-"`

	// Selector for a Team in oss to populate teamId.
	// +kubebuilder:validation:Optional
	TeamSelector *v1.Selector `json:"teamSelector,omitempty" tf:"-"`
}

type TeamExternalGroupObservation struct {

	// (Set of String) The team external groups list
	// The team external groups list
	// +listType=set
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The Team ID
	// The Team ID
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`
}

type TeamExternalGroupParameters struct {

	// (Set of String) The team external groups list
	// The team external groups list
	// +kubebuilder:validation:Optional
	// +listType=set
	Groups []*string `json:"groups,omitempty" tf:"groups,omitempty"`

	// (String) The Team ID
	// The Team ID
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oss/v1alpha1.Team
	// +crossplane:generate:reference:refFieldName=TeamRef
	// +crossplane:generate:reference:selectorFieldName=TeamSelector
	// +kubebuilder:validation:Optional
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// Reference to a Team in oss to populate teamId.
	// +kubebuilder:validation:Optional
	TeamRef *v1.Reference `json:"teamRef,omitempty" tf:"-"`

	// Selector for a Team in oss to populate teamId.
	// +kubebuilder:validation:Optional
	TeamSelector *v1.Selector `json:"teamSelector,omitempty" tf:"-"`
}

// TeamExternalGroupSpec defines the desired state of TeamExternalGroup
type TeamExternalGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TeamExternalGroupParameters `json:"forProvider"`
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
	InitProvider TeamExternalGroupInitParameters `json:"initProvider,omitempty"`
}

// TeamExternalGroupStatus defines the observed state of TeamExternalGroup.
type TeamExternalGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TeamExternalGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TeamExternalGroup is the Schema for the TeamExternalGroups API. Equivalent to the the team_sync attribute of the grafana_team resource. Use one or the other to configure a team's external groups syncing config.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type TeamExternalGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.groups) || (has(self.initProvider) && has(self.initProvider.groups))",message="spec.forProvider.groups is a required parameter"
	Spec   TeamExternalGroupSpec   `json:"spec"`
	Status TeamExternalGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TeamExternalGroupList contains a list of TeamExternalGroups
type TeamExternalGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TeamExternalGroup `json:"items"`
}

// Repository type metadata.
var (
	TeamExternalGroup_Kind             = "TeamExternalGroup"
	TeamExternalGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TeamExternalGroup_Kind}.String()
	TeamExternalGroup_KindAPIVersion   = TeamExternalGroup_Kind + "." + CRDGroupVersion.String()
	TeamExternalGroup_GroupVersionKind = CRDGroupVersion.WithKind(TeamExternalGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&TeamExternalGroup{}, &TeamExternalGroupList{})
}
