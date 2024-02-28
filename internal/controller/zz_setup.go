// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	contactpoint "github.com/grafana/crossplane-provider-grafana/internal/controller/alerting/contactpoint"
	messagetemplate "github.com/grafana/crossplane-provider-grafana/internal/controller/alerting/messagetemplate"
	mutetiming "github.com/grafana/crossplane-provider-grafana/internal/controller/alerting/mutetiming"
	notificationpolicy "github.com/grafana/crossplane-provider-grafana/internal/controller/alerting/notificationpolicy"
	rulegroup "github.com/grafana/crossplane-provider-grafana/internal/controller/alerting/rulegroup"
	apikey "github.com/grafana/crossplane-provider-grafana/internal/controller/cloud/apikey"
	stack "github.com/grafana/crossplane-provider-grafana/internal/controller/cloud/stack"
	stackserviceaccount "github.com/grafana/crossplane-provider-grafana/internal/controller/cloud/stackserviceaccount"
	stackserviceaccounttoken "github.com/grafana/crossplane-provider-grafana/internal/controller/cloud/stackserviceaccounttoken"
	datasourcepermission "github.com/grafana/crossplane-provider-grafana/internal/controller/enterprise/datasourcepermission"
	report "github.com/grafana/crossplane-provider-grafana/internal/controller/enterprise/report"
	role "github.com/grafana/crossplane-provider-grafana/internal/controller/enterprise/role"
	roleassignment "github.com/grafana/crossplane-provider-grafana/internal/controller/enterprise/roleassignment"
	teamexternalgroup "github.com/grafana/crossplane-provider-grafana/internal/controller/enterprise/teamexternalgroup"
	escalation "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/escalation"
	escalationchain "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/escalationchain"
	integration "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/integration"
	oncallshift "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/oncallshift"
	outgoingwebhook "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/outgoingwebhook"
	route "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/route"
	schedule "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/schedule"
	apikeyoss "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/apikey"
	dashboard "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/dashboard"
	dashboardpermission "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/dashboardpermission"
	datasource "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/datasource"
	folder "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/folder"
	folderpermission "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/folderpermission"
	organization "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/organization"
	organizationpreferences "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/organizationpreferences"
	serviceaccount "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/serviceaccount"
	serviceaccountpermission "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/serviceaccountpermission"
	serviceaccounttoken "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/serviceaccounttoken"
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
		stackserviceaccount.Setup,
		stackserviceaccounttoken.Setup,
		datasourcepermission.Setup,
		report.Setup,
		role.Setup,
		roleassignment.Setup,
		teamexternalgroup.Setup,
		escalation.Setup,
		escalationchain.Setup,
		integration.Setup,
		oncallshift.Setup,
		outgoingwebhook.Setup,
		route.Setup,
		schedule.Setup,
		apikeyoss.Setup,
		dashboard.Setup,
		dashboardpermission.Setup,
		datasource.Setup,
		folder.Setup,
		folderpermission.Setup,
		organization.Setup,
		organizationpreferences.Setup,
		serviceaccount.Setup,
		serviceaccountpermission.Setup,
		serviceaccounttoken.Setup,
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
