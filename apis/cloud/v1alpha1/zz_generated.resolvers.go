/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	grafana "github.com/grafana/crossplane-provider-grafana/config/grafana"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AccessPolicy.
func (mg *AccessPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Realm); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Realm[i3].Identifier),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Realm[i3].StackRef,
			Selector:     mg.Spec.ForProvider.Realm[i3].StackSelector,
			To: reference.To{
				List:    &StackList{},
				Managed: &Stack{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Realm[i3].Identifier")
		}
		mg.Spec.ForProvider.Realm[i3].Identifier = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Realm[i3].StackRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Realm); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Realm[i3].Identifier),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.Realm[i3].StackRef,
			Selector:     mg.Spec.InitProvider.Realm[i3].StackSelector,
			To: reference.To{
				List:    &StackList{},
				Managed: &Stack{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Realm[i3].Identifier")
		}
		mg.Spec.InitProvider.Realm[i3].Identifier = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Realm[i3].StackRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this AccessPolicyToken.
func (mg *AccessPolicyToken) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccessPolicyID),
		Extract:      grafana.PolicyIDExtractor(),
		Reference:    mg.Spec.ForProvider.AccessPolicyRef,
		Selector:     mg.Spec.ForProvider.AccessPolicySelector,
		To: reference.To{
			List:    &AccessPolicyList{},
			Managed: &AccessPolicy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccessPolicyID")
	}
	mg.Spec.ForProvider.AccessPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccessPolicyRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AccessPolicyID),
		Extract:      grafana.PolicyIDExtractor(),
		Reference:    mg.Spec.InitProvider.AccessPolicyRef,
		Selector:     mg.Spec.InitProvider.AccessPolicySelector,
		To: reference.To{
			List:    &AccessPolicyList{},
			Managed: &AccessPolicy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AccessPolicyID")
	}
	mg.Spec.InitProvider.AccessPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AccessPolicyRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PluginInstallation.
func (mg *PluginInstallation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StackSlug),
		Extract:      grafana.CloudStackSlugExtractor(),
		Reference:    mg.Spec.ForProvider.CloudStackRef,
		Selector:     mg.Spec.ForProvider.CloudStackSelector,
		To: reference.To{
			List:    &StackList{},
			Managed: &Stack{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StackSlug")
	}
	mg.Spec.ForProvider.StackSlug = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CloudStackRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StackSlug),
		Extract:      grafana.CloudStackSlugExtractor(),
		Reference:    mg.Spec.InitProvider.CloudStackRef,
		Selector:     mg.Spec.InitProvider.CloudStackSelector,
		To: reference.To{
			List:    &StackList{},
			Managed: &Stack{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.StackSlug")
	}
	mg.Spec.InitProvider.StackSlug = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CloudStackRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this StackServiceAccount.
func (mg *StackServiceAccount) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StackSlug),
		Extract:      grafana.CloudStackSlugExtractor(),
		Reference:    mg.Spec.ForProvider.CloudStackRef,
		Selector:     mg.Spec.ForProvider.CloudStackSelector,
		To: reference.To{
			List:    &StackList{},
			Managed: &Stack{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StackSlug")
	}
	mg.Spec.ForProvider.StackSlug = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CloudStackRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StackSlug),
		Extract:      grafana.CloudStackSlugExtractor(),
		Reference:    mg.Spec.InitProvider.CloudStackRef,
		Selector:     mg.Spec.InitProvider.CloudStackSelector,
		To: reference.To{
			List:    &StackList{},
			Managed: &Stack{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.StackSlug")
	}
	mg.Spec.InitProvider.StackSlug = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CloudStackRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this StackServiceAccountToken.
func (mg *StackServiceAccountToken) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceAccountRef,
		Selector:     mg.Spec.ForProvider.ServiceAccountSelector,
		To: reference.To{
			List:    &StackServiceAccountList{},
			Managed: &StackServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceAccountID")
	}
	mg.Spec.ForProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StackSlug),
		Extract:      grafana.CloudStackSlugExtractor(),
		Reference:    mg.Spec.ForProvider.CloudStackRef,
		Selector:     mg.Spec.ForProvider.CloudStackSelector,
		To: reference.To{
			List:    &StackList{},
			Managed: &Stack{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StackSlug")
	}
	mg.Spec.ForProvider.StackSlug = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CloudStackRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ServiceAccountRef,
		Selector:     mg.Spec.InitProvider.ServiceAccountSelector,
		To: reference.To{
			List:    &StackServiceAccountList{},
			Managed: &StackServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceAccountID")
	}
	mg.Spec.InitProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceAccountRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StackSlug),
		Extract:      grafana.CloudStackSlugExtractor(),
		Reference:    mg.Spec.InitProvider.CloudStackRef,
		Selector:     mg.Spec.InitProvider.CloudStackSelector,
		To: reference.To{
			List:    &StackList{},
			Managed: &Stack{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.StackSlug")
	}
	mg.Spec.InitProvider.StackSlug = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CloudStackRef = rsp.ResolvedReference

	return nil
}
