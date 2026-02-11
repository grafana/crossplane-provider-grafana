// Package schemeregistry provides mappings from API group names to their scheme registration functions.
package schemeregistry

import (
	"k8s.io/apimachinery/pkg/runtime"

	// Cluster-scoped API groups
	v1alpha1alerting "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster/alerting/v1alpha1"
	v1alpha1asserts "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster/asserts/v1alpha1"
	v1alpha1cloud "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster/cloud/v1alpha1"
	v1alpha1cloudprovider "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster/cloudprovider/v1alpha1"
	v1alpha1connections "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster/connections/v1alpha1"
	v1alpha1enterprise "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster/enterprise/v1alpha1"
	v1alpha1fleetmanagement "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster/fleetmanagement/v1alpha1"
	v1alpha1frontendobservability "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster/frontendobservability/v1alpha1"
	v1alpha1k6 "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster/k6/v1alpha1"
	v1alpha1ml "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster/ml/v1alpha1"
	v1alpha1oncall "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster/oncall/v1alpha1"
	v1alpha1oss "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster/oss/v1alpha1"
	v1alpha1slo "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster/slo/v1alpha1"
	v1alpha1sm "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster/sm/v1alpha1"
	v1alpha1cluster "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster/v1alpha1"
	v1beta1cluster "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster/v1beta1"

	// Namespaced API groups
	v1alpha1namespacedalerting "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/alerting/v1alpha1"
	v1alpha1namespacedasserts "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/asserts/v1alpha1"
	v1alpha1namespacedcloud "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/cloud/v1alpha1"
	v1alpha1namespacedcloudprovider "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/cloudprovider/v1alpha1"
	v1alpha1namespacedconnections "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/connections/v1alpha1"
	v1alpha1namespacedenterprise "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/enterprise/v1alpha1"
	v1alpha1namespacedfleetmanagement "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/fleetmanagement/v1alpha1"
	v1alpha1namespacedfrontendobservability "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/frontendobservability/v1alpha1"
	v1alpha1namespacedk6 "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/k6/v1alpha1"
	v1alpha1namespacedml "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/ml/v1alpha1"
	v1alpha1namespacedoncall "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/oncall/v1alpha1"
	v1alpha1namespacedoss "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/oss/v1alpha1"
	v1alpha1namespacedslo "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/slo/v1alpha1"
	v1alpha1namespacedsm "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/sm/v1alpha1"
	v1alpha1namespaced "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/v1alpha1"
	v1beta1namespaced "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/v1beta1"
)

// clusterGroupSchemes maps API group names to their SchemeBuilder.AddToScheme functions
// for cluster-scoped resources.
var clusterGroupSchemes = map[string]func(*runtime.Scheme) error{
	"alerting":              v1alpha1alerting.SchemeBuilder.AddToScheme,
	"asserts":               v1alpha1asserts.SchemeBuilder.AddToScheme,
	"cloud":                 v1alpha1cloud.SchemeBuilder.AddToScheme,
	"cloudprovider":         v1alpha1cloudprovider.SchemeBuilder.AddToScheme,
	"connections":           v1alpha1connections.SchemeBuilder.AddToScheme,
	"enterprise":            v1alpha1enterprise.SchemeBuilder.AddToScheme,
	"fleetmanagement":       v1alpha1fleetmanagement.SchemeBuilder.AddToScheme,
	"frontendobservability": v1alpha1frontendobservability.SchemeBuilder.AddToScheme,
	"k6":                    v1alpha1k6.SchemeBuilder.AddToScheme,
	"ml":                    v1alpha1ml.SchemeBuilder.AddToScheme,
	"oncall":                v1alpha1oncall.SchemeBuilder.AddToScheme,
	"oss":                   v1alpha1oss.SchemeBuilder.AddToScheme,
	"slo":                   v1alpha1slo.SchemeBuilder.AddToScheme,
	"sm":                    v1alpha1sm.SchemeBuilder.AddToScheme,
	"cluster":               v1alpha1cluster.SchemeBuilder.AddToScheme,
	"cluster-v1beta1":       v1beta1cluster.SchemeBuilder.AddToScheme,
}

// namespacedGroupSchemes maps API group names to their SchemeBuilder.AddToScheme functions
// for namespaced resources.
var namespacedGroupSchemes = map[string]func(*runtime.Scheme) error{
	"alerting":              v1alpha1namespacedalerting.SchemeBuilder.AddToScheme,
	"asserts":               v1alpha1namespacedasserts.SchemeBuilder.AddToScheme,
	"cloud":                 v1alpha1namespacedcloud.SchemeBuilder.AddToScheme,
	"cloudprovider":         v1alpha1namespacedcloudprovider.SchemeBuilder.AddToScheme,
	"connections":           v1alpha1namespacedconnections.SchemeBuilder.AddToScheme,
	"enterprise":            v1alpha1namespacedenterprise.SchemeBuilder.AddToScheme,
	"fleetmanagement":       v1alpha1namespacedfleetmanagement.SchemeBuilder.AddToScheme,
	"frontendobservability": v1alpha1namespacedfrontendobservability.SchemeBuilder.AddToScheme,
	"k6":                    v1alpha1namespacedk6.SchemeBuilder.AddToScheme,
	"ml":                    v1alpha1namespacedml.SchemeBuilder.AddToScheme,
	"oncall":                v1alpha1namespacedoncall.SchemeBuilder.AddToScheme,
	"oss":                   v1alpha1namespacedoss.SchemeBuilder.AddToScheme,
	"slo":                   v1alpha1namespacedslo.SchemeBuilder.AddToScheme,
	"sm":                    v1alpha1namespacedsm.SchemeBuilder.AddToScheme,
	"namespaced":            v1alpha1namespaced.SchemeBuilder.AddToScheme,
	"namespaced-v1beta1":    v1beta1namespaced.SchemeBuilder.AddToScheme,
}

// GetClusterGroupSchemeFunc returns the scheme registration function for a cluster-scoped API group.
// Returns nil if the group name is not found.
func GetClusterGroupSchemeFunc(groupName string) func(*runtime.Scheme) error {
	return clusterGroupSchemes[groupName]
}

// GetNamespacedGroupSchemeFunc returns the scheme registration function for a namespaced API group.
// Returns nil if the group name is not found.
func GetNamespacedGroupSchemeFunc(groupName string) func(*runtime.Scheme) error {
	return namespacedGroupSchemes[groupName]
}

// GetAllClusterGroupNames returns a list of all registered cluster API group names.
// This is primarily useful for testing and validation.
func GetAllClusterGroupNames() []string {
	names := make([]string, 0, len(clusterGroupSchemes))
	for name := range clusterGroupSchemes {
		names = append(names, name)
	}
	return names
}

// GetAllNamespacedGroupNames returns a list of all registered namespaced API group names.
// This is primarily useful for testing and validation.
func GetAllNamespacedGroupNames() []string {
	names := make([]string, 0, len(namespacedGroupSchemes))
	for name := range namespacedGroupSchemes {
		names = append(names, name)
	}
	return names
}
