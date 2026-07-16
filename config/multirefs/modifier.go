// Package multirefs adds strongly typed alternative references for a single
// Terraform field. It does so by creating synthetic schema fields and
// consolidating their resolved values into the original field before the
// Terraform configuration is used.
// Inspired by https://github.com/crossplane-contrib/provider-keycloak/blob/main/config/multitypes/modifier.go
// See also https://github.com/crossplane/upjet/issues/414
package multirefs

import (
	"fmt"
	"slices"
	"strings"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/types/name"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Alternative describes a synthetic field and its reference target.
type Alternative struct {
	// Name is the synthetic Terraform field name in snake_case. It is created
	// next to the original field, including when that field is nested.
	Name string

	// Reference configures the Ref and Selector generated for the synthetic
	// field.
	Reference ujconfig.Reference
}

// Add adds alternative strongly typed references for fieldPath while keeping
// the original field and its reference unchanged. Resolved synthetic values
// are consolidated into fieldPath by a TerraformConfigurationInjector.
func Add(r *ujconfig.Resource, fieldPath string, alternatives ...Alternative) {
	if len(alternatives) == 0 {
		return
	}

	parts := strings.Split(fieldPath, ".")
	parent := schemaMapAt(r.TerraformResource, parts[:len(parts)-1])
	originalName := parts[len(parts)-1]
	original, ok := parent[originalName]
	if !ok {
		panic(fmt.Sprintf("multirefs: Terraform field %q does not exist", fieldPath))
	}

	alternativeNames := make([]string, 0, len(alternatives))
	for _, alternative := range alternatives {
		if alternative.Name == "" || strings.Contains(alternative.Name, ".") {
			panic(fmt.Sprintf("multirefs: alternative name %q must be a non-empty field name", alternative.Name))
		}
		if _, exists := parent[alternative.Name]; exists {
			panic(fmt.Sprintf("multirefs: alternative field %q already exists next to %q", alternative.Name, fieldPath))
		}

		clone := *original
		parent[alternative.Name] = &clone

		alternativePath := strings.Join(append(append([]string{}, parts[:len(parts)-1]...), alternative.Name), ".")
		r.References[alternativePath] = alternative.Reference
		alternativeNames = append(alternativeNames, alternative.Name)
	}

	// A value injected through an alternative is observed under the original
	// Terraform field. Do not late-initialize that value back into the original
	// spec field, which would make two alternatives appear configured.
	if !slices.Contains(r.LateInitializer.IgnoredFields, fieldPath) {
		r.LateInitializer.IgnoredFields = append(r.LateInitializer.IgnoredFields, fieldPath)
	}

	r.TerraformConfigurationInjector = wrapInjector(
		r.TerraformConfigurationInjector,
		parts[:len(parts)-1],
		originalName,
		alternativeNames,
	)
}

func schemaMapAt(resource *schema.Resource, path []string) map[string]*schema.Schema {
	current := resource.Schema
	for _, segment := range path {
		field, ok := current[segment]
		if !ok {
			panic(fmt.Sprintf("multirefs: Terraform schema path segment %q does not exist", segment))
		}
		nested, ok := field.Elem.(*schema.Resource)
		if !ok {
			panic(fmt.Sprintf("multirefs: Terraform schema path segment %q is not an object", segment))
		}
		current = nested.Schema
	}
	return current
}

func wrapInjector(
	current ujconfig.ConfigurationInjector,
	parentPath []string,
	originalName string,
	alternativeNames []string,
) ujconfig.ConfigurationInjector {
	return func(jsonMap map[string]any, tfMap map[string]any) error {
		if current != nil {
			if err := current(jsonMap, tfMap); err != nil {
				return err
			}
		}

		return visitParents(jsonMap, tfMap, parentPath, func(jsonParent, tfParent map[string]any) error {
			configured := make([]string, 0, len(alternativeNames)+1)
			selectedName := ""
			var selectedValue any

			allNames := append([]string{originalName}, alternativeNames...)
			for _, fieldName := range allNames {
				jsonName := name.NewFromSnake(fieldName).LowerCamelComputed
				if value, exists := jsonParent[jsonName]; exists && value != nil {
					configured = append(configured, jsonName)
					selectedName = fieldName
					selectedValue = value
				}
			}

			if len(configured) > 1 {
				return fmt.Errorf("only one of %s may be configured", strings.Join(configured, ", "))
			}

			for _, alternativeName := range alternativeNames {
				delete(tfParent, alternativeName)
			}
			if selectedName != "" && selectedName != originalName {
				tfParent[originalName] = selectedValue
			}
			return nil
		})
	}
}

func visitParents(
	jsonNode any,
	tfNode any,
	path []string,
	visit func(jsonParent, tfParent map[string]any) error,
) error {
	jsonParent, jsonOK := jsonNode.(map[string]any)
	tfParent, tfOK := tfNode.(map[string]any)
	if !jsonOK || !tfOK {
		return nil
	}
	if len(path) == 0 {
		return visit(jsonParent, tfParent)
	}

	segment := path[0]
	jsonChild := jsonParent[name.NewFromSnake(segment).LowerCamelComputed]
	tfChild := tfParent[segment]

	switch children := jsonChild.(type) {
	case map[string]any:
		return visitParents(children, tfChild, path[1:], visit)
	case []any:
		tfChildren, ok := tfChild.([]any)
		if !ok {
			return nil
		}
		for i := range children {
			if i >= len(tfChildren) {
				break
			}
			if err := visitParents(children[i], tfChildren[i], path[1:], visit); err != nil {
				return err
			}
		}
	}
	return nil
}
