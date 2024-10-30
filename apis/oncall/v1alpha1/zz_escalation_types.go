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

type EscalationInitParameters struct {

	// (String) The ID of an Action for trigger_webhook type step.
	// The ID of an Action for trigger_webhook type step.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oncall/v1alpha1.OutgoingWebhook
	// +crossplane:generate:reference:refFieldName=ActionToTriggerRef
	// +crossplane:generate:reference:selectorFieldName=ActionToTriggerSelector
	ActionToTrigger *string `json:"actionToTrigger,omitempty" tf:"action_to_trigger,omitempty"`

	// Reference to a OutgoingWebhook in oncall to populate actionToTrigger.
	// +kubebuilder:validation:Optional
	ActionToTriggerRef *v1.Reference `json:"actionToTriggerRef,omitempty" tf:"-"`

	// Selector for a OutgoingWebhook in oncall to populate actionToTrigger.
	// +kubebuilder:validation:Optional
	ActionToTriggerSelector *v1.Selector `json:"actionToTriggerSelector,omitempty" tf:"-"`

	// 86400) seconds
	// The duration of delay for wait type step. (60-86400) seconds
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// (String) The ID of the escalation chain.
	// The ID of the escalation chain.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oncall/v1alpha1.EscalationChain
	// +crossplane:generate:reference:refFieldName=EscalationChainRef
	// +crossplane:generate:reference:selectorFieldName=EscalationChainSelector
	EscalationChainID *string `json:"escalationChainId,omitempty" tf:"escalation_chain_id,omitempty"`

	// Reference to a EscalationChain in oncall to populate escalationChainId.
	// +kubebuilder:validation:Optional
	EscalationChainRef *v1.Reference `json:"escalationChainRef,omitempty" tf:"-"`

	// Selector for a EscalationChain in oncall to populate escalationChainId.
	// +kubebuilder:validation:Optional
	EscalationChainSelector *v1.Selector `json:"escalationChainSelector,omitempty" tf:"-"`

	// (String) The ID of a User Group for notify_user_group type step.
	// The ID of a User Group for notify_user_group type step.
	GroupToNotify *string `json:"groupToNotify,omitempty" tf:"group_to_notify,omitempty"`

	// (Boolean) Will activate "important" personal notification rules. Actual for steps: notify_persons, notify_on_call_from_schedule and notify_user_group,notify_team_members
	// Will activate "important" personal notification rules. Actual for steps: notify_persons, notify_on_call_from_schedule and notify_user_group,notify_team_members
	Important *bool `json:"important,omitempty" tf:"important,omitempty"`

	// (String) The beginning of the time interval for notify_if_time_from_to type step in UTC (for example 08:00:00Z).
	// The beginning of the time interval for notify_if_time_from_to type step in UTC (for example 08:00:00Z).
	NotifyIfTimeFrom *string `json:"notifyIfTimeFrom,omitempty" tf:"notify_if_time_from,omitempty"`

	// (String) The end of the time interval for notify_if_time_from_to type step in UTC (for example 18:00:00Z).
	// The end of the time interval for notify_if_time_from_to type step in UTC (for example 18:00:00Z).
	NotifyIfTimeTo *string `json:"notifyIfTimeTo,omitempty" tf:"notify_if_time_to,omitempty"`

	// (String) ID of a Schedule for notify_on_call_from_schedule type step.
	// ID of a Schedule for notify_on_call_from_schedule type step.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oncall/v1alpha1.Schedule
	// +crossplane:generate:reference:refFieldName=NotifyOnCallFromScheduleRef
	// +crossplane:generate:reference:selectorFieldName=NotifyOnCallFromScheduleSelector
	NotifyOnCallFromSchedule *string `json:"notifyOnCallFromSchedule,omitempty" tf:"notify_on_call_from_schedule,omitempty"`

	// Reference to a Schedule in oncall to populate notifyOnCallFromSchedule.
	// +kubebuilder:validation:Optional
	NotifyOnCallFromScheduleRef *v1.Reference `json:"notifyOnCallFromScheduleRef,omitempty" tf:"-"`

	// Selector for a Schedule in oncall to populate notifyOnCallFromSchedule.
	// +kubebuilder:validation:Optional
	NotifyOnCallFromScheduleSelector *v1.Selector `json:"notifyOnCallFromScheduleSelector,omitempty" tf:"-"`

	// (String) The ID of a Team for a notify_team_members type step.
	// The ID of a Team for a notify_team_members type step.
	NotifyToTeamMembers *string `json:"notifyToTeamMembers,omitempty" tf:"notify_to_team_members,omitempty"`

	// (Set of String) The list of ID's of users for notify_persons type step.
	// The list of ID's of users for notify_persons type step.
	// +listType=set
	PersonsToNotify []*string `json:"personsToNotify,omitempty" tf:"persons_to_notify,omitempty"`

	// (Set of String) The list of ID's of users for notify_person_next_each_time type step.
	// The list of ID's of users for notify_person_next_each_time type step.
	// +listType=set
	PersonsToNotifyNextEachTime []*string `json:"personsToNotifyNextEachTime,omitempty" tf:"persons_to_notify_next_each_time,omitempty"`

	// (Number) The position of the escalation step (starts from 0).
	// The position of the escalation step (starts from 0).
	Position *float64 `json:"position,omitempty" tf:"position,omitempty"`

	// (String) The severity of the incident for declare_incident type step.
	// The severity of the incident for declare_incident type step.
	Severity *string `json:"severity,omitempty" tf:"severity,omitempty"`

	// (String) The type of escalation policy. Can be wait, notify_persons, notify_person_next_each_time, notify_on_call_from_schedule, trigger_webhook, notify_user_group, resolve, notify_whole_channel, notify_if_time_from_to, repeat_escalation, notify_team_members, declare_incident
	// The type of escalation policy. Can be wait, notify_persons, notify_person_next_each_time, notify_on_call_from_schedule, trigger_webhook, notify_user_group, resolve, notify_whole_channel, notify_if_time_from_to, repeat_escalation, notify_team_members, declare_incident
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EscalationObservation struct {

	// (String) The ID of an Action for trigger_webhook type step.
	// The ID of an Action for trigger_webhook type step.
	ActionToTrigger *string `json:"actionToTrigger,omitempty" tf:"action_to_trigger,omitempty"`

	// 86400) seconds
	// The duration of delay for wait type step. (60-86400) seconds
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// (String) The ID of the escalation chain.
	// The ID of the escalation chain.
	EscalationChainID *string `json:"escalationChainId,omitempty" tf:"escalation_chain_id,omitempty"`

	// (String) The ID of a User Group for notify_user_group type step.
	// The ID of a User Group for notify_user_group type step.
	GroupToNotify *string `json:"groupToNotify,omitempty" tf:"group_to_notify,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) Will activate "important" personal notification rules. Actual for steps: notify_persons, notify_on_call_from_schedule and notify_user_group,notify_team_members
	// Will activate "important" personal notification rules. Actual for steps: notify_persons, notify_on_call_from_schedule and notify_user_group,notify_team_members
	Important *bool `json:"important,omitempty" tf:"important,omitempty"`

	// (String) The beginning of the time interval for notify_if_time_from_to type step in UTC (for example 08:00:00Z).
	// The beginning of the time interval for notify_if_time_from_to type step in UTC (for example 08:00:00Z).
	NotifyIfTimeFrom *string `json:"notifyIfTimeFrom,omitempty" tf:"notify_if_time_from,omitempty"`

	// (String) The end of the time interval for notify_if_time_from_to type step in UTC (for example 18:00:00Z).
	// The end of the time interval for notify_if_time_from_to type step in UTC (for example 18:00:00Z).
	NotifyIfTimeTo *string `json:"notifyIfTimeTo,omitempty" tf:"notify_if_time_to,omitempty"`

	// (String) ID of a Schedule for notify_on_call_from_schedule type step.
	// ID of a Schedule for notify_on_call_from_schedule type step.
	NotifyOnCallFromSchedule *string `json:"notifyOnCallFromSchedule,omitempty" tf:"notify_on_call_from_schedule,omitempty"`

	// (String) The ID of a Team for a notify_team_members type step.
	// The ID of a Team for a notify_team_members type step.
	NotifyToTeamMembers *string `json:"notifyToTeamMembers,omitempty" tf:"notify_to_team_members,omitempty"`

	// (Set of String) The list of ID's of users for notify_persons type step.
	// The list of ID's of users for notify_persons type step.
	// +listType=set
	PersonsToNotify []*string `json:"personsToNotify,omitempty" tf:"persons_to_notify,omitempty"`

	// (Set of String) The list of ID's of users for notify_person_next_each_time type step.
	// The list of ID's of users for notify_person_next_each_time type step.
	// +listType=set
	PersonsToNotifyNextEachTime []*string `json:"personsToNotifyNextEachTime,omitempty" tf:"persons_to_notify_next_each_time,omitempty"`

	// (Number) The position of the escalation step (starts from 0).
	// The position of the escalation step (starts from 0).
	Position *float64 `json:"position,omitempty" tf:"position,omitempty"`

	// (String) The severity of the incident for declare_incident type step.
	// The severity of the incident for declare_incident type step.
	Severity *string `json:"severity,omitempty" tf:"severity,omitempty"`

	// (String) The type of escalation policy. Can be wait, notify_persons, notify_person_next_each_time, notify_on_call_from_schedule, trigger_webhook, notify_user_group, resolve, notify_whole_channel, notify_if_time_from_to, repeat_escalation, notify_team_members, declare_incident
	// The type of escalation policy. Can be wait, notify_persons, notify_person_next_each_time, notify_on_call_from_schedule, trigger_webhook, notify_user_group, resolve, notify_whole_channel, notify_if_time_from_to, repeat_escalation, notify_team_members, declare_incident
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EscalationParameters struct {

	// (String) The ID of an Action for trigger_webhook type step.
	// The ID of an Action for trigger_webhook type step.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oncall/v1alpha1.OutgoingWebhook
	// +crossplane:generate:reference:refFieldName=ActionToTriggerRef
	// +crossplane:generate:reference:selectorFieldName=ActionToTriggerSelector
	// +kubebuilder:validation:Optional
	ActionToTrigger *string `json:"actionToTrigger,omitempty" tf:"action_to_trigger,omitempty"`

	// Reference to a OutgoingWebhook in oncall to populate actionToTrigger.
	// +kubebuilder:validation:Optional
	ActionToTriggerRef *v1.Reference `json:"actionToTriggerRef,omitempty" tf:"-"`

	// Selector for a OutgoingWebhook in oncall to populate actionToTrigger.
	// +kubebuilder:validation:Optional
	ActionToTriggerSelector *v1.Selector `json:"actionToTriggerSelector,omitempty" tf:"-"`

	// 86400) seconds
	// The duration of delay for wait type step. (60-86400) seconds
	// +kubebuilder:validation:Optional
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// (String) The ID of the escalation chain.
	// The ID of the escalation chain.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oncall/v1alpha1.EscalationChain
	// +crossplane:generate:reference:refFieldName=EscalationChainRef
	// +crossplane:generate:reference:selectorFieldName=EscalationChainSelector
	// +kubebuilder:validation:Optional
	EscalationChainID *string `json:"escalationChainId,omitempty" tf:"escalation_chain_id,omitempty"`

	// Reference to a EscalationChain in oncall to populate escalationChainId.
	// +kubebuilder:validation:Optional
	EscalationChainRef *v1.Reference `json:"escalationChainRef,omitempty" tf:"-"`

	// Selector for a EscalationChain in oncall to populate escalationChainId.
	// +kubebuilder:validation:Optional
	EscalationChainSelector *v1.Selector `json:"escalationChainSelector,omitempty" tf:"-"`

	// (String) The ID of a User Group for notify_user_group type step.
	// The ID of a User Group for notify_user_group type step.
	// +kubebuilder:validation:Optional
	GroupToNotify *string `json:"groupToNotify,omitempty" tf:"group_to_notify,omitempty"`

	// (Boolean) Will activate "important" personal notification rules. Actual for steps: notify_persons, notify_on_call_from_schedule and notify_user_group,notify_team_members
	// Will activate "important" personal notification rules. Actual for steps: notify_persons, notify_on_call_from_schedule and notify_user_group,notify_team_members
	// +kubebuilder:validation:Optional
	Important *bool `json:"important,omitempty" tf:"important,omitempty"`

	// (String) The beginning of the time interval for notify_if_time_from_to type step in UTC (for example 08:00:00Z).
	// The beginning of the time interval for notify_if_time_from_to type step in UTC (for example 08:00:00Z).
	// +kubebuilder:validation:Optional
	NotifyIfTimeFrom *string `json:"notifyIfTimeFrom,omitempty" tf:"notify_if_time_from,omitempty"`

	// (String) The end of the time interval for notify_if_time_from_to type step in UTC (for example 18:00:00Z).
	// The end of the time interval for notify_if_time_from_to type step in UTC (for example 18:00:00Z).
	// +kubebuilder:validation:Optional
	NotifyIfTimeTo *string `json:"notifyIfTimeTo,omitempty" tf:"notify_if_time_to,omitempty"`

	// (String) ID of a Schedule for notify_on_call_from_schedule type step.
	// ID of a Schedule for notify_on_call_from_schedule type step.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oncall/v1alpha1.Schedule
	// +crossplane:generate:reference:refFieldName=NotifyOnCallFromScheduleRef
	// +crossplane:generate:reference:selectorFieldName=NotifyOnCallFromScheduleSelector
	// +kubebuilder:validation:Optional
	NotifyOnCallFromSchedule *string `json:"notifyOnCallFromSchedule,omitempty" tf:"notify_on_call_from_schedule,omitempty"`

	// Reference to a Schedule in oncall to populate notifyOnCallFromSchedule.
	// +kubebuilder:validation:Optional
	NotifyOnCallFromScheduleRef *v1.Reference `json:"notifyOnCallFromScheduleRef,omitempty" tf:"-"`

	// Selector for a Schedule in oncall to populate notifyOnCallFromSchedule.
	// +kubebuilder:validation:Optional
	NotifyOnCallFromScheduleSelector *v1.Selector `json:"notifyOnCallFromScheduleSelector,omitempty" tf:"-"`

	// (String) The ID of a Team for a notify_team_members type step.
	// The ID of a Team for a notify_team_members type step.
	// +kubebuilder:validation:Optional
	NotifyToTeamMembers *string `json:"notifyToTeamMembers,omitempty" tf:"notify_to_team_members,omitempty"`

	// (Set of String) The list of ID's of users for notify_persons type step.
	// The list of ID's of users for notify_persons type step.
	// +kubebuilder:validation:Optional
	// +listType=set
	PersonsToNotify []*string `json:"personsToNotify,omitempty" tf:"persons_to_notify,omitempty"`

	// (Set of String) The list of ID's of users for notify_person_next_each_time type step.
	// The list of ID's of users for notify_person_next_each_time type step.
	// +kubebuilder:validation:Optional
	// +listType=set
	PersonsToNotifyNextEachTime []*string `json:"personsToNotifyNextEachTime,omitempty" tf:"persons_to_notify_next_each_time,omitempty"`

	// (Number) The position of the escalation step (starts from 0).
	// The position of the escalation step (starts from 0).
	// +kubebuilder:validation:Optional
	Position *float64 `json:"position,omitempty" tf:"position,omitempty"`

	// (String) The severity of the incident for declare_incident type step.
	// The severity of the incident for declare_incident type step.
	// +kubebuilder:validation:Optional
	Severity *string `json:"severity,omitempty" tf:"severity,omitempty"`

	// (String) The type of escalation policy. Can be wait, notify_persons, notify_person_next_each_time, notify_on_call_from_schedule, trigger_webhook, notify_user_group, resolve, notify_whole_channel, notify_if_time_from_to, repeat_escalation, notify_team_members, declare_incident
	// The type of escalation policy. Can be wait, notify_persons, notify_person_next_each_time, notify_on_call_from_schedule, trigger_webhook, notify_user_group, resolve, notify_whole_channel, notify_if_time_from_to, repeat_escalation, notify_team_members, declare_incident
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// EscalationSpec defines the desired state of Escalation
type EscalationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EscalationParameters `json:"forProvider"`
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
	InitProvider EscalationInitParameters `json:"initProvider,omitempty"`
}

// EscalationStatus defines the observed state of Escalation.
type EscalationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EscalationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Escalation is the Schema for the Escalations API. Official documentation https://grafana.com/docs/oncall/latest/configure/escalation-chains-and-routes/HTTP API https://grafana.com/docs/oncall/latest/oncall-api-reference/escalation_policies/
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type Escalation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.position) || (has(self.initProvider) && has(self.initProvider.position))",message="spec.forProvider.position is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   EscalationSpec   `json:"spec"`
	Status EscalationStatus `json:"status,omitempty"`
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
