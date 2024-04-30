/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package playlist

import (
	"time"

	"github.com/crossplane/crossplane-runtime/pkg/connection"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/crossplane-runtime/pkg/statemetrics"
	tjcontroller "github.com/crossplane/upjet/pkg/controller"
	"github.com/crossplane/upjet/pkg/controller/handler"
	"github.com/crossplane/upjet/pkg/metrics"
	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"

	v1alpha1 "github.com/grafana/crossplane-provider-grafana/apis/oss/v1alpha1"
	features "github.com/grafana/crossplane-provider-grafana/internal/features"
)

// Setup adds a controller that reconciles Playlist managed resources.
func Setup(mgr ctrl.Manager, o tjcontroller.Options) error {
	name := managed.ControllerName(v1alpha1.Playlist_GroupVersionKind.String())
	var initializers managed.InitializerChain
	cps := []managed.ConnectionPublisher{managed.NewAPISecretPublisher(mgr.GetClient(), mgr.GetScheme())}
	if o.SecretStoreConfigGVK != nil {
		cps = append(cps, connection.NewDetailsManager(mgr.GetClient(), *o.SecretStoreConfigGVK, connection.WithTLSConfig(o.ESSOptions.TLSConfig)))
	}
	eventHandler := handler.NewEventHandler(handler.WithLogger(o.Logger.WithValues("gvk", v1alpha1.Playlist_GroupVersionKind)))
	opts := []managed.ReconcilerOption{
		managed.WithExternalConnecter(tjcontroller.NewTerraformPluginSDKConnector(mgr.GetClient(), o.SetupFn, o.Provider.Resources["grafana_playlist"], o.OperationTrackerStore,
			tjcontroller.WithTerraformPluginSDKLogger(o.Logger),
			tjcontroller.WithTerraformPluginSDKMetricRecorder(metrics.NewMetricRecorder(v1alpha1.Playlist_GroupVersionKind, mgr, o.PollInterval)),
			tjcontroller.WithTerraformPluginSDKManagementPolicies(o.Features.Enabled(features.EnableBetaManagementPolicies)))),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithFinalizer(tjcontroller.NewOperationTrackerFinalizer(o.OperationTrackerStore, xpresource.NewAPIFinalizer(mgr.GetClient(), managed.FinalizerName))),
		managed.WithTimeout(3 * time.Minute),
		managed.WithInitializers(initializers),
		managed.WithConnectionPublishers(cps...),
		managed.WithPollInterval(o.PollInterval),
	}
	if o.PollJitter != 0 {
		opts = append(opts, managed.WithPollJitterHook(o.PollJitter))
	}
	if o.Features.Enabled(features.EnableBetaManagementPolicies) {
		opts = append(opts, managed.WithManagementPolicies())
	}
	if o.MetricOptions != nil {
		opts = append(opts, managed.WithMetricRecorder(o.MetricOptions.MRMetrics))
	}

	// register webhooks for the kind v1alpha1.Playlist
	// if they're enabled.
	if o.StartWebhooks {
		if err := ctrl.NewWebhookManagedBy(mgr).
			For(&v1alpha1.Playlist{}).
			Complete(); err != nil {
			return errors.Wrap(err, "cannot register webhook for the kind v1alpha1.Playlist")
		}
	}

	if o.MetricOptions != nil && o.MetricOptions.MRStateMetrics != nil {
		stateMetricsRecorder := statemetrics.NewMRStateRecorder(
			mgr.GetClient(), o.Logger, o.MetricOptions.MRStateMetrics, &v1alpha1.PlaylistList{}, o.MetricOptions.PollStateMetricInterval,
		)
		if err := mgr.Add(stateMetricsRecorder); err != nil {
			return errors.Wrap(err, "cannot register MR state metrics recorder for kind v1alpha1.PlaylistList")
		}
	}

	r := managed.NewReconciler(mgr, xpresource.ManagedKind(v1alpha1.Playlist_GroupVersionKind), opts...)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		WithEventFilter(xpresource.DesiredStateChanged()).
		Watches(&v1alpha1.Playlist{}, eventHandler).
		Complete(ratelimiter.NewReconciler(name, r, o.GlobalRateLimiter))
}
