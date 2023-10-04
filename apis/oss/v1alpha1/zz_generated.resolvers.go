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

// ResolveReferences of this APIKey.
func (mg *APIKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrgID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.OrganizationRef,
		Selector:     mg.Spec.ForProvider.OrganizationSelector,
		To: reference.To{
			List:    &OrganizationList{},
			Managed: &Organization{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OrgID")
	}
	mg.Spec.ForProvider.OrgID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrganizationRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Dashboard.
func (mg *Dashboard) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Folder),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderRef,
		Selector:     mg.Spec.ForProvider.FolderSelector,
		To: reference.To{
			List:    &FolderList{},
			Managed: &Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Folder")
	}
	mg.Spec.ForProvider.Folder = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrgID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.OrganizationRef,
		Selector:     mg.Spec.ForProvider.OrganizationSelector,
		To: reference.To{
			List:    &OrganizationList{},
			Managed: &Organization{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OrgID")
	}
	mg.Spec.ForProvider.OrgID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrganizationRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DashboardPermission.
func (mg *DashboardPermission) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DashboardUID),
		Extract:      grafana.UIDExtractor(),
		Reference:    mg.Spec.ForProvider.DashboardRef,
		Selector:     mg.Spec.ForProvider.DashboardSelector,
		To: reference.To{
			List:    &DashboardList{},
			Managed: &Dashboard{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DashboardUID")
	}
	mg.Spec.ForProvider.DashboardUID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DashboardRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrgID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.OrganizationRef,
		Selector:     mg.Spec.ForProvider.OrganizationSelector,
		To: reference.To{
			List:    &OrganizationList{},
			Managed: &Organization{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OrgID")
	}
	mg.Spec.ForProvider.OrgID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrganizationRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DataSource.
func (mg *DataSource) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrgID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.OrganizationRef,
		Selector:     mg.Spec.ForProvider.OrganizationSelector,
		To: reference.To{
			List:    &OrganizationList{},
			Managed: &Organization{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OrgID")
	}
	mg.Spec.ForProvider.OrgID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrganizationRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Folder.
func (mg *Folder) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrgID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.OrganizationRef,
		Selector:     mg.Spec.ForProvider.OrganizationSelector,
		To: reference.To{
			List:    &OrganizationList{},
			Managed: &Organization{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OrgID")
	}
	mg.Spec.ForProvider.OrgID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrganizationRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FolderPermission.
func (mg *FolderPermission) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderUID),
		Extract:      grafana.UIDExtractor(),
		Reference:    mg.Spec.ForProvider.FolderRef,
		Selector:     mg.Spec.ForProvider.FolderSelector,
		To: reference.To{
			List:    &FolderList{},
			Managed: &Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderUID")
	}
	mg.Spec.ForProvider.FolderUID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrgID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.OrganizationRef,
		Selector:     mg.Spec.ForProvider.OrganizationSelector,
		To: reference.To{
			List:    &OrganizationList{},
			Managed: &Organization{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OrgID")
	}
	mg.Spec.ForProvider.OrgID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrganizationRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Team.
func (mg *Team) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Members),
		Extract:       grafana.UserEmailExtractor(),
		References:    mg.Spec.ForProvider.MemberRefs,
		Selector:      mg.Spec.ForProvider.MemberSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Members")
	}
	mg.Spec.ForProvider.Members = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.MemberRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.OrgID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.OrganizationRef,
		Selector:     mg.Spec.ForProvider.OrganizationSelector,
		To: reference.To{
			List:    &OrganizationList{},
			Managed: &Organization{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.OrgID")
	}
	mg.Spec.ForProvider.OrgID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.OrganizationRef = rsp.ResolvedReference

	return nil
}
