# Observed Resources (TF Data Sources as Crossplane Resources)

## Overview

Observed resources are read-only Crossplane resources backed by Terraform data
sources. They allow Crossplane compositions and references to resolve values
that are not managed by the provider but exist externally (e.g. an OnCall team
that was created outside of Crossplane).

Each observed resource lives under a `.o.` API group (for "observed"), parallel
to the `.m.` groups used for managed resources:

| Managed group | Observed group |
|---|---|
| `oncall.grafana.m.crossplane.io` | `oncall.grafana.o.crossplane.io` |
| `oss.grafana.m.crossplane.io` | `oss.grafana.o.crossplane.io` |
| `cloud.grafana.m.crossplane.io` | `cloud.grafana.o.crossplane.io` |
| ... | ... |

## How it works

An observed resource has the same lifecycle as any Crossplane managed resource,
but with observe-only semantics:

- **Observe**: Always returns `ResourceExists: true` (unless the CR is deleted).
- **Update** (called when not up-to-date): Executes the Terraform data source
  read and populates `status.atProvider` with the results. The data source's
  `id` is stored in the standard `crossplane.io/external-name` annotation.
- **Create**: Returns an error — observed resources cannot be created.
- **Delete**: No-op — deleting the CR just removes it from the cluster.

### Field classification

Fields from the Terraform data source schema are classified as:

- **`spec.forProvider`** — Required or Optional fields (inputs to the data
  source query, e.g. `name` for looking up a team).
- **`status.atProvider`** — Computed-only fields (outputs returned by the data
  source, e.g. `email`, `avatarUrl`).

## Example

```yaml
apiVersion: oncall.grafana.o.crossplane.io/v1alpha1
kind: Team
metadata:
  name: my-team
  namespace: default
spec:
  forProvider:
    name: "ops-team"
  providerConfigRef:
    name: grafana
```

After reconciliation:

```yaml
status:
  atProvider:
    email: "ops@example.com"
    avatarUrl: "https://..."
  conditions:
    - type: Ready
      status: "True"
    - type: Synced
      status: "True"
metadata:
  annotations:
    crossplane.io/external-name: "ABCDEF12345"  # OnCall team ID
```

The external name can then be referenced from other resources (e.g. an
Integration's `teamId` field).

## Code generation

All observed resource types, controller specs, and registration boilerplate are
generated from the Terraform provider's schema. The generator supports both
Legacy SDK (`provider.DataSourcesMap`) and Plugin Framework
(`FrameworkProvider().DataSources()`) data sources.

### Running the generator

```bash
go run ./cmd/generate-observed
```

This introspects the vendored Terraform provider at compile time and emits:

```
apis/observed/<category>/v1alpha1/
    groupversion_info.go          — CRD group constants and scheme builder
    <name>_types.go               — ForProvider/AtProvider structs, CRD type, GVK
internal/controller/namespaced/observed/<category>/
    <name>_spec.go                — tfdatasource.Spec with read callbacks
    factories.go                  — Plugin Framework data source factories (if needed)
    setup.go                      — per-group controller registration
apis/observed/register.go         — aggregated AddToScheme across all groups
internal/controller/namespaced/observed/setup.go — aggregated Setup/SetupGated
examples-generated/observed/      — example YAML manifests per data source
```

After running the generator, run the standard codegen tools:

```bash
# Generate deepcopy methods
controller-gen object paths=./apis/observed/...

# Generate managed resource interface implementations
angryjet generate-methodsets ./apis/observed/...
```

### Keeping up with provider upgrades

When `github.com/grafana/terraform-provider-grafana/v4` is bumped:

```bash
go get github.com/grafana/terraform-provider-grafana/v4@latest
go run ./cmd/generate-observed
git diff  # review new/changed/removed files
```

- **New data source** → new `_types.go` + `_spec.go` files appear as untracked.
- **Removed data source** → stale files remain (delete manually or add orphan
  detection to CI).
- **Schema change** → modified types/specs show up in the diff.

A CI check can enforce that generated code matches the vendored provider:

```yaml
- run: go run ./cmd/generate-observed
- run: git diff --exit-code -- apis/observed/ internal/controller/namespaced/observed/
```

## Architecture

### Generator library (`pkg/generateobserved/`)

The generation engine is provider-agnostic and reusable. It accepts a `Config`
struct with all provider-specific settings (module path, group suffix, category
rules, acronyms, etc.) and a pair of TF provider instances. See the
[package documentation](../pkg/generateobserved/config.go) for the full API.

### Runtime controller (`internal/controller/namespaced/observed/tfdatasource/`)

A generic controller that reconciles any observe-only resource given a `Spec`:

```go
type Spec struct {
    DataSourceName string
    ManagedKind    schema.GroupVersionKind
    NewManaged     func() resource.Managed
    Read           ReadFn
    IsUpToDate     func(resource.Managed) bool
}
```

Two `ReadFn` builders are provided:

- **`NewLegacyReadFn`** — for TF Plugin SDK v2 data sources. Uses
  `schema.Resource.Data()` + `ReadContext`.
- **`NewFrameworkReadFn`** — for TF Plugin Framework data sources. Constructs a
  `tfsdk.Config`, calls `ds.Read()`, and extracts the resulting `tfsdk.State`.

### Grafana-specific wrapper (`cmd/generate-observed/`)

A thin `main.go` (~60 lines) that defines the Grafana category rules, acronyms,
and provider constructors, then calls `generateobserved.Generate()`.

## Category mapping

Data sources are assigned to CRD groups by TF name prefix:

| TF prefix | CRD group | Example kind |
|---|---|---|
| `grafana_oncall_` | `oncall.grafana.o.crossplane.io` | `Team` |
| `grafana_cloud_provider_` | `cloudprovider.grafana.o.crossplane.io` | `AWSAccount` |
| `grafana_synthetic_monitoring_` | `sm.grafana.o.crossplane.io` | `Probe` |
| `grafana_cloud_` | `cloud.grafana.o.crossplane.io` | `Stack` |
| `grafana_k6_` | `k6.grafana.o.crossplane.io` | `LoadTest` |
| `grafana_slo` | `slo.grafana.o.crossplane.io` | `Slos` |
| `grafana_` (fallback) | `oss.grafana.o.crossplane.io` | `Dashboards` |

Individual overrides (e.g. `grafana_role` → `enterprise`) are configured in
`CategoryOverrides`.
