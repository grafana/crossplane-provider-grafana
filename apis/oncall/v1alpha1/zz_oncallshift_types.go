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

type OnCallShiftInitParameters struct {

	// (Set of String) This parameter takes a list of days in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	// This parameter takes a list of days in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	// +listType=set
	ByDay []*string `json:"byDay,omitempty" tf:"by_day,omitempty"`

	// (Set of Number) This parameter takes a list of months. Valid values are 1 to 12
	// This parameter takes a list of months. Valid values are 1 to 12
	// +listType=set
	ByMonth []*float64 `json:"byMonth,omitempty" tf:"by_month,omitempty"`

	// 31 to -1
	// This parameter takes a list of days of the month.  Valid values are 1 to 31 or -31 to -1
	// +listType=set
	ByMonthday []*float64 `json:"byMonthday,omitempty" tf:"by_monthday,omitempty"`

	// (Number) The duration of the event.
	// The duration of the event.
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// (String) The frequency of the event. Can be hourly, daily, weekly, monthly
	// The frequency of the event. Can be hourly, daily, weekly, monthly
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// (Number) The positive integer representing at which intervals the recurrence rule repeats.
	// The positive integer representing at which intervals the recurrence rule repeats.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// (Number) The priority level. The higher the value, the higher the priority.
	// The priority level. The higher the value, the higher the priority.
	Level *float64 `json:"level,omitempty" tf:"level,omitempty"`

	// (String) The shift's name.
	// The shift's name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// call users (for rolling_users event type)
	// The list of lists with on-call users (for rolling_users event type)
	RollingUsers [][]*string `json:"rollingUsers,omitempty" tf:"rolling_users,omitempty"`

	// call shift. This parameter takes a date format as yyyy-MM-dd'T'HH:mm:ss (for example "2020-09-05T08:00:00")
	// The start time of the on-call shift. This parameter takes a date format as yyyy-MM-dd'T'HH:mm:ss (for example "2020-09-05T08:00:00")
	Start *string `json:"start,omitempty" tf:"start,omitempty"`

	// call rotation starts.
	// The index of the list of users in rolling_users, from which on-call rotation starts.
	StartRotationFromUserIndex *float64 `json:"startRotationFromUserIndex,omitempty" tf:"start_rotation_from_user_index,omitempty"`

	// (String) The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the grafana_oncall_team datasource.
	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana_oncall_team` datasource.
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// (String) The shift's timezone.  Overrides schedule's timezone.
	// The shift's timezone.  Overrides schedule's timezone.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`

	// (String) The shift's type. Can be rolling_users, recurrent_event, single_event
	// The shift's type. Can be rolling_users, recurrent_event, single_event
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// call users (for single_event and recurrent_event event type).
	// The list of on-call users (for single_event and recurrent_event event type).
	// +listType=set
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`

	// (String) Start day of the week in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	// Start day of the week in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	WeekStart *string `json:"weekStart,omitempty" tf:"week_start,omitempty"`
}

type OnCallShiftObservation struct {

	// (Set of String) This parameter takes a list of days in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	// This parameter takes a list of days in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	// +listType=set
	ByDay []*string `json:"byDay,omitempty" tf:"by_day,omitempty"`

	// (Set of Number) This parameter takes a list of months. Valid values are 1 to 12
	// This parameter takes a list of months. Valid values are 1 to 12
	// +listType=set
	ByMonth []*float64 `json:"byMonth,omitempty" tf:"by_month,omitempty"`

	// 31 to -1
	// This parameter takes a list of days of the month.  Valid values are 1 to 31 or -31 to -1
	// +listType=set
	ByMonthday []*float64 `json:"byMonthday,omitempty" tf:"by_monthday,omitempty"`

	// (Number) The duration of the event.
	// The duration of the event.
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// (String) The frequency of the event. Can be hourly, daily, weekly, monthly
	// The frequency of the event. Can be hourly, daily, weekly, monthly
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Number) The positive integer representing at which intervals the recurrence rule repeats.
	// The positive integer representing at which intervals the recurrence rule repeats.
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// (Number) The priority level. The higher the value, the higher the priority.
	// The priority level. The higher the value, the higher the priority.
	Level *float64 `json:"level,omitempty" tf:"level,omitempty"`

	// (String) The shift's name.
	// The shift's name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// call users (for rolling_users event type)
	// The list of lists with on-call users (for rolling_users event type)
	RollingUsers [][]*string `json:"rollingUsers,omitempty" tf:"rolling_users,omitempty"`

	// call shift. This parameter takes a date format as yyyy-MM-dd'T'HH:mm:ss (for example "2020-09-05T08:00:00")
	// The start time of the on-call shift. This parameter takes a date format as yyyy-MM-dd'T'HH:mm:ss (for example "2020-09-05T08:00:00")
	Start *string `json:"start,omitempty" tf:"start,omitempty"`

	// call rotation starts.
	// The index of the list of users in rolling_users, from which on-call rotation starts.
	StartRotationFromUserIndex *float64 `json:"startRotationFromUserIndex,omitempty" tf:"start_rotation_from_user_index,omitempty"`

	// (String) The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the grafana_oncall_team datasource.
	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana_oncall_team` datasource.
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// (String) The shift's timezone.  Overrides schedule's timezone.
	// The shift's timezone.  Overrides schedule's timezone.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`

	// (String) The shift's type. Can be rolling_users, recurrent_event, single_event
	// The shift's type. Can be rolling_users, recurrent_event, single_event
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// call users (for single_event and recurrent_event event type).
	// The list of on-call users (for single_event and recurrent_event event type).
	// +listType=set
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`

	// (String) Start day of the week in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	// Start day of the week in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	WeekStart *string `json:"weekStart,omitempty" tf:"week_start,omitempty"`
}

type OnCallShiftParameters struct {

	// (Set of String) This parameter takes a list of days in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	// This parameter takes a list of days in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	// +kubebuilder:validation:Optional
	// +listType=set
	ByDay []*string `json:"byDay,omitempty" tf:"by_day,omitempty"`

	// (Set of Number) This parameter takes a list of months. Valid values are 1 to 12
	// This parameter takes a list of months. Valid values are 1 to 12
	// +kubebuilder:validation:Optional
	// +listType=set
	ByMonth []*float64 `json:"byMonth,omitempty" tf:"by_month,omitempty"`

	// 31 to -1
	// This parameter takes a list of days of the month.  Valid values are 1 to 31 or -31 to -1
	// +kubebuilder:validation:Optional
	// +listType=set
	ByMonthday []*float64 `json:"byMonthday,omitempty" tf:"by_monthday,omitempty"`

	// (Number) The duration of the event.
	// The duration of the event.
	// +kubebuilder:validation:Optional
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// (String) The frequency of the event. Can be hourly, daily, weekly, monthly
	// The frequency of the event. Can be hourly, daily, weekly, monthly
	// +kubebuilder:validation:Optional
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// (Number) The positive integer representing at which intervals the recurrence rule repeats.
	// The positive integer representing at which intervals the recurrence rule repeats.
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// (Number) The priority level. The higher the value, the higher the priority.
	// The priority level. The higher the value, the higher the priority.
	// +kubebuilder:validation:Optional
	Level *float64 `json:"level,omitempty" tf:"level,omitempty"`

	// (String) The shift's name.
	// The shift's name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// call users (for rolling_users event type)
	// The list of lists with on-call users (for rolling_users event type)
	// +kubebuilder:validation:Optional
	RollingUsers [][]*string `json:"rollingUsers,omitempty" tf:"rolling_users,omitempty"`

	// call shift. This parameter takes a date format as yyyy-MM-dd'T'HH:mm:ss (for example "2020-09-05T08:00:00")
	// The start time of the on-call shift. This parameter takes a date format as yyyy-MM-dd'T'HH:mm:ss (for example "2020-09-05T08:00:00")
	// +kubebuilder:validation:Optional
	Start *string `json:"start,omitempty" tf:"start,omitempty"`

	// call rotation starts.
	// The index of the list of users in rolling_users, from which on-call rotation starts.
	// +kubebuilder:validation:Optional
	StartRotationFromUserIndex *float64 `json:"startRotationFromUserIndex,omitempty" tf:"start_rotation_from_user_index,omitempty"`

	// (String) The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the grafana_oncall_team datasource.
	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana_oncall_team` datasource.
	// +kubebuilder:validation:Optional
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// (String) The shift's timezone.  Overrides schedule's timezone.
	// The shift's timezone.  Overrides schedule's timezone.
	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`

	// (String) The shift's type. Can be rolling_users, recurrent_event, single_event
	// The shift's type. Can be rolling_users, recurrent_event, single_event
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// call users (for single_event and recurrent_event event type).
	// The list of on-call users (for single_event and recurrent_event event type).
	// +kubebuilder:validation:Optional
	// +listType=set
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`

	// (String) Start day of the week in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	// Start day of the week in iCal format. Can be MO, TU, WE, TH, FR, SA, SU
	// +kubebuilder:validation:Optional
	WeekStart *string `json:"weekStart,omitempty" tf:"week_start,omitempty"`
}

// OnCallShiftSpec defines the desired state of OnCallShift
type OnCallShiftSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OnCallShiftParameters `json:"forProvider"`
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
	InitProvider OnCallShiftInitParameters `json:"initProvider,omitempty"`
}

// OnCallShiftStatus defines the observed state of OnCallShift.
type OnCallShiftStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OnCallShiftObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OnCallShift is the Schema for the OnCallShifts API. HTTP API https://grafana.com/docs/oncall/latest/oncall-api-reference/on_call_shifts/
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type OnCallShift struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.duration) || (has(self.initProvider) && has(self.initProvider.duration))",message="spec.forProvider.duration is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.start) || (has(self.initProvider) && has(self.initProvider.start))",message="spec.forProvider.start is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   OnCallShiftSpec   `json:"spec"`
	Status OnCallShiftStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OnCallShiftList contains a list of OnCallShifts
type OnCallShiftList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OnCallShift `json:"items"`
}

// Repository type metadata.
var (
	OnCallShift_Kind             = "OnCallShift"
	OnCallShift_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OnCallShift_Kind}.String()
	OnCallShift_KindAPIVersion   = OnCallShift_Kind + "." + CRDGroupVersion.String()
	OnCallShift_GroupVersionKind = CRDGroupVersion.WithKind(OnCallShift_Kind)
)

func init() {
	SchemeBuilder.Register(&OnCallShift{}, &OnCallShiftList{})
}
