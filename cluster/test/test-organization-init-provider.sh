#!/usr/bin/env bash
set -euo pipefail

KUBECTL=${KUBECTL:-kubectl}
GRAFANA_PORT=${GRAFANA_PORT:-3000}
GRAFANA_URL="http://127.0.0.1:${GRAFANA_PORT}"
OMITTED_ORG=organization-users-omitted
INIT_EMPTY_ORG=organization-users-init-empty
EXTERNAL_USER_PREFIX=organization-external-
PORT_FORWARD_PID=

trap 'if [[ -n "${PORT_FORWARD_PID}" ]]; then kill "${PORT_FORWARD_PID}" 2>/dev/null || true; wait "${PORT_FORWARD_PID}" 2>/dev/null || true; fi' EXIT

grafana_api() {
	curl --fail --silent --show-error --user admin:admin "$@"
}

wait_for_grafana() {
	for _ in $(seq 1 60); do
		if grafana_api "${GRAFANA_URL}/api/health" >/dev/null 2>&1; then
			return
		fi
		sleep 1
	done
	echo "Grafana did not become reachable at ${GRAFANA_URL}" >&2
	return 1
}

external_name() {
	"${KUBECTL}" get organizations.oss.grafana.crossplane.io "$1" \
		-o jsonpath='{.metadata.annotations.crossplane\.io/external-name}'
}

add_memberships() {
	local org_id=$1
	local role
	for role in Admin Editor Viewer; do
		grafana_api --request POST \
			--header 'Content-Type: application/json' \
			--data "{\"loginOrEmail\":\"${EXTERNAL_USER_PREFIX}${role,,}@example.com\",\"role\":\"${role}\"}" \
			"${GRAFANA_URL}/api/orgs/${org_id}/users" >/dev/null
	done
}

membership_count() {
	local org_id=$1
	grafana_api "${GRAFANA_URL}/api/orgs/${org_id}/users" | jq \
		--arg prefix "${EXTERNAL_USER_PREFIX}" \
		'[.[] | select(.email | startswith($prefix))] | length'
}

wait_for_organization_name() {
	local org_id=$1
	local expected=$2
	for _ in $(seq 1 60); do
		if [[ $(grafana_api "${GRAFANA_URL}/api/orgs/${org_id}" | jq --raw-output '.name') == "${expected}" ]]; then
			return
		fi
		sleep 1
	done
	echo "Organization ${org_id} was not updated to ${expected}" >&2
	return 1
}

assert_membership_count() {
	local org_id=$1
	local expected=$2
	local description=$3
	local actual
	actual=$(membership_count "${org_id}")
	if [[ ${actual} -ne ${expected} ]]; then
		echo "FAIL: ${description}: expected ${expected} external memberships, got ${actual}" >&2
		return 1
	fi
	echo "PASS: ${description}"
}

remove_memberships() {
	local org_id=$1
	local user_id
	while read -r user_id; do
		grafana_api --request DELETE "${GRAFANA_URL}/api/orgs/${org_id}/users/${user_id}" >/dev/null
	done < <(
		grafana_api "${GRAFANA_URL}/api/orgs/${org_id}/users" | jq --raw-output \
			--arg prefix "${EXTERNAL_USER_PREFIX}" \
			'.[] | select(.email | startswith($prefix)) | .userId'
	)
}

"${KUBECTL}" wait --for=condition=Test --timeout=5m \
	organizations.oss.grafana.crossplane.io/${OMITTED_ORG} \
	organizations.oss.grafana.crossplane.io/${INIT_EMPTY_ORG} \
	users.oss.grafana.crossplane.io/organization-external-admin \
	users.oss.grafana.crossplane.io/organization-external-editor \
	users.oss.grafana.crossplane.io/organization-external-viewer

"${KUBECTL}" --namespace grafana port-forward service/grafana "${GRAFANA_PORT}:80" >/dev/null 2>&1 &
PORT_FORWARD_PID=$!
wait_for_grafana

omitted_org_id=$(external_name "${OMITTED_ORG}")
init_empty_org_id=$(external_name "${INIT_EMPTY_ORG}")
add_memberships "${omitted_org_id}"
add_memberships "${init_empty_org_id}"
assert_membership_count "${omitted_org_id}" 3 "out-of-band users added to organization with omitted sets"
assert_membership_count "${init_empty_org_id}" 3 "out-of-band users added to organization with initProvider sets"

omitted_updated_name="Organization Users Omitted Updated"
init_empty_updated_name="Organization Users Init Empty Updated"
"${KUBECTL}" patch organizations.oss.grafana.crossplane.io "${OMITTED_ORG}" --type merge \
	--patch "{\"spec\":{\"forProvider\":{\"name\":\"${omitted_updated_name}\"}}}" >/dev/null
"${KUBECTL}" patch organizations.oss.grafana.crossplane.io "${INIT_EMPTY_ORG}" --type merge \
	--patch "{\"spec\":{\"forProvider\":{\"name\":\"${init_empty_updated_name}\"}}}" >/dev/null
wait_for_organization_name "${omitted_org_id}" "${omitted_updated_name}"
wait_for_organization_name "${init_empty_org_id}" "${init_empty_updated_name}"

result=0
assert_membership_count "${omitted_org_id}" 3 "Terraform preserves omitted membership sets" || result=1
assert_membership_count "${init_empty_org_id}" 3 "initProvider membership sets are ignored after creation" || result=1

# Keep Grafana organization deletion reliable when uptest deletes all resources.
remove_memberships "${omitted_org_id}"
remove_memberships "${init_empty_org_id}"
exit "${result}"
