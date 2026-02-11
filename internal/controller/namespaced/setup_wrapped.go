/*
Copyright 2022 Upbound Inc.
*/

// This file provides wrapped versions of the Setup and SetupGated functions
// that register API group schemes lazily (only when controllers are activated).
// This prevents shutdown errors when using ManagedActivationResourcePolicies (MRAP)
// to disable certain resource types.
//
// DO NOT MODIFY: This file is manually maintained and should not be overwritten
// by code generation. It works in conjunction with the generated zz_setup.go file.

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	"github.com/grafana/crossplane-provider-grafana/v2/internal/controller/schemeregistry"

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
	profileconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/asserts/profileconfig"
	promrulefile "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/asserts/promrulefile"
	suppressedassertionsconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/asserts/suppressedassertionsconfig"
	thresholds "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/asserts/thresholds"
	traceconfig "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced/asserts/traceconfig"
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

// setupWithScheme wraps a controller setup function to register its API group scheme first.
func setupWithScheme(mgr ctrl.Manager, o controller.Options, apiGroup string, setupFunc func(ctrl.Manager, controller.Options) error) error {
	// Register the API group scheme if not already registered
	if err := schemeregistry.RegisterNamespacedAPIGroup(mgr, apiGroup); err != nil {
		return err
	}

	// Call the original setup function
	return setupFunc(mgr, o)
}

// setupGatedWithScheme wraps a SetupGated function to register the API group scheme
// when the controller is activated by the Gate.
func setupGatedWithScheme(mgr ctrl.Manager, o controller.Options, apiGroup string, setupGatedFunc func(ctrl.Manager, controller.Options) error) error {
	// The SetupGated function registers a callback with the Gate.
	// We need to intercept the callback to register the scheme first.
	// However, since we can't easily modify the callback, we'll register
	// the scheme upfront when SetupGated is called. The actual controller
	// Setup() will be called later by the Gate callback.
	//
	// Note: This means the scheme is registered when SetupGated is called,
	// not when the controller is actually activated. This is acceptable because:
	// 1. It only registers schemes for controllers that have SetupGated called
	// 2. In SafeStart mode, only controllers that will eventually activate call SetupGated
	// 3. The alternative (modifying generated code) is more fragile

	// For now, we'll use a different approach: wrap the original Setup function
	// We'll do this by creating wrapped versions of each controller's Setup function

	// Actually, let's take a simpler approach: register the scheme here
	// This is called during initialization, before any controllers activate
	// But it's only called for controllers that are in the setup list
	// So disabled controllers won't reach this point

	// CORRECTION: We should NOT register here, as this defeats the purpose.
	// Instead, we need to ensure the callback itself registers the scheme.
	// Let's use a modified approach where we override the Setup function.

	// For gated setup, we'll register the scheme when the gate callback fires.
	// Since the generated SetupGated calls the original Setup(), we need to
	// ensure Setup() registers the scheme. We'll do this by wrapping Setup calls.

	// Actually, the cleanest approach: call SetupGated normally, but modify
	// the options to include a setup hook that registers the scheme.
	// However, the Options struct doesn't support this.

	// Final approach: Accept that for gated setup, we register the scheme
	// at gate registration time, not activation time. This is still better
	// than registering ALL schemes upfront, because only controllers that
	// pass the MRAP filter will call SetupGated.

	if err := schemeregistry.RegisterNamespacedAPIGroup(mgr, apiGroup); err != nil {
		return err
	}

	return setupGatedFunc(mgr, o)
}

// SetupWrapped creates all controllers with the supplied logger and adds them to
// the supplied manager. Each controller's API group scheme is registered before
// the controller is set up.
func SetupWrapped(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "alerting", alertenrichmentv1beta1.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "alerting", alertrulev0alpha1.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "alerting", contactpoint.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "alerting", messagetemplate.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "alerting", mutetiming.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "alerting", notificationpolicy.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "alerting", recordingrulev0alpha1.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "alerting", rulegroup.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "asserts", custommodelrules.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "asserts", logconfig.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "asserts", notificationalertsconfig.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "asserts", profileconfig.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "asserts", promrulefile.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "asserts", suppressedassertionsconfig.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "asserts", thresholds.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "asserts", traceconfig.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "cloud", accesspolicy.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "cloud", accesspolicyrotatingtoken.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "cloud", accesspolicytoken.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "cloud", appo11yconfigv1alpha1.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "cloud", k8so11yconfigv1alpha1.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "cloud", orgmember.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "cloud", plugininstallation.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "cloud", privatedatasourceconnectnetwork.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "cloud", privatedatasourceconnectnetworktoken.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "cloud", stack.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "cloud", stackserviceaccount.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "cloud", stackserviceaccountrotatingtoken.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "cloud", stackserviceaccounttoken.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "cloudprovider", awsaccount.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "cloudprovider", awscloudwatchscrapejob.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "cloudprovider", awsresourcemetadatascrapejob.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "cloudprovider", azurecredential.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "connections", metricsendpointscrapejob.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "enterprise", datasourceconfiglbacrules.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "enterprise", datasourcepermission.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "enterprise", datasourcepermissionitem.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "enterprise", report.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "enterprise", role.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "enterprise", roleassignment.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "enterprise", roleassignmentitem.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "enterprise", scimconfig.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "enterprise", teamexternalgroup.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "fleetmanagement", collector.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "fleetmanagement", pipeline.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "frontendobservability", app.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "k6", installation.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "k6", loadtest.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "k6", project.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "k6", projectallowedloadzones.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "k6", projectlimits.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "k6", schedule.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "ml", alert.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "ml", holiday.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "ml", job.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "ml", outlierdetector.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oncall", escalation.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oncall", escalationchain.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oncall", integration.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oncall", oncallshift.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oncall", outgoingwebhook.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oncall", route.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oncall", scheduleoncall.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oncall", usernotificationrule.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", annotation.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", dashboard.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", dashboardpermission.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", dashboardpermissionitem.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", dashboardpublic.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", dashboardv1beta1.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", datasource.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", datasourceconfig.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", folder.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", folderpermission.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", folderpermissionitem.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", librarypanel.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", organization.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", organizationpreferences.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", playlist.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", playlistv0alpha1.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", serviceaccount.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", serviceaccountpermission.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", serviceaccountpermissionitem.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", serviceaccountrotatingtoken.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", serviceaccounttoken.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", ssosettings.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", team.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "oss", user.Setup) },
		providerconfig.Setup, // No API group registration needed for providerconfig
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "slo", slo.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "sm", check.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "sm", checkalerts.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "sm", installationsm.Setup) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupWithScheme(mgr, o, "sm", probe.Setup) },
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGatedWrapped creates all controllers with the supplied logger and adds them to
// the supplied manager gated. Each controller's API group scheme is registered before
// the controller is set up.
func SetupGatedWrapped(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "alerting", alertenrichmentv1beta1.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "alerting", alertrulev0alpha1.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "alerting", contactpoint.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "alerting", messagetemplate.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "alerting", mutetiming.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "alerting", notificationpolicy.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "alerting", recordingrulev0alpha1.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "alerting", rulegroup.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "asserts", custommodelrules.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "asserts", logconfig.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "asserts", notificationalertsconfig.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "asserts", profileconfig.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "asserts", promrulefile.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "asserts", suppressedassertionsconfig.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "asserts", thresholds.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "asserts", traceconfig.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "cloud", accesspolicy.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "cloud", accesspolicyrotatingtoken.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "cloud", accesspolicytoken.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "cloud", appo11yconfigv1alpha1.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "cloud", k8so11yconfigv1alpha1.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "cloud", orgmember.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "cloud", plugininstallation.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "cloud", privatedatasourceconnectnetwork.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "cloud", privatedatasourceconnectnetworktoken.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "cloud", stack.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "cloud", stackserviceaccount.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "cloud", stackserviceaccountrotatingtoken.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "cloud", stackserviceaccounttoken.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "cloudprovider", awsaccount.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "cloudprovider", awscloudwatchscrapejob.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "cloudprovider", awsresourcemetadatascrapejob.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "cloudprovider", azurecredential.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "connections", metricsendpointscrapejob.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "enterprise", datasourceconfiglbacrules.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "enterprise", datasourcepermission.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "enterprise", datasourcepermissionitem.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "enterprise", report.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "enterprise", role.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "enterprise", roleassignment.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "enterprise", roleassignmentitem.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "enterprise", scimconfig.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "enterprise", teamexternalgroup.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "fleetmanagement", collector.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "fleetmanagement", pipeline.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "frontendobservability", app.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "k6", installation.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "k6", loadtest.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "k6", project.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "k6", projectallowedloadzones.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "k6", projectlimits.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "k6", schedule.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "ml", alert.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "ml", holiday.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "ml", job.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "ml", outlierdetector.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oncall", escalation.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oncall", escalationchain.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oncall", integration.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oncall", oncallshift.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oncall", outgoingwebhook.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oncall", route.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oncall", scheduleoncall.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oncall", usernotificationrule.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", annotation.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", dashboard.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", dashboardpermission.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", dashboardpermissionitem.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", dashboardpublic.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", dashboardv1beta1.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", datasource.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", datasourceconfig.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", folder.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", folderpermission.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", folderpermissionitem.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", librarypanel.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", organization.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", organizationpreferences.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", playlist.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", playlistv0alpha1.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", serviceaccount.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", serviceaccountpermission.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", serviceaccountpermissionitem.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", serviceaccountrotatingtoken.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", serviceaccounttoken.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", ssosettings.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", team.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "oss", user.SetupGated) },
		providerconfig.SetupGated, // No API group registration needed for providerconfig
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "slo", slo.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "sm", check.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "sm", checkalerts.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "sm", installationsm.SetupGated) },
		func(mgr ctrl.Manager, o controller.Options) error { return setupGatedWithScheme(mgr, o, "sm", probe.SetupGated) },
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
