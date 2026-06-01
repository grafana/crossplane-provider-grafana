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

## Signature Verification

Published packages are cryptographically signed using [cosign](https://docs.sigstore.dev/cosign/system_config/installation/) keyless signing with GitHub Actions OIDC. This lets you verify that a package was built by this repository's CI pipeline and hasn't been tampered with.

### Verify a package signature

```bash
cosign verify \
  ghcr.io/grafana/provider-grafana:v2.10.0 \
  --certificate-identity-regexp 'https://github.com/grafana/crossplane-provider-grafana/.github/workflows/ci.*' \
  --certificate-oidc-issuer 'https://token.actions.githubusercontent.com'
```

The same works for packages on the Upbound Marketplace:

```bash
cosign verify \
  xpkg.upbound.io/grafana/provider-grafana:v2.10.0 \
  --certificate-identity-regexp 'https://github.com/grafana/crossplane-provider-grafana/.github/workflows/ci.*' \
  --certificate-oidc-issuer 'https://token.actions.githubusercontent.com'
```

### Automatic verification in Crossplane

Crossplane 1.18+ supports automatic signature verification on package install via [ImageConfig](https://docs.crossplane.io/latest/packages/image-configs/#configuring-signature-verification):

```yaml
apiVersion: pkg.crossplane.io/v1beta1
kind: ImageConfig
metadata:
  name: verify-provider-grafana
spec:
  matchImages:
    - prefix: "ghcr.io/grafana/provider-grafana:"
    - prefix: "xpkg.upbound.io/grafana/provider-grafana:"
  verification:
    provider: Cosign
    cosign:
      authorities:
        - name: grafana-ci
          keyless:
            identities:
              - issuer: https://token.actions.githubusercontent.com
                subjectRegExp: https://github.com/grafana/crossplane-provider-grafana/.github/workflows/ci.*
```

## Contributing

See [docs/contributing.md](docs/contributing.md) for development setup and contribution guidelines.

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please open an [issue](https://github.com/grafana/crossplane-provider-grafana/issues).

Note that most resource logic lives in the upstream [Terraform provider for Grafana](https://github.com/grafana/terraform-provider-grafana). If the issue is related to resource behavior rather than Crossplane-specific functionality, it may need to be reported there instead.
