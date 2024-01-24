// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type IntervalsInitParameters struct {

	// 31, within a month, e.g. "1" or "14:16". Negative values can be used to represent days counting from the end of a month, e.g. "-1".
	// An inclusive range of days, 1-31, within a month, e.g. "1" or "14:16". Negative values can be used to represent days counting from the end of a month, e.g. "-1".
	DaysOfMonth []*string `json:"daysOfMonth,omitempty" tf:"days_of_month,omitempty"`

	// (String) Provides the time zone for the time interval. Must be a location in the IANA time zone database, e.g "America/New_York"
	// Provides the time zone for the time interval. Must be a location in the IANA time zone database, e.g "America/New_York"
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// (List of String) An inclusive range of months, either numerical or full calendar month, e.g. "1:3", "december", or "may:august".
	// An inclusive range of months, either numerical or full calendar month, e.g. "1:3", "december", or "may:august".
	Months []*string `json:"months,omitempty" tf:"months,omitempty"`

	// (Block List) The time ranges, represented in minutes, during which to mute in a given day. (see below for nested schema)
	// The time ranges, represented in minutes, during which to mute in a given day.
	Times []TimesInitParameters `json:"times,omitempty" tf:"times,omitempty"`

	// (List of String) An inclusive range of weekdays, e.g. "monday" or "tuesday:thursday".
	// An inclusive range of weekdays, e.g. "monday" or "tuesday:thursday".
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`

	// (List of String) A positive inclusive range of years, e.g. "2030" or "2025:2026".
	// A positive inclusive range of years, e.g. "2030" or "2025:2026".
	Years []*string `json:"years,omitempty" tf:"years,omitempty"`
}

type IntervalsObservation struct {

	// 31, within a month, e.g. "1" or "14:16". Negative values can be used to represent days counting from the end of a month, e.g. "-1".
	// An inclusive range of days, 1-31, within a month, e.g. "1" or "14:16". Negative values can be used to represent days counting from the end of a month, e.g. "-1".
	DaysOfMonth []*string `json:"daysOfMonth,omitempty" tf:"days_of_month,omitempty"`

	// (String) Provides the time zone for the time interval. Must be a location in the IANA time zone database, e.g "America/New_York"
	// Provides the time zone for the time interval. Must be a location in the IANA time zone database, e.g "America/New_York"
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// (List of String) An inclusive range of months, either numerical or full calendar month, e.g. "1:3", "december", or "may:august".
	// An inclusive range of months, either numerical or full calendar month, e.g. "1:3", "december", or "may:august".
	Months []*string `json:"months,omitempty" tf:"months,omitempty"`

	// (Block List) The time ranges, represented in minutes, during which to mute in a given day. (see below for nested schema)
	// The time ranges, represented in minutes, during which to mute in a given day.
	Times []TimesObservation `json:"times,omitempty" tf:"times,omitempty"`

	// (List of String) An inclusive range of weekdays, e.g. "monday" or "tuesday:thursday".
	// An inclusive range of weekdays, e.g. "monday" or "tuesday:thursday".
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`

	// (List of String) A positive inclusive range of years, e.g. "2030" or "2025:2026".
	// A positive inclusive range of years, e.g. "2030" or "2025:2026".
	Years []*string `json:"years,omitempty" tf:"years,omitempty"`
}

type IntervalsParameters struct {

	// 31, within a month, e.g. "1" or "14:16". Negative values can be used to represent days counting from the end of a month, e.g. "-1".
	// An inclusive range of days, 1-31, within a month, e.g. "1" or "14:16". Negative values can be used to represent days counting from the end of a month, e.g. "-1".
	// +kubebuilder:validation:Optional
	DaysOfMonth []*string `json:"daysOfMonth,omitempty" tf:"days_of_month,omitempty"`

	// (String) Provides the time zone for the time interval. Must be a location in the IANA time zone database, e.g "America/New_York"
	// Provides the time zone for the time interval. Must be a location in the IANA time zone database, e.g "America/New_York"
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// (List of String) An inclusive range of months, either numerical or full calendar month, e.g. "1:3", "december", or "may:august".
	// An inclusive range of months, either numerical or full calendar month, e.g. "1:3", "december", or "may:august".
	// +kubebuilder:validation:Optional
	Months []*string `json:"months,omitempty" tf:"months,omitempty"`

	// (Block List) The time ranges, represented in minutes, during which to mute in a given day. (see below for nested schema)
	// The time ranges, represented in minutes, during which to mute in a given day.
	// +kubebuilder:validation:Optional
	Times []TimesParameters `json:"times,omitempty" tf:"times,omitempty"`

	// (List of String) An inclusive range of weekdays, e.g. "monday" or "tuesday:thursday".
	// An inclusive range of weekdays, e.g. "monday" or "tuesday:thursday".
	// +kubebuilder:validation:Optional
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`

	// (List of String) A positive inclusive range of years, e.g. "2030" or "2025:2026".
	// A positive inclusive range of years, e.g. "2030" or "2025:2026".
	// +kubebuilder:validation:Optional
	Years []*string `json:"years,omitempty" tf:"years,omitempty"`
}

type MuteTimingInitParameters struct {

	// (Block List) The time intervals at which to mute notifications. Use an empty block to mute all the time. (see below for nested schema)
	// The time intervals at which to mute notifications. Use an empty block to mute all the time.
	Intervals []IntervalsInitParameters `json:"intervals,omitempty" tf:"intervals,omitempty"`

	// (String) The name of the mute timing.
	// The name of the mute timing.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type MuteTimingObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block List) The time intervals at which to mute notifications. Use an empty block to mute all the time. (see below for nested schema)
	// The time intervals at which to mute notifications. Use an empty block to mute all the time.
	Intervals []IntervalsObservation `json:"intervals,omitempty" tf:"intervals,omitempty"`

	// (String) The name of the mute timing.
	// The name of the mute timing.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The Organization ID. If not set, the Org ID defined in the provider block will be used.
	// The Organization ID. If not set, the Org ID defined in the provider block will be used.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`
}

type MuteTimingParameters struct {

	// (Block List) The time intervals at which to mute notifications. Use an empty block to mute all the time. (see below for nested schema)
	// The time intervals at which to mute notifications. Use an empty block to mute all the time.
	// +kubebuilder:validation:Optional
	Intervals []IntervalsParameters `json:"intervals,omitempty" tf:"intervals,omitempty"`

	// (String) The name of the mute timing.
	// The name of the mute timing.
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
}

type TimesInitParameters struct {

	// (String) The time, in hh:mm format, of when the interval should end exclusively.
	// The time, in hh:mm format, of when the interval should end exclusively.
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	// (String) The time, in hh:mm format, of when the interval should begin inclusively.
	// The time, in hh:mm format, of when the interval should begin inclusively.
	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type TimesObservation struct {

	// (String) The time, in hh:mm format, of when the interval should end exclusively.
	// The time, in hh:mm format, of when the interval should end exclusively.
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	// (String) The time, in hh:mm format, of when the interval should begin inclusively.
	// The time, in hh:mm format, of when the interval should begin inclusively.
	Start *string `json:"start,omitempty" tf:"start,omitempty"`
}

type TimesParameters struct {

	// (String) The time, in hh:mm format, of when the interval should end exclusively.
	// The time, in hh:mm format, of when the interval should end exclusively.
	// +kubebuilder:validation:Optional
	End *string `json:"end" tf:"end,omitempty"`

	// (String) The time, in hh:mm format, of when the interval should begin inclusively.
	// The time, in hh:mm format, of when the interval should begin inclusively.
	// +kubebuilder:validation:Optional
	Start *string `json:"start" tf:"start,omitempty"`
}

// MuteTimingSpec defines the desired state of MuteTiming
type MuteTimingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MuteTimingParameters `json:"forProvider"`
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
	InitProvider MuteTimingInitParameters `json:"initProvider,omitempty"`
}

// MuteTimingStatus defines the observed state of MuteTiming.
type MuteTimingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MuteTimingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MuteTiming is the Schema for the MuteTimings API. Manages Grafana Alerting mute timings. Official documentation https://grafana.com/docs/grafana/latest/alerting/manage-notifications/mute-timings/HTTP API https://grafana.com/docs/grafana/next/developers/http_api/alerting_provisioning/#mute-timings This resource requires Grafana 9.1.0 or later.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type MuteTiming struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   MuteTimingSpec   `json:"spec"`
	Status MuteTimingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MuteTimingList contains a list of MuteTimings
type MuteTimingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MuteTiming `json:"items"`
}

// Repository type metadata.
var (
	MuteTiming_Kind             = "MuteTiming"
	MuteTiming_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MuteTiming_Kind}.String()
	MuteTiming_KindAPIVersion   = MuteTiming_Kind + "." + CRDGroupVersion.String()
	MuteTiming_GroupVersionKind = CRDGroupVersion.WithKind(MuteTiming_Kind)
)

func init() {
	SchemeBuilder.Register(&MuteTiming{}, &MuteTimingList{})
}
