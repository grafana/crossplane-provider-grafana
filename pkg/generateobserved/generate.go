/*
Copyright 2026 Grafana Labs
*/

package generateobserved

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	fwschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	fwtypes "github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	sdkschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/tools/imports"
)

// dsInfo holds metadata about a single data source for code generation.
// Fields are exported for access by Go templates.
type dsInfo struct {
	TFName      string
	KindName    string
	FileName    string
	CatInfo     CategoryRule
	IsLegacySDK bool

	LegacySchema *sdkschema.Resource
	FwDS         datasource.DataSource

	ForProviderFields []fieldInfo
	AtProviderFields  []fieldInfo
}

// fieldInfo describes a single schema field for code generation.
// Fields are exported for access by Go templates.
type fieldInfo struct {
	TFName      string
	GoName      string
	JSONName    string
	GoType      string
	Required    bool
	Description string

	// NestedFields holds the fields of a nested struct type (for List/Set of objects).
	NestedFields []fieldInfo
	// NestedStructName is the Go type name for the nested struct (e.g. "UsersUser").
	NestedStructName string
	// IsSet is true when the field is a TypeSet (vs TypeList). Affects how data is extracted.
	IsSet bool
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
			kindSet[ds.KindName] = ds
		}
		for _, ds := range dsList {
			singular := strings.TrimSuffix(ds.KindName, "s")
			if singular != ds.KindName && kindSet[singular] != nil {
				ds.KindName = singular + "Set"
				ds.FileName = strings.ToLower(singular) + "set"
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
	acronyms := buildAcronymSet(cfg)
	for name, ds := range p.DataSourcesMap {
		ci := categorize(cfg, name)
		if ci == nil {
			log.Printf("WARN: skipping legacy data source %q (no category match)", name)
			continue
		}
		info := &dsInfo{
			TFName:       name,
			CatInfo:      *ci,
			IsLegacySDK:  true,
			LegacySchema: ds,
		}
		info.KindName, info.FileName = deriveNames(cfg, acronyms, name, ci.TFPrefix)
		parseLegacySchema(acronyms, info)
		grouped[ci.DirName] = append(grouped[ci.DirName], info)
	}
}

func collectFrameworkDataSources(cfg Config, fwp fwprovider.Provider, grouped map[string][]*dsInfo) {
	acronyms := buildAcronymSet(cfg)
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
			TFName:      name,
			CatInfo:     *ci,
			IsLegacySDK: false,
			FwDS:        ds,
		}
		info.KindName, info.FileName = deriveNames(cfg, acronyms, name, ci.TFPrefix)
		parseFWSchema(acronyms, info, ds)
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

func deriveNames(cfg Config, acronyms map[string]bool, tfName, tfPrefix string) (kindName, fileName string) {
	stripped := strings.TrimPrefix(tfName, tfPrefix)
	if stripped == tfName {
		stripped = strings.TrimPrefix(tfName, cfg.FallbackTFPrefix)
	}
	return snakeToCamel(acronyms, stripped), strings.ReplaceAll(stripped, "_", "")
}

func sortedGroupNames(grouped map[string][]*dsInfo) []string {
	names := make([]string, 0, len(grouped))
	for name := range grouped {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		sort.Slice(grouped[name], func(i, j int) bool {
			return grouped[name][i].TFName < grouped[name][j].TFName
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

func parseLegacySchema(acronyms map[string]bool, info *dsInfo) {
	if info.LegacySchema == nil || info.LegacySchema.Schema == nil {
		log.Fatalf("data source %q has nil schema", info.TFName)
	}
	for name, field := range info.LegacySchema.Schema {
		if name == "id" {
			continue
		}
		fi := buildLegacyFieldInfo(acronyms, info.KindName, name, field)
		if field.Required || field.Optional {
			info.ForProviderFields = append(info.ForProviderFields, fi)
		}
		// All fields are observable outputs in data sources: TF data sources
		// populate every field on read, including Optional inputs.
		if field.Computed || field.Optional {
			info.AtProviderFields = append(info.AtProviderFields, fi)
		}
	}
	sortFields(info)
}

func buildLegacyFieldInfo(acronyms map[string]bool, parentName, name string, field *sdkschema.Schema) fieldInfo {
	fi := fieldInfo{
		TFName:      name,
		GoName:      snakeToCamel(acronyms, name),
		JSONName:    snakeToCamelJSON(acronyms, name),
		GoType:      sdkTypeToGo(field),
		Required:    field.Required,
		Description: field.Description,
	}

	// Handle nested resource elements in List/Set types.
	if field.Type == sdkschema.TypeList || field.Type == sdkschema.TypeSet {
		if res, ok := field.Elem.(*sdkschema.Resource); ok {
			structName := parentName + snakeToCamel(acronyms, name)
			fi.NestedStructName = structName
			fi.GoType = "[]" + structName
			fi.IsSet = field.Type == sdkschema.TypeSet

			nested := make([]fieldInfo, 0, len(res.Schema))
			for nestedName, nestedField := range res.Schema {
				nfi := buildLegacyFieldInfo(acronyms, structName, nestedName, nestedField)
				nested = append(nested, nfi)
			}
			sort.Slice(nested, func(i, j int) bool { return nested[i].TFName < nested[j].TFName })
			fi.NestedFields = nested
		}
	}

	return fi
}

func parseFWSchema(acronyms map[string]bool, info *dsInfo, ds datasource.DataSource) {
	ctx := context.Background()
	type schemaer interface {
		Schema(context.Context, datasource.SchemaRequest, *datasource.SchemaResponse)
	}
	s, ok := ds.(schemaer)
	if !ok {
		log.Fatalf("data source %q does not implement Schema", info.TFName)
	}
	var resp datasource.SchemaResponse
	s.Schema(ctx, datasource.SchemaRequest{}, &resp)
	if resp.Diagnostics.HasError() {
		log.Fatalf("schema for %q has errors: %v", info.TFName, resp.Diagnostics.Errors())
	}

	for name, attr := range resp.Schema.Attributes {
		if name == "id" {
			continue
		}
		fi := fieldInfo{
			TFName:      name,
			GoName:      snakeToCamel(acronyms, name),
			JSONName:    snakeToCamelJSON(acronyms, name),
			GoType:      fwAttrTypeToGo(attr),
			Required:    attr.IsRequired(),
			Description: attr.GetMarkdownDescription(),
		}
		if fi.Description == "" {
			fi.Description = attr.GetDescription()
		}
		if attr.IsRequired() || attr.IsOptional() {
			info.ForProviderFields = append(info.ForProviderFields, fi)
		}
		// All fields are observable outputs in data sources: TF data sources
		// populate every field on read, including Optional inputs.
		if attr.IsComputed() || attr.IsOptional() {
			info.AtProviderFields = append(info.AtProviderFields, fi)
		}
	}
	sortFields(info)
}

func sortFields(info *dsInfo) {
	sort.Slice(info.ForProviderFields, func(i, j int) bool { return info.ForProviderFields[i].TFName < info.ForProviderFields[j].TFName })
	sort.Slice(info.AtProviderFields, func(i, j int) bool { return info.AtProviderFields[i].TFName < info.AtProviderFields[j].TFName })
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
				return goTypeSliceStr
			case sdkschema.TypeInt:
				return "[]int64"
			default:
				return goTypeSliceStr
			}
		}
		return goTypeSliceStr
	case sdkschema.TypeMap:
		return "map[string]string"
	default:
		return goTypeString
	}
}

func fwAttrTypeToGo(attr fwschema.Attribute) string {
	ctx := context.Background()
	attrType := attr.GetType()
	tfType := attrType.TerraformType(ctx)
	required := attr.IsRequired()

	switch {
	case tfType.Is(tftypes.List{}) || tfType.Is(tftypes.Set{}):
		return goTypeSliceStr
	case tfType.Is(tftypes.Map{}):
		return "map[string]string"
	case attrType.Equal(fwtypes.StringType):
		if required {
			return goTypeString
		}
		return goTypePtrString
	case attrType.Equal(fwtypes.Int64Type):
		if required {
			return goTypeInt64
		}
		return goTypePtrInt64
	case attrType.Equal(fwtypes.Float64Type):
		if required {
			return goTypeFloat64
		}
		return goTypePtrFloat
	case attrType.Equal(fwtypes.BoolType):
		if required {
			return goTypeBool
		}
		return goTypePtrBool
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
		ci := dsList[0].CatInfo

		apiDir := filepath.Join(apisBase, groupName, cfg.APIVersion)
		ctrlDir := filepath.Join(ctrlBase, groupName)
		examplesDir := filepath.Join(examplesBase, groupName, cfg.APIVersion)

		mustMkdirAll(apiDir)
		mustMkdirAll(ctrlDir)
		mustMkdirAll(examplesDir)

		writeFormatted(filepath.Join(apiDir, "zz_groupversion_info.go"), generateGroupVersionInfo(cfg, ci))

		for _, ds := range dsList {
			writeFormatted(filepath.Join(apiDir, "zz_"+ds.FileName+"_types.go"), generateTypes(cfg, ds))
			writeFormatted(filepath.Join(ctrlDir, "zz_"+ds.FileName+"_spec.go"), generateSpec(cfg, ds, ci))
			writeRaw(filepath.Join(examplesDir, ds.FileName+".yaml"), generateExample(cfg, ds, ci))
		}

		writeFormatted(filepath.Join(ctrlDir, "zz_setup.go"), generateGroupSetup(cfg, dsList, ci))

		if hasFrameworkDS(dsList) {
			writeFormatted(filepath.Join(ctrlDir, "zz_factories.go"), generateFactories(cfg, dsList, ci))
		}

		if hasLegacyDS(dsList) {
			writeFormatted(filepath.Join(ctrlDir, "zz_legacy_factories.go"), generateLegacyFactories(cfg, dsList, ci))
		}
	}

	mustMkdirAll(apisBase)
	writeFormatted(filepath.Join(apisBase, "zz_register.go"), generateRegister(cfg, groupNames))

	mustMkdirAll(ctrlBase)
	writeFormatted(filepath.Join(ctrlBase, "zz_setup.go"), generateTopSetup(cfg, groupNames))
	writeFormatted(filepath.Join(ctrlBase, "zz_connect.go"), generateConnectFn(cfg))
}

func mustMkdirAll(path string) {
	if err := os.MkdirAll(path, 0o750); err != nil {
		log.Fatalf("mkdir %s: %v", path, err)
	}
}

func writeFormatted(path, content string) {
	formatted, err := imports.Process(path, []byte(content), nil)
	if err != nil {
		log.Printf("WARN: goimports failed for %s: %v (writing unformatted)", path, err)
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
