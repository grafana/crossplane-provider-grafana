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

type LibraryPanelInitParameters struct {

	// (String) Unique ID (UID) of the folder containing the library panel.
	// Unique ID (UID) of the folder containing the library panel.
	FolderUID *string `json:"folderUid,omitempty" tf:"folder_uid,omitempty"`

	// (String) The JSON model for the library panel.
	// The JSON model for the library panel.
	ModelJSON *string `json:"modelJson,omitempty" tf:"model_json,omitempty"`

	// (String) Name of the library panel.
	// Name of the library panel.
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

	// (String) The unique identifier (UID) of a library panel uniquely identifies library panels between multiple Grafana installs. It’s automatically generated unless you specify it during library panel creation.The UID provides consistent URLs for accessing library panels and when syncing library panels between multiple Grafana installs.
	// The unique identifier (UID) of a library panel uniquely identifies library panels between multiple Grafana installs. It’s automatically generated unless you specify it during library panel creation.The UID provides consistent URLs for accessing library panels and when syncing library panels between multiple Grafana installs.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`
}

type LibraryPanelObservation struct {

	// (String) Timestamp when the library panel was created.
	// Timestamp when the library panel was created.
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	// (List of Number) Numerical IDs of Grafana dashboards containing the library panel.
	// Numerical IDs of Grafana dashboards containing the library panel.
	DashboardIds []*float64 `json:"dashboardIds,omitempty" tf:"dashboard_ids,omitempty"`

	// (String) Description of the library panel.
	// Description of the library panel.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Name of the folder containing the library panel.
	// Name of the folder containing the library panel.
	FolderName *string `json:"folderName,omitempty" tf:"folder_name,omitempty"`

	// (String) Unique ID (UID) of the folder containing the library panel.
	// Unique ID (UID) of the folder containing the library panel.
	FolderUID *string `json:"folderUid,omitempty" tf:"folder_uid,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The JSON model for the library panel.
	// The JSON model for the library panel.
	ModelJSON *string `json:"modelJson,omitempty" tf:"model_json,omitempty"`

	// (String) Name of the library panel.
	// Name of the library panel.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The Organization ID. If not set, the Org ID defined in the provider block will be used.
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// (Number) The numeric ID of the library panel computed by Grafana.
	// The numeric ID of the library panel computed by Grafana.
	PanelID *float64 `json:"panelId,omitempty" tf:"panel_id,omitempty"`

	// (String) Type of the library panel (eg. text).
	// Type of the library panel (eg. text).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) The unique identifier (UID) of a library panel uniquely identifies library panels between multiple Grafana installs. It’s automatically generated unless you specify it during library panel creation.The UID provides consistent URLs for accessing library panels and when syncing library panels between multiple Grafana installs.
	// The unique identifier (UID) of a library panel uniquely identifies library panels between multiple Grafana installs. It’s automatically generated unless you specify it during library panel creation.The UID provides consistent URLs for accessing library panels and when syncing library panels between multiple Grafana installs.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// (String) Timestamp when the library panel was last modified.
	// Timestamp when the library panel was last modified.
	Updated *string `json:"updated,omitempty" tf:"updated,omitempty"`

	// (Number) Version of the library panel.
	// Version of the library panel.
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type LibraryPanelParameters struct {

	// (String) Unique ID (UID) of the folder containing the library panel.
	// Unique ID (UID) of the folder containing the library panel.
	// +kubebuilder:validation:Optional
	FolderUID *string `json:"folderUid,omitempty" tf:"folder_uid,omitempty"`

	// (String) The JSON model for the library panel.
	// The JSON model for the library panel.
	// +kubebuilder:validation:Optional
	ModelJSON *string `json:"modelJson,omitempty" tf:"model_json,omitempty"`

	// (String) Name of the library panel.
	// Name of the library panel.
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

	// (String) The unique identifier (UID) of a library panel uniquely identifies library panels between multiple Grafana installs. It’s automatically generated unless you specify it during library panel creation.The UID provides consistent URLs for accessing library panels and when syncing library panels between multiple Grafana installs.
	// The unique identifier (UID) of a library panel uniquely identifies library panels between multiple Grafana installs. It’s automatically generated unless you specify it during library panel creation.The UID provides consistent URLs for accessing library panels and when syncing library panels between multiple Grafana installs.
	// +kubebuilder:validation:Optional
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`
}

// LibraryPanelSpec defines the desired state of LibraryPanel
type LibraryPanelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LibraryPanelParameters `json:"forProvider"`
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
	InitProvider LibraryPanelInitParameters `json:"initProvider,omitempty"`
}

// LibraryPanelStatus defines the observed state of LibraryPanel.
type LibraryPanelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LibraryPanelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LibraryPanel is the Schema for the LibraryPanels API. Manages Grafana library panels. Official documentation https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/manage-library-panels/HTTP API https://grafana.com/docs/grafana/latest/developers/http_api/library_element/
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type LibraryPanel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.modelJson) || (has(self.initProvider) && has(self.initProvider.modelJson))",message="spec.forProvider.modelJson is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   LibraryPanelSpec   `json:"spec"`
	Status LibraryPanelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LibraryPanelList contains a list of LibraryPanels
type LibraryPanelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LibraryPanel `json:"items"`
}

// Repository type metadata.
var (
	LibraryPanel_Kind             = "LibraryPanel"
	LibraryPanel_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LibraryPanel_Kind}.String()
	LibraryPanel_KindAPIVersion   = LibraryPanel_Kind + "." + CRDGroupVersion.String()
	LibraryPanel_GroupVersionKind = CRDGroupVersion.WithKind(LibraryPanel_Kind)
)

func init() {
	SchemeBuilder.Register(&LibraryPanel{}, &LibraryPanelList{})
}
