#!/usr/bin/env bash
# Cleanup k3d test cluster

set -euo pipefail

CLUSTER_NAME="${CLUSTER_NAME:-grafana-provider-test}"

echo "==> Deleting k3d cluster: ${CLUSTER_NAME}"
k3d cluster delete "${CLUSTER_NAME}" || true

echo "==> Cleanup complete!"
