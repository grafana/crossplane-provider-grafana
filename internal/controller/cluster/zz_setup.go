/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	notificationalertsconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/asserts/notificationalertsconfig"
suppressedassertionsconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/asserts/suppressedassertionsconfig"
stackcloud "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/stack"
awsresourcemetadatascrapejob "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloudprovider/awsresourcemetadatascrapejob"
keeperactivationv1beta1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/keeperactivationv1beta1"
pipeline "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/fleetmanagement/pipeline"
alert "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/ml/alert"
annotation "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/annotation"
contactpoint "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/alerting/contactpoint"
inhibitionrulev0alpha1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/alerting/inhibitionrulev0alpha1"
orgmember "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/orgmember"
datasourcepermission "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/datasourcepermission"
serviceaccount "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/serviceaccount"
serviceaccountpermission "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/serviceaccountpermission"
thresholds "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/asserts/thresholds"
scimconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/scimconfig"
serviceaccounttoken "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/serviceaccounttoken"
accesspolicytoken "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/accesspolicytoken"
stackserviceaccountrotatingtoken "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/stackserviceaccountrotatingtoken"
stackserviceaccounttoken "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/stackserviceaccounttoken"
dashboardpermission "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/dashboardpermission"
dashboardv2beta1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/dashboardv2beta1"
datasource "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/datasource"
notificationpolicy "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/alerting/notificationpolicy"
traceconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/asserts/traceconfig"
app "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/frontendobservability/app"
datasourceconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/datasourceconfig"
playlist "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/playlist"
accesspolicy "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/accesspolicy"
k8so11yconfigv1alpha1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/k8so11yconfigv1alpha1"
datasourceconfiglbacrules "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/datasourceconfiglbacrules"
collector "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/fleetmanagement/collector"
escalation "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oncall/escalation"
integration "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oncall/integration"
folderpermission "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/folderpermission"
organizationpreferences "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/organizationpreferences"
stackserviceaccount "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/stackserviceaccount"
installation "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/k6/installation"
projectallowedloadzones "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/k6/projectallowedloadzones"
outgoingwebhook "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oncall/outgoingwebhook"
serviceaccountpermissionitem "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/serviceaccountpermissionitem"
slo "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/slo/slo"
logconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/asserts/logconfig"
securevaluev1beta1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/securevaluev1beta1"
project "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/k6/project"
projectlimits "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/k6/projectlimits"
user "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/user"
check "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/sm/check"
probe "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/sm/probe"
stack "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/asserts/stack"
loadtest "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/k6/loadtest"
providerconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/providerconfig"
alertenrichmentv1beta1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/alerting/alertenrichmentv1beta1"
alertrulev0alpha1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/alerting/alertrulev0alpha1"
awsaccount "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloudprovider/awsaccount"
awscloudwatchscrapejob "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloudprovider/awscloudwatchscrapejob"
oncallshift "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oncall/oncallshift"
route "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oncall/route"
scheduleoncall "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oncall/schedule"
dashboardpermissionitem "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/dashboardpermissionitem"
mutetiming "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/alerting/mutetiming"
recordingrulev0alpha1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/alerting/recordingrulev0alpha1"
rulegroup "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/alerting/rulegroup"
folder "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/folder"
folderpermissionitem "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/folderpermissionitem"
organization "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/organization"
privatedatasourceconnectnetworktoken "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/privatedatasourceconnectnetworktoken"
role "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/role"
roleassignment "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/roleassignment"
teamexternalgroup "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/teamexternalgroup"
job "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/ml/job"
usernotificationrule "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oncall/usernotificationrule"
dashboard "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/dashboard"
librarypanel "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/librarypanel"
custommodelrules "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/asserts/custommodelrules"
profileconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/asserts/profileconfig"
datasourcecacheconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/datasourcecacheconfig"
outlierdetector "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/ml/outlierdetector"
escalationchain "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oncall/escalationchain"
dashboardv1beta1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/dashboardv1beta1"
playlistv0alpha1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/playlistv0alpha1"
ssosettings "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/ssosettings"
promrulefile "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/asserts/promrulefile"
keeperv1beta1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/keeperv1beta1"
holiday "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/ml/holiday"
dashboardpublic "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/dashboardpublic"
accesspolicyrotatingtoken "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/accesspolicyrotatingtoken"
appo11yconfigv1alpha1 "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/appo11yconfigv1alpha1"
plugininstallation "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/plugininstallation"
privatedatasourceconnectnetwork "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloud/privatedatasourceconnectnetwork"
metricsendpointscrapejob "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/connections/metricsendpointscrapejob"
report "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/report"
serviceaccountrotatingtoken "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/serviceaccountrotatingtoken"
installationsm "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/sm/installation"
messagetemplate "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/alerting/messagetemplate"
azurecredential "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/cloudprovider/azurecredential"
datasourcepermissionitem "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/datasourcepermissionitem"
roleassignmentitem "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/enterprise/roleassignmentitem"
schedule "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/k6/schedule"
team "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/oss/team"
checkalerts "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster/sm/checkalerts"

)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alertenrichmentv1beta1.Setup,
		alertrulev0alpha1.Setup,
		contactpoint.Setup,
		inhibitionrulev0alpha1.Setup,
		messagetemplate.Setup,
		mutetiming.Setup,
		notificationpolicy.Setup,
		recordingrulev0alpha1.Setup,
		rulegroup.Setup,
		custommodelrules.Setup,
		logconfig.Setup,
		notificationalertsconfig.Setup,
		profileconfig.Setup,
		promrulefile.Setup,
		stack.Setup,
		suppressedassertionsconfig.Setup,
		thresholds.Setup,
		traceconfig.Setup,
		accesspolicy.Setup,
		accesspolicyrotatingtoken.Setup,
		accesspolicytoken.Setup,
		appo11yconfigv1alpha1.Setup,
		k8so11yconfigv1alpha1.Setup,
		orgmember.Setup,
		plugininstallation.Setup,
		privatedatasourceconnectnetwork.Setup,
		privatedatasourceconnectnetworktoken.Setup,
		stackcloud.Setup,
		stackserviceaccount.Setup,
		stackserviceaccountrotatingtoken.Setup,
		stackserviceaccounttoken.Setup,
		awsaccount.Setup,
		awscloudwatchscrapejob.Setup,
		awsresourcemetadatascrapejob.Setup,
		azurecredential.Setup,
		metricsendpointscrapejob.Setup,
		datasourcecacheconfig.Setup,
		datasourceconfiglbacrules.Setup,
		datasourcepermission.Setup,
		datasourcepermissionitem.Setup,
		keeperactivationv1beta1.Setup,
		keeperv1beta1.Setup,
		report.Setup,
		role.Setup,
		roleassignment.Setup,
		roleassignmentitem.Setup,
		scimconfig.Setup,
		securevaluev1beta1.Setup,
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
		dashboardv2beta1.Setup,
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
		inhibitionrulev0alpha1.SetupGated,
		messagetemplate.SetupGated,
		mutetiming.SetupGated,
		notificationpolicy.SetupGated,
		recordingrulev0alpha1.SetupGated,
		rulegroup.SetupGated,
		custommodelrules.SetupGated,
		logconfig.SetupGated,
		notificationalertsconfig.SetupGated,
		profileconfig.SetupGated,
		promrulefile.SetupGated,
		stack.SetupGated,
		suppressedassertionsconfig.SetupGated,
		thresholds.SetupGated,
		traceconfig.SetupGated,
		accesspolicy.SetupGated,
		accesspolicyrotatingtoken.SetupGated,
		accesspolicytoken.SetupGated,
		appo11yconfigv1alpha1.SetupGated,
		k8so11yconfigv1alpha1.SetupGated,
		orgmember.SetupGated,
		plugininstallation.SetupGated,
		privatedatasourceconnectnetwork.SetupGated,
		privatedatasourceconnectnetworktoken.SetupGated,
		stackcloud.SetupGated,
		stackserviceaccount.SetupGated,
		stackserviceaccountrotatingtoken.SetupGated,
		stackserviceaccounttoken.SetupGated,
		awsaccount.SetupGated,
		awscloudwatchscrapejob.SetupGated,
		awsresourcemetadatascrapejob.SetupGated,
		azurecredential.SetupGated,
		metricsendpointscrapejob.SetupGated,
		datasourcecacheconfig.SetupGated,
		datasourceconfiglbacrules.SetupGated,
		datasourcepermission.SetupGated,
		datasourcepermissionitem.SetupGated,
		keeperactivationv1beta1.SetupGated,
		keeperv1beta1.SetupGated,
		report.SetupGated,
		role.SetupGated,
		roleassignment.SetupGated,
		roleassignmentitem.SetupGated,
		scimconfig.SetupGated,
		securevaluev1beta1.SetupGated,
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
		dashboardv2beta1.SetupGated,
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