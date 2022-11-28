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

type EscalationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EscalationParameters struct {

	// The ID of an Action for trigger_action type step.
	// +kubebuilder:validation:Optional
	ActionToTrigger *string `json:"actionToTrigger,omitempty" tf:"action_to_trigger,omitempty"`

	// The duration of delay for wait type step.
	// +kubebuilder:validation:Optional
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// The ID of the escalation chain.
	// +kubebuilder:validation:Required
	EscalationChainID *string `json:"escalationChainId" tf:"escalation_chain_id,omitempty"`

	// The ID of a User Group for notify_user_group type step.
	// +kubebuilder:validation:Optional
	GroupToNotify *string `json:"groupToNotify,omitempty" tf:"group_to_notify,omitempty"`

	// Will activate "important" personal notification rules. Actual for steps: notify_persons, notify_on_call_from_schedule and notify_user_group
	// +kubebuilder:validation:Optional
	Important *bool `json:"important,omitempty" tf:"important,omitempty"`

	// The beginning of the time interval for notify_if_time_from_to type step in UTC (for example 08:00:00Z).
	// +kubebuilder:validation:Optional
	NotifyIfTimeFrom *string `json:"notifyIfTimeFrom,omitempty" tf:"notify_if_time_from,omitempty"`

	// The end of the time interval for notify_if_time_from_to type step in UTC (for example 18:00:00Z).
	// +kubebuilder:validation:Optional
	NotifyIfTimeTo *string `json:"notifyIfTimeTo,omitempty" tf:"notify_if_time_to,omitempty"`

	// ID of a Schedule for notify_on_call_from_schedule type step.
	// +kubebuilder:validation:Optional
	NotifyOnCallFromSchedule *string `json:"notifyOnCallFromSchedule,omitempty" tf:"notify_on_call_from_schedule,omitempty"`

	// The list of ID's of users for notify_persons type step.
	// +kubebuilder:validation:Optional
	PersonsToNotify []*string `json:"personsToNotify,omitempty" tf:"persons_to_notify,omitempty"`

	// The list of ID's of users for notify_person_next_each_time type step.
	// +kubebuilder:validation:Optional
	PersonsToNotifyNextEachTime []*string `json:"personsToNotifyNextEachTime,omitempty" tf:"persons_to_notify_next_each_time,omitempty"`

	// The position of the escalation step (starts from 0).
	// +kubebuilder:validation:Required
	Position *float64 `json:"position" tf:"position,omitempty"`

	// The type of escalation policy. Can be wait, notify_persons, notify_person_next_each_time, notify_on_call_from_schedule, trigger_action, notify_user_group, resolve, notify_whole_channel, notify_if_time_from_to, repeat_escalation
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// EscalationSpec defines the desired state of Escalation
type EscalationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EscalationParameters `json:"forProvider"`
}

// EscalationStatus defines the observed state of Escalation.
type EscalationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EscalationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Escalation is the Schema for the Escalations API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type Escalation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EscalationSpec   `json:"spec"`
	Status            EscalationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EscalationList contains a list of Escalations
type EscalationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Escalation `json:"items"`
}

// Repository type metadata.
var (
	Escalation_Kind             = "Escalation"
	Escalation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Escalation_Kind}.String()
	Escalation_KindAPIVersion   = Escalation_Kind + "." + CRDGroupVersion.String()
	Escalation_GroupVersionKind = CRDGroupVersion.WithKind(Escalation_Kind)
)

func init() {
	SchemeBuilder.Register(&Escalation{}, &EscalationList{})
}