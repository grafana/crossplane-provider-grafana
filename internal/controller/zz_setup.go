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
	orgmember "github.com/grafana/crossplane-provider-grafana/internal/controller/cloud/orgmember"
	plugininstallation "github.com/grafana/crossplane-provider-grafana/internal/controller/cloud/plugininstallation"
	privatedatasourceconnectnetwork "github.com/grafana/crossplane-provider-grafana/internal/controller/cloud/privatedatasourceconnectnetwork"
	privatedatasourceconnectnetworktoken "github.com/grafana/crossplane-provider-grafana/internal/controller/cloud/privatedatasourceconnectnetworktoken"
	stack "github.com/grafana/crossplane-provider-grafana/internal/controller/cloud/stack"
	stackserviceaccount "github.com/grafana/crossplane-provider-grafana/internal/controller/cloud/stackserviceaccount"
	stackserviceaccounttoken "github.com/grafana/crossplane-provider-grafana/internal/controller/cloud/stackserviceaccounttoken"
	awsaccount "github.com/grafana/crossplane-provider-grafana/internal/controller/cloudprovider/awsaccount"
	awscloudwatchscrapejob "github.com/grafana/crossplane-provider-grafana/internal/controller/cloudprovider/awscloudwatchscrapejob"
	azurecredential "github.com/grafana/crossplane-provider-grafana/internal/controller/cloudprovider/azurecredential"
	metricsendpointscrapejob "github.com/grafana/crossplane-provider-grafana/internal/controller/connections/metricsendpointscrapejob"
	datasourceconfiglbacrules "github.com/grafana/crossplane-provider-grafana/internal/controller/enterprise/datasourceconfiglbacrules"
	datasourcepermission "github.com/grafana/crossplane-provider-grafana/internal/controller/enterprise/datasourcepermission"
	datasourcepermissionitem "github.com/grafana/crossplane-provider-grafana/internal/controller/enterprise/datasourcepermissionitem"
	report "github.com/grafana/crossplane-provider-grafana/internal/controller/enterprise/report"
	role "github.com/grafana/crossplane-provider-grafana/internal/controller/enterprise/role"
	roleassignment "github.com/grafana/crossplane-provider-grafana/internal/controller/enterprise/roleassignment"
	roleassignmentitem "github.com/grafana/crossplane-provider-grafana/internal/controller/enterprise/roleassignmentitem"
	teamexternalgroup "github.com/grafana/crossplane-provider-grafana/internal/controller/enterprise/teamexternalgroup"
	collector "github.com/grafana/crossplane-provider-grafana/internal/controller/fleetmanagement/collector"
	pipeline "github.com/grafana/crossplane-provider-grafana/internal/controller/fleetmanagement/pipeline"
	app "github.com/grafana/crossplane-provider-grafana/internal/controller/frontendobservability/app"
	alert "github.com/grafana/crossplane-provider-grafana/internal/controller/ml/alert"
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
	usernotificationrule "github.com/grafana/crossplane-provider-grafana/internal/controller/oncall/usernotificationrule"
	annotation "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/annotation"
	dashboard "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/dashboard"
	dashboardpermission "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/dashboardpermission"
	dashboardpermissionitem "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/dashboardpermissionitem"
	dashboardpublic "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/dashboardpublic"
	datasource "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/datasource"
	datasourceconfig "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/datasourceconfig"
	folder "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/folder"
	folderpermission "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/folderpermission"
	folderpermissionitem "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/folderpermissionitem"
	librarypanel "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/librarypanel"
	organization "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/organization"
	organizationpreferences "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/organizationpreferences"
	playlist "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/playlist"
	serviceaccount "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/serviceaccount"
	serviceaccountpermission "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/serviceaccountpermission"
	serviceaccountpermissionitem "github.com/grafana/crossplane-provider-grafana/internal/controller/oss/serviceaccountpermissionitem"
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
		orgmember.Setup,
		plugininstallation.Setup,
		privatedatasourceconnectnetwork.Setup,
		privatedatasourceconnectnetworktoken.Setup,
		stack.Setup,
		stackserviceaccount.Setup,
		stackserviceaccounttoken.Setup,
		awsaccount.Setup,
		awscloudwatchscrapejob.Setup,
		azurecredential.Setup,
		metricsendpointscrapejob.Setup,
		datasourceconfiglbacrules.Setup,
		datasourcepermission.Setup,
		datasourcepermissionitem.Setup,
		report.Setup,
		role.Setup,
		roleassignment.Setup,
		roleassignmentitem.Setup,
		teamexternalgroup.Setup,
		collector.Setup,
		pipeline.Setup,
		app.Setup,
		alert.Setup,
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
		usernotificationrule.Setup,
		annotation.Setup,
		dashboard.Setup,
		dashboardpermission.Setup,
		dashboardpermissionitem.Setup,
		dashboardpublic.Setup,
		datasource.Setup,
		datasourceconfig.Setup,
		folder.Setup,
		folderpermission.Setup,
		folderpermissionitem.Setup,
		librarypanel.Setup,
		organization.Setup,
		organizationpreferences.Setup,
		playlist.Setup,
		serviceaccount.Setup,
		serviceaccountpermission.Setup,
		serviceaccountpermissionitem.Setup,
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
