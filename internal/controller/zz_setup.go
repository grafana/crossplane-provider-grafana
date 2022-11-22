/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	apikey "github.com/grafana/crossplane-provider-grafana/internal/controller/cloud/apikey"
	stack "github.com/grafana/crossplane-provider-grafana/internal/controller/cloud/stack"
	apikeyoss "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/apikey"
	dashboard "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/dashboard"
	folder "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/folder"
	team "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/team"
	user "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/user"
	providerconfig "github.com/grafana/crossplane-provider-grafana/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		apikey.Setup,
		stack.Setup,
		apikeyoss.Setup,
		dashboard.Setup,
		folder.Setup,
		team.Setup,
		user.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
