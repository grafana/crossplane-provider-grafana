#!/usr/bin/env bash
set -euo pipefail

# Patch SM Check types to add ProbeNames field
# This script is automatically run during 'make generate' after upjet generates types

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"

echo "Patching SM Check types to add ProbeNames field..."

for check_types in \
  "${PROJECT_ROOT}/apis/namespaced/sm/v1alpha1/zz_check_types.go" \
  "${PROJECT_ROOT}/apis/cluster/sm/v1alpha1/zz_check_types.go"; do

  if [[ ! -f "${check_types}" ]]; then
    echo "Warning: ${check_types} not found, skipping"
    continue
  fi

  # Add ProbeNames field to CheckParameters struct (after Probes field)
  # Using sed for reliable in-place replacement with proper whitespace
  sed -i 's/\(	Probes \[\]\*float64 `json:"probes,omitempty" tf:"probes,omitempty"`\)/\1\
\
	\/\/ CUSTOM FIELD: List of probe names to resolve to probe IDs.\
	\/\/ If set, this will override the Probes field with the resolved probe IDs on every reconciliation.\
	\/\/ This field is excluded from Terraform via tf:"-" tag.\
	\/\/ +kubebuilder:validation:Optional\
	ProbeNames []string `json:"probeNames,omitempty" tf:"-"`/' "${check_types}"

  echo "  Patched: ${check_types}"
done

echo "SM Check types patched successfully"
