/*
Copyright 2025 Grafana
*/

// Package tfdatasource provides a generic Crossplane controller that reconciles
// observe-only resources backed by Terraform data sources.
package tfdatasource

import (
	"context"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/v2/pkg/event"
	"github.com/crossplane/crossplane-runtime/v2/pkg/meta"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	tjcontroller "github.com/crossplane/upjet/v2/pkg/controller"

	grafanaProvider "github.com/grafana/terraform-provider-grafana/v4/pkg/provider"
	terraformSDK "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/grafana/crossplane-provider-grafana/v2/internal/clients"
)

// ReadFn wraps either a legacy SDK or Plugin Framework data source read call.
// It is responsible for populating AtProvider fields and setting the external
// name annotation on mg.
type ReadFn func(ctx context.Context, mg resource.Managed, providerMeta any) error

// Spec parameterizes one data-source-backed observe controller.
type Spec struct {
	// DataSourceName is the Terraform data source name, e.g. "grafana_oncall_team".
	DataSourceName string

	// ManagedKind is the GVK of the Crossplane resource.
	ManagedKind schema.GroupVersionKind

	// NewManaged returns a zero-value instance of the Crossplane resource.
	NewManaged func() resource.Managed

	// Read performs the data source read and populates the resource status.
	Read ReadFn

	// IsUpToDate optionally checks whether the resource status is already
	// current. If nil, the resource is always re-read.
	IsUpToDate func(resource.Managed) bool
}

// Setup adds a controller that reconciles an observe-only resource.
func Setup(mgr ctrl.Manager, o tjcontroller.Options, spec Spec) error {
	name := managed.ControllerName(spec.ManagedKind.String())
	r := managed.NewReconciler(mgr,
		resource.ManagedKind(spec.ManagedKind),
		managed.WithExternalConnecter(&connector{kube: mgr.GetClient(), spec: spec}),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithPollInterval(o.PollInterval),
	)
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(spec.NewManaged()).
		Complete(r)
}

// SetupGated registers the controller behind the SafeStart gate.
func SetupGated(mgr ctrl.Manager, o tjcontroller.Options, spec Spec) error {
	o.Gate.Register(func() {
		if err := Setup(mgr, o, spec); err != nil {
			mgr.GetLogger().Error(err, "unable to setup reconciler", "gvk", spec.ManagedKind.String())
		}
	}, spec.ManagedKind)
	return nil
}

// connector implements managed.ExternalConnector.
type connector struct {
	kube client.Client
	spec Spec
}

func (c *connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(resource.ModernManaged)
	if !ok {
		return nil, fmt.Errorf("managed resource %T is not a ModernManaged", mg)
	}

	cfg, creds, err := clients.ExtractModernConfig(ctx, c.kube, cr)
	if err != nil {
		return nil, errors.Wrap(err, "cannot extract provider config")
	}

	tfConfig := clients.BuildTFConfig(cfg, creds)

	p := grafanaProvider.Provider("crossplane")
	diags := p.Configure(ctx, terraformSDK.NewResourceConfigRaw(tfConfig))
	if diags.HasError() {
		return nil, fmt.Errorf("failed to configure the Grafana provider: %v", diags)
	}

	return &external{spec: c.spec, providerMeta: p.Meta()}, nil
}

// external implements managed.ExternalClient for observe-only resources.
type external struct {
	spec         Spec
	providerMeta any
}

func (e *external) Observe(_ context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	if meta.WasDeleted(mg) {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}

	upToDate := false
	if e.spec.IsUpToDate != nil {
		upToDate = e.spec.IsUpToDate(mg)
	}

	if upToDate {
		mg.(interface{ SetConditions(...xpv1.Condition) }).SetConditions(xpv1.Available())
	}

	return managed.ExternalObservation{
		ResourceExists:   true,
		ResourceUpToDate: upToDate,
	}, nil
}

func (e *external) Create(_ context.Context, _ resource.Managed) (managed.ExternalCreation, error) {
	return managed.ExternalCreation{}, errors.New("observe-only resource cannot be created")
}

func (e *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	if err := e.spec.Read(ctx, mg, e.providerMeta); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "data source read failed")
	}
	mg.(interface{ SetConditions(...xpv1.Condition) }).SetConditions(xpv1.Available())
	return managed.ExternalUpdate{}, nil
}

func (e *external) Delete(_ context.Context, _ resource.Managed) (managed.ExternalDelete, error) {
	return managed.ExternalDelete{}, nil
}

func (e *external) Disconnect(_ context.Context) error {
	return nil
}

// DeepEqualStatus is a convenience helper that checks whether two values are
// deeply equal. Spec implementations can use it to build IsUpToDate callbacks.
var DeepEqualStatus = reflect.DeepEqual
