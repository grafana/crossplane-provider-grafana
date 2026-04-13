# ProviderConfig Secret Fields

This document describes the fields that can be configured in the Kubernetes secret referenced by a ProviderConfig. These fields represent the Terraform provider configuration for the Grafana provider.

## Overview

The ProviderConfig uses a Kubernetes secret to store sensitive credentials and configuration. The secret should contain a JSON object with the configuration fields. Some of these fields can be overridden by specifying them directly in the ProviderConfig spec.

## Secret Structure

The secret should contain a key (typically `credentials`) with a JSON object containing the configuration fields:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: grafana-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "url": "https://grafana.example.com",
      "auth": "admin:admin",
      "cloud_access_policy_token": "token",
      ...
    }
```

## Configuration Fields

The following table lists all fields that can be configured in the secret, along with their purpose and whether they can be overridden by the ProviderConfig object.

| Field | Description | Overridable by ProviderConfig | ProviderConfig Field |
|-------|-------------|-------------------------------|---------------------|
| `auth` | Authentication credentials for Grafana (e.g., "username:password" or API token) | ❌ No | N/A |
| `url` | Base URL of the Grafana instance | ✅ Yes | `spec.url` |
| `cloud_access_policy_token` | Access token for Grafana Cloud API | ❌ No | N/A |
| `cloud_api_url` | URL for Grafana Cloud API | ✅ Yes | `spec.cloudApiUrl` |
| `cloud_provider_access_token` | Access token for Grafana Cloud Provider API | ❌ No | N/A |
| `cloud_provider_url` | URL for Grafana Cloud Provider API | ✅ Yes | `spec.cloudProviderUrl` |
| `connections_api_access_token` | Access token for Grafana Connections API | ❌ No | N/A |
| `connections_api_url` | URL for Grafana Connections API | ✅ Yes | `spec.connectionsApiUrl` |
| `fleet_management_auth` | Authentication for Fleet Management API | ❌ No | N/A |
| `fleet_management_url` | URL for Fleet Management API | ✅ Yes | `spec.fleetManagementUrl` |
| `frontend_o11y_api_access_token` | Access token for Frontend Observability API | ❌ No | N/A |
| `oncall_access_token` | Access token for Grafana OnCall (not required if `auth` is a service account token) | ❌ No | N/A |
| `oncall_url` | URL for Grafana OnCall API | ✅ Yes | `spec.oncallUrl` |
| `sm_access_token` | Access token for Synthetic Monitoring | ❌ No | N/A |
| `sm_url` | URL for Synthetic Monitoring API | ✅ Yes | `spec.smUrl` |
| `cloud_api_key` | Legacy Cloud API key (deprecated, use access tokens) | ❌ No | N/A |
| `org_id` | Grafana organization ID | ✅ Yes | `spec.orgId` |
| `stack_id` | Grafana Cloud stack ID (required for k6 resources) | ✅ Yes | `spec.stackId` |
| `k6_access_token` | Access token for k6 Cloud API | ❌ No | N/A |

## Stack Secret Reference

The ProviderConfig supports an optional `stackSecretRef` field that references a Kubernetes Secret produced by a `grafana_cloud_stack` resource (managed or observed) via its `writeConnectionSecretToRef`. This secret contains all the Stack's `atProvider` fields as individual keys using their Terraform attribute names.

When `stackSecretRef` is set, the secret's keys are merged into the credential map with the following key remapping applied:

| Stack secret key | Remapped to |
|------------------|-------------|
| `oncall_api_url` | `oncall_url` |
| `id` | `stack_id` |

All other keys are passed through unchanged (e.g., `url`, `fleet_management_url`, `org_id`).

### Precedence (lowest to highest)

1. **Primary credential secret** (`credentials.secretRef`) — base credentials
2. **Stack secret** (`stackSecretRef`) — overrides the primary secret
3. **ProviderConfig spec fields** (`url`, `oncallUrl`, etc.) — overrides both secrets

### Example: Using Stack Secret with ProviderConfig

**Step 1: Create a Stack that writes its details to a Secret:**
```yaml
apiVersion: cloud.grafana.m.crossplane.io/v1alpha1
kind: Stack
metadata:
  name: my-stack
  namespace: default
spec:
  forProvider:
    name: my-stack
    slug: my-stack
  writeConnectionSecretToRef:
    name: my-stack-details
```

**Step 2: Create a StackServiceAccountToken that writes auth to a Secret:**
```yaml
apiVersion: cloud.grafana.m.crossplane.io/v1alpha1
kind: StackServiceAccountToken
metadata:
  name: my-stack-sa-token
  namespace: default
spec:
  forProvider:
    stackSlugRef:
      name: my-stack
    serviceAccountRef:
      name: my-stack-sa
    name: crossplane
  writeConnectionSecretToRef:
    name: my-stack-token
```

**Step 3: Reference both secrets in a ProviderConfig:**
```yaml
apiVersion: grafana.m.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: my-stack-config
  namespace: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: my-stack-token
      namespace: default
      key: instanceCredentials
  stackSecretRef:
    name: my-stack-details
    namespace: default
```

The ProviderConfig will read:
- `auth` and `url` from the `instanceCredentials` key in `my-stack-token`
- `oncall_url` (remapped from `oncall_api_url`), `fleet_management_url`, `org_id`, `stack_id` (remapped from `id`), and all other Stack fields from `my-stack-details`
- Stack secret values override primary credential values where keys overlap (e.g., `url`)

### Stack secret coverage of ProviderConfig fields

The table below shows every ProviderConfig spec field and whether the stack connection secret can provide it via `stackSecretRef`.

| ProviderConfig field | Covered by `stackSecretRef`? | Stack secret key | Future source if not covered |
|---|---|---|---|
| `spec.url` | Yes | `url` | — |
| `spec.oncallUrl` | Yes | `oncall_api_url` (remapped to `oncall_url`) | — |
| `spec.fleetManagementUrl` | Yes | `fleet_management_url` | — |
| `spec.orgId` | Yes | `org_id` | — |
| `spec.stackId` | Yes | `id` (remapped to `stack_id`) | — |
| `spec.smUrl` | No | — | `grafana_synthetic_monitoring_installation` exposes `sm_url` in its state |
| `spec.cloudApiUrl` | No | — | Defaults to Grafana Cloud; can't be sourced from the stack (chicken/egg: the Cloud API is needed to read the stack) |
| `spec.cloudProviderUrl` | No | — | Not currently exposed by any Grafana resource |
| `spec.connectionsApiUrl` | No | — | Not currently exposed by any Grafana resource |

**Summary:** `stackSecretRef` covers **5 of 9** ProviderConfig spec fields. The remaining 4 are not stack attributes — they are either environment-specific URLs or require separate resources.

Additional informational keys in the stack secret include: `alertmanager_url`, `prometheus_url`, `logs_url`, `traces_url`, `graphite_url`, `profiles_url`, `otlp_url`, `influx_url`, `slug`, `name`, `status`, `region_slug`, and all service-specific names, statuses, and user IDs.

## Override Behavior

When a field is marked as "Overridable by ProviderConfig":

1. The value from the secret is used as the **default** configuration
2. If the corresponding field is set in the ProviderConfig spec, it **overrides** the secret value
3. This allows you to store common configuration in the secret while customizing specific values per ProviderConfig

### Example: Overriding URL

**Secret:**
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: grafana-creds
stringData:
  credentials: |
    {
      "url": "https://default.grafana.com",
      "auth": "admin:admin"
    }
```

**ProviderConfig (using default URL):**
```yaml
apiVersion: grafana.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default-grafana
spec:
  credentials:
    source: Secret
    secretRef:
      name: grafana-creds
      key: credentials
  # url is not specified, will use "https://default.grafana.com" from secret
```

**ProviderConfig (overriding URL):**
```yaml
apiVersion: grafana.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: custom-grafana
spec:
  url: "https://custom.grafana.com"  # Overrides the URL from secret
  credentials:
    source: Secret
    secretRef:
      name: grafana-creds
      key: credentials
```

## Field Categories

### Core Grafana Configuration
- `url` - The main Grafana instance URL
- `auth` - Authentication credentials (username:password or API token)
- `org_id` - Organization ID for multi-tenant Grafana instances

### Grafana Cloud APIs
- `cloud_access_policy_token` - For managing cloud resources
- `cloud_api_url` - Cloud API endpoint
- `cloud_provider_access_token` - For cloud provider operations
- `cloud_provider_url` - Cloud provider API endpoint
- `stack_id` - Stack identifier for cloud instances

### Product-Specific APIs
- **OnCall**: `oncall_access_token`, `oncall_url`
  - Note: `oncall_access_token` is not required if the `auth` field contains a Grafana service account token with OnCall permissions
- **Synthetic Monitoring**: `sm_access_token`, `sm_url`
- **Frontend Observability**: `frontend_o11y_api_access_token`
- **Fleet Management**: `fleet_management_auth`, `fleet_management_url`
- **Connections API**: `connections_api_access_token`, `connections_api_url`
- **k6 Cloud**: `k6_access_token`

## Authentication Options

### OnCall Authentication

Grafana OnCall resources can be authenticated in two ways:

1. **Dedicated OnCall Token** (recommended for OnCall-only access):
   ```json
   {
     "oncall_access_token": "your-oncall-token",
     "oncall_url": "https://oncall-prod-us-central-0.grafana.net/oncall"
   }
   ```

2. **Service Account Token** (when using a Grafana service account with OnCall permissions):
   ```json
   {
     "auth": "glsa_your_service_account_token",
     "url": "https://your-instance.grafana.net"
   }
   ```

   In this case, `oncall_access_token` is not required. The Grafana provider will use the service account token to authenticate OnCall API requests. The OnCall URL will be automatically derived from the Grafana instance URL unless explicitly overridden with `oncall_url`.

## Security Considerations

1. **Sensitive Fields**: All `*_token`, `*_auth`, and `auth` fields contain sensitive credentials and should never be exposed or logged
2. **Secret Management**: Use proper Kubernetes RBAC to restrict access to secrets containing these credentials
3. **URL Overrides**: While URLs can be overridden in ProviderConfig, tokens cannot, ensuring credentials remain secure in secrets
4. **Access Tokens vs API Keys**: Prefer using access tokens over the deprecated `cloud_api_key` for better security and rotation capabilities
5. **Service Account Tokens**: When using Grafana service account tokens, ensure they have the appropriate permissions for all resources you intend to manage

## Reference

For more information on the Grafana Terraform provider configuration, see the [official documentation](https://registry.terraform.io/providers/grafana/grafana/latest/docs).

The implementation can be found in:
- Secret field processing: `internal/clients/grafana.go`
- Stack secret merging: `internal/clients/grafana.go` (`mergeStackSecret`)
- ProviderConfig spec definition: `apis/cluster/v1beta1/types.go` and `apis/namespaced/v1beta1/types.go`
- Managed Stack connection details: `config/grafana/cloud.go` (`AdditionalConnectionDetailsFn`)
- Observed Stack connection details: generated in `internal/controller/namespaced/observed/cloud/zz_stack_spec.go` (`ConnectionDetailsFn`)
- Connection details controller support: `pkg/tfdatasource/controller.go`
