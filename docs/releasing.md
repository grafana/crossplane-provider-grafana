# Releasing

## Versioning

This project follows [Semantic Versioning](https://semver.org/) (`MAJOR.MINOR.PATCH`).
Version bumps are determined automatically from
[conventional commit](https://www.conventionalcommits.org/) messages using
[git-cliff](https://git-cliff.org).

- **Major** (`vX.0.0`): Reserved for intentional breaking changes. Major
  releases should be rare and well-documented with an upgrade guide. Examples:
  - Removing or renaming existing managed resources
  - Removing or renaming CRD fields
  - Changing field types in incompatible ways
  - Changing provider configuration in incompatible ways
  - Any change that requires users to rewrite their compositions, claims, or
    manually migrate managed resources

- **Minor** (`vX.Y.0`): New features and improvements that are
  backwards-compatible. **This is the most common release type.** Examples:
  - Adding new managed resources
  - Adding new optional fields to existing resources
  - Upstream Terraform provider updates (these bring new/updated resources)
  - Refactoring internals
  - Performance improvements

- **Patch** (`vX.Y.Z`): Backwards-compatible bug fixes and dependency updates.
  These are typically follow-up releases after a minor release. Examples:
  - Fixing incorrect API calls or state handling
  - Fixing controller reconciliation logic
  - Go dependency updates (security patches, bug fixes)
  - Other small, targeted fixes

### Bump rules

The version bump is computed by git-cliff from the conventional commit messages
since the last tag. The rules are configured in [`cliff.toml`](../cliff.toml)
under the `[bump]` section.

| Commit type(s) since last tag                             | Version bump          |
|-----------------------------------------------------------|-----------------------|
| Only `fix` (including `fix(deps)`)                        | **patch** (vX.Y.Z)   |
| `feat`, `refactor`, or `perf` (with or without others)    | **minor** (vX.Y.0)   |
| `BREAKING CHANGE` footer or `!` suffix on any commit type | **major** (vX.0.0)   |
| Only `ci`, `docs`, `test`, `build`, `style`, `chore`      | No bump (nothing to release) |

### Commit types and release impact

| Commit type | Affects binary? | Triggers release? | Bump    |
|-------------|-----------------|-------------------|---------|
| `feat`      | yes             | yes               | minor   |
| `fix`       | yes             | yes               | patch   |
| `refactor`  | yes             | yes               | minor   |
| `perf`      | yes             | yes               | minor   |
| `docs`      | no              | no                | —       |
| `ci`        | no              | no                | —       |
| `test`      | no              | no                | —       |
| `build`     | no              | no                | —       |
| `style`     | no              | no                | —       |
| `chore`     | no              | no                | —       |

### Dependency updates

Terraform provider updates (`github.com/grafana/terraform-provider-grafana`)
are managed by Renovate and use `feat(deps)` as their commit type, because
they bring new and updated upstream resources. These trigger a **minor**
release.

Other Go module dependency updates use `fix(deps)`, because changes to
`go.mod`/`go.sum` affect the compiled provider binary. These trigger a
**patch** release.

GitHub Actions and other non-Go dependency updates use `chore(deps)` and do
**not** trigger a release, since they have no effect on the shipped binary.

This behavior is configured in [`renovate.json5`](../renovate.json5).

## Creating a release

### Prerequisites

- [git-cliff](https://git-cliff.org/docs/installation) installed locally
- Push access to the repository

### Steps

1. Switch to the `main` branch:

   ```sh
   git checkout main
   ```

2. Run the release target:

   ```sh
   make release
   ```

   The target will:
   - Verify you are on `main` (refuses to run from other branches)
   - Fetch from origin and prompt to pull if your local branch is behind
   - Compute the next version from conventional commits using git-cliff

   To override the computed version:

   ```sh
   RELEASE_VERSION=v3.0.0 make release
   ```

3. The tag push triggers the [`release`](../.github/workflows/release.yaml)
   GitHub Actions workflow (see [CI pipeline](#ci-pipeline) below).

4. Review the release on the
   [Releases page](https://github.com/grafana/crossplane-provider-grafana/releases).

### What `make release` does

`make release` runs [`scripts/release.sh`](../scripts/release.sh), which
performs the following checks and actions:

```
scripts/release.sh
  │
  ├─ on main branch?       → error if not
  ├─ fetch origin/main
  │   └─ local behind?     → prompt to pull (--ff-only)
  │
  ├─ RELEASE_VERSION set?  → use it
  ├─ git-cliff installed?  → compute version via `git cliff --bumped-version`
  │   └─ version == latest tag?  → error: nothing to release
  └─ neither?              → error: install git-cliff or set RELEASE_VERSION
  │
  ├─ git tag <version>
  └─ git push origin <version>
```

## CI pipeline

The [`release.yaml`](../.github/workflows/release.yaml) workflow triggers on
any tag push matching `v*`. It runs the following steps:

1. **Generate changelog** — git-cliff produces release notes from conventional
   commits since the previous tag (`--latest --strip header`).

2. **Validate semver bump** — git-cliff computes the minimum required version
   (`--bumped-version`) and compares it against the pushed tag. If the tag is
   lower than what the commits require (e.g., tagging a patch when there are
   `feat` commits), the workflow **fails** before building. This is a safety
   net that catches incorrect manual overrides.

3. **Merge CRDs** — combines individual CRD files for the release artifacts.

4. **Create GitHub release** — creates a GitHub Release with the git-cliff
   changelog as release notes and any artifacts from `dist/`.

5. **Trigger deployment** — an Argo Workflow is triggered to deploy the new
   version to the platform-monitoring infrastructure.

### Configuration files

| File                                                        | Purpose                                    |
|-------------------------------------------------------------|--------------------------------------------|
| [`scripts/release.sh`](../scripts/release.sh)               | Local release script (branch check, pull prompt, version computation, tag + push) |
| [`cliff.toml`](../cliff.toml)                               | Changelog format, commit parsing, bump rules |
| [`.github/workflows/release.yaml`](../.github/workflows/release.yaml) | CI workflow that orchestrates the release |
| [`renovate.json5`](../renovate.json5)                       | Renovate config for dependency update commit types |
