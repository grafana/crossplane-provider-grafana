# Renovate Post-Upgrade Tasks

A generic, language-agnostic composite action that detects which packages
changed in a Renovate PR and runs corresponding post-upgrade commands. Results
are committed and pushed back to the PR branch.

> [!NOTE]
> This action should probably live in
> [grafana/shared-workflows](https://github.com/grafana/shared-workflows).

## Inputs

| Name       | Required | Description                                |
|------------|----------|--------------------------------------------|
| `rules`    | Yes      | JSON array of rule objects (see below)     |
| `base-ref` | Yes      | Base branch to diff against (e.g. `main`) |

### Rules format

```json
[
  {
    "packages": ["github.com/grafana/terraform-provider-grafana/v4"],
    "commands": ["make submodules", "make generate"],
    "files": ["go.mod"]
  }
]
```

Each rule must have all three fields:

- **`packages`**: Package names to match in the diff. If any package appears in
  the diff of the specified files, the rule matches.
- **`commands`**: Commands to run when the rule matches.
- **`files`**: Files to diff against the base branch for package detection.

Multiple rules are supported. Commands are deduplicated across rules.

## Outputs

| Name           | Description                                      |
|----------------|--------------------------------------------------|
| `has-commands` | `true` if any rules matched, `false` otherwise   |
| `commands`     | Newline-separated list of commands that were run  |

## Prerequisites

The action assumes:

1. The repository is already checked out with `fetch-depth: 0` (full history
   needed for diffing against the base branch).
2. The checkout token has push permissions (so the commit can be pushed back).
3. Any language-specific tooling needed by the commands is already set up
   (e.g. Go, Node, Python).

## Error handling

- All steps use `set -euo pipefail` for strict error handling.
- If a command fails, the job fails immediately. The "Commit and push" step is
  skipped, so no broken state gets pushed.
- The failure is visible on the PR as a failed check.

## Example usage

```yaml
name: Renovate Post-Upgrade Tasks

on:
  pull_request:
    types: [opened, synchronize]

jobs:
  post-upgrade:
    runs-on: ubuntu-24.04
    if: github.event.pull_request.user.login == 'renovate[bot]'

    permissions:
      contents: read
      id-token: write

    steps:
      - name: Get GitHub App Token
        id: get-github-token
        uses: grafana/shared-workflows/actions/create-github-app-token@create-github-app-token/v0.2.2
        with:
          github_app: your-app-name

      - name: Checkout
        uses: actions/checkout@v4
        with:
          ref: ${{ github.head_ref }}
          token: ${{ steps.get-github-token.outputs.token }}
          fetch-depth: 0

      # Set up any tooling your commands need
      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version-file: go.mod

      - name: Run post-upgrade tasks
        uses: ./.github/actions/renovate-post-upgrade
        with:
          base-ref: ${{ github.base_ref }}
          rules: |
            [
              {
                "packages": ["github.com/grafana/terraform-provider-grafana/v4"],
                "commands": ["make submodules", "make generate"],
                "files": ["go.mod"]
              }
            ]
```

### Node.js example

```yaml
      - uses: actions/setup-node@v4
        with:
          node-version-file: .nvmrc

      - uses: ./.github/actions/renovate-post-upgrade
        with:
          base-ref: ${{ github.base_ref }}
          rules: |
            [
              {
                "packages": ["@grafana/ui", "@grafana/data"],
                "commands": ["npm run generate"],
                "files": ["package.json"]
              }
            ]
```
