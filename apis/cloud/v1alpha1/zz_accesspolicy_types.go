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

type AccessPolicyInitParameters struct {

	// (Block Set) Conditions for the access policy. (see below for nested schema)
	// Conditions for the access policy.
	Conditions []ConditionsInitParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// (String) Display name of the access policy. Defaults to the name.
	// Display name of the access policy. Defaults to the name.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (String) Name of the access policy.
	// Name of the access policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block Set, Min: 1) (see below for nested schema)
	Realm []RealmInitParameters `json:"realm,omitempty" tf:"realm,omitempty"`

	// cloud/developer-resources/api-reference/cloud-api/#list-regions.
	// Region where the API is deployed. Generally where the stack is deployed. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// cloud/security-and-account-management/authentication-and-permissions/access-policies/#scopes for possible values.
	// Scopes of the access policy. See https://grafana.com/docs/grafana-cloud/security-and-account-management/authentication-and-permissions/access-policies/#scopes for possible values.
	// +listType=set
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`
}

type AccessPolicyObservation struct {

	// (Block Set) Conditions for the access policy. (see below for nested schema)
	// Conditions for the access policy.
	Conditions []ConditionsObservation `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// (String) Creation date of the access policy.
	// Creation date of the access policy.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// (String) Display name of the access policy. Defaults to the name.
	// Display name of the access policy. Defaults to the name.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Name of the access policy.
	// Name of the access policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) ID of the access policy.
	// ID of the access policy.
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// (Block Set, Min: 1) (see below for nested schema)
	Realm []RealmObservation `json:"realm,omitempty" tf:"realm,omitempty"`

	// cloud/developer-resources/api-reference/cloud-api/#list-regions.
	// Region where the API is deployed. Generally where the stack is deployed. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// cloud/security-and-account-management/authentication-and-permissions/access-policies/#scopes for possible values.
	// Scopes of the access policy. See https://grafana.com/docs/grafana-cloud/security-and-account-management/authentication-and-permissions/access-policies/#scopes for possible values.
	// +listType=set
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// (String) Last update date of the access policy.
	// Last update date of the access policy.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type AccessPolicyParameters struct {

	// (Block Set) Conditions for the access policy. (see below for nested schema)
	// Conditions for the access policy.
	// +kubebuilder:validation:Optional
	Conditions []ConditionsParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// (String) Display name of the access policy. Defaults to the name.
	// Display name of the access policy. Defaults to the name.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (String) Name of the access policy.
	// Name of the access policy.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Realm []RealmParameters `json:"realm,omitempty" tf:"realm,omitempty"`

	// cloud/developer-resources/api-reference/cloud-api/#list-regions.
	// Region where the API is deployed. Generally where the stack is deployed. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#list-regions.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// cloud/security-and-account-management/authentication-and-permissions/access-policies/#scopes for possible values.
	// Scopes of the access policy. See https://grafana.com/docs/grafana-cloud/security-and-account-management/authentication-and-permissions/access-policies/#scopes for possible values.
	// +kubebuilder:validation:Optional
	// +listType=set
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`
}

type ConditionsInitParameters struct {

	// (Set of String) Conditions that apply to the access policy,such as IP Allow lists.
	// Conditions that apply to the access policy,such as IP Allow lists.
	// +listType=set
	AllowedSubnets []*string `json:"allowedSubnets,omitempty" tf:"allowed_subnets,omitempty"`
}

type ConditionsObservation struct {

	// (Set of String) Conditions that apply to the access policy,such as IP Allow lists.
	// Conditions that apply to the access policy,such as IP Allow lists.
	// +listType=set
	AllowedSubnets []*string `json:"allowedSubnets,omitempty" tf:"allowed_subnets,omitempty"`
}

type ConditionsParameters struct {

	// (Set of String) Conditions that apply to the access policy,such as IP Allow lists.
	// Conditions that apply to the access policy,such as IP Allow lists.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedSubnets []*string `json:"allowedSubnets" tf:"allowed_subnets,omitempty"`
}

type LabelPolicyInitParameters struct {

	// (String) The label selector to match in metrics or logs query. Should be in PromQL or LogQL format.
	// The label selector to match in metrics or logs query. Should be in PromQL or LogQL format.
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`
}

type LabelPolicyObservation struct {

	// (String) The label selector to match in metrics or logs query. Should be in PromQL or LogQL format.
	// The label selector to match in metrics or logs query. Should be in PromQL or LogQL format.
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`
}

type LabelPolicyParameters struct {

	// (String) The label selector to match in metrics or logs query. Should be in PromQL or LogQL format.
	// The label selector to match in metrics or logs query. Should be in PromQL or LogQL format.
	// +kubebuilder:validation:Optional
	Selector *string `json:"selector" tf:"selector,omitempty"`
}

type RealmInitParameters struct {

	// (String) The identifier of the org or stack. For orgs, this is the slug, for stacks, this is the stack ID.
	// The identifier of the org or stack. For orgs, this is the slug, for stacks, this is the stack ID.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/cloud/v1alpha1.Stack
	// +crossplane:generate:reference:extractor=github.com/grafana/crossplane-provider-grafana/config/grafana.ComputedFieldExtractor("id")
	// +crossplane:generate:reference:refFieldName=StackRef
	// +crossplane:generate:reference:selectorFieldName=StackSelector
	Identifier *string `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// (Block Set) (see below for nested schema)
	LabelPolicy []LabelPolicyInitParameters `json:"labelPolicy,omitempty" tf:"label_policy,omitempty"`

	// Reference to a Stack in cloud to populate identifier.
	// +kubebuilder:validation:Optional
	StackRef *v1.Reference `json:"stackRef,omitempty" tf:"-"`

	// Selector for a Stack in cloud to populate identifier.
	// +kubebuilder:validation:Optional
	StackSelector *v1.Selector `json:"stackSelector,omitempty" tf:"-"`

	// (String) Whether a policy applies to a Cloud org or a specific stack. Should be one of org or stack.
	// Whether a policy applies to a Cloud org or a specific stack. Should be one of `org` or `stack`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RealmObservation struct {

	// (String) The identifier of the org or stack. For orgs, this is the slug, for stacks, this is the stack ID.
	// The identifier of the org or stack. For orgs, this is the slug, for stacks, this is the stack ID.
	Identifier *string `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// (Block Set) (see below for nested schema)
	LabelPolicy []LabelPolicyObservation `json:"labelPolicy,omitempty" tf:"label_policy,omitempty"`

	// (String) Whether a policy applies to a Cloud org or a specific stack. Should be one of org or stack.
	// Whether a policy applies to a Cloud org or a specific stack. Should be one of `org` or `stack`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RealmParameters struct {

	// (String) The identifier of the org or stack. For orgs, this is the slug, for stacks, this is the stack ID.
	// The identifier of the org or stack. For orgs, this is the slug, for stacks, this is the stack ID.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/cloud/v1alpha1.Stack
	// +crossplane:generate:reference:extractor=github.com/grafana/crossplane-provider-grafana/config/grafana.ComputedFieldExtractor("id")
	// +crossplane:generate:reference:refFieldName=StackRef
	// +crossplane:generate:reference:selectorFieldName=StackSelector
	// +kubebuilder:validation:Optional
	Identifier *string `json:"identifier,omitempty" tf:"identifier,omitempty"`

	// (Block Set) (see below for nested schema)
	// +kubebuilder:validation:Optional
	LabelPolicy []LabelPolicyParameters `json:"labelPolicy,omitempty" tf:"label_policy,omitempty"`

	// Reference to a Stack in cloud to populate identifier.
	// +kubebuilder:validation:Optional
	StackRef *v1.Reference `json:"stackRef,omitempty" tf:"-"`

	// Selector for a Stack in cloud to populate identifier.
	// +kubebuilder:validation:Optional
	StackSelector *v1.Selector `json:"stackSelector,omitempty" tf:"-"`

	// (String) Whether a policy applies to a Cloud org or a specific stack. Should be one of org or stack.
	// Whether a policy applies to a Cloud org or a specific stack. Should be one of `org` or `stack`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// AccessPolicySpec defines the desired state of AccessPolicy
type AccessPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccessPolicyParameters `json:"forProvider"`
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
	InitProvider AccessPolicyInitParameters `json:"initProvider,omitempty"`
}

// AccessPolicyStatus defines the observed state of AccessPolicy.
type AccessPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccessPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AccessPolicy is the Schema for the AccessPolicys API. Official documentation https://grafana.com/docs/grafana-cloud/security-and-account-management/authentication-and-permissions/access-policies/API documentation https://grafana.com/docs/grafana-cloud/developer-resources/api-reference/cloud-api/#create-an-access-policy Required access policy scopes: accesspolicies:readaccesspolicies:writeaccesspolicies:delete
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type AccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.realm) || (has(self.initProvider) && has(self.initProvider.realm))",message="spec.forProvider.realm is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.region) || (has(self.initProvider) && has(self.initProvider.region))",message="spec.forProvider.region is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scopes) || (has(self.initProvider) && has(self.initProvider.scopes))",message="spec.forProvider.scopes is a required parameter"
	Spec   AccessPolicySpec   `json:"spec"`
	Status AccessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessPolicyList contains a list of AccessPolicys
type AccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessPolicy `json:"items"`
}

// Repository type metadata.
var (
	AccessPolicy_Kind             = "AccessPolicy"
	AccessPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccessPolicy_Kind}.String()
	AccessPolicy_KindAPIVersion   = AccessPolicy_Kind + "." + CRDGroupVersion.String()
	AccessPolicy_GroupVersionKind = CRDGroupVersion.WithKind(AccessPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&AccessPolicy{}, &AccessPolicyList{})
}
