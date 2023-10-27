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

	// (String) The name of the message template.
	// The name of the message template.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The content of the message template.
	// The content of the message template.
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

type MessageTemplateObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the message template.
	// The name of the message template.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The content of the message template.
	// The content of the message template.
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

type MessageTemplateParameters struct {

	// (String) The name of the message template.
	// The name of the message template.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The content of the message template.
	// The content of the message template.
	// +kubebuilder:validation:Optional
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

// MessageTemplateSpec defines the desired state of MessageTemplate
type MessageTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MessageTemplateParameters `json:"forProvider"`
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
	InitProvider MessageTemplateInitParameters `json:"initProvider,omitempty"`
}

// MessageTemplateStatus defines the observed state of MessageTemplate.
type MessageTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MessageTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MessageTemplate is the Schema for the MessageTemplates API. Manages Grafana Alerting message templates. Official documentation https://grafana.com/docs/grafana/latest/alerting/manage-notifications/template-notifications/create-notification-templates/HTTP API https://grafana.com/docs/grafana/next/developers/http_api/alerting_provisioning/#templates This resource requires Grafana 9.1.0 or later.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type MessageTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.template) || has(self.initProvider.template)",message="template is a required parameter"
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
