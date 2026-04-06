#!/usr/bin/env bash
set -euo pipefail

echo "Running teardown-cloud.sh"

# --- Delete Claims first (XR instances) ---
echo "Deleting Claims..."
${KUBECTL} delete oncallshiftrollingusers.oncall.grafana.crossplane.grafana.net --all \
	-n upbound-system --timeout=2m --wait=true 2>/dev/null || true

# Wait for composed resources to be cleaned up
echo "Waiting for composed resources to be deleted..."
sleep 10

# --- Delete ProviderConfigs and secrets ---
echo "Deleting ProviderConfigs..."
${KUBECTL} delete providerconfig.grafana.crossplane.io e2e-cloud-instance 2>/dev/null || true
${KUBECTL} delete clusterproviderconfig.grafana.m.crossplane.io e2e-cloud-instance 2>/dev/null || true
${KUBECTL} delete secret e2e-cloud-instance-creds -n upbound-system 2>/dev/null || true

# --- Delete XRDs and Compositions ---
echo "Deleting Compositions and XRDs..."
${KUBECTL} delete composition --all 2>/dev/null || true
${KUBECTL} delete xrd --all 2>/dev/null || true

# --- Delete Functions ---
echo "Deleting Functions..."
${KUBECTL} delete function.pkg function-patch-and-transform function-kcl function-auto-ready 2>/dev/null || true

echo "teardown-cloud.sh completed."
