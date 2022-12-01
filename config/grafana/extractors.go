package grafana

import (
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

// nolint: golint
func CloudStackSlugExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("spec.forProvider.slug")
		if err != nil {
			return ""
		}
		return r
	}
}

// nolint: golint
func DashboardIDExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("status.atProvider.dashboard_id")
		if err != nil {
			return ""
		}
		return r
	}
}

// nolint: golint
func NameExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("spec.forProvider.name")
		if err != nil {
			return ""
		}
		return r
	}
}

// nolint: golint
func UIDExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("spec.forProvider.uid")
		if err != nil {
			return ""
		}
		// UID is optional, so it can be in atProvider if it's not in forProvider
		if r == "" {
			r, err = paved.GetString("status.atProvider.uid")
			if err != nil {
				return ""
			}
		}
		return r
	}
}

// nolint: golint
func UserEmailExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("spec.forProvider.email")
		if err != nil {
			return ""
		}
		return r
	}
}
