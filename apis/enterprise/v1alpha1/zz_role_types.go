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

type RoleInitParameters struct {

	// (Boolean) Whether the role version should be incremented automatically on updates (and set to 1 on creation). This field or version should be set.
	// Whether the role version should be incremented automatically on updates (and set to 1 on creation). This field or `version` should be set.
	AutoIncrementVersion *bool `json:"autoIncrementVersion,omitempty" tf:"auto_increment_version,omitempty"`

	// (String) Description of the role.
	// Description of the role.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Display name of the role. Available with Grafana 8.5+.
	// Display name of the role. Available with Grafana 8.5+.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Boolean) Boolean to state whether the role is available across all organizations or not. Defaults to false.
	// Boolean to state whether the role is available across all organizations or not. Defaults to `false`.
	Global *bool `json:"global,omitempty" tf:"global,omitempty"`

	// (String) Group of the role. Available with Grafana 8.5+.
	// Group of the role. Available with Grafana 8.5+.
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// (Boolean) Boolean to state whether the role should be visible in the Grafana UI or not. Available with Grafana 8.5+. Defaults to false.
	// Boolean to state whether the role should be visible in the Grafana UI or not. Available with Grafana 8.5+. Defaults to `false`.
	Hidden *bool `json:"hidden,omitempty" tf:"hidden,omitempty"`

	// (String) Name of the role
	// Name of the role
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

	// (Block Set) Specific set of actions granted by the role. (see below for nested schema)
	// Specific set of actions granted by the role.
	Permissions []RolePermissionsInitParameters `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// (String) Unique identifier of the role. Used for assignments.
	// Unique identifier of the role. Used for assignments.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// (Number) Version of the role. A role is updated only on version increase. This field or auto_increment_version should be set.
	// Version of the role. A role is updated only on version increase. This field or `auto_increment_version` should be set.
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type RoleObservation struct {

	// (Boolean) Whether the role version should be incremented automatically on updates (and set to 1 on creation). This field or version should be set.
	// Whether the role version should be incremented automatically on updates (and set to 1 on creation). This field or `version` should be set.
	AutoIncrementVersion *bool `json:"autoIncrementVersion,omitempty" tf:"auto_increment_version,omitempty"`

	// (String) Description of the role.
	// Description of the role.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Display name of the role. Available with Grafana 8.5+.
	// Display name of the role. Available with Grafana 8.5+.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Boolean) Boolean to state whether the role is available across all organizations or not. Defaults to false.
	// Boolean to state whether the role is available across all organizations or not. Defaults to `false`.
	Global *bool `json:"global,omitempty" tf:"global,omitempty"`

	// (String) Group of the role. Available with Grafana 8.5+.
	// Group of the role. Available with Grafana 8.5+.
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// (Boolean) Boolean to state whether the role should be visible in the Grafana UI or not. Available with Grafana 8.5+. Defaults to false.
	// Boolean to state whether the role should be visible in the Grafana UI or not. Available with Grafana 8.5+. Defaults to `false`.
	Hidden *bool `json:"hidden,omitempty" tf:"hidden,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Name of the role
	// Name of the role
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The Organization ID. If not set, the Org ID defined in the provider block will be used.
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// (Block Set) Specific set of actions granted by the role. (see below for nested schema)
	// Specific set of actions granted by the role.
	Permissions []RolePermissionsObservation `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// (String) Unique identifier of the role. Used for assignments.
	// Unique identifier of the role. Used for assignments.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// (Number) Version of the role. A role is updated only on version increase. This field or auto_increment_version should be set.
	// Version of the role. A role is updated only on version increase. This field or `auto_increment_version` should be set.
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type RoleParameters struct {

	// (Boolean) Whether the role version should be incremented automatically on updates (and set to 1 on creation). This field or version should be set.
	// Whether the role version should be incremented automatically on updates (and set to 1 on creation). This field or `version` should be set.
	// +kubebuilder:validation:Optional
	AutoIncrementVersion *bool `json:"autoIncrementVersion,omitempty" tf:"auto_increment_version,omitempty"`

	// (String) Description of the role.
	// Description of the role.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Display name of the role. Available with Grafana 8.5+.
	// Display name of the role. Available with Grafana 8.5+.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Boolean) Boolean to state whether the role is available across all organizations or not. Defaults to false.
	// Boolean to state whether the role is available across all organizations or not. Defaults to `false`.
	// +kubebuilder:validation:Optional
	Global *bool `json:"global,omitempty" tf:"global,omitempty"`

	// (String) Group of the role. Available with Grafana 8.5+.
	// Group of the role. Available with Grafana 8.5+.
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// (Boolean) Boolean to state whether the role should be visible in the Grafana UI or not. Available with Grafana 8.5+. Defaults to false.
	// Boolean to state whether the role should be visible in the Grafana UI or not. Available with Grafana 8.5+. Defaults to `false`.
	// +kubebuilder:validation:Optional
	Hidden *bool `json:"hidden,omitempty" tf:"hidden,omitempty"`

	// (String) Name of the role
	// Name of the role
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

	// (Block Set) Specific set of actions granted by the role. (see below for nested schema)
	// Specific set of actions granted by the role.
	// +kubebuilder:validation:Optional
	Permissions []RolePermissionsParameters `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// (String) Unique identifier of the role. Used for assignments.
	// Unique identifier of the role. Used for assignments.
	// +kubebuilder:validation:Optional
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// (Number) Version of the role. A role is updated only on version increase. This field or auto_increment_version should be set.
	// Version of the role. A role is updated only on version increase. This field or `auto_increment_version` should be set.
	// +kubebuilder:validation:Optional
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type RolePermissionsInitParameters struct {

	// (String) Specific action users granted with the role will be allowed to perform (for example: users:read)
	// Specific action users granted with the role will be allowed to perform (for example: `users:read`)
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// (String) Scope to restrict the action to a set of resources (for example: users:* or roles:customrole1) Defaults to “.
	// Scope to restrict the action to a set of resources (for example: `users:*` or `roles:customrole1`) Defaults to “.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type RolePermissionsObservation struct {

	// (String) Specific action users granted with the role will be allowed to perform (for example: users:read)
	// Specific action users granted with the role will be allowed to perform (for example: `users:read`)
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// (String) Scope to restrict the action to a set of resources (for example: users:* or roles:customrole1) Defaults to “.
	// Scope to restrict the action to a set of resources (for example: `users:*` or `roles:customrole1`) Defaults to “.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type RolePermissionsParameters struct {

	// (String) Specific action users granted with the role will be allowed to perform (for example: users:read)
	// Specific action users granted with the role will be allowed to perform (for example: `users:read`)
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// (String) Scope to restrict the action to a set of resources (for example: users:* or roles:customrole1) Defaults to “.
	// Scope to restrict the action to a set of resources (for example: `users:*` or `roles:customrole1`) Defaults to “.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

// RoleSpec defines the desired state of Role
type RoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleParameters `json:"forProvider"`
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
	InitProvider RoleInitParameters `json:"initProvider,omitempty"`
}

// RoleStatus defines the observed state of Role.
type RoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Role is the Schema for the Roles API. Note: This resource is available only with Grafana Enterprise 8.+. Official documentation https://grafana.com/docs/grafana/latest/administration/roles-and-permissions/access-control/HTTP API https://grafana.com/docs/grafana/latest/developers/http_api/access_control/
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type Role struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   RoleSpec   `json:"spec"`
	Status RoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleList contains a list of Roles
type RoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Role `json:"items"`
}

// Repository type metadata.
var (
	Role_Kind             = "Role"
	Role_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Role_Kind}.String()
	Role_KindAPIVersion   = Role_Kind + "." + CRDGroupVersion.String()
	Role_GroupVersionKind = CRDGroupVersion.WithKind(Role_Kind)
)

func init() {
	SchemeBuilder.Register(&Role{}, &RoleList{})
}
