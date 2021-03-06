/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CloudPluginInstallationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CloudPluginInstallationParameters struct {

	// Slug of the plugin to be installed.
	// +kubebuilder:validation:Required
	Slug *string `json:"slug" tf:"slug,omitempty"`

	// +kubebuilder:validation:Optional
	StackRef *v1.Reference `json:"stackRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StackSelector *v1.Selector `json:"stackSelector,omitempty" tf:"-"`

	// The stack id to which the plugin should be installed.
	// +crossplane:generate:reference:type=CloudStack
	// +crossplane:generate:reference:extractor=github.com/grafana/crossplane-provider-grafana/config.SlugExtractor()
	// +crossplane:generate:reference:refFieldName=StackRef
	// +crossplane:generate:reference:selectorFieldName=StackSelector
	// +kubebuilder:validation:Optional
	StackSlug *string `json:"stackSlug,omitempty" tf:"stack_slug,omitempty"`

	// Version of the plugin to be installed.
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

// CloudPluginInstallationSpec defines the desired state of CloudPluginInstallation
type CloudPluginInstallationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CloudPluginInstallationParameters `json:"forProvider"`
}

// CloudPluginInstallationStatus defines the observed state of CloudPluginInstallation.
type CloudPluginInstallationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CloudPluginInstallationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudPluginInstallation is the Schema for the CloudPluginInstallations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafanajet}
type CloudPluginInstallation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudPluginInstallationSpec   `json:"spec"`
	Status            CloudPluginInstallationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudPluginInstallationList contains a list of CloudPluginInstallations
type CloudPluginInstallationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudPluginInstallation `json:"items"`
}

// Repository type metadata.
var (
	CloudPluginInstallation_Kind             = "CloudPluginInstallation"
	CloudPluginInstallation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CloudPluginInstallation_Kind}.String()
	CloudPluginInstallation_KindAPIVersion   = CloudPluginInstallation_Kind + "." + CRDGroupVersion.String()
	CloudPluginInstallation_GroupVersionKind = CRDGroupVersion.WithKind(CloudPluginInstallation_Kind)
)

func init() {
	SchemeBuilder.Register(&CloudPluginInstallation{}, &CloudPluginInstallationList{})
}
