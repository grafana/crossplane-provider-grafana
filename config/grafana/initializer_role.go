package grafana

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func createroleInitializer(client client.Client) managed.Initializer {
	return &roleInitializer{
		kube: client,
	}
}

// Based on the Tagger: https://github.com/crossplane/upjet/blob/v1.1.0/pkg/config/resource.go#L268
type roleInitializer struct {
	kube client.Client
}

// If `autoIncrementVersion` is set to true, the `version` field will be set to the current value of `status.atProvider.version`.
// If the version field was never set, it will be deleted from the forProvider spec.
func (i *roleInitializer) Initialize(ctx context.Context, mg resource.Managed) error {
	paved, err := fieldpath.PaveObject(mg)
	if err != nil {
		return err
	}
	v, err := paved.GetBool("spec.forProvider.autoIncrementVersion")
	if err != nil {
		return fmt.Errorf("could not get autoIncrementVersion: %w", err)
	}
	if !v {
		return nil
	}
	atProviderValue, err := paved.GetValue("status.atProvider.version")
	if err == nil {
		if err2 := paved.SetValue("spec.forProvider.version", atProviderValue); err2 != nil {
			return fmt.Errorf("could not set version: %w", err2)
		}
	} else if err3 := paved.DeleteField("spec.forProvider.version"); err3 != nil {
		return fmt.Errorf("could not delete version: %w", err3)
	}
	pavedByte, err := paved.MarshalJSON()
	if err != nil {
		return fmt.Errorf("could not marshal modified role spec into JSON: %w", err)
	}
	if err := json.Unmarshal(pavedByte, mg); err != nil {
		return fmt.Errorf("could not unmarshal modified role spec into managed resource interface: %w", err)
	}
	if err := i.kube.Update(ctx, mg); err != nil {
		return fmt.Errorf("could not update managed resource: %w", err)
	}
	return nil
}
