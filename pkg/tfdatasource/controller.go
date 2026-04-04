/*
Copyright 2026 Grafana Labs
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
	tjresource "github.com/crossplane/upjet/v2/pkg/resource"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ConnectFn establishes a connection to the upstream provider and returns
// provider metadata (e.g. a configured client). The returned value is passed
// to ReadFn as providerMeta.
type ConnectFn func(ctx context.Context, kube client.Client, mg resource.Managed) (providerMeta any, err error)

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

	// ConnectFn establishes a connection to the upstream provider.
	// It is injected at controller setup time.
	ConnectFn ConnectFn

	// IsUpToDate optionally checks whether the resource status is already
	// current. If nil, the resource is always re-read.
	IsUpToDate func(resource.Managed) bool
}

// Setup adds a controller that reconciles an observe-only resource.
func Setup(mgr ctrl.Manager, o tjcontroller.Options, spec Spec) error {
	if spec.NewManaged == nil {
		return fmt.Errorf("tfdatasource: Spec.NewManaged must not be nil (kind %s)", spec.ManagedKind)
	}
	if spec.ConnectFn == nil {
		return fmt.Errorf("tfdatasource: Spec.ConnectFn must not be nil (kind %s)", spec.ManagedKind)
	}
	if spec.Read == nil {
		return fmt.Errorf("tfdatasource: Spec.Read must not be nil (kind %s)", spec.ManagedKind)
	}
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
	providerMeta, err := c.spec.ConnectFn(ctx, c.kube, mg)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to provider")
	}
	return &external{kube: c.kube, spec: c.spec, providerMeta: providerMeta}, nil
}

// external implements managed.ExternalClient for observe-only resources.
type external struct {
	kube         client.Client
	spec         Spec
	providerMeta any
}

// Observe checks whether the resource needs to be refreshed but does not
// perform the data source read itself. When ResourceUpToDate=false, the
// reconciler calls Update, which is where the actual Read happens. This means
// Observe effectively controls the poll interval: if IsUpToDate returns false
// (or is nil, the default), the resource is re-read on every poll cycle.
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
		tjresource.SetUpToDateCondition(mg, true)
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
	extNameBefore := meta.GetExternalName(mg)
	if err := e.spec.Read(ctx, mg, e.providerMeta); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "data source read failed")
	}

	// The managed reconciler only persists annotation changes after Create(),
	// not after Update(). Since observe-only resources never go through Create,
	// we must persist annotation changes (specifically external-name) ourselves.
	if meta.GetExternalName(mg) != extNameBefore {
		if err := e.kube.Update(ctx, mg.(client.Object)); err != nil {
			return managed.ExternalUpdate{}, errors.Wrap(err, "cannot persist external-name annotation")
		}
	}

	mg.(interface{ SetConditions(...xpv1.Condition) }).SetConditions(xpv1.Available())
	tjresource.SetUpToDateCondition(mg, true)
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
