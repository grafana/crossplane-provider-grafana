#!/usr/bin/env bash
# Setup k3d test cluster for testing the Grafana provider with MRAP

set -euo pipefail

CLUSTER_NAME="${CLUSTER_NAME:-grafana-provider-test}"
CROSSPLANE_VERSION="${CROSSPLANE_VERSION:-1.18.0}"

echo "==> Creating k3d cluster: ${CLUSTER_NAME}"
k3d cluster create "${CLUSTER_NAME}" \
  --agents 1 \
  --wait \
  --timeout 5m

echo "==> Waiting for cluster to be ready"
kubectl wait --for=condition=Ready nodes --all --timeout=2m

echo "==> Installing Crossplane ${CROSSPLANE_VERSION}"
helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update

helm install crossplane \
  crossplane-stable/crossplane \
  --namespace crossplane-system \
  --create-namespace \
  --version "${CROSSPLANE_VERSION}" \
  --wait \
  --timeout 5m

echo "==> Waiting for Crossplane to be ready"
kubectl wait --for=condition=Available deployment/crossplane -n crossplane-system --timeout=5m

echo "==> k3d cluster setup complete!"
echo "    Cluster: ${CLUSTER_NAME}"
echo "    Crossplane: ${CROSSPLANE_VERSION}"
echo ""
echo "To use this cluster:"
echo "    export KUBECONFIG=\$(k3d kubeconfig write ${CLUSTER_NAME})"
