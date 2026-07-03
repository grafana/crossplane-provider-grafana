#!/usr/bin/env bash
#
# Create a verified (signed) commit on a remote branch from the staged changes
# in the working tree, using GitHub's createCommitOnBranch GraphQL mutation.
#
# Commits created through the API with a GitHub App / Actions token are signed
# by GitHub and therefore satisfy branch rules that require verified signatures
# (which plain `git push` of locally-created commits does not).
#
# After the remote branch is advanced, the local branch is reset to match it so
# subsequent steps continue from the signed commit.
#
# Usage:
#   create-signed-commit.sh <branch> <commit-message>
#
# Requirements:
#   - GH_TOKEN must be set to a token that can write to the repo.
#   - GITHUB_REPOSITORY must be set (owner/repo).
#   - Changes to include must already be staged with `git add`.
#
# Exit status:
#   0  commit created (or nothing staged, in which case it is a no-op)
#   1  an error occurred
set -euo pipefail

branch="${1:?branch required}"
message="${2:?commit message required}"

owner="${GITHUB_REPOSITORY%%/*}"
repo="${GITHUB_REPOSITORY##*/}"

# Collect staged additions/modifications and deletions separately.
mapfile -t changed < <(git diff --cached --name-only --diff-filter=ACMR)
mapfile -t deleted < <(git diff --cached --name-only --diff-filter=D)

if [ "${#changed[@]}" -eq 0 ] && [ "${#deleted[@]}" -eq 0 ]; then
  echo "Nothing staged; skipping commit."
  exit 0
fi

# Build the FileChanges object as JSON: { additions: [...], deletions: [...] }.
#
# The base64-encoded contents of a file (and the accumulated JSON that holds
# them) can be large - many megabytes across a full regeneration. Passing such
# values via `jq --arg`/`--argjson` places the whole blob on the command line,
# which can exceed the OS ARG_MAX limit and fail with "Argument list too long".
# To stay within argv limits we stream large data through temp files with
# `--rawfile`/`--slurpfile` instead of putting it on argv.
tmp_b64=$(mktemp)
tmp_additions=$(mktemp)
tmp_deletions=$(mktemp)
tmp_changes=$(mktemp)
tmp_input=$(mktemp)
trap 'rm -f "$tmp_b64" "$tmp_additions" "$tmp_deletions" "$tmp_changes" "$tmp_input"' EXIT

printf '[]' >"$tmp_additions"
for path in "${changed[@]}"; do
  base64 -w0 "$path" >"$tmp_b64"
  jq -c --arg p "$path" --rawfile c "$tmp_b64" \
    '. += [{path: $p, contents: $c}]' "$tmp_additions" >"${tmp_additions}.new"
  mv "${tmp_additions}.new" "$tmp_additions"
done

printf '[]' >"$tmp_deletions"
for path in "${deleted[@]}"; do
  jq -c --arg p "$path" '. += [{path: $p}]' "$tmp_deletions" >"${tmp_deletions}.new"
  mv "${tmp_deletions}.new" "$tmp_deletions"
done

# Combine additions/deletions into the FileChanges object, reading both large
# arrays from files (--slurpfile wraps each in an outer array we index with [0]).
jq -cn --slurpfile a "$tmp_additions" --slurpfile d "$tmp_deletions" \
  '{additions: $a[0], deletions: $d[0]}' >"$tmp_changes"

# The mutation requires the current head OID of the branch as an optimistic lock.
head_oid=$(git rev-parse HEAD)

# $input here is a GraphQL variable, not a shell variable; keep single quotes.
# shellcheck disable=SC2016
mutation='mutation($input: CreateCommitOnBranchInput!) {
  createCommitOnBranch(input: $input) { commit { oid url } }
}'

jq -cn \
  --arg repo "${owner}/${repo}" \
  --arg branch "refs/heads/${branch}" \
  --arg oid "$head_oid" \
  --arg headline "$message" \
  --slurpfile changes "$tmp_changes" \
  '{
    branch: {repositoryNameWithOwner: $repo, branchName: $branch},
    expectedHeadOid: $oid,
    message: {headline: $headline},
    fileChanges: $changes[0]
  }' >"$tmp_input"

# Pass the query and variables as a single JSON request body. `gh api graphql`
# with -F cannot coerce a JSON string into a nested input object, so build the
# full GraphQL payload ({query, variables}) and feed it on stdin via --input -.
request=$(jq -cn --arg query "$mutation" --slurpfile input "$tmp_input" \
  '{query: $query, variables: {input: $input[0]}}')

new_oid=$(gh api graphql --input - <<<"$request" \
  --jq '.data.createCommitOnBranch.commit.oid')

echo "Created signed commit ${new_oid} on ${branch}."

# Sync the local branch to the new remote commit so later steps build on it.
git fetch origin "$branch"
git reset --hard "$new_oid"
