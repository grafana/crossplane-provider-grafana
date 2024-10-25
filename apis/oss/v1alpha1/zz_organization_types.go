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

type OrganizationInitParameters struct {

	// (String) The login name of the configured default admin user for the Grafana
	// installation. If unset, this value defaults to admin, the Grafana default.
	// Defaults to admin.
	// The login name of the configured default admin user for the Grafana
	// installation. If unset, this value defaults to admin, the Grafana default.
	// Defaults to `admin`.
	AdminUser *string `json:"adminUser,omitempty" tf:"admin_user,omitempty"`

	// (Set of String) A list of email addresses corresponding to users who should be given admin
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	// A list of email addresses corresponding to users who should be given admin
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	// +listType=set
	Admins []*string `json:"admins,omitempty" tf:"admins,omitempty"`

	// (Boolean) Whether or not to create Grafana users specified in the organization's
	// membership if they don't already exist in Grafana. If unspecified, this
	// parameter defaults to true, creating placeholder users with the name, login,
	// and email set to the email of the user, and a random password. Setting this
	// option to false will cause an error to be thrown for any users that do not
	// already exist in Grafana.
	// Defaults to true.
	// Whether or not to create Grafana users specified in the organization's
	// membership if they don't already exist in Grafana. If unspecified, this
	// parameter defaults to true, creating placeholder users with the name, login,
	// and email set to the email of the user, and a random password. Setting this
	// option to false will cause an error to be thrown for any users that do not
	// already exist in Grafana.
	// Defaults to `true`.
	CreateUsers *bool `json:"createUsers,omitempty" tf:"create_users,omitempty"`

	// (Set of String) A list of email addresses corresponding to users who should be given editor
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	// A list of email addresses corresponding to users who should be given editor
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	// +listType=set
	Editors []*string `json:"editors,omitempty" tf:"editors,omitempty"`

	// (String) The display name for the Grafana organization created.
	// The display name for the Grafana organization created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) A list of email addresses corresponding to users who should be given none access to the organization.
	// Note: users specified here must already exist in Grafana, unless 'create_users' is
	// set to true. This feature is only available in Grafana 10.2+.
	// A list of email addresses corresponding to users who should be given none access to the organization.
	// Note: users specified here must already exist in Grafana, unless 'create_users' is
	// set to true. This feature is only available in Grafana 10.2+.
	// +listType=set
	UsersWithoutAccess []*string `json:"usersWithoutAccess,omitempty" tf:"users_without_access,omitempty"`

	// (Set of String) A list of email addresses corresponding to users who should be given viewer
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	// A list of email addresses corresponding to users who should be given viewer
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	// +listType=set
	Viewers []*string `json:"viewers,omitempty" tf:"viewers,omitempty"`
}

type OrganizationObservation struct {

	// (String) The login name of the configured default admin user for the Grafana
	// installation. If unset, this value defaults to admin, the Grafana default.
	// Defaults to admin.
	// The login name of the configured default admin user for the Grafana
	// installation. If unset, this value defaults to admin, the Grafana default.
	// Defaults to `admin`.
	AdminUser *string `json:"adminUser,omitempty" tf:"admin_user,omitempty"`

	// (Set of String) A list of email addresses corresponding to users who should be given admin
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	// A list of email addresses corresponding to users who should be given admin
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	// +listType=set
	Admins []*string `json:"admins,omitempty" tf:"admins,omitempty"`

	// (Boolean) Whether or not to create Grafana users specified in the organization's
	// membership if they don't already exist in Grafana. If unspecified, this
	// parameter defaults to true, creating placeholder users with the name, login,
	// and email set to the email of the user, and a random password. Setting this
	// option to false will cause an error to be thrown for any users that do not
	// already exist in Grafana.
	// Defaults to true.
	// Whether or not to create Grafana users specified in the organization's
	// membership if they don't already exist in Grafana. If unspecified, this
	// parameter defaults to true, creating placeholder users with the name, login,
	// and email set to the email of the user, and a random password. Setting this
	// option to false will cause an error to be thrown for any users that do not
	// already exist in Grafana.
	// Defaults to `true`.
	CreateUsers *bool `json:"createUsers,omitempty" tf:"create_users,omitempty"`

	// (Set of String) A list of email addresses corresponding to users who should be given editor
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	// A list of email addresses corresponding to users who should be given editor
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	// +listType=set
	Editors []*string `json:"editors,omitempty" tf:"editors,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The display name for the Grafana organization created.
	// The display name for the Grafana organization created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number) The organization id assigned to this organization by Grafana.
	// The organization id assigned to this organization by Grafana.
	OrgID *float64 `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// (Set of String) A list of email addresses corresponding to users who should be given none access to the organization.
	// Note: users specified here must already exist in Grafana, unless 'create_users' is
	// set to true. This feature is only available in Grafana 10.2+.
	// A list of email addresses corresponding to users who should be given none access to the organization.
	// Note: users specified here must already exist in Grafana, unless 'create_users' is
	// set to true. This feature is only available in Grafana 10.2+.
	// +listType=set
	UsersWithoutAccess []*string `json:"usersWithoutAccess,omitempty" tf:"users_without_access,omitempty"`

	// (Set of String) A list of email addresses corresponding to users who should be given viewer
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	// A list of email addresses corresponding to users who should be given viewer
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	// +listType=set
	Viewers []*string `json:"viewers,omitempty" tf:"viewers,omitempty"`
}

type OrganizationParameters struct {

	// (String) The login name of the configured default admin user for the Grafana
	// installation. If unset, this value defaults to admin, the Grafana default.
	// Defaults to admin.
	// The login name of the configured default admin user for the Grafana
	// installation. If unset, this value defaults to admin, the Grafana default.
	// Defaults to `admin`.
	// +kubebuilder:validation:Optional
	AdminUser *string `json:"adminUser,omitempty" tf:"admin_user,omitempty"`

	// (Set of String) A list of email addresses corresponding to users who should be given admin
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	// A list of email addresses corresponding to users who should be given admin
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	// +kubebuilder:validation:Optional
	// +listType=set
	Admins []*string `json:"admins,omitempty" tf:"admins,omitempty"`

	// (Boolean) Whether or not to create Grafana users specified in the organization's
	// membership if they don't already exist in Grafana. If unspecified, this
	// parameter defaults to true, creating placeholder users with the name, login,
	// and email set to the email of the user, and a random password. Setting this
	// option to false will cause an error to be thrown for any users that do not
	// already exist in Grafana.
	// Defaults to true.
	// Whether or not to create Grafana users specified in the organization's
	// membership if they don't already exist in Grafana. If unspecified, this
	// parameter defaults to true, creating placeholder users with the name, login,
	// and email set to the email of the user, and a random password. Setting this
	// option to false will cause an error to be thrown for any users that do not
	// already exist in Grafana.
	// Defaults to `true`.
	// +kubebuilder:validation:Optional
	CreateUsers *bool `json:"createUsers,omitempty" tf:"create_users,omitempty"`

	// (Set of String) A list of email addresses corresponding to users who should be given editor
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	// A list of email addresses corresponding to users who should be given editor
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	// +kubebuilder:validation:Optional
	// +listType=set
	Editors []*string `json:"editors,omitempty" tf:"editors,omitempty"`

	// (String) The display name for the Grafana organization created.
	// The display name for the Grafana organization created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) A list of email addresses corresponding to users who should be given none access to the organization.
	// Note: users specified here must already exist in Grafana, unless 'create_users' is
	// set to true. This feature is only available in Grafana 10.2+.
	// A list of email addresses corresponding to users who should be given none access to the organization.
	// Note: users specified here must already exist in Grafana, unless 'create_users' is
	// set to true. This feature is only available in Grafana 10.2+.
	// +kubebuilder:validation:Optional
	// +listType=set
	UsersWithoutAccess []*string `json:"usersWithoutAccess,omitempty" tf:"users_without_access,omitempty"`

	// (Set of String) A list of email addresses corresponding to users who should be given viewer
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	// A list of email addresses corresponding to users who should be given viewer
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	// +kubebuilder:validation:Optional
	// +listType=set
	Viewers []*string `json:"viewers,omitempty" tf:"viewers,omitempty"`
}

// OrganizationSpec defines the desired state of Organization
type OrganizationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationParameters `json:"forProvider"`
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
	InitProvider OrganizationInitParameters `json:"initProvider,omitempty"`
}

// OrganizationStatus defines the observed state of Organization.
type OrganizationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Organization is the Schema for the Organizations API. Official documentation https://grafana.com/docs/grafana/latest/administration/organization-management/HTTP API https://grafana.com/docs/grafana/latest/developers/http_api/org/ This resource represents an instance-scoped resource and uses Grafana's admin APIs. It does not work with API tokens or service accounts which are org-scoped. You must use basic auth. This resource is also not compatible with Grafana Cloud, as it does not allow basic auth.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type Organization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   OrganizationSpec   `json:"spec"`
	Status OrganizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationList contains a list of Organizations
type OrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Organization `json:"items"`
}

// Repository type metadata.
var (
	Organization_Kind             = "Organization"
	Organization_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Organization_Kind}.String()
	Organization_KindAPIVersion   = Organization_Kind + "." + CRDGroupVersion.String()
	Organization_GroupVersionKind = CRDGroupVersion.WithKind(Organization_Kind)
)

func init() {
	SchemeBuilder.Register(&Organization{}, &OrganizationList{})
}
