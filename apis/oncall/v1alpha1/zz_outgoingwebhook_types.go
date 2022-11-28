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

type OutgoingWebhookObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OutgoingWebhookParameters struct {

	// The auth data of the webhook. Used in Authorization header instead of user/password auth.
	// +kubebuilder:validation:Optional
	AuthorizationHeader *string `json:"authorizationHeader,omitempty" tf:"authorization_header,omitempty"`

	// The data of the webhook.
	// +kubebuilder:validation:Optional
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// Forwards whole payload of the alert to the webhook's url as POST data.
	// +kubebuilder:validation:Optional
	ForwardWholePayload *bool `json:"forwardWholePayload,omitempty" tf:"forward_whole_payload,omitempty"`

	// The name of the outgoing webhook.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The auth data of the webhook. Used for Basic authentication
	// +kubebuilder:validation:Optional
	Password *string `json:"password,omitempty" tf:"password,omitempty"`

	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana_oncall_team` datasource.
	// +kubebuilder:validation:Optional
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// The webhook URL.
	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`

	// The auth data of the webhook. Used for Basic authentication.
	// +kubebuilder:validation:Optional
	User *string `json:"user,omitempty" tf:"user,omitempty"`
}

// OutgoingWebhookSpec defines the desired state of OutgoingWebhook
type OutgoingWebhookSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OutgoingWebhookParameters `json:"forProvider"`
}

// OutgoingWebhookStatus defines the observed state of OutgoingWebhook.
type OutgoingWebhookStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OutgoingWebhookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OutgoingWebhook is the Schema for the OutgoingWebhooks API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type OutgoingWebhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OutgoingWebhookSpec   `json:"spec"`
	Status            OutgoingWebhookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OutgoingWebhookList contains a list of OutgoingWebhooks
type OutgoingWebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OutgoingWebhook `json:"items"`
}

// Repository type metadata.
var (
	OutgoingWebhook_Kind             = "OutgoingWebhook"
	OutgoingWebhook_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OutgoingWebhook_Kind}.String()
	OutgoingWebhook_KindAPIVersion   = OutgoingWebhook_Kind + "." + CRDGroupVersion.String()
	OutgoingWebhook_GroupVersionKind = CRDGroupVersion.WithKind(OutgoingWebhook_Kind)
)

func init() {
	SchemeBuilder.Register(&OutgoingWebhook{}, &OutgoingWebhookList{})
}