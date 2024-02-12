//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardsInitParameters) DeepCopyInto(out *DashboardsInitParameters) {
	*out = *in
	if in.TimeRange != nil {
		in, out := &in.TimeRange, &out.TimeRange
		*out = make([]TimeRangeInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UID != nil {
		in, out := &in.UID, &out.UID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardsInitParameters.
func (in *DashboardsInitParameters) DeepCopy() *DashboardsInitParameters {
	if in == nil {
		return nil
	}
	out := new(DashboardsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardsObservation) DeepCopyInto(out *DashboardsObservation) {
	*out = *in
	if in.TimeRange != nil {
		in, out := &in.TimeRange, &out.TimeRange
		*out = make([]TimeRangeObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UID != nil {
		in, out := &in.UID, &out.UID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardsObservation.
func (in *DashboardsObservation) DeepCopy() *DashboardsObservation {
	if in == nil {
		return nil
	}
	out := new(DashboardsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardsParameters) DeepCopyInto(out *DashboardsParameters) {
	*out = *in
	if in.TimeRange != nil {
		in, out := &in.TimeRange, &out.TimeRange
		*out = make([]TimeRangeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UID != nil {
		in, out := &in.UID, &out.UID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardsParameters.
func (in *DashboardsParameters) DeepCopy() *DashboardsParameters {
	if in == nil {
		return nil
	}
	out := new(DashboardsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Report) DeepCopyInto(out *Report) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Report.
func (in *Report) DeepCopy() *Report {
	if in == nil {
		return nil
	}
	out := new(Report)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Report) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportInitParameters) DeepCopyInto(out *ReportInitParameters) {
	*out = *in
	if in.DashboardRef != nil {
		in, out := &in.DashboardRef, &out.DashboardRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DashboardSelector != nil {
		in, out := &in.DashboardSelector, &out.DashboardSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.DashboardUID != nil {
		in, out := &in.DashboardUID, &out.DashboardUID
		*out = new(string)
		**out = **in
	}
	if in.Dashboards != nil {
		in, out := &in.Dashboards, &out.Dashboards
		*out = make([]DashboardsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Formats != nil {
		in, out := &in.Formats, &out.Formats
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IncludeDashboardLink != nil {
		in, out := &in.IncludeDashboardLink, &out.IncludeDashboardLink
		*out = new(bool)
		**out = **in
	}
	if in.IncludeTableCsv != nil {
		in, out := &in.IncludeTableCsv, &out.IncludeTableCsv
		*out = new(bool)
		**out = **in
	}
	if in.Layout != nil {
		in, out := &in.Layout, &out.Layout
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.OrganizationRef != nil {
		in, out := &in.OrganizationRef, &out.OrganizationRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.OrganizationSelector != nil {
		in, out := &in.OrganizationSelector, &out.OrganizationSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Orientation != nil {
		in, out := &in.Orientation, &out.Orientation
		*out = new(string)
		**out = **in
	}
	if in.Recipients != nil {
		in, out := &in.Recipients, &out.Recipients
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ReplyTo != nil {
		in, out := &in.ReplyTo, &out.ReplyTo
		*out = new(string)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = make([]ScheduleInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TimeRange != nil {
		in, out := &in.TimeRange, &out.TimeRange
		*out = make([]ReportTimeRangeInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportInitParameters.
func (in *ReportInitParameters) DeepCopy() *ReportInitParameters {
	if in == nil {
		return nil
	}
	out := new(ReportInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportList) DeepCopyInto(out *ReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Report, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportList.
func (in *ReportList) DeepCopy() *ReportList {
	if in == nil {
		return nil
	}
	out := new(ReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportObservation) DeepCopyInto(out *ReportObservation) {
	*out = *in
	if in.DashboardUID != nil {
		in, out := &in.DashboardUID, &out.DashboardUID
		*out = new(string)
		**out = **in
	}
	if in.Dashboards != nil {
		in, out := &in.Dashboards, &out.Dashboards
		*out = make([]DashboardsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Formats != nil {
		in, out := &in.Formats, &out.Formats
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IncludeDashboardLink != nil {
		in, out := &in.IncludeDashboardLink, &out.IncludeDashboardLink
		*out = new(bool)
		**out = **in
	}
	if in.IncludeTableCsv != nil {
		in, out := &in.IncludeTableCsv, &out.IncludeTableCsv
		*out = new(bool)
		**out = **in
	}
	if in.Layout != nil {
		in, out := &in.Layout, &out.Layout
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.Orientation != nil {
		in, out := &in.Orientation, &out.Orientation
		*out = new(string)
		**out = **in
	}
	if in.Recipients != nil {
		in, out := &in.Recipients, &out.Recipients
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ReplyTo != nil {
		in, out := &in.ReplyTo, &out.ReplyTo
		*out = new(string)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = make([]ScheduleObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TimeRange != nil {
		in, out := &in.TimeRange, &out.TimeRange
		*out = make([]ReportTimeRangeObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportObservation.
func (in *ReportObservation) DeepCopy() *ReportObservation {
	if in == nil {
		return nil
	}
	out := new(ReportObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportParameters) DeepCopyInto(out *ReportParameters) {
	*out = *in
	if in.DashboardRef != nil {
		in, out := &in.DashboardRef, &out.DashboardRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DashboardSelector != nil {
		in, out := &in.DashboardSelector, &out.DashboardSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.DashboardUID != nil {
		in, out := &in.DashboardUID, &out.DashboardUID
		*out = new(string)
		**out = **in
	}
	if in.Dashboards != nil {
		in, out := &in.Dashboards, &out.Dashboards
		*out = make([]DashboardsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Formats != nil {
		in, out := &in.Formats, &out.Formats
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IncludeDashboardLink != nil {
		in, out := &in.IncludeDashboardLink, &out.IncludeDashboardLink
		*out = new(bool)
		**out = **in
	}
	if in.IncludeTableCsv != nil {
		in, out := &in.IncludeTableCsv, &out.IncludeTableCsv
		*out = new(bool)
		**out = **in
	}
	if in.Layout != nil {
		in, out := &in.Layout, &out.Layout
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.OrganizationRef != nil {
		in, out := &in.OrganizationRef, &out.OrganizationRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.OrganizationSelector != nil {
		in, out := &in.OrganizationSelector, &out.OrganizationSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Orientation != nil {
		in, out := &in.Orientation, &out.Orientation
		*out = new(string)
		**out = **in
	}
	if in.Recipients != nil {
		in, out := &in.Recipients, &out.Recipients
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ReplyTo != nil {
		in, out := &in.ReplyTo, &out.ReplyTo
		*out = new(string)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = make([]ScheduleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TimeRange != nil {
		in, out := &in.TimeRange, &out.TimeRange
		*out = make([]ReportTimeRangeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportParameters.
func (in *ReportParameters) DeepCopy() *ReportParameters {
	if in == nil {
		return nil
	}
	out := new(ReportParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportSpec) DeepCopyInto(out *ReportSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportSpec.
func (in *ReportSpec) DeepCopy() *ReportSpec {
	if in == nil {
		return nil
	}
	out := new(ReportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportStatus) DeepCopyInto(out *ReportStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportStatus.
func (in *ReportStatus) DeepCopy() *ReportStatus {
	if in == nil {
		return nil
	}
	out := new(ReportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportTimeRangeInitParameters) DeepCopyInto(out *ReportTimeRangeInitParameters) {
	*out = *in
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(string)
		**out = **in
	}
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportTimeRangeInitParameters.
func (in *ReportTimeRangeInitParameters) DeepCopy() *ReportTimeRangeInitParameters {
	if in == nil {
		return nil
	}
	out := new(ReportTimeRangeInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportTimeRangeObservation) DeepCopyInto(out *ReportTimeRangeObservation) {
	*out = *in
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(string)
		**out = **in
	}
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportTimeRangeObservation.
func (in *ReportTimeRangeObservation) DeepCopy() *ReportTimeRangeObservation {
	if in == nil {
		return nil
	}
	out := new(ReportTimeRangeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportTimeRangeParameters) DeepCopyInto(out *ReportTimeRangeParameters) {
	*out = *in
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(string)
		**out = **in
	}
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportTimeRangeParameters.
func (in *ReportTimeRangeParameters) DeepCopy() *ReportTimeRangeParameters {
	if in == nil {
		return nil
	}
	out := new(ReportTimeRangeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleInitParameters) DeepCopyInto(out *ScheduleInitParameters) {
	*out = *in
	if in.CustomInterval != nil {
		in, out := &in.CustomInterval, &out.CustomInterval
		*out = new(string)
		**out = **in
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(string)
		**out = **in
	}
	if in.LastDayOfMonth != nil {
		in, out := &in.LastDayOfMonth, &out.LastDayOfMonth
		*out = new(bool)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	if in.WorkdaysOnly != nil {
		in, out := &in.WorkdaysOnly, &out.WorkdaysOnly
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleInitParameters.
func (in *ScheduleInitParameters) DeepCopy() *ScheduleInitParameters {
	if in == nil {
		return nil
	}
	out := new(ScheduleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleObservation) DeepCopyInto(out *ScheduleObservation) {
	*out = *in
	if in.CustomInterval != nil {
		in, out := &in.CustomInterval, &out.CustomInterval
		*out = new(string)
		**out = **in
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(string)
		**out = **in
	}
	if in.LastDayOfMonth != nil {
		in, out := &in.LastDayOfMonth, &out.LastDayOfMonth
		*out = new(bool)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	if in.WorkdaysOnly != nil {
		in, out := &in.WorkdaysOnly, &out.WorkdaysOnly
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleObservation.
func (in *ScheduleObservation) DeepCopy() *ScheduleObservation {
	if in == nil {
		return nil
	}
	out := new(ScheduleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleParameters) DeepCopyInto(out *ScheduleParameters) {
	*out = *in
	if in.CustomInterval != nil {
		in, out := &in.CustomInterval, &out.CustomInterval
		*out = new(string)
		**out = **in
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(string)
		**out = **in
	}
	if in.LastDayOfMonth != nil {
		in, out := &in.LastDayOfMonth, &out.LastDayOfMonth
		*out = new(bool)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	if in.WorkdaysOnly != nil {
		in, out := &in.WorkdaysOnly, &out.WorkdaysOnly
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleParameters.
func (in *ScheduleParameters) DeepCopy() *ScheduleParameters {
	if in == nil {
		return nil
	}
	out := new(ScheduleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeRangeInitParameters) DeepCopyInto(out *TimeRangeInitParameters) {
	*out = *in
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(string)
		**out = **in
	}
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeRangeInitParameters.
func (in *TimeRangeInitParameters) DeepCopy() *TimeRangeInitParameters {
	if in == nil {
		return nil
	}
	out := new(TimeRangeInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeRangeObservation) DeepCopyInto(out *TimeRangeObservation) {
	*out = *in
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(string)
		**out = **in
	}
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeRangeObservation.
func (in *TimeRangeObservation) DeepCopy() *TimeRangeObservation {
	if in == nil {
		return nil
	}
	out := new(TimeRangeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeRangeParameters) DeepCopyInto(out *TimeRangeParameters) {
	*out = *in
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(string)
		**out = **in
	}
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeRangeParameters.
func (in *TimeRangeParameters) DeepCopy() *TimeRangeParameters {
	if in == nil {
		return nil
	}
	out := new(TimeRangeParameters)
	in.DeepCopyInto(out)
	return out
}
