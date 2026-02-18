# End-to-End Tests for MRAP Support

This directory contains end-to-end tests for verifying that the Grafana provider correctly handles ManagedResourceActivationPolicies (MRAP) in Crossplane v2.

## What is MRAP?

ManagedResourceActivationPolicy (MRAP) is a Crossplane v2 feature that allows you to control which ManagedResourceDefinitions (MRDs) become active in your cluster. Instead of installing all 100+ resources that a provider ships with, you can activate only the specific resources you need (e.g., just dashboards and datasources).

## Problem Being Tested

The test validates that:
1. The provider works correctly with restrictive MRAPs
2. Only activated resources are available in the cluster
3. Disabled resources don't cause issues during provider lifecycle
4. The provider can start, run, and shut down cleanly with MRAP enabled

## Test Suite

### Files

- `setup-k3d.sh` - Creates a k3d cluster and installs Crossplane
- `cleanup-k3d.sh` - Deletes the k3d test cluster
- `test-mrap.yaml` - Sample MRAP that activates only a few resources
- `provider-config.yaml` - Provider deployment manifest with SafeStart enabled
- `run-mrap-test.sh` - Full test orchestration script

### Running Tests

#### Full Test Suite

Run the complete end-to-end test with a published provider image:

```bash
cd test/e2e
IMAGE_TAG=ghcr.io/grafana/provider-grafana:v2.3.0 ./run-mrap-test.sh
```

This will:
1. Create a k3d cluster with Crossplane v2.1.4
2. Validate configuration (check IMAGE_TAG is set)
3. Delete the default MRAP that activates all resources
4. Apply a custom MRAP that activates only dashboards
5. Install the provider with SafeStart enabled
6. Wait for the provider pod to become ready
7. Check provider logs for controller starts
8. Trigger a provider restart to test shutdown
9. Check previous pod logs for shutdown errors
10. Clean up the cluster

#### Keep Cluster for Debugging

To keep the cluster after the test for manual inspection:

```bash
KEEP_CLUSTER=true IMAGE_TAG=ghcr.io/grafana/provider-grafana:v2.3.0 ./run-mrap-test.sh
```

#### Manual Testing

Set up the cluster manually:

```bash
# Create cluster with Crossplane v2
./setup-k3d.sh

# Set provider image (will be pulled from GHCR)
PROVIDER_IMAGE=ghcr.io/grafana/provider-grafana:v2.3.0

# Delete default MRAP and apply custom one
kubectl delete managedresourceactivationpolicy default
kubectl apply -f test-mrap.yaml

# Install provider with SafeStart
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-grafana
spec:
  package: ${PROVIDER_IMAGE}
  packagePullPolicy: IfNotPresent
  runtimeConfigRef:
    name: grafana-provider-config
---
apiVersion: pkg.crossplane.io/v1beta1
kind: DeploymentRuntimeConfig
metadata:
  name: grafana-provider-config
spec:
  deploymentTemplate:
    spec:
      selector: {}
      template:
        spec:
          containers:
            - name: package-runtime
              args:
                - --debug
                - --safe-start
EOF

# Watch logs
kubectl logs -f -n crossplane-system -l pkg.crossplane.io/provider=provider-grafana

# Clean up when done
./cleanup-k3d.sh
```

## Expected Results

### Success Criteria

✅ Crossplane v2.1.4 installs successfully
✅ Default MRAP is deleted
✅ Custom MRAP activates only specified resources (dashboards)
✅ Provider deploys to k3d cluster with SafeStart enabled
✅ Provider pod becomes ready
✅ Only activated MRDs are created in the cluster
✅ Provider shuts down cleanly on pod restart
✅ No errors in provider logs

### Verification

Check which resources are activated by the MRAP:

```bash
kubectl get managedresourceactivationpolicy grafana-test-policy -o yaml
```

Check which MRDs are active in the cluster:

```bash
kubectl get managedresourcedefinitions.apiextensions.crossplane.io | grep grafana
```

You should see only the resources specified in the MRAP (e.g., dashboards.oss.grafana.crossplane.io).

## Prerequisites

- Docker
- k3d
- kubectl
- Helm
- Access to published provider images in GHCR

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
./run-mrap-test.sh
```

Keep the cluster for manual inspection:
```bash
KEEP_CLUSTER=true ./run-mrap-test.sh
```

## How It Works

### Crossplane v2 MRAP Flow

1. **Crossplane Installation**:
   - Crossplane v2.1.4 is installed via Helm
   - Creates a default MRAP that activates all resources (`*`)

2. **Custom MRAP Configuration**:
   - Test deletes the default MRAP
   - Applies custom MRAP activating only specific resources (e.g., dashboards)
   - Crossplane processes the MRAP and creates only matching MRDs

3. **Provider Installation**:
   - Provider package (xpkg) is installed from GHCR
   - Provider pod starts with `--safe-start` flag
   - SafeStart waits for MRDs to exist before starting controllers

4. **Controller Activation**:
   - Only controllers for activated MRDs start
   - Disabled resources remain inactive
   - Provider becomes healthy with subset of resources

5. **Lifecycle Testing**:
   - Provider pod is deleted to trigger restart
   - Test verifies clean shutdown and startup
   - Checks logs for any errors

### Test Environment

- **k3d cluster**: Lightweight Kubernetes for local testing
- **Crossplane v2.1.4**: Latest stable v2 release
- **MRAP**: Controls which resources are activated
- **SafeStart**: Provider feature that waits for CRDs before starting controllers
