# End-to-End Tests for MRAP Support

This directory contains end-to-end tests for verifying that the Grafana provider correctly handles ManagedActivationResourcePolicies (MRAP) with deferred scheme registration.

## Problem Being Tested

When using MRAP to disable certain CRDs, the provider previously crashed during shutdown with errors like:
```
failed to list managed resources: no matches for kind "SuppressedAssertionsConfig"
```

This was caused by all API types being registered with the manager's scheme upfront, even for resources that would never be activated. During shutdown, the controller-runtime manager tried to list all schemed resources, including those without CRDs.

## Solution

The implementation now defers scheme registration until controllers are actually activated. This ensures disabled resources never get registered, preventing shutdown errors.

## Test Suite

### Files

- `setup-k3d.sh` - Creates a k3d cluster and installs Crossplane
- `cleanup-k3d.sh` - Deletes the k3d test cluster
- `test-mrap.yaml` - Sample MRAP that activates only a few resources
- `provider-config.yaml` - Provider deployment manifest with SafeStart enabled
- `run-test.sh` - Full test orchestration script

### Running Tests

#### Full Test Suite

Run the complete end-to-end test:

```bash
cd test/e2e
./run-test.sh
```

This will:
1. Create a k3d cluster
2. Build the provider image
3. Load the image into k3d
4. Install the provider with SafeStart enabled
5. Apply an MRAP to disable most resources
6. Verify only activated resources register their schemes
7. Trigger a provider restart
8. Check logs for shutdown errors
9. Clean up the cluster

#### Keep Cluster for Debugging

To keep the cluster after the test:

```bash
KEEP_CLUSTER=true ./run-test.sh
```

#### Manual Testing

Set up the cluster manually:

```bash
# Create cluster
./setup-k3d.sh

# Build and load provider
docker build -t localhost/grafana-provider:test ../..
k3d image import localhost/grafana-provider:test -c grafana-provider-test

# Install provider
kubectl apply -f provider-config.yaml

# Apply MRAP
kubectl apply -f test-mrap.yaml

# Watch logs
kubectl logs -f -n crossplane-system -l pkg.crossplane.io/provider=provider-grafana

# Clean up when done
./cleanup-k3d.sh
```

## Expected Results

### Success Criteria

✅ Provider builds without errors
✅ Provider deploys to k3d cluster with SafeStart + MRAP
✅ Only activated controllers register their API schemes
✅ Provider starts successfully with only subset of CRDs activated
✅ Provider shuts down cleanly without "no matches for kind" errors
✅ Disabled resources don't cause shutdown failures
✅ Provider logs show scheme registration only for activated groups

### Log Examples

**Good** - Scheme registration for activated groups:
```
Registered API group scheme  group="oss"
Registered API group scheme  group="alerting"
```

**Bad** - Shutdown errors (should NOT appear):
```
failed to list managed resources: no matches for kind "SuppressedAssertionsConfig"
```

## Prerequisites

- Docker
- k3d
- kubectl
- Helm
- Go 1.21+

## Troubleshooting

### Provider Won't Start

Check the provider logs:
```bash
kubectl logs -n crossplane-system -l pkg.crossplane.io/provider=provider-grafana
```

Check the provider package status:
```bash
kubectl describe provider provider-grafana -n crossplane-system
```

### Image Load Fails

Ensure the image is built:
```bash
docker images | grep grafana-provider
```

Verify k3d cluster exists:
```bash
k3d cluster list
```

### Tests Fail

Run with debugging enabled:
```bash
set -x
./run-test.sh
```

Keep the cluster for manual inspection:
```bash
KEEP_CLUSTER=true ./run-test.sh
```

## Architecture

The deferred scheme registration works as follows:

1. **Main entry point** (`cmd/provider/main.go`):
   - Removes upfront `AddToScheme` calls for API groups
   - Keeps only `apiextensionsv1` registration (needed for CRD gate)
   - Calls `SetupWrapped` / `SetupGatedWrapped` instead of generated setup

2. **Wrapped setup functions** (`internal/controller/*/setup_wrapped.go`):
   - Intercepts each controller's setup call
   - Registers the controller's API group scheme (once per group)
   - Calls the original generated setup function

3. **Scheme registry** (`internal/controller/schemeregistry/`):
   - Thread-safe tracking of registered API groups
   - Maps group names to their SchemeBuilder.AddToScheme functions
   - Ensures each group is only registered once

4. **Controller activation**:
   - With SafeStart: Gate activates controllers when CRDs are created
   - Without SafeStart: All controllers start immediately
   - In both cases: Scheme is registered only when controller starts

5. **Shutdown**:
   - Manager only tries to list resources for registered schemes
   - Disabled resources (not registered) don't cause errors
   - Clean shutdown without "no matches for kind" errors
