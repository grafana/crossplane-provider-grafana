/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/pipeline"
	"github.com/crossplane/upjet/v2/pkg/schema/traverser"
	"github.com/crossplane/upjet/v2/pkg/types/conversion/tfjson"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	kyaml "k8s.io/apimachinery/pkg/util/yaml"
	"sigs.k8s.io/yaml"

	"github.com/grafana/crossplane-provider-grafana/v2/config"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		panic("root directory is required to be given as argument")
	}
	rootDir := os.Args[1]
	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path with %s", rootDir))
	}
	clusterProvider, err := config.GetProvider(true)
	if err != nil {
		panic(fmt.Sprintf("cannot get cluster provider configuration: %s", err))
	}
	namespacedProvider, err := config.GetProviderNamespaced(true)
	if err != nil {
		panic(fmt.Sprintf("cannot get namespaced provider configuration: %s", err))
	}
	pipeline.Run(clusterProvider, namespacedProvider, absRootDir)

	// Fix generated example manifests: NestingModeSingle Terraform blocks
	// (SchemaTypeObject) are embedded objects in the CRD but the HCL-to-JSON
	// scraper wraps them in arrays. Flatten them to plain objects.
	//
	// Note: upstream ApplyAPIConverters cannot handle SchemaTypeObject paths
	// because they lack [*] wildcards that fieldpath needs to traverse into
	// arrays. We process these ourselves by collecting the CRD paths via the
	// schema traverser and flattening single-element arrays at those paths.
	//
	// See: https://github.com/grafana/crossplane-provider-grafana/issues/549
	if err := flattenObjectBlocks(clusterProvider, filepath.Join(absRootDir, "examples-generated", "cluster")); err != nil {
		panic(fmt.Sprintf("cannot flatten object blocks in cluster examples: %s", err))
	}
	if err := flattenObjectBlocks(namespacedProvider, filepath.Join(absRootDir, "examples-generated", "namespaced")); err != nil {
		panic(fmt.Sprintf("cannot flatten object blocks in namespaced examples: %s", err))
	}
}

// flattenObjectBlocks walks all YAML example files under dir and flattens
// single-element arrays at SchemaTypeObject field paths to plain objects.
func flattenObjectBlocks(pc *ujconfig.Provider, dir string) error {
	kindPaths := collectObjectBlockPaths(pc)

	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(info.Name(), ".yaml") {
			return err
		}
		content, err := os.ReadFile(filepath.Clean(path)) //nolint:gosec
		if err != nil {
			return errors.Wrapf(err, "read %s", path)
		}

		docs, err := decodeDocuments(content)
		if err != nil {
			return errors.Wrapf(err, "decode %s", path)
		}
		if len(docs) == 0 {
			return nil
		}

		anyChanged := false
		for _, obj := range docs {
			kind, _ := obj["kind"].(string)
			paths, ok := kindPaths[kind]
			if !ok {
				continue
			}
			for _, p := range paths {
				if flattenPath(obj, strings.Split("spec.forProvider."+p, ".")) {
					anyChanged = true
				}
			}
		}
		if !anyChanged {
			return nil
		}

		if err := writeDocuments(path, docs); err != nil {
			return errors.Wrapf(err, "write %s", path)
		}
		log.Printf("Flattened object blocks in: %s", path)
		return nil
	})
}

// collectObjectBlockPaths uses the upjet schema traverser to collect CRD
// field paths for all SchemaTypeObject (NestingModeSingle) fields, keyed
// by Kind.
func collectObjectBlockPaths(pc *ujconfig.Provider) map[string][]string {
	kindPaths := map[string][]string{}
	for _, r := range pc.Resources {
		c := &objectPathCollector{}
		if err := ujconfig.TraverseSchemas(r.Name, r, c); err != nil {
			log.Printf("WARN: failed to traverse schema for %s: %v", r.Name, err)
			continue
		}
		if len(c.paths) > 0 {
			kindPaths[r.Kind] = c.paths
		}
	}
	return kindPaths
}

// objectPathCollector is a schema traverser that collects CRD paths for
// SchemaTypeObject fields.
type objectPathCollector struct {
	paths []string
	traverser.NoopTraverser
}

func (c *objectPathCollector) VisitResource(r *traverser.ResourceNode) error {
	if r.Schema.Type != tfjson.SchemaTypeObject {
		return nil
	}
	c.paths = append(c.paths, traverser.FieldPathWithWildcard(r.CRDPath))
	return nil
}

// flattenPath descends the object following the path segments. At the last
// segment, if the value is a single-element []interface{} containing a map,
// it replaces it with that map. Returns true if a change was made.
func flattenPath(obj map[string]interface{}, parts []string) bool {
	if len(parts) == 0 {
		return false
	}
	key := parts[0]
	val, ok := obj[key]
	if !ok {
		return false
	}
	if len(parts) == 1 {
		if arr, ok := val.([]interface{}); ok && len(arr) == 1 {
			if m, ok := arr[0].(map[string]interface{}); ok {
				obj[key] = m
				return true
			}
		}
		return false
	}
	// Descend: handle both maps and single-element arrays (not yet flattened).
	switch v := val.(type) {
	case map[string]interface{}:
		return flattenPath(v, parts[1:])
	case []interface{}:
		changed := false
		for _, item := range v {
			if m, ok := item.(map[string]interface{}); ok {
				if flattenPath(m, parts[1:]) {
					changed = true
				}
			}
		}
		return changed
	}
	return false
}

// decodeDocuments parses all YAML documents from content.
func decodeDocuments(content []byte) ([]map[string]interface{}, error) {
	var docs []map[string]interface{}
	decoder := kyaml.NewYAMLOrJSONDecoder(bytes.NewReader(content), 4096)
	for {
		u := &unstructured.Unstructured{}
		if err := decoder.Decode(u); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}
		if u.Object != nil {
			docs = append(docs, u.Object)
		}
	}
	return docs, nil
}

// writeDocuments writes YAML documents back to the file, separated by "---".
func writeDocuments(path string, docs []map[string]interface{}) error {
	var buf bytes.Buffer
	for i, doc := range docs {
		out, err := yaml.Marshal(doc)
		if err != nil {
			return errors.Wrapf(err, "marshal document %d", i)
		}
		if i > 0 {
			buf.WriteString("\n---\n\n")
		}
		buf.Write(out)
	}
	return os.WriteFile(path, buf.Bytes(), 0o600)
}
