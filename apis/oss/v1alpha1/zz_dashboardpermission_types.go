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

type DashboardPermissionInitParameters struct {

	// Reference to a Dashboard in oss to populate dashboardUid.
	// +kubebuilder:validation:Optional
	DashboardRef *v1.Reference `json:"dashboardRef,omitempty" tf:"-"`

	// Selector for a Dashboard in oss to populate dashboardUid.
	// +kubebuilder:validation:Optional
	DashboardSelector *v1.Selector `json:"dashboardSelector,omitempty" tf:"-"`

	// (String) UID of the dashboard to apply permissions to.
	// UID of the dashboard to apply permissions to.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oss/v1alpha1.Dashboard
	// +crossplane:generate:reference:extractor=github.com/grafana/crossplane-provider-grafana/config/grafana.OptionalFieldExtractor("uid")
	// +crossplane:generate:reference:refFieldName=DashboardRef
	// +crossplane:generate:reference:selectorFieldName=DashboardSelector
	DashboardUID *string `json:"dashboardUid,omitempty" tf:"dashboard_uid,omitempty"`

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

	// (Block Set) The permission items to add/update. Items that are omitted from the list will be removed. (see below for nested schema)
	// The permission items to add/update. Items that are omitted from the list will be removed.
	Permissions []PermissionsInitParameters `json:"permissions,omitempty" tf:"permissions,omitempty"`
}

type DashboardPermissionObservation struct {

	// (String) UID of the dashboard to apply permissions to.
	// UID of the dashboard to apply permissions to.
	DashboardUID *string `json:"dashboardUid,omitempty" tf:"dashboard_uid,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The Organization ID. If not set, the Org ID defined in the provider block will be used.
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// (Block Set) The permission items to add/update. Items that are omitted from the list will be removed. (see below for nested schema)
	// The permission items to add/update. Items that are omitted from the list will be removed.
	Permissions []PermissionsObservation `json:"permissions,omitempty" tf:"permissions,omitempty"`
}

type DashboardPermissionParameters struct {

	// Reference to a Dashboard in oss to populate dashboardUid.
	// +kubebuilder:validation:Optional
	DashboardRef *v1.Reference `json:"dashboardRef,omitempty" tf:"-"`

	// Selector for a Dashboard in oss to populate dashboardUid.
	// +kubebuilder:validation:Optional
	DashboardSelector *v1.Selector `json:"dashboardSelector,omitempty" tf:"-"`

	// (String) UID of the dashboard to apply permissions to.
	// UID of the dashboard to apply permissions to.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oss/v1alpha1.Dashboard
	// +crossplane:generate:reference:extractor=github.com/grafana/crossplane-provider-grafana/config/grafana.OptionalFieldExtractor("uid")
	// +crossplane:generate:reference:refFieldName=DashboardRef
	// +crossplane:generate:reference:selectorFieldName=DashboardSelector
	// +kubebuilder:validation:Optional
	DashboardUID *string `json:"dashboardUid,omitempty" tf:"dashboard_uid,omitempty"`

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

	// (Block Set) The permission items to add/update. Items that are omitted from the list will be removed. (see below for nested schema)
	// The permission items to add/update. Items that are omitted from the list will be removed.
	// +kubebuilder:validation:Optional
	Permissions []PermissionsParameters `json:"permissions,omitempty" tf:"permissions,omitempty"`
}

type PermissionsInitParameters struct {

	// (String) Permission to associate with item. Must be one of View, Edit, or Admin.
	// Permission to associate with item. Must be one of `View`, `Edit`, or `Admin`.
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

	// (String) Manage permissions for Viewer or Editor roles.
	// Manage permissions for `Viewer` or `Editor` roles.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// (String) ID of the team to manage permissions for. Defaults to 0.
	// ID of the team to manage permissions for. Defaults to `0`.
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

	// (String) ID of the user or service account to manage permissions for. Defaults to 0.
	// ID of the user or service account to manage permissions for. Defaults to `0`.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oss/v1alpha1.User
	// +crossplane:generate:reference:refFieldName=UserRef
	// +crossplane:generate:reference:selectorFieldName=UserSelector
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a User in oss to populate userId.
	// +kubebuilder:validation:Optional
	UserRef *v1.Reference `json:"userRef,omitempty" tf:"-"`

	// Selector for a User in oss to populate userId.
	// +kubebuilder:validation:Optional
	UserSelector *v1.Selector `json:"userSelector,omitempty" tf:"-"`
}

type PermissionsObservation struct {

	// (String) Permission to associate with item. Must be one of View, Edit, or Admin.
	// Permission to associate with item. Must be one of `View`, `Edit`, or `Admin`.
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

	// (String) Manage permissions for Viewer or Editor roles.
	// Manage permissions for `Viewer` or `Editor` roles.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// (String) ID of the team to manage permissions for. Defaults to 0.
	// ID of the team to manage permissions for. Defaults to `0`.
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// (String) ID of the user or service account to manage permissions for. Defaults to 0.
	// ID of the user or service account to manage permissions for. Defaults to `0`.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type PermissionsParameters struct {

	// (String) Permission to associate with item. Must be one of View, Edit, or Admin.
	// Permission to associate with item. Must be one of `View`, `Edit`, or `Admin`.
	// +kubebuilder:validation:Optional
	Permission *string `json:"permission" tf:"permission,omitempty"`

	// (String) Manage permissions for Viewer or Editor roles.
	// Manage permissions for `Viewer` or `Editor` roles.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// (String) ID of the team to manage permissions for. Defaults to 0.
	// ID of the team to manage permissions for. Defaults to `0`.
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

	// (String) ID of the user or service account to manage permissions for. Defaults to 0.
	// ID of the user or service account to manage permissions for. Defaults to `0`.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oss/v1alpha1.User
	// +crossplane:generate:reference:refFieldName=UserRef
	// +crossplane:generate:reference:selectorFieldName=UserSelector
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a User in oss to populate userId.
	// +kubebuilder:validation:Optional
	UserRef *v1.Reference `json:"userRef,omitempty" tf:"-"`

	// Selector for a User in oss to populate userId.
	// +kubebuilder:validation:Optional
	UserSelector *v1.Selector `json:"userSelector,omitempty" tf:"-"`
}

// DashboardPermissionSpec defines the desired state of DashboardPermission
type DashboardPermissionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DashboardPermissionParameters `json:"forProvider"`
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
	InitProvider DashboardPermissionInitParameters `json:"initProvider,omitempty"`
}

// DashboardPermissionStatus defines the observed state of DashboardPermission.
type DashboardPermissionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DashboardPermissionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DashboardPermission is the Schema for the DashboardPermissions API. Manages the entire set of permissions for a dashboard. Permissions that aren't specified when applying this resource will be removed.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type DashboardPermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DashboardPermissionSpec   `json:"spec"`
	Status            DashboardPermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DashboardPermissionList contains a list of DashboardPermissions
type DashboardPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DashboardPermission `json:"items"`
}

// Repository type metadata.
var (
	DashboardPermission_Kind             = "DashboardPermission"
	DashboardPermission_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DashboardPermission_Kind}.String()
	DashboardPermission_KindAPIVersion   = DashboardPermission_Kind + "." + CRDGroupVersion.String()
	DashboardPermission_GroupVersionKind = CRDGroupVersion.WithKind(DashboardPermission_Kind)
)

func init() {
	SchemeBuilder.Register(&DashboardPermission{}, &DashboardPermissionList{})
}
