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

type CollectorInitParameters struct {

	// (Boolean) Whether the collector is enabled or not
	// Whether the collector is enabled or not
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Map of String) Remote attributes for the collector
	// Remote attributes for the collector
	// +mapType=granular
	RemoteAttributes map[string]*string `json:"remoteAttributes,omitempty" tf:"remote_attributes,omitempty"`
}

type CollectorObservation struct {

	// (Boolean) Whether the collector is enabled or not
	// Whether the collector is enabled or not
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) ID of the collector
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Map of String) Remote attributes for the collector
	// Remote attributes for the collector
	// +mapType=granular
	RemoteAttributes map[string]*string `json:"remoteAttributes,omitempty" tf:"remote_attributes,omitempty"`
}

type CollectorParameters struct {

	// (Boolean) Whether the collector is enabled or not
	// Whether the collector is enabled or not
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Map of String) Remote attributes for the collector
	// Remote attributes for the collector
	// +kubebuilder:validation:Optional
	// +mapType=granular
	RemoteAttributes map[string]*string `json:"remoteAttributes,omitempty" tf:"remote_attributes,omitempty"`
}

// CollectorSpec defines the desired state of Collector
type CollectorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CollectorParameters `json:"forProvider"`
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
	InitProvider CollectorInitParameters `json:"initProvider,omitempty"`
}

// CollectorStatus defines the observed state of Collector.
type CollectorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CollectorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Collector is the Schema for the Collectors API. Manages Grafana Fleet Management collectors. Official documentation https://grafana.com/docs/grafana-cloud/send-data/fleet-management/API documentation https://grafana.com/docs/grafana-cloud/send-data/fleet-management/api-reference/collector-api/ Required access policy scopes: fleet-management:readfleet-management:write
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type Collector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CollectorSpec   `json:"spec"`
	Status            CollectorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CollectorList contains a list of Collectors
type CollectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Collector `json:"items"`
}

// Repository type metadata.
var (
	Collector_Kind             = "Collector"
	Collector_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Collector_Kind}.String()
	Collector_KindAPIVersion   = Collector_Kind + "." + CRDGroupVersion.String()
	Collector_GroupVersionKind = CRDGroupVersion.WithKind(Collector_Kind)
)

func init() {
	SchemeBuilder.Register(&Collector{}, &CollectorList{})
}
