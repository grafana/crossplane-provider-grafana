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
