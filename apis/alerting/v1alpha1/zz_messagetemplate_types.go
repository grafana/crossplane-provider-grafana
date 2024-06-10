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

type MessageTemplateInitParameters struct {

	// Defaults to false. Defaults to `false`.
	DisableProvenance *bool `json:"disableProvenance,omitempty" tf:"disable_provenance,omitempty"`

	// (String) The name of the message template.
	// The name of the message template.
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

	// (String) The content of the message template.
	// The content of the message template.
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

type MessageTemplateObservation struct {

	// Defaults to false. Defaults to `false`.
	DisableProvenance *bool `json:"disableProvenance,omitempty" tf:"disable_provenance,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the message template.
	// The name of the message template.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The Organization ID. If not set, the Org ID defined in the provider block will be used.
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// (String) The content of the message template.
	// The content of the message template.
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

type MessageTemplateParameters struct {

	// Defaults to false. Defaults to `false`.
	// +kubebuilder:validation:Optional
	DisableProvenance *bool `json:"disableProvenance,omitempty" tf:"disable_provenance,omitempty"`

	// (String) The name of the message template.
	// The name of the message template.
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

	// (String) The content of the message template.
	// The content of the message template.
	// +kubebuilder:validation:Optional
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

// MessageTemplateSpec defines the desired state of MessageTemplate
type MessageTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MessageTemplateParameters `json:"forProvider"`
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
	InitProvider MessageTemplateInitParameters `json:"initProvider,omitempty"`
}

// MessageTemplateStatus defines the observed state of MessageTemplate.
type MessageTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MessageTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MessageTemplate is the Schema for the MessageTemplates API. Manages Grafana Alerting message templates. Official documentation https://grafana.com/docs/grafana/latest/alerting/configure-notifications/template-notifications/create-notification-templates/HTTP API https://grafana.com/docs/grafana/next/developers/http_api/alerting_provisioning/#templates This resource requires Grafana 9.1.0 or later.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type MessageTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.template) || (has(self.initProvider) && has(self.initProvider.template))",message="spec.forProvider.template is a required parameter"
	Spec   MessageTemplateSpec   `json:"spec"`
	Status MessageTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MessageTemplateList contains a list of MessageTemplates
type MessageTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MessageTemplate `json:"items"`
}

// Repository type metadata.
var (
	MessageTemplate_Kind             = "MessageTemplate"
	MessageTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MessageTemplate_Kind}.String()
	MessageTemplate_KindAPIVersion   = MessageTemplate_Kind + "." + CRDGroupVersion.String()
	MessageTemplate_GroupVersionKind = CRDGroupVersion.WithKind(MessageTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&MessageTemplate{}, &MessageTemplateList{})
}
