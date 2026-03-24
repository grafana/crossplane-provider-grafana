/*
Copyright 2025 Grafana
*/

package teams

import (
	"context"
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

	v1alpha1observe "github.com/grafana/crossplane-provider-grafana/v2/apis/namespaced/observe/v1alpha1"
	"github.com/grafana/crossplane-provider-grafana/v2/internal/clients"

	oateams "github.com/grafana/grafana-openapi-client-go/client/teams"
)

// Setup adds a controller that reconciles Teams observe resources.
func Setup(mgr ctrl.Manager, o tjcontroller.Options) error {
	name := managed.ControllerName(v1alpha1observe.Teams_GroupVersionKind.String())
	r := managed.NewReconciler(mgr,
		resource.ManagedKind(v1alpha1observe.Teams_GroupVersionKind),
		managed.WithExternalConnecter(&connector{kube: mgr.GetClient()}),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithPollInterval(o.PollInterval),
	)
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&v1alpha1observe.Teams{}).
		Complete(r)
}

// SetupGated registers the Teams controller behind the SafeStart gate.
func SetupGated(mgr ctrl.Manager, o tjcontroller.Options) error {
	o.Gate.Register(func() {
		if err := Setup(mgr, o); err != nil {
			mgr.GetLogger().Error(err, "unable to setup reconciler", "gvk", v1alpha1observe.Teams_GroupVersionKind.String())
		}
	}, v1alpha1observe.Teams_GroupVersionKind)
	return nil
}

// connector implements managed.ExternalConnector.
type connector struct {
	kube client.Client
}

func (c *connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*v1alpha1observe.Teams)
	if !ok {
		return nil, errors.New("managed resource is not a Teams")
	}

	cfg, creds, err := clients.ExtractModernConfig(ctx, c.kube, cr)
	if err != nil {
		return nil, errors.Wrap(err, "cannot extract provider config")
	}

	auth := creds["auth"]
	grafanaURL := cfg.URL
	if urlOverride, ok := creds["url"]; ok && urlOverride != "" {
		grafanaURL = urlOverride
	}

	grafanaClient, err := clients.NewOAPIClient(grafanaURL, auth)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create Grafana OAPI client")
	}

	return &external{teamsClient: grafanaClient.Teams}, nil
}

// external implements managed.ExternalClient.
type external struct {
	teamsClient  oateams.ClientService
	lastObserved []v1alpha1observe.TeamSummary // populated by Observe, reused by Update
}

func (e *external) Observe(_ context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*v1alpha1observe.Teams)
	if !ok {
		return managed.ExternalObservation{}, errors.New("managed resource is not a Teams")
	}

	// When being deleted, report absent so the managed reconciler skips Delete
	// and removes the finalizer immediately. Without this, the reconciler loops
	// forever because Observe keeps returning ResourceExists: true.
	if meta.WasDeleted(cr) {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}

	allTeams, err := e.searchAllTeams(cr)
	if err != nil {
		return managed.ExternalObservation{}, err
	}
	// Normalize nil to empty slice so reflect.DeepEqual matches a status
	// that round-tripped through JSON (which deserializes null as []).
	if allTeams == nil {
		allTeams = []v1alpha1observe.TeamSummary{}
	}

	// Cache the result so Update can reuse it without a second API call.
	e.lastObserved = allTeams

	upToDate := reflect.DeepEqual(allTeams, cr.Status.AtProvider.Teams)
	if upToDate {
		cr.SetConditions(xpv1.Available())
	}

	return managed.ExternalObservation{
		ResourceExists:   true,
		ResourceUpToDate: upToDate,
	}, nil
}

func (e *external) Create(_ context.Context, _ resource.Managed) (managed.ExternalCreation, error) {
	return managed.ExternalCreation{}, errors.New("Teams is observe-only")
}

func (e *external) Update(_ context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*v1alpha1observe.Teams)
	if !ok {
		return managed.ExternalUpdate{}, errors.New("managed resource is not a Teams")
	}

	// Reuse the teams fetched during Observe to avoid a redundant API call.
	cr.Status.AtProvider.Teams = e.lastObserved
	return managed.ExternalUpdate{}, nil
}

func (e *external) Delete(_ context.Context, _ resource.Managed) (managed.ExternalDelete, error) {
	// No-op: Teams is a read-only view of the Grafana API; there is nothing to
	// delete in Grafana when this resource is removed from Kubernetes.
	return managed.ExternalDelete{}, nil
}

func (e *external) Disconnect(_ context.Context) error {
	return nil
}

// searchAllTeams pages through all teams matching the spec filters.
func (e *external) searchAllTeams(cr *v1alpha1observe.Teams) ([]v1alpha1observe.TeamSummary, error) {
	var allTeams []v1alpha1observe.TeamSummary
	page, perPage := int64(1), int64(1000)

	for {
		params := oateams.NewSearchTeamsParams().WithPage(&page).WithPerpage(&perPage)
		if cr.Spec.ForProvider.Name != nil {
			params = params.WithName(cr.Spec.ForProvider.Name)
		}
		if cr.Spec.ForProvider.Query != nil {
			params = params.WithQuery(cr.Spec.ForProvider.Query)
		}

		resp, err := e.teamsClient.SearchTeams(params)
		if err != nil {
			return nil, errors.Wrap(err, "cannot search teams")
		}

		payload := resp.GetPayload()
		for _, t := range payload.Teams {
			if t == nil {
				continue
			}
			allTeams = append(allTeams, v1alpha1observe.TeamSummary{
				ID:          t.ID,
				UID:         t.UID,
				Name:        t.Name,
				Email:       t.Email,
				MemberCount: t.MemberCount,
				OrgID:       t.OrgID,
			})
		}

		if int64(len(allTeams)) >= payload.TotalCount {
			break
		}
		page++
	}

	return allTeams, nil
}
