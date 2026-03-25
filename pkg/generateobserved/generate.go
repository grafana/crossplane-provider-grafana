/*
Copyright 2025 Grafana
*/

package generateobserved

import (
	"context"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	fwschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	sdkschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dsInfo holds metadata about a single data source for code generation.
type dsInfo struct {
	tfName      string
	kindName    string
	fileName    string
	catInfo     CategoryRule
	isLegacySDK bool

	legacySchema *sdkschema.Resource
	fwDS         datasource.DataSource

	forProviderFields []fieldInfo
	atProviderFields  []fieldInfo
}

type fieldInfo struct {
	tfName      string
	goName      string
	jsonName    string
	goType      string
	required    bool
	description string

	// nestedFields holds the fields of a nested struct type (for List/Set of objects).
	nestedFields []fieldInfo
	// nestedStructName is the Go type name for the nested struct (e.g. "UsersUser").
	nestedStructName string
	// isSet is true when the field is a TypeSet (vs TypeList). Affects how data is extracted.
	isSet bool
}

// Generate introspects the given TF providers and emits all observed resource
// code (types, specs, factories, registration boilerplate).
//
// legacyProvider may be nil if the TF provider has no legacy SDK data sources.
// frameworkProvider may be nil if it has no Plugin Framework data sources.
func Generate(cfg Config, legacyProvider *sdkschema.Provider, frameworkProvider fwprovider.Provider) {
	grouped := map[string][]*dsInfo{}

	if legacyProvider != nil {
		collectLegacyDataSources(cfg, legacyProvider, grouped)
	}
	if frameworkProvider != nil {
		collectFrameworkDataSources(cfg, frameworkProvider, grouped)
	}

	// Detect singular/plural kind collisions and rename the plural kind.
	// When both e.g. "User" and "Users" exist, K8s pluralizes both to "users",
	// causing CRD overwrites. Rename "Users" → "UserSet" to avoid this.
	for _, dsList := range grouped {
		kindSet := map[string]*dsInfo{}
		for _, ds := range dsList {
			kindSet[ds.kindName] = ds
		}
		for _, ds := range dsList {
			singular := strings.TrimSuffix(ds.kindName, "s")
			if singular != ds.kindName && kindSet[singular] != nil {
				ds.kindName = singular + "Set"
				ds.fileName = strings.ToLower(singular) + "set"
			}
		}
	}

	groupNames := sortedGroupNames(grouped)
	emitFiles(cfg, grouped, groupNames)

	log.Printf("Generated observed resources for %d categories, %d total data sources", len(groupNames), countTotal(grouped))
}

// =============================================================================
// Data source collection
// =============================================================================

func collectLegacyDataSources(cfg Config, p *sdkschema.Provider, grouped map[string][]*dsInfo) {
	for name, ds := range p.DataSourcesMap {
		ci := categorize(cfg, name)
		if ci == nil {
			log.Printf("WARN: skipping legacy data source %q (no category match)", name)
			continue
		}
		info := &dsInfo{
			tfName:       name,
			catInfo:      *ci,
			isLegacySDK:  true,
			legacySchema: ds,
		}
		info.kindName, info.fileName = deriveNames(cfg, name, ci.TFPrefix)
		parseLegacySchema(cfg, info)
		grouped[ci.DirName] = append(grouped[ci.DirName], info)
	}
}

func collectFrameworkDataSources(cfg Config, fwp fwprovider.Provider, grouped map[string][]*dsInfo) {
	ctx := context.Background()
	var fwpMetaResp fwprovider.MetadataResponse
	fwp.Metadata(ctx, fwprovider.MetadataRequest{}, &fwpMetaResp)
	providerTypeName := fwpMetaResp.TypeName

	for _, newDS := range fwp.DataSources(ctx) {
		ds := newDS()
		var metaResp datasource.MetadataResponse
		ds.Metadata(ctx, datasource.MetadataRequest{ProviderTypeName: providerTypeName}, &metaResp)
		name := metaResp.TypeName

		ci := categorize(cfg, name)
		if ci == nil {
			log.Printf("WARN: skipping framework data source %q (no category match)", name)
			continue
		}

		info := &dsInfo{
			tfName:      name,
			catInfo:     *ci,
			isLegacySDK: false,
			fwDS:        ds,
		}
		info.kindName, info.fileName = deriveNames(cfg, name, ci.TFPrefix)
		parseFWSchema(cfg, info, ds)
		grouped[ci.DirName] = append(grouped[ci.DirName], info)
	}
}

// =============================================================================
// Categorization and naming
// =============================================================================

func categorize(cfg Config, name string) *CategoryRule {
	if override, ok := cfg.CategoryOverrides[name]; ok {
		return &override
	}
	for _, rule := range cfg.CategoryRules {
		if strings.HasPrefix(name, rule.TFPrefix) {
			cp := rule
			return &cp
		}
	}
	return nil
}

func deriveNames(cfg Config, tfName, tfPrefix string) (kindName, fileName string) {
	stripped := strings.TrimPrefix(tfName, tfPrefix)
	if stripped == tfName {
		stripped = strings.TrimPrefix(tfName, cfg.FallbackTFPrefix)
	}
	return snakeToCamel(cfg, stripped), strings.ReplaceAll(stripped, "_", "")
}

func sortedGroupNames(grouped map[string][]*dsInfo) []string {
	names := make([]string, 0, len(grouped))
	for name := range grouped {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		sort.Slice(grouped[name], func(i, j int) bool {
			return grouped[name][i].tfName < grouped[name][j].tfName
		})
	}
	return names
}

func countTotal(grouped map[string][]*dsInfo) int {
	n := 0
	for _, v := range grouped {
		n += len(v)
	}
	return n
}

// =============================================================================
// Schema parsing
// =============================================================================

func parseLegacySchema(cfg Config, info *dsInfo) {
	if info.legacySchema == nil || info.legacySchema.Schema == nil {
		return
	}
	for name, field := range info.legacySchema.Schema {
		if name == "id" {
			continue
		}
		fi := buildLegacyFieldInfo(cfg, info.kindName, name, field)
		if field.Computed && !field.Required && !field.Optional {
			info.atProviderFields = append(info.atProviderFields, fi)
		} else {
			info.forProviderFields = append(info.forProviderFields, fi)
		}
	}
	sortFields(info)
}

func buildLegacyFieldInfo(cfg Config, parentName, name string, field *sdkschema.Schema) fieldInfo {
	fi := fieldInfo{
		tfName:      name,
		goName:      snakeToCamel(cfg, name),
		jsonName:    snakeToCamelJSON(cfg, name),
		goType:      sdkTypeToGo(field),
		required:    field.Required,
		description: field.Description,
	}

	// Handle nested resource elements in List/Set types.
	if field.Type == sdkschema.TypeList || field.Type == sdkschema.TypeSet {
		if res, ok := field.Elem.(*sdkschema.Resource); ok {
			structName := parentName + snakeToCamel(cfg, name)
			fi.nestedStructName = structName
			fi.goType = "[]" + structName
			fi.isSet = field.Type == sdkschema.TypeSet

			nested := make([]fieldInfo, 0, len(res.Schema))
			for nestedName, nestedField := range res.Schema {
				nfi := buildLegacyFieldInfo(cfg, structName, nestedName, nestedField)
				nested = append(nested, nfi)
			}
			sort.Slice(nested, func(i, j int) bool { return nested[i].tfName < nested[j].tfName })
			fi.nestedFields = nested
		}
	}

	return fi
}

func parseFWSchema(cfg Config, info *dsInfo, ds datasource.DataSource) {
	ctx := context.Background()
	type schemaer interface {
		Schema(context.Context, datasource.SchemaRequest, *datasource.SchemaResponse)
	}
	s, ok := ds.(schemaer)
	if !ok {
		log.Printf("WARN: data source %q does not implement Schema", info.tfName)
		return
	}
	var resp datasource.SchemaResponse
	s.Schema(ctx, datasource.SchemaRequest{}, &resp)
	if resp.Diagnostics.HasError() {
		log.Printf("WARN: schema for %q has errors: %v", info.tfName, resp.Diagnostics.Errors())
		return
	}

	for name, attr := range resp.Schema.Attributes {
		if name == "id" {
			continue
		}
		fi := fieldInfo{
			tfName:      name,
			goName:      snakeToCamel(cfg, name),
			jsonName:    snakeToCamelJSON(cfg, name),
			goType:      fwAttrTypeToGo(attr),
			required:    attr.IsRequired(),
			description: attr.GetMarkdownDescription(),
		}
		if fi.description == "" {
			fi.description = attr.GetDescription()
		}
		if attr.IsComputed() && !attr.IsRequired() && !attr.IsOptional() {
			info.atProviderFields = append(info.atProviderFields, fi)
		} else {
			info.forProviderFields = append(info.forProviderFields, fi)
		}
	}
	sortFields(info)
}

func sortFields(info *dsInfo) {
	sort.Slice(info.forProviderFields, func(i, j int) bool { return info.forProviderFields[i].tfName < info.forProviderFields[j].tfName })
	sort.Slice(info.atProviderFields, func(i, j int) bool { return info.atProviderFields[i].tfName < info.atProviderFields[j].tfName })
}

func sdkTypeToGo(field *sdkschema.Schema) string {
	switch field.Type { //nolint:exhaustive // TypeInvalid is not a valid field type.
	case sdkschema.TypeString:
		if field.Required {
			return goTypeString
		}
		return goTypePtrString
	case sdkschema.TypeInt:
		if field.Required {
			return goTypeInt64
		}
		return goTypePtrInt64
	case sdkschema.TypeFloat:
		if field.Required {
			return goTypeFloat64
		}
		return goTypePtrFloat
	case sdkschema.TypeBool:
		if field.Required {
			return goTypeBool
		}
		return goTypePtrBool
	case sdkschema.TypeList, sdkschema.TypeSet:
		if elem, ok := field.Elem.(*sdkschema.Schema); ok {
			switch elem.Type { //nolint:exhaustive // Only string and int element types are relevant.
			case sdkschema.TypeString:
				return "[]string"
			case sdkschema.TypeInt:
				return "[]int64"
			default:
				return "[]string"
			}
		}
		return "[]string"
	case sdkschema.TypeMap:
		return "map[string]string"
	default:
		return goTypeString
	}
}

func fwAttrTypeToGo(attr fwschema.Attribute) string {
	typeName := attr.GetType().String()
	required := attr.IsRequired()
	switch {
	case strings.Contains(typeName, "StringType"):
		if required {
			return goTypeString
		}
		return goTypePtrString
	case strings.Contains(typeName, "Int64Type"):
		if required {
			return goTypeInt64
		}
		return goTypePtrInt64
	case strings.Contains(typeName, "Float64Type"):
		if required {
			return goTypeFloat64
		}
		return goTypePtrFloat
	case strings.Contains(typeName, "BoolType"):
		if required {
			return goTypeBool
		}
		return goTypePtrBool
	case strings.Contains(typeName, "ListType"), strings.Contains(typeName, "SetType"):
		return "[]string"
	case strings.Contains(typeName, "MapType"):
		return "map[string]string"
	default:
		if required {
			return goTypeString
		}
		return goTypePtrString
	}
}

// =============================================================================
// File emission
// =============================================================================

func emitFiles(cfg Config, grouped map[string][]*dsInfo, groupNames []string) {
	apisBase := "apis/observed"
	ctrlBase := "internal/controller/namespaced/observed"
	examplesBase := "examples-generated/observed"

	for _, groupName := range groupNames {
		dsList := grouped[groupName]
		ci := dsList[0].catInfo

		apiDir := filepath.Join(apisBase, groupName, cfg.APIVersion)
		ctrlDir := filepath.Join(ctrlBase, groupName)
		examplesDir := filepath.Join(examplesBase, groupName, cfg.APIVersion)

		mustMkdirAll(apiDir)
		mustMkdirAll(ctrlDir)
		mustMkdirAll(examplesDir)

		writeFormatted(filepath.Join(apiDir, "zz_groupversion_info.go"), generateGroupVersionInfo(cfg, ci))

		for _, ds := range dsList {
			writeFormatted(filepath.Join(apiDir, "zz_"+ds.fileName+"_types.go"), generateTypes(cfg, ds))
			writeFormatted(filepath.Join(ctrlDir, "zz_"+ds.fileName+"_spec.go"), generateSpec(cfg, ds, ci))
			writeRaw(filepath.Join(examplesDir, ds.fileName+".yaml"), generateExample(cfg, ds, ci))
		}

		writeFormatted(filepath.Join(ctrlDir, "zz_setup.go"), generateGroupSetup(cfg, dsList, ci))

		if hasFrameworkDS(dsList) {
			writeFormatted(filepath.Join(ctrlDir, "zz_factories.go"), generateFactories(cfg, dsList, ci))
		}
	}

	mustMkdirAll(apisBase)
	writeFormatted(filepath.Join(apisBase, "zz_register.go"), generateRegister(cfg, groupNames))

	mustMkdirAll(ctrlBase)
	writeFormatted(filepath.Join(ctrlBase, "zz_setup.go"), generateTopSetup(cfg, groupNames))
}

func mustMkdirAll(path string) {
	if err := os.MkdirAll(path, 0o750); err != nil {
		log.Fatalf("mkdir %s: %v", path, err)
	}
}

func writeFormatted(path, content string) {
	formatted, err := format.Source([]byte(content))
	if err != nil {
		log.Printf("WARN: gofmt failed for %s: %v (writing unformatted)", path, err)
		formatted = []byte(content)
	}
	if err := os.WriteFile(path, formatted, 0o600); err != nil {
		log.Fatalf("write %s: %v", path, err)
	}
	log.Printf("  wrote %s", path)
}

func writeRaw(path, content string) {
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		log.Fatalf("write %s: %v", path, err)
	}
	log.Printf("  wrote %s", path)
}
