#!/usr/bin/env bash
set -aeuo pipefail

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)

echo "Running setup-cloud.sh"

# --- Validate required environment variables ---
for var in GRAFANA_SA_TOKEN GRAFANA_CLOUD_ACCESS_POLICY_TOKEN GRAFANA_STACK_SLUG; do
	if [[ -z "${!var:-}" ]]; then
		echo "ERROR: ${var} is required but not set."
		exit 1
	fi
done

echo "Using stack slug: ${GRAFANA_STACK_SLUG}"

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

# --- Create cloud API credentials secret ---
# This secret is used to authenticate against the Grafana Cloud API
# (e.g. for observing stacks). It only contains the cloud access policy token.
echo "Creating cloud API credentials secret..."
${KUBECTL} apply -f - <<EOF
apiVersion: v1
kind: Secret
metadata:
  name: e2e-cloud-api-creds
  namespace: upbound-system
type: Opaque
stringData:
  credentials: |
    {
      "cloud_access_policy_token": "${GRAFANA_CLOUD_ACCESS_POLICY_TOKEN}"
    }
EOF

# --- Create cloud API ClusterProviderConfig ---
echo "Waiting for ClusterProviderConfig CRD to be established..."
${KUBECTL} wait crd/clusterproviderconfigs.grafana.m.crossplane.io \
	--for=condition=Established --timeout=5m

echo "Creating cloud API ClusterProviderConfig..."
${KUBECTL} apply -f - <<'EOF'
apiVersion: grafana.m.crossplane.io/v1beta1
kind: ClusterProviderConfig
metadata:
  name: e2e-cloud-api
spec:
  credentials:
    source: Secret
    secretRef:
      name: e2e-cloud-api-creds
      namespace: upbound-system
      key: credentials
EOF

# --- Observe the existing stack ---
# The observed Stack publishes all its atProvider fields as connection details
# into a Secret. This secret is then used by the instance ProviderConfigs via
# stackSecretRef to automatically resolve url, oncall_url, org_id, etc.
echo "Creating observed Stack..."
${KUBECTL} apply -f - <<EOF
apiVersion: cloud.grafana.o.crossplane.io/v1alpha1
kind: Stack
metadata:
  name: e2e-observed-stack
  namespace: upbound-system
spec:
  forProvider:
    slug: "${GRAFANA_STACK_SLUG}"
  providerConfigRef:
    kind: ClusterProviderConfig
    name: e2e-cloud-api
  writeConnectionSecretToRef:
    name: e2e-stack-details
EOF

echo "Waiting for observed Stack to be ready..."
${KUBECTL} wait stack.cloud.grafana.o.crossplane.io/e2e-observed-stack \
	-n upbound-system --for=condition=Ready --timeout=5m

echo "Verifying stack connection secret exists..."
${KUBECTL} get secret e2e-stack-details -n upbound-system -o jsonpath='{.data.url}' >/dev/null

# --- Create auth-only instance credentials secret ---
# This secret only contains the SA token for authentication. The URL and other
# endpoint fields come from the stack connection secret via stackSecretRef.
echo "Creating auth-only instance credentials secret..."
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
      "auth": "${GRAFANA_SA_TOKEN}"
    }
EOF

# --- Create instance ProviderConfig (legacy, cluster-scoped) ---
# Uses stackSecretRef to get url, oncall_url, org_id, etc. from the observed
# stack's connection secret. No explicit URL or oncallUrl needed.
echo "Creating instance ProviderConfig..."
${KUBECTL} apply -f - <<'EOF'
apiVersion: grafana.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: e2e-cloud-instance
spec:
  credentials:
    source: Secret
    secretRef:
      name: e2e-cloud-instance-creds
      namespace: upbound-system
      key: instanceCredentials
  stackSecretRef:
    name: e2e-stack-details
    namespace: upbound-system
EOF

# --- Create ClusterProviderConfig for namespaced resources ---
# Same pattern: auth-only credentials + stackSecretRef for endpoints.
echo "Creating ClusterProviderConfig..."
${KUBECTL} apply -f - <<'EOF'
apiVersion: grafana.m.crossplane.io/v1beta1
kind: ClusterProviderConfig
metadata:
  name: e2e-cloud-instance
spec:
  credentials:
    source: Secret
    secretRef:
      name: e2e-cloud-instance-creds
      namespace: upbound-system
      key: instanceCredentials
  stackSecretRef:
    name: e2e-stack-details
    namespace: upbound-system
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
