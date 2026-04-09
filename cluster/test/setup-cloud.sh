#!/usr/bin/env bash
set -aeuo pipefail

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)

echo "Running setup-cloud.sh"

# --- Validate required environment variables ---
# These point to a pre-existing Grafana Cloud stack.
for var in GRAFANA_URL GRAFANA_SA_TOKEN GRAFANA_ONCALL_URL; do
	if [[ -z "${!var:-}" ]]; then
		echo "ERROR: ${var} is required but not set."
		exit 1
	fi
done

# Strip trailing slashes from URLs to avoid double-slash path issues
GRAFANA_URL="${GRAFANA_URL%/}"
GRAFANA_ONCALL_URL="${GRAFANA_ONCALL_URL%/}"

echo "Using Grafana instance: ${GRAFANA_URL}"
echo "Using OnCall API URL: ${GRAFANA_ONCALL_URL}"

echo "Waiting until provider is healthy..."
${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m

echo "Waiting for all pods to come online..."
${KUBECTL} -n upbound-system wait --for=condition=Available deployment --all --timeout=5m

# --- Install Crossplane functions ---
echo "Installing Crossplane functions..."
${KUBECTL} apply -f - <<'EOF'
apiVersion: pkg.crossplane.io/v1beta1
kind: Function
metadata:
  name: function-patch-and-transform
spec:
  package: xpkg.upbound.io/crossplane-contrib/function-patch-and-transform:v0.8.1
---
apiVersion: pkg.crossplane.io/v1beta1
kind: Function
metadata:
  name: function-kcl
spec:
  package: xpkg.upbound.io/crossplane-contrib/function-kcl:v0.11.2
---
apiVersion: pkg.crossplane.io/v1beta1
kind: Function
metadata:
  name: function-auto-ready
spec:
  package: xpkg.upbound.io/crossplane-contrib/function-auto-ready:v0.4.0
EOF

echo "Waiting for functions to be healthy..."
${KUBECTL} wait function.pkg --all --for condition=Healthy --timeout 5m

# --- Create instance credentials secret ---
echo "Creating instance credentials secret..."
${KUBECTL} apply -f - <<EOF
apiVersion: v1
kind: Secret
metadata:
  name: e2e-cloud-instance-creds
  namespace: upbound-system
type: Opaque
stringData:
  instanceCredentials: |
    {
      "url": "${GRAFANA_URL}",
      "auth": "${GRAFANA_SA_TOKEN}"
    }
EOF

# --- Create instance ProviderConfig with oncallUrl override ---
echo "Creating instance ProviderConfig..."
${KUBECTL} apply -f - <<EOF
apiVersion: grafana.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: e2e-cloud-instance
spec:
  oncallUrl: "${GRAFANA_ONCALL_URL}"
  credentials:
    source: Secret
    secretRef:
      name: e2e-cloud-instance-creds
      namespace: upbound-system
      key: instanceCredentials
EOF

# --- Create ClusterProviderConfig for namespaced resources ---
echo "Waiting for ClusterProviderConfig CRD to be established..."
${KUBECTL} wait crd/clusterproviderconfigs.grafana.m.crossplane.io \
	--for=condition=Established --timeout=5m

echo "Creating ClusterProviderConfig..."
${KUBECTL} apply -f - <<EOF
apiVersion: grafana.m.crossplane.io/v1beta1
kind: ClusterProviderConfig
metadata:
  name: e2e-cloud-instance
spec:
  oncallUrl: "${GRAFANA_ONCALL_URL}"
  credentials:
    source: Secret
    secretRef:
      name: e2e-cloud-instance-creds
      namespace: upbound-system
      key: instanceCredentials
EOF

REPO_ROOT=$(cd -- "${SCRIPT_DIR}/../.." &>/dev/null && pwd)

# --- Apply XRDs and Compositions (without the example Claims) ---
echo "Applying XRDs and Compositions..."
for f in "${REPO_ROOT}/examples/compositions/oncall-shift-rolling-users.yaml"; do
	if [[ -f "${f}" ]]; then
		echo "  Applying XRDs/Compositions from ${f}..."
		# Extract only CompositeResourceDefinition and Composition documents.
		awk '
		BEGIN { doc=""; emit=0; first=1 }
		/^---/ {
			if (emit && doc != "") {
				if (!first) print "---";
				print doc;
				first=0
			}
			doc=""; emit=0; next
		}
		{
			doc = doc $0 "\n"
			if ($0 ~ /^kind: (CompositeResourceDefinition|Composition)$/) emit=1
		}
		END {
			if (emit && doc != "") {
				if (!first) print "---";
				print doc
			}
		}
		' "${f}" | ${KUBECTL} apply -f -
	fi
done

echo "Waiting for XRDs to be established..."
sleep 5
${KUBECTL} wait xrd --all --for condition=Established --timeout 2m

echo "setup-cloud.sh completed successfully."
