/*
Copyright 2025 Grafana
*/

// Package namespaced contains Kubernetes API for the provider.
package namespaced

import (
	observev1alpha1 "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/observe/v1alpha1"
)

func init() {
	AddToSchemes = append(AddToSchemes, observev1alpha1.SchemeBuilder.AddToScheme)
}
