#!/usr/bin/env bash
set -euo pipefail

# Patch OnCall contact point types to add URLSecretRef field
# This script is automatically run during 'make generate' after upjet generates types

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"

echo "Patching OnCall contact point types to add URLSecretRef field..."

# Patch cluster-scoped contact point types (uses v1.SecretKeySelector for cross-namespace refs)
CLUSTER_TYPES="${PROJECT_ROOT}/apis/cluster/alerting/v1alpha1/zz_contactpoint_types.go"
if [[ -f "${CLUSTER_TYPES}" ]]; then
  # Check if already patched (idempotency)
  if grep -q "// The OnCall webhook URL (from secret)" "${CLUSTER_TYPES}"; then
    echo "  Already patched: ${CLUSTER_TYPES}"
  else
    # Use a unique marker that only exists in Oncall structs: OncallIntegrationSelector
    # This ensures we only patch the Oncall URL field, not other contact point types

    # Patch OncallInitParameters and OncallParameters: Find the URL field that comes after OncallIntegrationSelector comment
    sed -i '/OncallIntegrationSelector$/,/^	URL \*string.*url,omitempty/ {
      /^	URL \*string.*url,omitempty/ {
        a\
\
	// The OnCall webhook URL (from secret).\
	// +kubebuilder:validation:Optional\
	URLSecretRef *v1.SecretKeySelector `json:"urlSecretRef,omitempty" tf:"-"`
      }
    }' "${CLUSTER_TYPES}"

    echo "  Patched: ${CLUSTER_TYPES}"
  fi
else
  echo "  Warning: ${CLUSTER_TYPES} not found, skipping"
fi

# Patch namespaced contact point types (uses v1.LocalSecretKeySelector for same-namespace only)
NAMESPACED_TYPES="${PROJECT_ROOT}/apis/namespaced/alerting/v1alpha1/zz_contactpoint_types.go"
if [[ -f "${NAMESPACED_TYPES}" ]]; then
  # Check if already patched (idempotency)
  if grep -q "// The OnCall webhook URL (from secret)" "${NAMESPACED_TYPES}"; then
    echo "  Already patched: ${NAMESPACED_TYPES}"
  else
    # Use the same unique marker approach

    # Patch both OncallInitParameters and OncallParameters
    sed -i '/OncallIntegrationSelector$/,/^	URL \*string.*url,omitempty/ {
      /^	URL \*string.*url,omitempty/ {
        a\
\
	// The OnCall webhook URL (from secret).\
	// +kubebuilder:validation:Optional\
	URLSecretRef *v1.LocalSecretKeySelector `json:"urlSecretRef,omitempty" tf:"-"`
      }
    }' "${NAMESPACED_TYPES}"

    echo "  Patched: ${NAMESPACED_TYPES}"
  fi
else
  echo "  Warning: ${NAMESPACED_TYPES} not found, skipping"
fi

echo "OnCall contact point types patched successfully"
