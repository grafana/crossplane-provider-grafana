/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	contactpoint "github.com/grafana/crossplane-provider-grafana/internal/controller/alerting/contactpoint"
	messagetemplate "github.com/grafana/crossplane-provider-grafana/internal/controller/alerting/messagetemplate"
	mutetiming "github.com/grafana/crossplane-provider-grafana/internal/controller/alerting/mutetiming"
	notificationpolicy "github.com/grafana/crossplane-provider-grafana/internal/controller/alerting/notificationpolicy"
	rulegroup "github.com/grafana/crossplane-provider-grafana/internal/controller/alerting/rulegroup"
	accesspolicy "github.com/grafana/crossplane-provider-grafana/internal/controller/cloud/accesspolicy"
	accesspolicytoken "github.com/grafana/crossplane-provider-grafana/internal/controller/cloud/accesspolicytoken"
	plugininstallation "github.com/grafana/crossplane-provider-grafana/internal/controller/cloud/plugininstallation"
	stack "github.com/grafana/crossplane-provider-grafana/internal/controller/cloud/stack"
	stackserviceaccount "github.com/grafana/crossplane-provider-grafana/internal/controller/cloud/stackserviceaccount"
	stackserviceaccounttoken "github.com/grafana/crossplane-provider-grafana/internal/controller/cloud/stackserviceaccounttoken"
	datasourcepermission "github.com/grafana/crossplane-provider-grafana/internal/controller/enterprise/datasourcepermission"
	report "github.com/grafana/crossplane-provider-grafana/internal/controller/enterprise/report"
	role "github.com/grafana/crossplane-provider-grafana/internal/controller/enterprise/role"
	roleassignment "github.com/grafana/crossplane-provider-grafana/internal/controller/enterprise/roleassignment"
	teamexternalgroup "github.com/grafana/crossplane-provider-grafana/internal/controller/enterprise/teamexternalgroup"
	holiday "github.com/grafana/crossplane-provider-grafana/internal/controller/ml/holiday"
	job "github.com/grafana/crossplane-provider-grafana/internal/controller/ml/job"
	outlierdetector "github.com/grafana/crossplane-provider-grafana/internal/controller/ml/outlierdetector"
	escalation "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/escalation"
	escalationchain "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/escalationchain"
	integration "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/integration"
	oncallshift "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/oncallshift"
	outgoingwebhook "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/outgoingwebhook"
	route "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/route"
	schedule "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/schedule"
	annotation "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/annotation"
	dashboard "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/dashboard"
	dashboardpermission "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/dashboardpermission"
	dashboardpublic "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/dashboardpublic"
	datasource "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/datasource"
	folder "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/folder"
	folderpermission "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/folderpermission"
	librarypanel "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/librarypanel"
	organization "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/organization"
	organizationpreferences "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/organizationpreferences"
	playlist "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/playlist"
	serviceaccount "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/serviceaccount"
	serviceaccountpermission "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/serviceaccountpermission"
	serviceaccounttoken "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/serviceaccounttoken"
	ssosettings "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/ssosettings"
	team "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/team"
	user "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/user"
	providerconfig "github.com/grafana/crossplane-provider-grafana/internal/controller/providerconfig"
	slo "github.com/grafana/crossplane-provider-grafana/internal/controller/slo/slo"
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
		accesspolicy.Setup,
		accesspolicytoken.Setup,
		plugininstallation.Setup,
		stack.Setup,
		stackserviceaccount.Setup,
		stackserviceaccounttoken.Setup,
		datasourcepermission.Setup,
		report.Setup,
		role.Setup,
		roleassignment.Setup,
		teamexternalgroup.Setup,
		holiday.Setup,
		job.Setup,
		outlierdetector.Setup,
		escalation.Setup,
		escalationchain.Setup,
		integration.Setup,
		oncallshift.Setup,
		outgoingwebhook.Setup,
		route.Setup,
		schedule.Setup,
		annotation.Setup,
		dashboard.Setup,
		dashboardpermission.Setup,
		dashboardpublic.Setup,
		datasource.Setup,
		folder.Setup,
		folderpermission.Setup,
		librarypanel.Setup,
		organization.Setup,
		organizationpreferences.Setup,
		playlist.Setup,
		serviceaccount.Setup,
		serviceaccountpermission.Setup,
		serviceaccounttoken.Setup,
		ssosettings.Setup,
		team.Setup,
		user.Setup,
		providerconfig.Setup,
		slo.Setup,
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
