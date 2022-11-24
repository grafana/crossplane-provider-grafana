package grafana

import (
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

func UserEmailExtractor() reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		r, err := paved.GetString("status.forProvider.email")
		if err != nil {
			// todo(hasan): should we log this error?
			return ""
		}
		return r
	}
}
