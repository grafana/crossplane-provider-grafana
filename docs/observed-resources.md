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
- **`status.atProvider`** — Computed or Optional fields (outputs returned by the
  data source, e.g. `email`, `avatarUrl`). Optional fields appear in both
  `forProvider` and `atProvider` since data sources populate all fields on read.

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

## Set resources (plural data sources)

When a Terraform provider exposes both a singular and plural data source (e.g.
`grafana_user` and `grafana_users`), Kubernetes would pluralize both kinds to
the same CRD name (`users`), causing one to silently overwrite the other. To
avoid this, the generator detects these collisions and renames the plural kind
with a `Set` suffix:

| Singular kind | Plural kind (renamed) | CRD plural |
|---|---|---|
| `User` | `UserSet` | `usersets` |
| `Folder` | `FolderSet` | `foldersets` |
| `Dashboard` | `DashboardSet` | `dashboardsets` |
| `LibraryPanel` | `LibraryPanelSet` | `librarypanelsets` |
| `Probe` | `ProbeSet` | `probesets` |
| `LoadTest` | `LoadTestSet` | `loadtestsets` |
| `Project` | `ProjectSet` | `projectsets` |
| `Schedule` | `ScheduleSet` | `schedulesets` |
| `Collector` | `CollectorSet` | `collectorsets` |
| `AWSCloudwatchScrapeJob` | `AWSCloudwatchScrapeJobSet` | `awscloudwatchscrapejobsets` |

Set resources typically require no `forProvider` inputs and return all items in
`status.atProvider`:

```yaml
apiVersion: oss.grafana.o.crossplane.io/v1alpha1
kind: UserSet
metadata:
  name: all-users
  namespace: default
spec:
  forProvider: {}
  providerConfigRef:
    name: default
    kind: ClusterProviderConfig
```

After reconciliation:

```yaml
status:
  atProvider:
    users:
      - email: "admin@localhost"
        id: 1
        isAdmin: true
        login: "admin"
      - email: "john@example.com"
        id: 2
        isAdmin: false
        login: "john"
        name: "John"
```

## Testing

An uptest example is provided at
`examples/namespaced/v1alpha1/observed-sets.yaml`. It creates managed resources
(User, Folder) via `.m.` APIs and then observes them with UserSet and FolderSet
via `.o.` APIs.

To run all e2e tests (including observed resources):

```bash
make e2e
```

Example YAMLs under `examples/` are auto-discovered. To run a specific example:

```bash
UPTEST_EXAMPLE_LIST=examples/namespaced/v1alpha1/observed-sets.yaml make e2e
```

Observed resources use `uptest.upbound.io/disable-import: "true"` since they
are read-only and cannot be imported.

## Code generation

All observed resource types, controller specs, and registration boilerplate are
generated from the Terraform provider's schema. The generator supports both
Legacy SDK (`provider.DataSourcesMap`) and Plugin Framework
(`FrameworkProvider().DataSources()`) data sources.

### Running the generator

The generator runs automatically as part of `make generate`:

```bash
make generate
```

This deletes all existing `zz_*` files first, then re-runs the generator and
regenerates deepcopy methods, CRDs, and managed resource interface
implementations in one step. Generated files use the `zz_` prefix so the
cleanup step handles them automatically.

The generator emits:

```
apis/observed/<category>/v1alpha1/
    zz_groupversion_info.go       — CRD group constants and scheme builder
    zz_<name>_types.go            — ForProvider/AtProvider structs, CRD type, GVK
internal/controller/namespaced/observed/<category>/
    zz_<name>_spec.go             — tfdatasource.Spec with read callbacks
    zz_factories.go               — Plugin Framework data source factories (if needed)
    zz_legacy_factories.go        — Legacy SDK data source factories (if needed)
    zz_setup.go                   — per-group controller registration
apis/observed/zz_register.go      — aggregated AddToScheme across all groups
internal/controller/namespaced/observed/zz_setup.go  — aggregated Setup/SetupGated
internal/controller/namespaced/observed/zz_connect.go — provider ConnectFn
examples-generated/observed/      — example YAML manifests per data source
```

### Keeping up with provider upgrades

When `github.com/grafana/terraform-provider-grafana/v4` is bumped, run:

```bash
go get github.com/grafana/terraform-provider-grafana/v4@latest
make generate
git diff  # review new/changed/removed files
```

- **New data source** → new `_types.go` + `_spec.go` files appear.
- **Removed data source** → stale files are deleted automatically by the
  `make generate` cleanup step.
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

### Runtime controller (`pkg/tfdatasource/`)

A provider-agnostic controller (in `pkg/` so external providers can reuse it)
that reconciles any observe-only resource given a `Spec`:

```go
type Spec struct {
    DataSourceName string
    ManagedKind    schema.GroupVersionKind
    NewManaged     func() resource.Managed
    ConnectFn      ConnectFn  // injected at setup time from zz_connect.go
    Read           ReadFn
    IsUpToDate     func(resource.Managed) bool
}
```

Two `ReadFn` builders are provided:

- **`NewLegacyReadFn`** — for TF Plugin SDK v2 data sources. Takes a resolved
  `*sdkschema.Resource` (from `zz_legacy_factories.go`) and uses
  `schema.Resource.Data()` + `ReadContext`.
- **`NewFrameworkReadFn`** — for TF Plugin Framework data sources. Takes a
  factory func (from `zz_factories.go`), constructs a `tfsdk.Config`, calls
  `ds.Read()`, and extracts the resulting `tfsdk.State`.

#### Shared controller vs. upjet's per-resource approach

Observed resources use a **single shared controller** (`tfdatasource.Setup` in
`pkg/tfdatasource/`) that handles all observed resources. Each resource is
registered with that controller by adding a `Spec` to a list; the `Spec`
contains the resource-specific logic (which Terraform data source to read, how
to deserialize it, etc.). The reconciliation loop itself is shared.

This differs from upjet's managed resource controllers, which create one
controller instance per resource. Each resource has its own `Setup()` function
that instantiates a separate controller.

The shared controller approach is more efficient for observed resources because
they all follow the same simple semantics: read the data source, populate
`status.atProvider`, done. There are no resource-specific lifecycle asymmetries
(like create/update/delete differences) that would require per-resource
controllers. Managed resources, conversely, have varied CRUD semantics,
validation rules, and lifecycle hooks per resource, so per-resource controllers
make sense for them.

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
| `grafana_` (fallback) | `oss.grafana.o.crossplane.io` | `DashboardSet` |

Individual overrides (e.g. `grafana_role` → `enterprise`) are configured in
`CategoryOverrides`.
