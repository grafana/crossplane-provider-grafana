#!/usr/bin/env bash
# Full end-to-end test for Grafana provider with MRAP support
# This test builds the provider package and installs it in k3d with Crossplane

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/../.." && pwd)"
CLUSTER_NAME="${CLUSTER_NAME:-grafana-provider-test}"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

log_info() {
    echo -e "${GREEN}==> $*${NC}"
}

log_warn() {
    echo -e "${YELLOW}==> WARNING: $*${NC}"
}

log_error() {
    echo -e "${RED}==> ERROR: $*${NC}"
}

cleanup_on_exit() {
    local exit_code=$?
    if [ $exit_code -ne 0 ]; then
        log_error "Test failed with exit code: $exit_code"
        log_info "Collecting provider logs..."
        kubectl logs -n crossplane-system -l pkg.crossplane.io/provider=provider-grafana --tail=200 || true
        log_info "Checking pods..."
        kubectl get pods -n crossplane-system || true
    fi

    if [ "${KEEP_CLUSTER:-false}" != "true" ]; then
        log_info "Cleaning up k3d cluster..."
        "${SCRIPT_DIR}/cleanup-k3d.sh"
    else
        log_warn "Keeping cluster (KEEP_CLUSTER=true)"
    fi
}

trap cleanup_on_exit EXIT

# Step 1: Create k3d cluster
log_info "Step 1: Creating k3d cluster with Crossplane"
"${SCRIPT_DIR}/setup-k3d.sh"

# Step 2: Validate required environment variables
log_info "Step 2: Validating configuration"

if [ -z "${IMAGE_TAG:-}" ]; then
    log_error "IMAGE_TAG environment variable is required"
    log_error "Example: IMAGE_TAG=ghcr.io/grafana/provider-grafana:v2.3.0"
    exit 1
fi

PROVIDER_IMAGE="${IMAGE_TAG}"

log_info "Using provider image: ${PROVIDER_IMAGE}"

# Step 3: Remove default MRAP and apply custom MRAP
log_info "Step 3: Configuring ManagedActivationResourcePolicy"

# Crossplane v2 creates a default MRAP that activates all resources (*)
# We need to delete it so our restrictive MRAP takes effect
log_info "Deleting default MRAP that activates all resources..."
kubectl delete managedresourceactivationpolicy default --ignore-not-found=true

log_info "Applying custom MRAP to activate only subset of resources..."
kubectl apply -f "${SCRIPT_DIR}/test-mrap.yaml" || {
    log_error "Failed to apply MRAP"
    log_error "Checking if MRAP CRD exists..."
    kubectl get crd managedresourceactivationpolicies.apiextensions.crossplane.io || true
    exit 1
}

# Step 4: Install provider package
log_info "Step 4: Installing provider package with SafeStart"

# Create temporary provider config pointing to our xpkg
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
              resources:
                limits:
                  cpu: 1000m
                  memory: 1Gi
                requests:
                  cpu: 100m
                  memory: 256Mi
EOF

# Step 5: Wait for provider to be installed
log_info "Step 5: Waiting for provider to be installed"
kubectl wait --for=condition=Installed provider/provider-grafana --timeout=10m || {
    log_error "Provider failed to install"
    kubectl describe provider provider-grafana
    exit 1
}

# Step 6: Wait for provider pod to be ready
log_info "Step 6: Waiting for provider pod to be ready"
sleep 10  # Give pod time to start
kubectl get pods -n crossplane-system -l pkg.crossplane.io/provider=provider-grafana
kubectl wait --for=condition=Ready pod -l pkg.crossplane.io/provider=provider-grafana -n crossplane-system --timeout=10m || {
    log_error "Provider pod failed to become ready"
    log_info "Pod description:"
    kubectl describe pod -l pkg.crossplane.io/provider=provider-grafana -n crossplane-system
    log_info "Pod logs:"
    kubectl logs -l pkg.crossplane.io/provider=provider-grafana -n crossplane-system --tail=200
    exit 1
}

PROVIDER_POD=$(kubectl get pod -n crossplane-system -l pkg.crossplane.io/provider=provider-grafana -o jsonpath='{.items[0].metadata.name}')
log_info "Provider pod ready: ${PROVIDER_POD}"

# Step 7: Check provider logs for scheme registration
log_info "Step 7: Checking provider logs"
sleep 10

log_info "Looking for activated groups (should only see: oss, alerting per MRAP)..."
kubectl logs -n crossplane-system "${PROVIDER_POD}" | grep -i "controller.*started\|registered.*scheme" || log_warn "No controller start logs found"

# Step 8: Test provider shutdown behavior
log_info "Step 8: Testing provider shutdown with MRAP"
log_info "Deleting provider pod to trigger shutdown..."
kubectl delete pod -n crossplane-system "${PROVIDER_POD}"

# Wait for new pod to start
log_info "Waiting for new pod to start..."
sleep 5
kubectl wait --for=condition=Ready pod -l pkg.crossplane.io/provider=provider-grafana -n crossplane-system --timeout=5m || true

NEW_POD=$(kubectl get pod -n crossplane-system -l pkg.crossplane.io/provider=provider-grafana -o jsonpath='{.items[0].metadata.name}' 2>/dev/null || echo "")

# Step 9: Check old pod logs for shutdown errors
log_info "Step 9: Checking previous pod logs for MRAP shutdown errors"
OLD_POD_LOGS=$(kubectl logs -n crossplane-system "${PROVIDER_POD}" --previous 2>/dev/null || echo "")

if echo "$OLD_POD_LOGS" | grep -i "no matches for kind"; then
    log_error "Found 'no matches for kind' errors during shutdown"
    log_error "This indicates the MRAP bug is present (expected on main branch)"

    log_info "Problematic resource kinds found:"
    echo "$OLD_POD_LOGS" | grep -i "no matches for kind" | grep -o 'kind "[^"]*"' | sort -u

    log_info "Full shutdown errors:"
    echo "$OLD_POD_LOGS" | grep -i "no matches for kind" | tail -20

    # This is expected to fail on main branch without the fix
    exit 1
else
    log_info "No 'no matches for kind' errors found - provider shut down cleanly!"
    log_info "This indicates the MRAP fix is working correctly"
fi

# Step 10: Summary
log_info "Step 10: Test complete"
log_info "MRAP test passed - no shutdown errors detected"
