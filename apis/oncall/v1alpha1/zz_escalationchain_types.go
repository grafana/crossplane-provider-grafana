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

type EscalationChainInitParameters struct {

	// (String) The name of the escalation chain.
	// The name of the escalation chain.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the grafana_oncall_team datasource.
	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana_oncall_team` datasource.
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`
}

type EscalationChainObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the escalation chain.
	// The name of the escalation chain.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the grafana_oncall_team datasource.
	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana_oncall_team` datasource.
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`
}

type EscalationChainParameters struct {

	// (String) The name of the escalation chain.
	// The name of the escalation chain.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the grafana_oncall_team datasource.
	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana_oncall_team` datasource.
	// +kubebuilder:validation:Optional
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`
}

// EscalationChainSpec defines the desired state of EscalationChain
type EscalationChainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EscalationChainParameters `json:"forProvider"`
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
	InitProvider EscalationChainInitParameters `json:"initProvider,omitempty"`
}

// EscalationChainStatus defines the observed state of EscalationChain.
type EscalationChainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EscalationChainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EscalationChain is the Schema for the EscalationChains API. HTTP API https://grafana.com/docs/oncall/latest/oncall-api-reference/escalation_chains/
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type EscalationChain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   EscalationChainSpec   `json:"spec"`
	Status EscalationChainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EscalationChainList contains a list of EscalationChains
type EscalationChainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EscalationChain `json:"items"`
}

// Repository type metadata.
var (
	EscalationChain_Kind             = "EscalationChain"
	EscalationChain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EscalationChain_Kind}.String()
	EscalationChain_KindAPIVersion   = EscalationChain_Kind + "." + CRDGroupVersion.String()
	EscalationChain_GroupVersionKind = CRDGroupVersion.WithKind(EscalationChain_Kind)
)

func init() {
	SchemeBuilder.Register(&EscalationChain{}, &EscalationChainList{})
}
