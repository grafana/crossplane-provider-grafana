/*
Copyright 2025
*/

package namespaced

import (
    "k8s.io/apimachinery/pkg/runtime"

    v1beta1 "github.com/grafana/crossplane-provider-grafana/apis/namespaced/v1beta1"
)

// AddToScheme adds all namespaced API types to the Scheme.
func AddToScheme(s *runtime.Scheme) error {
    if err := v1beta1.SchemeBuilder.AddToScheme(s); err != nil {
        return err
    }
    return nil
}
