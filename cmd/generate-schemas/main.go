/*
Copyright 2026 Grafana Labs

generate-schemas converts CRD YAML manifests into JSON Schema files
compatible with kubeconform.

It reads all CRD YAML files from a source directory, extracts the
openAPIV3Schema for each version, applies kubeconform-compatible
transformations, and writes the resulting JSON Schema files to an
output directory organized by API group.

Output filename format: {group}/{kind}_{version}.json (lowercased)

Usage:

	go run ./cmd/generate-schemas <crd-dir> <output-dir>
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"sigs.k8s.io/yaml"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <crd-dir> <output-dir>", os.Args[0]) //nolint:gosec // G706: build tool, not a server.
	}

	crdDir := os.Args[1]
	outDir := os.Args[2]

	files, err := filepath.Glob(filepath.Join(crdDir, "*.yaml"))
	if err != nil {
		log.Fatalf("Failed to glob CRD files: %v", err)
	}

	count := 0
	for _, f := range files {
		n, err := processCRD(f, outDir)
		if err != nil {
			log.Fatalf("Failed to process %s: %v", f, err) //nolint:gosec // G706: build tool, not a server.
		}
		count += n
	}

	fmt.Printf("Generated %d JSON schemas in %s\n", count, outDir)
}

// processCRD reads a single CRD YAML file and writes JSON Schema files
// for each version it defines. Returns the number of schemas written.
func processCRD(path, outDir string) (int, error) {
	data, err := os.ReadFile(path) //nolint:gosec // G304: path comes from filepath.Glob, not user input.
	if err != nil {
		return 0, fmt.Errorf("read file: %w", err)
	}

	// Parse YAML into a generic map so we can extract and transform
	// the openAPIV3Schema without losing any fields.
	var crd map[string]interface{}
	if err := yaml.Unmarshal(data, &crd); err != nil {
		return 0, fmt.Errorf("parse YAML: %w", err)
	}

	spec, _ := crd["spec"].(map[string]interface{})
	if spec == nil {
		return 0, nil
	}

	group, _ := spec["group"].(string)

	names, _ := spec["names"].(map[string]interface{})
	if names == nil {
		return 0, nil
	}
	kind, _ := names["kind"].(string)

	versions, _ := spec["versions"].([]interface{})

	count := 0
	for _, v := range versions {
		ver, _ := v.(map[string]interface{})
		if ver == nil {
			continue
		}

		versionName, _ := ver["name"].(string)
		schemaWrapper, _ := ver["schema"].(map[string]interface{})
		if schemaWrapper == nil {
			continue
		}

		openAPISchema, _ := schemaWrapper["openAPIV3Schema"].(map[string]interface{})
		if openAPISchema == nil {
			continue
		}

		// Apply kubeconform-compatible transformations.
		// 1. Set additionalProperties: false on objects with properties (skip root).
		setAdditionalProperties(openAPISchema, true)
		// 2. Replace format: int-or-string with oneOf.
		openAPISchema = replaceIntOrString(openAPISchema).(map[string]interface{})

		// Write JSON schema file.
		dir := filepath.Join(outDir, group)
		if err := os.MkdirAll(dir, 0o750); err != nil { //nolint:gosec // G703: paths are build-time constants from go:generate.
			return count, fmt.Errorf("create directory %s: %w", dir, err)
		}

		filename := fmt.Sprintf("%s_%s.json",
			strings.ToLower(kind),
			strings.ToLower(versionName),
		)
		jsonData, err := json.MarshalIndent(openAPISchema, "", "  ")
		if err != nil {
			return count, fmt.Errorf("marshal schema for %s/%s %s: %w", group, kind, versionName, err)
		}

		outPath := filepath.Join(dir, filename)
		if err := os.WriteFile(outPath, append(jsonData, '\n'), 0o600); err != nil { //nolint:gosec // G703: paths are build-time constants from go:generate.
			return count, fmt.Errorf("write %s: %w", outPath, err)
		}
		count++
	}

	return count, nil
}

// setAdditionalProperties walks the schema tree and sets
// "additionalProperties": false on every object that has a "properties"
// key. The root level is skipped (matching kubeconform's openapi2jsonschema.py
// default behavior).
func setAdditionalProperties(data interface{}, skipRoot bool) {
	m, ok := data.(map[string]interface{})
	if !ok {
		return
	}

	if _, hasProps := m["properties"]; hasProps && !skipRoot {
		if _, hasAdditional := m["additionalProperties"]; !hasAdditional {
			m["additionalProperties"] = false
		}
	}

	for _, val := range m {
		switch v := val.(type) {
		case map[string]interface{}:
			setAdditionalProperties(v, false)
		case []interface{}:
			for _, item := range v {
				setAdditionalProperties(item, false)
			}
		}
	}
}

// replaceIntOrString recursively replaces any map value that has
// {"format": "int-or-string"} with {"oneOf": [{"type": "string"}, {"type": "integer"}]}.
// This matches the behavior of kubeconform's openapi2jsonschema.py.
func replaceIntOrString(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{}, len(v))
		for key, val := range v {
			if valMap, ok := val.(map[string]interface{}); ok {
				if format, hasFormat := valMap["format"]; hasFormat && format == "int-or-string" {
					result[key] = map[string]interface{}{
						"oneOf": []interface{}{
							map[string]interface{}{"type": "string"},
							map[string]interface{}{"type": "integer"},
						},
					}
				} else {
					result[key] = replaceIntOrString(val)
				}
			} else {
				result[key] = replaceIntOrString(val)
			}
		}
		return result
	case []interface{}:
		result := make([]interface{}, len(v))
		for i, item := range v {
			result[i] = replaceIntOrString(item)
		}
		return result
	default:
		return data
	}
}
