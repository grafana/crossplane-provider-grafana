# Testing

## Unit tests

Unit tests use an in-process HTTP mock and have no external dependencies.

```bash
go test ./...
# or via make
make test
```

Tests in `internal/controller/namespaced/observe/teams/controller_test.go` cover:

- `searchAllTeams` with no filter, name filter, query filter, empty result, and pagination
- `Observe` reporting up-to-date and not-up-to-date
- `Update` populating `status.atProvider.teams`
- `Delete` being a no-op

## Integration tests

Integration tests run against a real Grafana instance and are skipped automatically when `GRAFANA_AUTH` is not set.

### Against an existing instance

Set `GRAFANA_AUTH` to an API token or `user:password` basic-auth string. `GRAFANA_URL` defaults to `http://localhost:3000`.

```bash
GRAFANA_AUTH=admin:admin \
  go test -v -run Integration ./internal/controller/namespaced/observe/teams/...

GRAFANA_URL=https://my-instance.grafana.net GRAFANA_AUTH=glsa_mytoken \
  go test -v -run Integration ./internal/controller/namespaced/observe/teams/...

# or via make
GRAFANA_AUTH=admin:admin make go.test.integration
```

### With a local Docker container (local development only)

Setting `GRAFANA_DOCKER=1` starts a `grafana/grafana:11.0.0` container on port 13000, seeds it with four known teams (`alpha-team`, `beta-team`, `ops-east`, `ops-west`), runs the integration tests against it, and stops the container on exit. Docker must be available.

```bash
GRAFANA_DOCKER=1 go test -v -run Integration ./internal/controller/namespaced/observe/teams/...
# or via make
make go.test.integration.docker
```

These tests exercise the Grafana API client code directly — no Kubernetes cluster or Crossplane reconciler is involved.

This target is **not** part of the CI suite. In CI, inject `GRAFANA_AUTH` as a secret and use `go.test.integration` instead.

### Integration test coverage

| Test | What it checks |
|------|---------------|
| `TestIntegration_SearchAllTeams` | Unfiltered search returns a valid list; in Docker mode asserts all 4 seeded teams are present |
| `TestIntegration_Observe` | Full observe cycle: first call returns `ResourceUpToDate=false`, after `Update` returns `true` |
| `TestIntegration_NameFilter` | Exact-name filter reaches the API and returns only matching teams |
| `TestIntegration_QueryFilter` | Query filter returns teams whose name contains the search string; in Docker mode asserts exactly `ops-east` and `ops-west` |

## E2E tests (uptest/chainsaw)

E2E tests exercise the full reconciliation loop: a real kind cluster runs the provider against an in-cluster Grafana instance, and uptest verifies that Kubernetes resources reach `Ready=True` and `Synced=True`.

The test environment is set up by `cluster/test/setup.sh`, which deploys Grafana and creates both a `ProviderConfig` (for OSS/cluster resources) and a `ClusterProviderConfig` (for `grafana.m.crossplane.io` namespaced resources), both named `grafana-provider`.

### Run observe e2e tests

```bash
make e2e.observe
```

This builds the provider, spins up a kind cluster with Crossplane, deploys the provider, then applies `examples/namespaced/observe/teams-all.yaml` which:

1. Creates an `oss.grafana.m.crossplane.io/v1alpha1/Team` (so there is at least one team to observe)
2. Creates an `observe.grafana.m.crossplane.io/v1alpha1/Teams` and waits for `Ready=True, Synced=True`

### Run all e2e tests

```bash
make e2e UPTEST_EXAMPLE_LIST="examples/namespaced/observe/teams-all.yaml examples/oss/..."
```

### What e2e covers that integration tests do not

| Aspect | Integration tests | E2E tests |
|--------|-----------------|-----------|
| Grafana API client code | Yes | Yes (indirectly) |
| ProviderConfig resolution | No | Yes |
| Kubernetes status write-back | No | Yes |
| `Ready`/`Synced` conditions | No | Yes |
| Full reconciler loop | No | Yes |
