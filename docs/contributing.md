# Contributing

## Submitting changes

1. **Fork the repo** and create a branch from `main`.
2. **Make your changes.** Follow the existing code style and patterns.
3. **Run the linter** before opening a PR: `make lint`.
4. **Open a pull request** against `main` — see [PR title format](#pr-title-format) below.

## PR title format

This repository uses **squash merges**, so all commits in a PR are combined
into a single commit when merged. The **PR title becomes the commit message**,
so it must follow the
[Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) format:

```
<type>(<scope>): <subject>
```

A CI check validates the PR title and will block merging if it doesn't conform.

### Types

| Type | Purpose |
|------|---------|
| `feat` | New feature |
| `fix` | Bug fix |
| `docs` | Documentation only |
| `style` | Formatting, no logic change |
| `refactor` | Code change that neither fixes a bug nor adds a feature |
| `perf` | Performance improvement |
| `test` | Adding or updating tests |
| `build` | Build system or external dependencies |
| `ci` | CI/CD configuration |
| `chore` | Maintenance tasks, dependency updates, housekeeping |

### Scope

Scope is optional but recommended. Use the affected resource name when
applicable. Scopes must be lowercase with no spaces.

### Subject

The subject should be lowercase, use imperative mood ("add" not "added"), and
not end with a period.

### Breaking changes

Breaking changes should be avoided whenever possible. Crossplane users depend
on stable CRD schemas, and breaking changes force manual migration of managed
resources and compositions.

When a breaking change is unavoidable, append `!` after the type (and scope, if
present):

```
feat(folder)!: rename uid to folder_uid
```

Breaking changes must be explained in the PR description body. When squash
merging, expand the commit message body to include a `BREAKING CHANGE:` footer
describing what changed and how users should migrate. For example:

```
feat(folder)!: rename uid to folder_uid

BREAKING CHANGE: The `uid` field on the Folder managed resource has been
renamed to `folder_uid`. Update your compositions and claims accordingly.
```

### Examples

- `feat(dashboard): add uid attribute`
- `fix(folder): handle missing parent on import`
- `refactor(team): simplify group mapping logic`
- `feat(folder)!: rename uid to folder_uid`
- `chore: update Go dependencies`
- `docs: update contributing guide`

## Update resources

Resources are regenerated from the upstream Terraform provider whenever its
version changes.

### Automated updates

The [update-terraform-provider](../.github/workflows/update-terraform-provider.yaml)
workflow handles updates automatically. It is triggered when a new
`terraform-provider-grafana` release is published (via `repository_dispatch`)
and can also be run manually from the Actions tab with a version input. The
workflow:

1. Bumps the provider version and commits the `go.mod`/`go.sum` change.
2. Runs `go mod tidy`.
3. Runs `make submodules && make generate`.
4. Commits the regenerated output.
5. Opens a pull request on the
   `automated/update-terraform-provider-<version>` branch.

The PR is **always** created, even when steps 2-4 fail, so that a human can
fix things up on top of the automated commits.

#### When manual intervention is required

If `go mod tidy` or `make generate` fails, the PR is opened **as a draft** with
a "⚠️ Manual intervention required" section describing which step failed and a
link to the workflow run. Look for the open draft PR on the
`automated/update-terraform-provider-<version>` branch.

To resolve it, check out the branch, fix the issue, and push additional
commits, then mark the PR ready for review:

```console
git fetch origin
git checkout automated/update-terraform-provider-<version>
# fix the issue (see manual steps below), then:
make submodules && make generate
git commit -am "fix generated resources"
git push
```

The most common cause is a new resource category added upstream, which makes
`make generate` panic with a message telling you to add an entry to
`categoryConfig` in [groups.go](../config/groups.go).

If the workflow fails *before* a PR can be created (e.g. dependency resolution
or the branch push fails), it opens a GitHub issue labeled
`automated-update-failure` instead.

### Manual updates

If you need to update resources by hand:

1. Update the Terraform provider version in the [go.mod](../go.mod) file.
2. Generate the resources:

   ```console
   make submodules && make generate
   ```

   * Resources are automatically picked up from the upstream provider's category metadata.
   * If a new category was added upstream, `make generate` will panic with a message telling you to add an entry to `categoryConfig` in [groups.go](../config/groups.go).

3. Create a PR with the result.

## Developing

### Initial setup

```console
make submodules
```

Run code-generation pipeline:

```console
make generate
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

### Installing Provider/CRDs into your local k8s cluster

1. Ensure Crossplane is installed on your local cluster (instructions [here](https://docs.crossplane.io/latest/software/install/))
2. Run the following:

   ```console
   kubectl config use-context <name-of-your-local-k8s-context>
   kubectl apply -f ./package/crossplane.yaml
   kubectl apply -f ./package/crds
   ```

### Possible issues when running locally

Below are some issues that have been encountered and may be helpful in the future to others.

```console
# make generate
14:35:30 [ .. ] generating provider schema for grafana/grafana 2.19.1
make[1]: *** [config/schema.json] Error 1
make: *** [generate] Error 2
```

**Solution**: ensure that you do not have a `.terraformrc` defined somewhere. For example `~/.terraformrc`:

```hcl
provider_installation {
    dev_overrides {
        "grafana/grafana" = "/Users/joeyorlando/coding/grafana/terraform-provider-grafana"
    }
}
```

**Solution 2**: run `make distclean` to remove generated `.cache`, `.work` and `_output` folders and try again.

Additionally, you can check the `terraform` logs via:

```console
cat ./.work/terraform/terraform-logs.txt
...
```

Lastly, make sure that you have the following defined in your `.bashrc` (or `.zshrc`):

```bash
export PATH="$PATH:$HOME/go/bin"
```
