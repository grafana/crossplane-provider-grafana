/*
Copyright 2025 Grafana
*/

package controller

import (
	tjcontroller "github.com/crossplane/upjet/v2/pkg/controller"
	ctrl "sigs.k8s.io/controller-runtime"

	observeteams "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/observe/teams"
)

// SetupCustom registers all custom namespaced controllers.
func SetupCustom(mgr ctrl.Manager, o tjcontroller.Options) error {
	return observeteams.Setup(mgr, o)
}

// SetupCustomGated registers all custom namespaced controllers behind the SafeStart gate.
func SetupCustomGated(mgr ctrl.Manager, o tjcontroller.Options) error {
	return observeteams.SetupGated(mgr, o)
}
