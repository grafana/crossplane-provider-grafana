package grafana

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func createDashboardConfigInitializer(client client.Client) managed.Initializer {
	return &dashboardConfigInitializer{
		kube: client,
	}
}

// Based on the Tagger: https://github.com/crossplane/upjet/blob/v1.1.0/pkg/config/resource.go#L268
type dashboardConfigInitializer struct {
	kube client.Client
}

// Replaces ${ with $${ to avoid TF interpolation
func (i *dashboardConfigInitializer) Initialize(ctx context.Context, mg resource.Managed) error {
	paved, err := fieldpath.PaveObject(mg)
	if err != nil {
		return err
	}
	v, err := paved.GetString("spec.forProvider.configJson")
	if err != nil {
		return fmt.Errorf("could not get configJson: %w", err)
	}
	if err := paved.SetString("spec.forProvider.configJson", replaceInterpolation(v)); err != nil {
		return fmt.Errorf("could not set configJson: %w", err)
	}
	pavedByte, err := paved.MarshalJSON()
	if err != nil {
		return fmt.Errorf("could not marshal modified dashboard spec into JSON: %w", err)
	}
	if err := json.Unmarshal(pavedByte, mg); err != nil {
		return fmt.Errorf("could not unmarshal modified dashboard spec into managed resource interface: %w", err)
	}
	if err := i.kube.Update(ctx, mg); err != nil {
		return fmt.Errorf("could not update managed resource: %w", err)
	}
	return nil
}

// Replaces ${ with $${ to avoid TF interpolation
// If a string is already escaped, it should not be escaped again
func replaceInterpolation(s string) string {
	replaced := strings.ReplaceAll(s, "${", "$${")
	return strings.ReplaceAll(replaced, "$$${", "$${") // Unescape already escaped strings
}
