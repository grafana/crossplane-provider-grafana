# Provider Grafana

`provider-grafana` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Grafana API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/grafana/provider-grafana):

```
up ctp provider install xpkg.upbound.io/grafana/provider-grafana:v0.24.0
```

Alternatively, you can use declarative installation:

```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-grafana
spec:
  package: xpkg.upbound.io/grafana/provider-grafana:v0.24.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/grafana/crossplane-provider-grafana).

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

Below are some issues that have been encountered and may be helpful in the future
to others.

```bash
❯ make generate
14:35:30 [ .. ] generating provider schema for grafana/grafana 2.19.1
make[1]: *** [config/schema.json] Error 1
make: *** [generate] Error 2
```

**Solution**: ensure that you do not have a `.terraformrc` defined somewhere. For example
`~/.terraformrc`:

```bash
❯ cat ~/.terraformrc
provider_installation {
    dev_overrides {
        "grafana/grafana" = "/Users/joeyorlando/coding/grafana/terraform-provider-grafana"
    }
}
```

Additionally, you can check the `terraform` logs via:

```bash
❯ cat ./.work/terraform/terraform-logs.txt
...
```

Lastly, make sure that you have the following defined in your `.bashrc` (or `.zshrc`):

```bash
export PATH="$PATH:$HOME/go/bin"
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/grafana/crossplane-provider-grafana/issues).
