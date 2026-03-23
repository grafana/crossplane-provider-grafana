/*
Copyright 2025 Grafana
*/

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	v2 "github.com/crossplane/crossplane-runtime/v2/apis/common/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TeamsParameters defines filter criteria for listing Grafana teams.
type TeamsParameters struct {
	// Name filters teams by exact name match.
	// +optional
	Name *string `json:"name,omitempty"`
	// Query filters teams by a search string (matches name/email).
	// +optional
	Query *string `json:"query,omitempty"`
}

// TeamSummary holds a summary of a single Grafana team.
type TeamSummary struct {
	ID          int64  `json:"id,omitempty"`
	UID         string `json:"uid,omitempty"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	MemberCount int64  `json:"memberCount,omitempty"`
	OrgID       int64  `json:"orgId,omitempty"`
}

// TeamsObservation holds the observed list of teams.
type TeamsObservation struct {
	Teams []TeamSummary `json:"teams,omitempty"`
}

// TeamsSpec defines the desired state of a Teams observe resource.
type TeamsSpec struct {
	v2.ManagedResourceSpec `json:",inline"`
	ForProvider            TeamsParameters `json:"forProvider"`
}

// TeamsStatus defines the observed state of a Teams observe resource.
type TeamsStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          TeamsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,categories=crossplane
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"

// Teams is an observe-only resource that lists all Grafana teams matching the
// given filter criteria in a single reconcile cycle.
type Teams struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TeamsSpec   `json:"spec"`
	Status TeamsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TeamsList contains a list of Teams.
type TeamsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Teams `json:"items"`
}

// Teams_GroupVersionKind is the GVK for Teams resources.
var Teams_GroupVersionKind = schema.GroupVersionKind{
	Group:   CRDGroup,
	Version: CRDVersion,
	Kind:    "Teams",
}

func init() {
	SchemeBuilder.Register(&Teams{}, &TeamsList{})
}
