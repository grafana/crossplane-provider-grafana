package grafana

import (
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/grafana/crossplane-provider-grafana/config/grafana"
)

func computedFieldExtractor(field string) string {
	return fmt.Sprintf("%s.ComputedFieldExtractor(%q)", SelfPackagePath, field)
}

// nolint: golint
func ComputedFieldExtractor(field string) reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("status.atProvider." + field)
		if err != nil {
			return ""
		}
		return r
	}
}

func fieldExtractor(field string) string {
	return fmt.Sprintf("%s.FieldExtractor(%q)", SelfPackagePath, field)
}

// nolint: golint
func FieldExtractor(field string) reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("spec.forProvider." + field)
		if err != nil {
			return ""
		}
		return r
	}
}

// nolint: unparam
func optionalFieldExtractor(field string) string {
	return fmt.Sprintf("%s.OptionalFieldExtractor(%q)", SelfPackagePath, field)
}

// nolint: golint
func OptionalFieldExtractor(field string) reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		res := FieldExtractor(field)(mg)
		if res != "" {
			return res
		}
		return ComputedFieldExtractor(field)(mg)
	}
}
