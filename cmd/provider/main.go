/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/crossplane/crossplane-runtime/v2/pkg/gate"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/customresourcesgate"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/v2/pkg/statemetrics"
	"go.uber.org/zap/zapcore"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/metrics"

	"github.com/alecthomas/kingpin/v2"
	xpcontroller "github.com/crossplane/crossplane-runtime/v2/pkg/controller"
	xpfeature "github.com/crossplane/crossplane-runtime/v2/pkg/feature"
	"github.com/crossplane/crossplane-runtime/v2/pkg/logging"
	"github.com/crossplane/crossplane-runtime/v2/pkg/ratelimiter"
	tjcontroller "github.com/crossplane/upjet/v2/pkg/controller"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	apisCluster "github.com/grafana/crossplane-provider-grafana/v2/apis/cluster"
	apisNamespaced "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced"
	"github.com/grafana/crossplane-provider-grafana/v2/config"
	"github.com/grafana/crossplane-provider-grafana/v2/internal/clients"
	clustercontroller "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/cluster"
	namespacedcontroller "github.com/grafana/crossplane-provider-grafana/v2/internal/controller/namespaced"
	"github.com/grafana/crossplane-provider-grafana/v2/internal/features"
)

// Inspired by the Azure provider: https://github.com/crossplane-contrib/provider-upjet-azure/blob/d6a52c46e243fd70d6a2859ec97f29da0d67efa2/cmd/provider/dbformysql/zz_main.go
func main() {
	var (
		app                     = kingpin.New(filepath.Base(os.Args[0]), "Terraform based Crossplane provider for Grafana").DefaultEnvars()
		debug                   = app.Flag("debug", "Run with debug logging.").Short('d').Bool()
		syncInterval            = app.Flag("sync", "Controller manager sync period such as 300ms, 1.5h, or 2h45m").Short('s').Default("1h").Duration()
		pollInterval            = app.Flag("poll", "Poll interval controls how often an individual resource should be checked for drift.").Default("10m").Duration()
		pollStateMetricInterval = app.Flag("poll-state-metric", "State metric recording interval").Default("5s").Duration()
		leaderElection          = app.Flag("leader-election", "Use leader election for the controller manager.").Short('l').Default("false").OverrideDefaultFromEnvar("LEADER_ELECTION").Bool()
		maxReconcileRate        = app.Flag("max-reconcile-rate", "The global maximum rate per second at which resources may checked for drift from the desired state.").Default("100").Int()

		enableManagementPolicies = app.Flag("enable-management-policies", "Enable support for Management Policies.").Default("true").Envar("ENABLE_MANAGEMENT_POLICIES").Bool()
		enableSafeStart          = app.Flag("safe-start", "Enable SafeStart gated controller initialization (phased rollout). ").Default("false").Envar("ENABLE_SAFE_START").Bool()
	)

	kingpin.MustParse(app.Parse(os.Args[1:]))
	log.Default().SetOutput(io.Discard)

	level := zap.Level(zapcore.InfoLevel)
	if *debug {
		level = zap.Level(zapcore.DebugLevel)
	}
	zl := zap.New(zap.UseDevMode(*debug), level).WithName("provider-grafana")
	ctrl.SetLogger(zl)
	logr := logging.NewLogrLogger(zl)

	// currently, we configure the jitter to be the 5% of the poll interval
	pollJitter := time.Duration(float64(*pollInterval) * 0.05)
	logr.Debug("Starting", "sync-interval", syncInterval.String(),
		"poll-interval", pollInterval.String(), "poll-jitter", pollJitter, "max-reconcile-rate", *maxReconcileRate)

	cfg, err := ctrl.GetConfig()
	kingpin.FatalIfError(err, "Cannot get API server rest config")

	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		LeaderElection:   *leaderElection,
		LeaderElectionID: "crossplane-leader-election-provider-grafana",
		Cache: cache.Options{
			SyncPeriod: syncInterval,
		},
		LeaderElectionResourceLock: resourcelock.LeasesResourceLock,
		LeaseDuration:              func() *time.Duration { d := 60 * time.Second; return &d }(),
		RenewDeadline:              func() *time.Duration { d := 50 * time.Second; return &d }(),
	})
	kingpin.FatalIfError(err, "Cannot create controller manager")
	kingpin.FatalIfError(apisCluster.AddToScheme(mgr.GetScheme()), "Cannot add Grafana cluster APIs to scheme")
	kingpin.FatalIfError(apisNamespaced.AddToScheme(mgr.GetScheme()), "Cannot add Grafana namespaced APIs to scheme")
	kingpin.FatalIfError(apiextensionsv1.AddToScheme(mgr.GetScheme()), "Cannot add CustomResourceDefinition to scheme")

	mm := managed.NewMRMetricRecorder()
	sm := statemetrics.NewMRStateMetrics()

	metrics.Registry.MustRegister(mm)
	metrics.Registry.MustRegister(sm)

	mo := &xpcontroller.MetricOptions{
		PollStateMetricInterval: *pollStateMetricInterval,
		MRMetrics:               mm,
		MRStateMetrics:          sm,
	}

	featureFlags := &xpfeature.Flags{}

	provider, err := config.GetProvider(false)
	kingpin.FatalIfError(err, "Cannot get provider configuration")

	o := tjcontroller.Options{
		Options: xpcontroller.Options{
			Logger:                  logr,
			GlobalRateLimiter:       ratelimiter.NewGlobal(*maxReconcileRate),
			PollInterval:            *pollInterval,
			MaxConcurrentReconciles: *maxReconcileRate,
			Features:                featureFlags,
			MetricOptions:           mo,
			Gate:                    new(gate.Gate[schema.GroupVersionKind]),
		},

		Provider:              provider,
		SetupFn:               clients.TerraformSetupBuilder(),
		PollJitter:            pollJitter,
		OperationTrackerStore: tjcontroller.NewOperationStore(logr),
	}
	if *enableManagementPolicies {
		o.Features.Enable(features.EnableBetaManagementPolicies)
		logr.Info("Beta feature enabled", "flag", features.EnableBetaManagementPolicies)
	}
	if *enableSafeStart {
		o.Features.Enable(features.EnableSafeStart)
		logr.Info("SafeStart enabled", "flag", features.EnableSafeStart)
	}

	// Setup controllers for both legacy cluster-scoped and modern namespaced API groups.
	// SafeStart gating selects the SetupGated variants which internally defer activation
	// until feature flags permit.
	if *enableSafeStart {
		kingpin.FatalIfError(clustercontroller.SetupGated(mgr, o), "Cannot setup gated cluster Grafana controllers")
		kingpin.FatalIfError(namespacedcontroller.SetupGated(mgr, o), "Cannot setup gated namespaced Grafana controllers")
		kingpin.FatalIfError(customresourcesgate.Setup(mgr, o.Options), "Cannot setup CRD gate controller")
	} else {
		kingpin.FatalIfError(clustercontroller.Setup(mgr, o), "Cannot setup cluster Grafana controllers")
		kingpin.FatalIfError(namespacedcontroller.Setup(mgr, o), "Cannot setup namespaced Grafana controllers")
	}
	kingpin.FatalIfError(mgr.Start(ctrl.SetupSignalHandler()), "Cannot start controller manager")
}
