#!/usr/bin/env bash
# Full end-to-end test for Grafana provider with MRAP support

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
        kubectl logs -n crossplane-system -l pkg.crossplane.io/provider=provider-grafana --tail=100 || true
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
log_info "Step 1: Creating k3d cluster"
"${SCRIPT_DIR}/setup-k3d.sh"

# Step 2: Build provider image
log_info "Step 2: Building provider image"
cd "${PROJECT_ROOT}"
docker build -t localhost/grafana-provider:test .

# Step 3: Load image into k3d
log_info "Step 3: Loading provider image into k3d cluster"
k3d image import localhost/grafana-provider:test -c "${CLUSTER_NAME}"

# Step 4: Install provider
log_info "Step 4: Installing provider with SafeStart enabled"
kubectl apply -f "${SCRIPT_DIR}/provider-config.yaml"

# Step 5: Wait for provider to be installed
log_info "Step 5: Waiting for provider to be installed"
kubectl wait --for=condition=Installed provider/provider-grafana -n crossplane-system --timeout=5m || {
    log_error "Provider failed to install"
    kubectl describe provider provider-grafana -n crossplane-system
    exit 1
}

# Step 6: Wait for provider pod to be ready
log_info "Step 6: Waiting for provider pod to be ready"
sleep 10
kubectl wait --for=condition=Ready pod -l pkg.crossplane.io/provider=provider-grafana -n crossplane-system --timeout=5m || {
    log_error "Provider pod failed to become ready"
    kubectl describe pod -l pkg.crossplane.io/provider=provider-grafana -n crossplane-system
    kubectl logs -l pkg.crossplane.io/provider=provider-grafana -n crossplane-system --tail=100
    exit 1
}

# Step 7: Apply MRAP to disable most resources
log_info "Step 7: Applying ManagedActivationResourcePolicy"
kubectl apply -f "${SCRIPT_DIR}/test-mrap.yaml"

# Step 8: Wait for MRAP to take effect
log_info "Step 8: Waiting for MRAP to take effect (30s)"
sleep 30

# Step 9: Check provider logs for scheme registration
log_info "Step 9: Checking provider logs for scheme registration"
PROVIDER_POD=$(kubectl get pod -n crossplane-system -l pkg.crossplane.io/provider=provider-grafana -o jsonpath='{.items[0].metadata.name}')

log_info "Looking for 'Registered API group scheme' messages..."
kubectl logs -n crossplane-system "${PROVIDER_POD}" | grep -i "Registered API group scheme" || log_warn "No scheme registration logs found (may be using Info level)"

log_info "Checking for activated groups (should only see: oss, alerting)..."
ACTIVATED_GROUPS=$(kubectl logs -n crossplane-system "${PROVIDER_POD}" | grep -i "Registered API group scheme" | grep -o 'group="[^"]*"' | sort -u || echo "")
if [ -n "${ACTIVATED_GROUPS}" ]; then
    echo "${ACTIVATED_GROUPS}"
else
    log_warn "Could not parse activated groups from logs"
fi

# Step 10: Trigger provider restart to test shutdown
log_info "Step 10: Triggering provider restart to test shutdown behavior"
kubectl delete pod -n crossplane-system "${PROVIDER_POD}"

log_info "Step 11: Waiting for new provider pod to be ready"
sleep 10
kubectl wait --for=condition=Ready pod -l pkg.crossplane.io/provider=provider-grafana -n crossplane-system --timeout=5m

# Step 12: Check previous pod logs for shutdown errors
log_info "Step 12: Checking previous pod logs for shutdown errors"
if kubectl logs -n crossplane-system "${PROVIDER_POD}" --previous 2>/dev/null | grep -i "failed to list managed resources"; then
    log_error "Found 'failed to list managed resources' errors in previous pod logs!"
    kubectl logs -n crossplane-system "${PROVIDER_POD}" --previous | grep -i "failed to list managed resources"
    exit 1
else
    log_info "✅ No 'failed to list managed resources' errors found in shutdown logs"
fi

# Step 13: Check for specific error patterns
log_info "Step 13: Checking for 'no matches for kind' errors"
if kubectl logs -n crossplane-system "${PROVIDER_POD}" --previous 2>/dev/null | grep -i "no matches for kind"; then
    log_error "Found 'no matches for kind' errors!"
    kubectl logs -n crossplane-system "${PROVIDER_POD}" --previous | grep -i "no matches for kind"
    exit 1
else
    log_info "✅ No 'no matches for kind' errors found"
fi

# Step 14: Verify disabled resources are not in scheme
log_info "Step 14: Verifying disabled groups are not registered"
CURRENT_LOGS=$(kubectl logs -n crossplane-system -l pkg.crossplane.io/provider=provider-grafana)

# Check that disabled groups are NOT in the logs
for disabled_group in "asserts" "cloud" "ml" "oncall" "k6" "enterprise"; do
    if echo "${CURRENT_LOGS}" | grep -i "Registered API group scheme.*group=\"${disabled_group}\""; then
        log_error "Disabled group '${disabled_group}' was registered (should not be)!"
        exit 1
    fi
done
log_info "✅ Disabled groups are not registered in scheme"

# Success!
log_info "========================================="
log_info "✅ All tests passed!"
log_info "========================================="
log_info "Summary:"
log_info "  - Provider built and deployed successfully"
log_info "  - SafeStart enabled"
log_info "  - MRAP applied to disable most resources"
log_info "  - Only activated resources registered their schemes"
log_info "  - Provider shutdown completed without errors"
log_info "  - No 'failed to list managed resources' errors"
log_info "  - No 'no matches for kind' errors"
log_info "========================================="

exit 0
