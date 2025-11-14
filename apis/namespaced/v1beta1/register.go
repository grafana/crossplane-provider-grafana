/*
Copyright 2025 Grafana
*/

package v1beta1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata for modern namespaced API group.
const (
	Group   = "grafana.m.crossplane.io"
	Version = "v1beta1"
)

var (
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}
	SchemeBuilder      = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

// ProviderConfig type metadata.
var (
	ProviderConfigKind             = reflect.TypeOf(ProviderConfig{}).Name()
	ProviderConfigGroupKind        = schema.GroupKind{Group: Group, Kind: ProviderConfigKind}.String()
	ProviderConfigKindAPIVersion   = ProviderConfigKind + "." + SchemeGroupVersion.String()
	ProviderConfigGroupVersionKind = SchemeGroupVersion.WithKind(ProviderConfigKind)
)

// ProviderConfigUsage type metadata.
var (
	ProviderConfigUsageKind             = reflect.TypeOf(ProviderConfigUsage{}).Name()
	ProviderConfigUsageGroupKind        = schema.GroupKind{Group: Group, Kind: ProviderConfigUsageKind}.String()
	ProviderConfigUsageKindAPIVersion   = ProviderConfigUsageKind + "." + SchemeGroupVersion.String()
	ProviderConfigUsageGroupVersionKind = SchemeGroupVersion.WithKind(ProviderConfigUsageKind)

	ProviderConfigUsageListKind             = reflect.TypeOf(ProviderConfigUsageList{}).Name()
	ProviderConfigUsageListGroupKind        = schema.GroupKind{Group: Group, Kind: ProviderConfigUsageListKind}.String()
	ProviderConfigUsageListKindAPIVersion   = ProviderConfigUsageListKind + "." + SchemeGroupVersion.String()
	ProviderConfigUsageListGroupVersionKind = SchemeGroupVersion.WithKind(ProviderConfigUsageListKind)
)

// ClusterProviderConfig type metadata.
var (
	ClusterProviderConfigKind             = reflect.TypeOf(ClusterProviderConfig{}).Name()
	ClusterProviderConfigGroupKind        = schema.GroupKind{Group: Group, Kind: ClusterProviderConfigKind}.String()
	ClusterProviderConfigKindAPIVersion   = ClusterProviderConfigKind + "." + SchemeGroupVersion.String()
	ClusterProviderConfigGroupVersionKind = SchemeGroupVersion.WithKind(ClusterProviderConfigKind)

	ClusterProviderConfigListKind             = reflect.TypeOf(ClusterProviderConfigList{}).Name()
	ClusterProviderConfigListGroupKind        = schema.GroupKind{Group: Group, Kind: ClusterProviderConfigListKind}.String()
	ClusterProviderConfigListKindAPIVersion   = ClusterProviderConfigListKind + "." + SchemeGroupVersion.String()
	ClusterProviderConfigListGroupVersionKind = SchemeGroupVersion.WithKind(ClusterProviderConfigListKind)
)

func init() {
	SchemeBuilder.Register(&ProviderConfig{}, &ProviderConfigList{})
	SchemeBuilder.Register(&ClusterProviderConfig{}, &ClusterProviderConfigList{})
	SchemeBuilder.Register(&ProviderConfigUsage{}, &ProviderConfigUsageList{})
}

// AddToScheme registers the types in the namespaced v1beta1 API group into the supplied scheme.
var AddToScheme = SchemeBuilder.AddToScheme
