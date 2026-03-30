# Provider Grafana

`provider-grafana` is a [Crossplane](https://crossplane.io/) provider that is built using [Upjet](https://github.com/upbound/upjet) code generation tools and exposes XRM-conformant managed resources for the Grafana API.

> **This tool is experimental**
>
> The code in this repository should be considered experimental. Documentation is only
> available alongside the code. It comes without support, but we are keen to receive
> feedback on the product and suggestions on how to improve it, though we cannot commit to
> resolution of any particular issue. No SLAs are available. It is not meant to be used in
> production environments, and the risks are unknown/high.
>
> Additional information can be found in [Release life cycle for Grafana Labs](https://grafana.com/docs/release-life-cycle/).

## Getting Started

Install the provider by using the following command after changing the image tag to the [latest release](https://marketplace.upbound.io/providers/grafana/provider-grafana):

```
up ctp provider install xpkg.upbound.io/grafana/provider-grafana:v2.2.0
```

Alternatively, you can use declarative installation:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-grafana
spec:
  package: xpkg.upbound.io/grafana/provider-grafana:v2.2.0
```

You can optionally customize the provider's runtime configuration using a [DeploymentRuntimeConfig](https://docs.crossplane.io/latest/packages/providers/#runtime-configuration) if you need to modify pod settings, add arguments, or configure resource limits.

You can see the API reference [here](https://marketplace.upbound.io/providers/grafana/provider-grafana).

For information on configuring provider credentials and ProviderConfig secret fields, see the [ProviderConfig Secret Fields documentation](docs/providerconfig-secret-fields.md).

## Contributing

See [docs/contributing.md](docs/contributing.md) for development setup and contribution guidelines.

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please open an [issue](https://github.com/grafana/crossplane-provider-grafana/issues).

Note that most resource logic lives in the upstream [Terraform provider for Grafana](https://github.com/grafana/terraform-provider-grafana). If the issue is related to resource behavior rather than Crossplane-specific functionality, it may need to be reported there instead.
