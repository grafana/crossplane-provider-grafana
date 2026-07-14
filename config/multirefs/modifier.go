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
	addAlternativeReferences(r, fieldPath, false, alternatives...)
}

// AddList adds alternative strongly typed references for a list or set field
// while keeping the original field and its reference unchanged. Values
// configured through the original field and any alternatives are concatenated
// into the original Terraform field by a TerraformConfigurationInjector.
func AddList(r *ujconfig.Resource, fieldPath string, alternatives ...Alternative) {
	addAlternativeReferences(r, fieldPath, true, alternatives...)
}

// addAlternativeReferences creates synthetic schema fields beside fieldPath,
// registers their Crossplane references, and installs the runtime translation
// that maps their resolved values back to the original Terraform field.
func addAlternativeReferences(r *ujconfig.Resource, fieldPath string, mergeCollections bool, alternatives ...Alternative) {
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
	if mergeCollections && original.Type != schema.TypeList && original.Type != schema.TypeSet {
		panic(fmt.Sprintf("multirefs: Terraform field %q is not a list or set", fieldPath))
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

	r.TerraformConfigurationInjector = chainAlternativeReferenceInjector(
		r.TerraformConfigurationInjector,
		parts[:len(parts)-1],
		originalName,
		alternativeNames,
		mergeCollections,
	)
}

// schemaMapAt follows a Terraform schema path and returns the schema map that
// contains the final field. Each path segment must describe a nested object.
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

// parentVisitor processes the JSON and Terraform representations of the same
// object after visitParents has navigated to the configured field's parent.
type parentVisitor func(jsonParent, tfParent map[string]any) error

// alternativeReferenceInjector translates synthetic Crossplane fields back to
// the original field understood by the Terraform provider. Upjet resolves each
// Ref or Selector into the field beside it, including the synthetic alternative
// fields created by Add and AddList. Before Terraform receives the configuration,
// this injector selects or merges those resolved values, writes the result under
// originalName, and removes the synthetic Terraform keys.
type alternativeReferenceInjector struct {
	previous         ujconfig.ConfigurationInjector
	parentPath       []string
	originalName     string
	alternativeNames []string
	mergeCollections bool
}

// chainAlternativeReferenceInjector preserves an injector already configured
// on the resource and appends alternative-reference translation to it.
// Configuration features can install injectors independently, so the previous
// injector must run first and its error must stop further modification.
func chainAlternativeReferenceInjector(
	previous ujconfig.ConfigurationInjector,
	parentPath []string,
	originalName string,
	alternativeNames []string,
	mergeCollections bool,
) ujconfig.ConfigurationInjector {
	injector := &alternativeReferenceInjector{
		previous:         previous,
		parentPath:       parentPath,
		originalName:     originalName,
		alternativeNames: alternativeNames,
		mergeCollections: mergeCollections,
	}
	return injector.injectConfiguration
}

// injectConfiguration runs the existing injector and then translates
// alternative references at every object reached through parentPath.
func (i *alternativeReferenceInjector) injectConfiguration(jsonMap, tfMap map[string]any) error {
	if i.previous != nil {
		if err := i.previous(jsonMap, tfMap); err != nil {
			return err
		}
	}
	return visitParents(jsonMap, tfMap, i.parentPath, i.injectParent)
}

// injectParent applies scalar-selection or collection-merge semantics to one
// pair of corresponding JSON and Terraform parent objects.
func (i *alternativeReferenceInjector) injectParent(jsonParent, tfParent map[string]any) error {
	if i.mergeCollections {
		return i.mergeCollectionValues(jsonParent, tfParent)
	}
	return i.selectScalarValue(jsonParent, tfParent)
}

// mergeCollectionValues concatenates values from the original collection and
// every configured alternative. The original values come first, followed by
// alternatives in registration order.
func (i *alternativeReferenceInjector) mergeCollectionValues(jsonParent, tfParent map[string]any) error {
	merged := make([]any, 0)
	configured := false
	for _, fieldName := range i.allFieldNames() {
		jsonName := name.NewFromSnake(fieldName).LowerCamelComputed
		value, exists := jsonParent[jsonName]
		if !exists || value == nil {
			continue
		}
		values, ok := value.([]any)
		if !ok {
			return fmt.Errorf("field %s must be a list or set", jsonName)
		}
		configured = true
		merged = append(merged, values...)
	}

	i.removeAlternatives(tfParent)
	if configured {
		tfParent[i.originalName] = merged
	}
	return nil
}

// selectScalarValue requires the original field and its alternatives to be
// mutually exclusive. A selected alternative is moved to the original
// Terraform field; a directly configured original value is already in place.
func (i *alternativeReferenceInjector) selectScalarValue(jsonParent, tfParent map[string]any) error {
	configured := make([]string, 0, len(i.alternativeNames)+1)
	selectedName := ""
	var selectedValue any

	for _, fieldName := range i.allFieldNames() {
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

	i.removeAlternatives(tfParent)
	if selectedName != "" && selectedName != i.originalName {
		tfParent[i.originalName] = selectedValue
	}
	return nil
}

// allFieldNames returns the original Terraform field followed by its synthetic
// alternatives in registration order.
func (i *alternativeReferenceInjector) allFieldNames() []string {
	return append([]string{i.originalName}, i.alternativeNames...)
}

// removeAlternatives prevents synthetic fields unknown to the Terraform
// provider from being included in its configuration.
func (i *alternativeReferenceInjector) removeAlternatives(tfParent map[string]any) {
	for _, alternativeName := range i.alternativeNames {
		delete(tfParent, alternativeName)
	}
}

// visitParents walks matching JSON and Terraform object trees until path is
// exhausted, then calls visit for each corresponding parent object. JSON uses
// lowerCamel field names while Terraform uses snake_case names. Nested object
// collections are traversed element by element so references inside repeated
// blocks are translated independently.
func visitParents(
	jsonNode any,
	tfNode any,
	path []string,
	visit parentVisitor,
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
