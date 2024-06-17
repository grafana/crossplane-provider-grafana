#!/usr/bin/env bash
set -aeuo pipefail

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

echo "Running setup.sh"

echo "Starting grafana on the cluster..."
${KUBECTL} apply -f "${SCRIPT_DIR}/grafana.yaml"

echo "Creating provider..."
${KUBECTL} apply -f "${SCRIPT_DIR}/provider.yaml"

echo "Waiting for grafana to come online..."
${KUBECTL} -n grafana wait --for=condition=Available deployment/grafana --timeout=5m

echo "Waiting until provider is healthy..."
${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m

echo "Waiting for all pods to come online..."
${KUBECTL} -n upbound-system wait --for=condition=Available deployment --all --timeout=5m

