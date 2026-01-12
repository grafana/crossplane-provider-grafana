/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	alertenrichmentv1beta1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/alerting/alertenrichmentv1beta1"
	contactpoint "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/alerting/contactpoint"
	messagetemplate "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/alerting/messagetemplate"
	mutetiming "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/alerting/mutetiming"
	notificationpolicy "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/alerting/notificationpolicy"
	rulegroup "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/alerting/rulegroup"
	custommodelrules "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/asserts/custommodelrules"
	logconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/asserts/logconfig"
	notificationalertsconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/asserts/notificationalertsconfig"
	suppressedassertionsconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/asserts/suppressedassertionsconfig"
	accesspolicy "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/accesspolicy"
	accesspolicytoken "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/accesspolicytoken"
	appo11yconfigv1alpha1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/appo11yconfigv1alpha1"
	k8so11yconfigv1alpha1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/k8so11yconfigv1alpha1"
	orgmember "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/orgmember"
	plugininstallation "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/plugininstallation"
	privatedatasourceconnectnetwork "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/privatedatasourceconnectnetwork"
	privatedatasourceconnectnetworktoken "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/privatedatasourceconnectnetworktoken"
	stack "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/stack"
	stackserviceaccount "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/stackserviceaccount"
	stackserviceaccounttoken "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/stackserviceaccounttoken"
	awsaccount "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloudprovider/awsaccount"
	awscloudwatchscrapejob "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloudprovider/awscloudwatchscrapejob"
	awsresourcemetadatascrapejob "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloudprovider/awsresourcemetadatascrapejob"
	azurecredential "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloudprovider/azurecredential"
	metricsendpointscrapejob "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/connections/metricsendpointscrapejob"
	datasourceconfiglbacrules "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/datasourceconfiglbacrules"
	datasourcepermission "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/datasourcepermission"
	datasourcepermissionitem "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/datasourcepermissionitem"
	report "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/report"
	role "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/role"
	roleassignment "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/roleassignment"
	roleassignmentitem "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/roleassignmentitem"
	scimconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/scimconfig"
	teamexternalgroup "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/teamexternalgroup"
	collector "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/fleetmanagement/collector"
	pipeline "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/fleetmanagement/pipeline"
	app "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/frontendobservability/app"
	installation "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/k6/installation"
	loadtest "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/k6/loadtest"
	project "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/k6/project"
	projectallowedloadzones "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/k6/projectallowedloadzones"
	projectlimits "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/k6/projectlimits"
	schedule "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/k6/schedule"
	alert "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/ml/alert"
	holiday "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/ml/holiday"
	job "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/ml/job"
	outlierdetector "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/ml/outlierdetector"
	escalation "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oncall/escalation"
	escalationchain "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oncall/escalationchain"
	integration "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oncall/integration"
	oncallshift "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oncall/oncallshift"
	outgoingwebhook "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oncall/outgoingwebhook"
	route "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oncall/route"
	scheduleoncall "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oncall/schedule"
	usernotificationrule "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oncall/usernotificationrule"
	annotation "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/annotation"
	dashboard "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/dashboard"
	dashboardpermission "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/dashboardpermission"
	dashboardpermissionitem "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/dashboardpermissionitem"
	dashboardpublic "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/dashboardpublic"
	dashboardv1beta1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/dashboardv1beta1"
	datasource "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/datasource"
	datasourceconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/datasourceconfig"
	folder "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/folder"
	folderpermission "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/folderpermission"
	folderpermissionitem "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/folderpermissionitem"
	librarypanel "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/librarypanel"
	organization "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/organization"
	organizationpreferences "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/organizationpreferences"
	playlist "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/playlist"
	playlistv0alpha1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/playlistv0alpha1"
	serviceaccount "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/serviceaccount"
	serviceaccountpermission "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/serviceaccountpermission"
	serviceaccountpermissionitem "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/serviceaccountpermissionitem"
	serviceaccounttoken "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/serviceaccounttoken"
	ssosettings "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/ssosettings"
	team "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/team"
	user "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/user"
	providerconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/providerconfig"
	slo "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/slo/slo"
	check "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/sm/check"
	checkalerts "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/sm/checkalerts"
	installationsm "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/sm/installation"
	probe "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/sm/probe"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alertenrichmentv1beta1.Setup,
		contactpoint.Setup,
		messagetemplate.Setup,
		mutetiming.Setup,
		notificationpolicy.Setup,
		rulegroup.Setup,
		custommodelrules.Setup,
		logconfig.Setup,
		notificationalertsconfig.Setup,
		suppressedassertionsconfig.Setup,
		accesspolicy.Setup,
		accesspolicytoken.Setup,
		appo11yconfigv1alpha1.Setup,
		k8so11yconfigv1alpha1.Setup,
		orgmember.Setup,
		plugininstallation.Setup,
		privatedatasourceconnectnetwork.Setup,
		privatedatasourceconnectnetworktoken.Setup,
		stack.Setup,
		stackserviceaccount.Setup,
		stackserviceaccounttoken.Setup,
		awsaccount.Setup,
		awscloudwatchscrapejob.Setup,
		awsresourcemetadatascrapejob.Setup,
		azurecredential.Setup,
		metricsendpointscrapejob.Setup,
		datasourceconfiglbacrules.Setup,
		datasourcepermission.Setup,
		datasourcepermissionitem.Setup,
		report.Setup,
		role.Setup,
		roleassignment.Setup,
		roleassignmentitem.Setup,
		scimconfig.Setup,
		teamexternalgroup.Setup,
		collector.Setup,
		pipeline.Setup,
		app.Setup,
		installation.Setup,
		loadtest.Setup,
		project.Setup,
		projectallowedloadzones.Setup,
		projectlimits.Setup,
		schedule.Setup,
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
		scheduleoncall.Setup,
		usernotificationrule.Setup,
		annotation.Setup,
		dashboard.Setup,
		dashboardpermission.Setup,
		dashboardpermissionitem.Setup,
		dashboardpublic.Setup,
		dashboardv1beta1.Setup,
		datasource.Setup,
		datasourceconfig.Setup,
		folder.Setup,
		folderpermission.Setup,
		folderpermissionitem.Setup,
		librarypanel.Setup,
		organization.Setup,
		organizationpreferences.Setup,
		playlist.Setup,
		playlistv0alpha1.Setup,
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
		checkalerts.Setup,
		installationsm.Setup,
		probe.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alertenrichmentv1beta1.SetupGated,
		contactpoint.SetupGated,
		messagetemplate.SetupGated,
		mutetiming.SetupGated,
		notificationpolicy.SetupGated,
		rulegroup.SetupGated,
		custommodelrules.SetupGated,
		logconfig.SetupGated,
		notificationalertsconfig.SetupGated,
		suppressedassertionsconfig.SetupGated,
		accesspolicy.SetupGated,
		accesspolicytoken.SetupGated,
		appo11yconfigv1alpha1.SetupGated,
		k8so11yconfigv1alpha1.SetupGated,
		orgmember.SetupGated,
		plugininstallation.SetupGated,
		privatedatasourceconnectnetwork.SetupGated,
		privatedatasourceconnectnetworktoken.SetupGated,
		stack.SetupGated,
		stackserviceaccount.SetupGated,
		stackserviceaccounttoken.SetupGated,
		awsaccount.SetupGated,
		awscloudwatchscrapejob.SetupGated,
		awsresourcemetadatascrapejob.SetupGated,
		azurecredential.SetupGated,
		metricsendpointscrapejob.SetupGated,
		datasourceconfiglbacrules.SetupGated,
		datasourcepermission.SetupGated,
		datasourcepermissionitem.SetupGated,
		report.SetupGated,
		role.SetupGated,
		roleassignment.SetupGated,
		roleassignmentitem.SetupGated,
		scimconfig.SetupGated,
		teamexternalgroup.SetupGated,
		collector.SetupGated,
		pipeline.SetupGated,
		app.SetupGated,
		installation.SetupGated,
		loadtest.SetupGated,
		project.SetupGated,
		projectallowedloadzones.SetupGated,
		projectlimits.SetupGated,
		schedule.SetupGated,
		alert.SetupGated,
		holiday.SetupGated,
		job.SetupGated,
		outlierdetector.SetupGated,
		escalation.SetupGated,
		escalationchain.SetupGated,
		integration.SetupGated,
		oncallshift.SetupGated,
		outgoingwebhook.SetupGated,
		route.SetupGated,
		scheduleoncall.SetupGated,
		usernotificationrule.SetupGated,
		annotation.SetupGated,
		dashboard.SetupGated,
		dashboardpermission.SetupGated,
		dashboardpermissionitem.SetupGated,
		dashboardpublic.SetupGated,
		dashboardv1beta1.SetupGated,
		datasource.SetupGated,
		datasourceconfig.SetupGated,
		folder.SetupGated,
		folderpermission.SetupGated,
		folderpermissionitem.SetupGated,
		librarypanel.SetupGated,
		organization.SetupGated,
		organizationpreferences.SetupGated,
		playlist.SetupGated,
		playlistv0alpha1.SetupGated,
		serviceaccount.SetupGated,
		serviceaccountpermission.SetupGated,
		serviceaccountpermissionitem.SetupGated,
		serviceaccounttoken.SetupGated,
		ssosettings.SetupGated,
		team.SetupGated,
		user.SetupGated,
		providerconfig.SetupGated,
		slo.SetupGated,
		check.SetupGated,
		checkalerts.SetupGated,
		installationsm.SetupGated,
		probe.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
