# Observe-Only Resources

Observe-only resources read data from Grafana without creating, updating, or deleting anything. They are useful for referencing existing Grafana state within a Crossplane composition, or for inspecting Grafana objects that are managed outside of Crossplane.

## API group

Observe-only resources live in the `observe.grafana.m.crossplane.io` API group, separate from managed resources (e.g. `oss.grafana.m.crossplane.io`). The `.m.` infix indicates namespaced scope — all resources in this group are namespace-scoped.

## Available resources

| Kind | API Version | Description |
|------|-------------|-------------|
| `Teams` | `observe.grafana.m.crossplane.io/v1alpha1` | Lists Grafana teams matching optional filter criteria |

---

## Teams

`Teams` lists all Grafana teams visible to the configured ProviderConfig, optionally filtered by name or a search query. On every reconcile the full matching team list is written to `status.atProvider.teams`.

### Spec fields (`spec.forProvider`)

| Field | Type | Description |
|-------|------|-------------|
| `orgId` | string (optional) | Restrict to a specific org ID. Defaults to the org in the ProviderConfig credentials. |
| `name` | string (optional) | Exact team name match. |
| `query` | string (optional) | Search string — returns teams whose name contains this value. |

`name` and `query` are independent filters; only one should be set at a time (the Grafana API applies whichever is provided).

### Status fields (`status.atProvider.teams[]`)

Each entry in the list contains:

| Field | Type | Description |
|-------|------|-------------|
| `id` | integer | Grafana numeric team ID |
| `uid` | string | Grafana team UID |
| `name` | string | Team name |
| `email` | string | Team email address |
| `memberCount` | integer | Number of members |
| `orgId` | integer | Org the team belongs to |

### Examples

List every team in the org:

```yaml
apiVersion: observe.grafana.m.crossplane.io/v1alpha1
kind: Teams
metadata:
  name: all-teams
  namespace: default
spec:
  providerConfigRef:
    kind: ProviderConfig   # or ClusterProviderConfig
    name: my-provider-config
  forProvider: {}
```

List teams whose name contains "ops":

```yaml
apiVersion: observe.grafana.m.crossplane.io/v1alpha1
kind: Teams
metadata:
  name: ops-teams
  namespace: default
spec:
  providerConfigRef:
    kind: ProviderConfig
    name: my-provider-config
  forProvider:
    query: ops
```

Look up a team by exact name:

```yaml
apiVersion: observe.grafana.m.crossplane.io/v1alpha1
kind: Teams
metadata:
  name: platform-team-lookup
  namespace: default
spec:
  providerConfigRef:
    kind: ProviderConfig
    name: my-provider-config
  forProvider:
    name: platform-team
```

### Reading the result

```bash
kubectl get teams.observe.grafana.m.crossplane.io -n default all-teams -o jsonpath='{.status.atProvider.teams}' | jq .
```

Example output:

```json
[
  {
    "id": 1,
    "uid": "abc123",
    "name": "platform-team",
    "email": "platform@example.com",
    "memberCount": 5,
    "orgId": 1
  }
]
```

### Conditions

| Condition | Meaning |
|-----------|---------|
| `Ready=True` | Last reconcile succeeded; `status.atProvider.teams` is up to date |
| `Ready=False` | Reconcile failed (e.g. Grafana API error); see the condition message |
| `Synced=True` | Controller is actively reconciling the resource |

### ProviderConfig

`Teams` supports both `ProviderConfig` (namespace-scoped) and `ClusterProviderConfig` (cluster-scoped). Specify the kind explicitly in `spec.providerConfigRef`:

```yaml
spec:
  providerConfigRef:
    kind: ClusterProviderConfig
    name: shared-grafana
```

The default kind when omitted is `ClusterProviderConfig` (matching the Crossplane convention for the namespaced API group).

### Pagination

The controller fetches all matching teams from the Grafana API in pages of 1,000, so large orgs are handled correctly without truncation.

### Management policies

The default management policy (`["*"]`) is the correct choice for `Teams` resources:

- **Create** is never triggered — `Observe` always reports the external resource as existing (there is always a list, even if empty), so the reconciler never calls `Create`.
- **Update** is called whenever the observed team list differs from what is in `status.atProvider.teams`. This is what populates the status field.
- **Delete** is a no-op — removing the Kubernetes object does not delete anything in Grafana.

Do **not** set `managementPolicies: ["Observe"]`. That policy skips `Update`, which means `status.atProvider.teams` would never be populated.
