/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	alertenrichmentv1beta1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/alerting/alertenrichmentv1beta1"
	alertrulev0alpha1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/alerting/alertrulev0alpha1"
	contactpoint "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/alerting/contactpoint"
	messagetemplate "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/alerting/messagetemplate"
	mutetiming "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/alerting/mutetiming"
	notificationpolicy "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/alerting/notificationpolicy"
	recordingrulev0alpha1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/alerting/recordingrulev0alpha1"
	rulegroup "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/alerting/rulegroup"
	custommodelrules "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/asserts/custommodelrules"
	logconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/asserts/logconfig"
	notificationalertsconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/asserts/notificationalertsconfig"
	suppressedassertionsconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/asserts/suppressedassertionsconfig"
	thresholds "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/asserts/thresholds"
	accesspolicy "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/cloud/accesspolicy"
	accesspolicyrotatingtoken "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/cloud/accesspolicyrotatingtoken"
	accesspolicytoken "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/cloud/accesspolicytoken"
	appo11yconfigv1alpha1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/cloud/appo11yconfigv1alpha1"
	k8so11yconfigv1alpha1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/cloud/k8so11yconfigv1alpha1"
	orgmember "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/cloud/orgmember"
	plugininstallation "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/cloud/plugininstallation"
	privatedatasourceconnectnetwork "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/cloud/privatedatasourceconnectnetwork"
	privatedatasourceconnectnetworktoken "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/cloud/privatedatasourceconnectnetworktoken"
	stack "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/cloud/stack"
	stackserviceaccount "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/cloud/stackserviceaccount"
	stackserviceaccountrotatingtoken "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/cloud/stackserviceaccountrotatingtoken"
	stackserviceaccounttoken "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/cloud/stackserviceaccounttoken"
	awsaccount "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/cloudprovider/awsaccount"
	awscloudwatchscrapejob "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/cloudprovider/awscloudwatchscrapejob"
	awsresourcemetadatascrapejob "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/cloudprovider/awsresourcemetadatascrapejob"
	azurecredential "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/cloudprovider/azurecredential"
	metricsendpointscrapejob "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/connections/metricsendpointscrapejob"
	datasourceconfiglbacrules "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/enterprise/datasourceconfiglbacrules"
	datasourcepermission "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/enterprise/datasourcepermission"
	datasourcepermissionitem "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/enterprise/datasourcepermissionitem"
	report "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/enterprise/report"
	role "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/enterprise/role"
	roleassignment "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/enterprise/roleassignment"
	roleassignmentitem "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/enterprise/roleassignmentitem"
	scimconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/enterprise/scimconfig"
	teamexternalgroup "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/enterprise/teamexternalgroup"
	collector "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/fleetmanagement/collector"
	pipeline "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/fleetmanagement/pipeline"
	app "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/frontendobservability/app"
	installation "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/k6/installation"
	loadtest "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/k6/loadtest"
	project "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/k6/project"
	projectallowedloadzones "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/k6/projectallowedloadzones"
	projectlimits "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/k6/projectlimits"
	schedule "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/k6/schedule"
	alert "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/ml/alert"
	holiday "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/ml/holiday"
	job "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/ml/job"
	outlierdetector "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/ml/outlierdetector"
	escalation "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oncall/escalation"
	escalationchain "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oncall/escalationchain"
	integration "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oncall/integration"
	oncallshift "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oncall/oncallshift"
	outgoingwebhook "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oncall/outgoingwebhook"
	route "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oncall/route"
	scheduleoncall "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oncall/schedule"
	usernotificationrule "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oncall/usernotificationrule"
	annotation "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/annotation"
	dashboard "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/dashboard"
	dashboardpermission "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/dashboardpermission"
	dashboardpermissionitem "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/dashboardpermissionitem"
	dashboardpublic "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/dashboardpublic"
	dashboardv1beta1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/dashboardv1beta1"
	datasource "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/datasource"
	datasourceconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/datasourceconfig"
	folder "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/folder"
	folderpermission "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/folderpermission"
	folderpermissionitem "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/folderpermissionitem"
	librarypanel "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/librarypanel"
	organization "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/organization"
	organizationpreferences "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/organizationpreferences"
	playlist "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/playlist"
	playlistv0alpha1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/playlistv0alpha1"
	serviceaccount "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/serviceaccount"
	serviceaccountpermission "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/serviceaccountpermission"
	serviceaccountpermissionitem "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/serviceaccountpermissionitem"
	serviceaccountrotatingtoken "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/serviceaccountrotatingtoken"
	serviceaccounttoken "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/serviceaccounttoken"
	ssosettings "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/ssosettings"
	team "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/team"
	user "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/oss/user"
	providerconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/providerconfig"
	slo "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/slo/slo"
	check "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/sm/check"
	checkalerts "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/sm/checkalerts"
	installationsm "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/sm/installation"
	probe "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/sm/probe"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alertenrichmentv1beta1.Setup,
		alertrulev0alpha1.Setup,
		contactpoint.Setup,
		messagetemplate.Setup,
		mutetiming.Setup,
		notificationpolicy.Setup,
		recordingrulev0alpha1.Setup,
		rulegroup.Setup,
		custommodelrules.Setup,
		logconfig.Setup,
		notificationalertsconfig.Setup,
		suppressedassertionsconfig.Setup,
		thresholds.Setup,
		accesspolicy.Setup,
		accesspolicyrotatingtoken.Setup,
		accesspolicytoken.Setup,
		appo11yconfigv1alpha1.Setup,
		k8so11yconfigv1alpha1.Setup,
		orgmember.Setup,
		plugininstallation.Setup,
		privatedatasourceconnectnetwork.Setup,
		privatedatasourceconnectnetworktoken.Setup,
		stack.Setup,
		stackserviceaccount.Setup,
		stackserviceaccountrotatingtoken.Setup,
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
		serviceaccountrotatingtoken.Setup,
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
		alertrulev0alpha1.SetupGated,
		contactpoint.SetupGated,
		messagetemplate.SetupGated,
		mutetiming.SetupGated,
		notificationpolicy.SetupGated,
		recordingrulev0alpha1.SetupGated,
		rulegroup.SetupGated,
		custommodelrules.SetupGated,
		logconfig.SetupGated,
		notificationalertsconfig.SetupGated,
		suppressedassertionsconfig.SetupGated,
		thresholds.SetupGated,
		accesspolicy.SetupGated,
		accesspolicyrotatingtoken.SetupGated,
		accesspolicytoken.SetupGated,
		appo11yconfigv1alpha1.SetupGated,
		k8so11yconfigv1alpha1.SetupGated,
		orgmember.SetupGated,
		plugininstallation.SetupGated,
		privatedatasourceconnectnetwork.SetupGated,
		privatedatasourceconnectnetworktoken.SetupGated,
		stack.SetupGated,
		stackserviceaccount.SetupGated,
		stackserviceaccountrotatingtoken.SetupGated,
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
		serviceaccountrotatingtoken.SetupGated,
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
