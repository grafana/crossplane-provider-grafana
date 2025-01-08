/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/grafana/crossplane-provider-grafana/apis/alerting/v1alpha1"
	v1alpha1cloud "github.com/grafana/crossplane-provider-grafana/apis/cloud/v1alpha1"
	v1alpha1cloudprovider "github.com/grafana/crossplane-provider-grafana/apis/cloudprovider/v1alpha1"
	v1alpha1connections "github.com/grafana/crossplane-provider-grafana/apis/connections/v1alpha1"
	v1alpha1enterprise "github.com/grafana/crossplane-provider-grafana/apis/enterprise/v1alpha1"
	v1alpha1ml "github.com/grafana/crossplane-provider-grafana/apis/ml/v1alpha1"
	v1alpha1oncall "github.com/grafana/crossplane-provider-grafana/apis/oncall/v1alpha1"
	v1alpha1oss "github.com/grafana/crossplane-provider-grafana/apis/oss/v1alpha1"
	v1alpha1slo "github.com/grafana/crossplane-provider-grafana/apis/slo/v1alpha1"
	v1alpha1sm "github.com/grafana/crossplane-provider-grafana/apis/sm/v1alpha1"
	v1alpha1apis "github.com/grafana/crossplane-provider-grafana/apis/v1alpha1"
	v1beta1 "github.com/grafana/crossplane-provider-grafana/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1cloud.SchemeBuilder.AddToScheme,
		v1alpha1cloudprovider.SchemeBuilder.AddToScheme,
		v1alpha1connections.SchemeBuilder.AddToScheme,
		v1alpha1enterprise.SchemeBuilder.AddToScheme,
		v1alpha1ml.SchemeBuilder.AddToScheme,
		v1alpha1oncall.SchemeBuilder.AddToScheme,
		v1alpha1oss.SchemeBuilder.AddToScheme,
		v1alpha1slo.SchemeBuilder.AddToScheme,
		v1alpha1sm.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
