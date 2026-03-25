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
	var names []string
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
		fi := fieldInfo{
			tfName:      name,
			goName:      snakeToCamel(cfg, name),
			jsonName:    snakeToCamelJSON(cfg, name),
			goType:      sdkTypeToGo(field),
			required:    field.Required,
			description: field.Description,
		}
		if field.Computed && !field.Required && !field.Optional {
			info.atProviderFields = append(info.atProviderFields, fi)
		} else {
			info.forProviderFields = append(info.forProviderFields, fi)
		}
	}
	sortFields(info)
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
	switch field.Type {
	case sdkschema.TypeString:
		if field.Required {
			return "string"
		}
		return "*string"
	case sdkschema.TypeInt:
		if field.Required {
			return "int64"
		}
		return "*int64"
	case sdkschema.TypeFloat:
		if field.Required {
			return "float64"
		}
		return "*float64"
	case sdkschema.TypeBool:
		if field.Required {
			return "bool"
		}
		return "*bool"
	case sdkschema.TypeList, sdkschema.TypeSet:
		if elem, ok := field.Elem.(*sdkschema.Schema); ok {
			switch elem.Type {
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
		return "string"
	}
}

func fwAttrTypeToGo(attr fwschema.Attribute) string {
	typeName := attr.GetType().String()
	required := attr.IsRequired()
	switch {
	case strings.Contains(typeName, "StringType"):
		if required {
			return "string"
		}
		return "*string"
	case strings.Contains(typeName, "Int64Type"):
		if required {
			return "int64"
		}
		return "*int64"
	case strings.Contains(typeName, "Float64Type"):
		if required {
			return "float64"
		}
		return "*float64"
	case strings.Contains(typeName, "BoolType"):
		if required {
			return "bool"
		}
		return "*bool"
	case strings.Contains(typeName, "ListType"), strings.Contains(typeName, "SetType"):
		return "[]string"
	case strings.Contains(typeName, "MapType"):
		return "map[string]string"
	default:
		if required {
			return "string"
		}
		return "*string"
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

		writeFormatted(filepath.Join(apiDir, "groupversion_info.go"), generateGroupVersionInfo(cfg, ci))

		for _, ds := range dsList {
			writeFormatted(filepath.Join(apiDir, ds.fileName+"_types.go"), generateTypes(cfg, ds))
			writeFormatted(filepath.Join(ctrlDir, ds.fileName+"_spec.go"), generateSpec(cfg, ds, ci))
			writeRaw(filepath.Join(examplesDir, ds.fileName+".yaml"), generateExample(cfg, ds, ci))
		}

		writeFormatted(filepath.Join(ctrlDir, "setup.go"), generateGroupSetup(cfg, dsList, ci))

		if hasFrameworkDS(dsList) {
			writeFormatted(filepath.Join(ctrlDir, "factories.go"), generateFactories(cfg, dsList, ci))
		}
	}

	mustMkdirAll(apisBase)
	writeFormatted(filepath.Join(apisBase, "register.go"), generateRegister(cfg, groupNames))

	mustMkdirAll(ctrlBase)
	writeFormatted(filepath.Join(ctrlBase, "setup.go"), generateTopSetup(cfg, groupNames))

	writeFormatted(filepath.Join(apisBase, "generate.go"), generateGenerateGo(cfg))
}

func mustMkdirAll(path string) {
	if err := os.MkdirAll(path, 0o755); err != nil {
		log.Fatalf("mkdir %s: %v", path, err)
	}
}

func writeFormatted(path, content string) {
	formatted, err := format.Source([]byte(content))
	if err != nil {
		log.Printf("WARN: gofmt failed for %s: %v (writing unformatted)", path, err)
		formatted = []byte(content)
	}
	if err := os.WriteFile(path, formatted, 0o644); err != nil {
		log.Fatalf("write %s: %v", path, err)
	}
	log.Printf("  wrote %s", path)
}

func writeRaw(path, content string) {
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		log.Fatalf("write %s: %v", path, err)
	}
	log.Printf("  wrote %s", path)
}
