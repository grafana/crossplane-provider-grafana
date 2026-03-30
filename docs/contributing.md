# Contributing

## Update resources

Steps to update resources from the latest Terraform provider version:

> **Note:** Renovate will generally update the Terraform provider dependency
> automatically, so there is no need to do this manually in normal circumstances.

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

```bash
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

   ```bash
   kubectl config use-context <name-of-your-local-k8s-context>
   kubectl apply -f ./package/crossplane.yaml
   kubectl apply -f ./package/crds
   ```

### Possible issues when running locally

Below are some issues that have been encountered and may be helpful in the future to others.

```bash
make generate
14:35:30 [ .. ] generating provider schema for grafana/grafana 2.19.1
make[1]: *** [config/schema.json] Error 1
make: *** [generate] Error 2
```

**Solution**: ensure that you do not have a `.terraformrc` defined somewhere. For example `~/.terraformrc`:

```bash
cat ~/.terraformrc
provider_installation {
    dev_overrides {
        "grafana/grafana" = "/Users/joeyorlando/coding/grafana/terraform-provider-grafana"
    }
}
```

**Solution 2**: delete generated `.cache`, `.work` and `_output` folders and try again.

Additionally, you can check the `terraform` logs via:

```bash
cat ./.work/terraform/terraform-logs.txt
...
```

Lastly, make sure that you have the following defined in your `.bashrc` (or `.zshrc`):

```bash
export PATH="$PATH:$HOME/go/bin"
```
