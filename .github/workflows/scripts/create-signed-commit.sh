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
additions="[]"
for path in "${changed[@]}"; do
  contents=$(base64 -w0 "$path")
  additions=$(jq -c --arg p "$path" --arg c "$contents" \
    '. += [{path: $p, contents: $c}]' <<<"$additions")
done

deletions="[]"
for path in "${deleted[@]}"; do
  deletions=$(jq -c --arg p "$path" '. += [{path: $p}]' <<<"$deletions")
done

file_changes=$(jq -cn --argjson a "$additions" --argjson d "$deletions" \
  '{additions: $a, deletions: $d}')

# The mutation requires the current head OID of the branch as an optimistic lock.
head_oid=$(git rev-parse HEAD)

# $input here is a GraphQL variable, not a shell variable; keep single quotes.
# shellcheck disable=SC2016
mutation='mutation($input: CreateCommitOnBranchInput!) {
  createCommitOnBranch(input: $input) { commit { oid url } }
}'

input=$(jq -cn \
  --arg repo "${owner}/${repo}" \
  --arg branch "refs/heads/${branch}" \
  --arg oid "$head_oid" \
  --arg headline "$message" \
  --argjson changes "$file_changes" \
  '{
    branch: {repositoryNameWithOwner: $repo, branchName: $branch},
    expectedHeadOid: $oid,
    message: {headline: $headline},
    fileChanges: $changes
  }')

new_oid=$(gh api graphql -f query="$mutation" -F input="$input" \
  --jq '.data.createCommitOnBranch.commit.oid')

echo "Created signed commit ${new_oid} on ${branch}."

# Sync the local branch to the new remote commit so later steps build on it.
git fetch origin "$branch"
git reset --hard "$new_oid"
