# Validating manifests with kubeconform

This repository ships JSON Schema files for all CRDs so that
[kubeconform](https://github.com/yannh/kubeconform) can validate Crossplane
manifests offline, without a running cluster.

The schemas live in the `schemas/` directory, organised by API group:

```
schemas/
  alerting.grafana.crossplane.io/
    contactpoint_v1alpha1.json
    rulegroup_v1alpha1.json
    ...
  cloud.grafana.crossplane.io/
    stack_v1alpha1.json
    ...
```

## Usage

### From the repository (latest on main)

```bash
kubeconform \
  -schema-location default \
  -schema-location 'https://raw.githubusercontent.com/grafana/crossplane-provider-grafana/refs/heads/main/schemas/{{.Group}}/{{.ResourceKind}}_{{.ResourceAPIVersion}}.json' \
  my-manifest.yaml
```

### Pinned to a release tag

```bash
kubeconform \
  -schema-location default \
  -schema-location 'https://raw.githubusercontent.com/grafana/crossplane-provider-grafana/refs/tags/v2.9.1/schemas/{{.Group}}/{{.ResourceKind}}_{{.ResourceAPIVersion}}.json' \
  my-manifest.yaml
```

### Using a local checkout

```bash
kubeconform \
  -schema-location default \
  -schema-location 'schemas/{{.Group}}/{{.ResourceKind}}_{{.ResourceAPIVersion}}.json' \
  my-manifest.yaml
```

## How schemas are generated

Schemas are generated as part of `make generate` (via `go generate`). The
`cmd/generate-schemas` program reads every CRD YAML from `package/crds/`,
extracts the `openAPIV3Schema` for each version, applies kubeconform-compatible
transformations, and writes the result as JSON.

The transformations match kubeconform's
[openapi2jsonschema.py](https://github.com/yannh/kubeconform/blob/master/scripts/openapi2jsonschema.py):

- `additionalProperties: false` is set on every object that has `properties`
  (the root object is excluded).
- `format: int-or-string` fields are replaced with
  `oneOf: [{type: string}, {type: integer}]`.

To regenerate schemas manually:

```bash
go run ./cmd/generate-schemas package/crds schemas
```
