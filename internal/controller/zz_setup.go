/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	contactpoint "github.com/grafana/crossplane-provider-grafana/internal/controller/alerting/contactpoint"
	messagetemplate "github.com/grafana/crossplane-provider-grafana/internal/controller/alerting/messagetemplate"
	mutetiming "github.com/grafana/crossplane-provider-grafana/internal/controller/alerting/mutetiming"
	notificationpolicy "github.com/grafana/crossplane-provider-grafana/internal/controller/alerting/notificationpolicy"
	rulegroup "github.com/grafana/crossplane-provider-grafana/internal/controller/alerting/rulegroup"
	apikey "github.com/grafana/crossplane-provider-grafana/internal/controller/cloud/apikey"
	stack "github.com/grafana/crossplane-provider-grafana/internal/controller/cloud/stack"
	escalation "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/escalation"
	escalationchain "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/escalationchain"
	integration "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/integration"
	oncallshift "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/oncallshift"
	outgoingwebhook "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/outgoingwebhook"
	route "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/route"
	schedule "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/schedule"
	apikeyoss "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/apikey"
	dashboard "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/dashboard"
	datasource "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/datasource"
	folder "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/folder"
	team "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/team"
	user "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/user"
	providerconfig "github.com/grafana/crossplane-provider-grafana/internal/controller/providerconfig"
	check "github.com/grafana/crossplane-provider-grafana/internal/controller/sm/check"
	installation "github.com/grafana/crossplane-provider-grafana/internal/controller/sm/installation"
	probe "github.com/grafana/crossplane-provider-grafana/internal/controller/sm/probe"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		contactpoint.Setup,
		messagetemplate.Setup,
		mutetiming.Setup,
		notificationpolicy.Setup,
		rulegroup.Setup,
		apikey.Setup,
		stack.Setup,
		escalation.Setup,
		escalationchain.Setup,
		integration.Setup,
		oncallshift.Setup,
		outgoingwebhook.Setup,
		route.Setup,
		schedule.Setup,
		apikeyoss.Setup,
		dashboard.Setup,
		datasource.Setup,
		folder.Setup,
		team.Setup,
		user.Setup,
		providerconfig.Setup,
		check.Setup,
		installation.Setup,
		probe.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
