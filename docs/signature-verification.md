# Signature Verification

Published packages are cryptographically signed and include an SBOM (software bill of materials) attestation. Both are produced using [cosign](https://docs.sigstore.dev/cosign/system_config/installation/) keyless signing with GitHub Actions OIDC. This lets you verify that a package was built by this repository's CI pipeline and hasn't been tampered with.

## How it works

After each package publish, the CI pipeline:

1. **Signs** the xpkg artifact using cosign's keyless mode
2. **Generates an SBOM** in SPDX format using [syft](https://github.com/anchore/syft)
3. **Attaches the SBOM** as a signed in-toto attestation using `cosign attest`

Keyless signing uses the GitHub Actions OIDC identity token to obtain a short-lived certificate from Sigstore's Fulcio CA, and records the signature in Sigstore's Rekor transparency log. No long-lived signing keys are involved.

Packages are signed on both registries:

- `ghcr.io/grafana/provider-grafana`
- `xpkg.upbound.io/grafana/provider-grafana`

## Verify a package signature

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

## Verify the SBOM attestation

```bash
cosign verify-attestation \
  ghcr.io/grafana/provider-grafana:v2.10.0 \
  --certificate-identity-regexp 'https://github.com/grafana/crossplane-provider-grafana/.github/workflows/ci.*' \
  --certificate-oidc-issuer 'https://token.actions.githubusercontent.com' \
  --type spdxjson
```

To extract the SBOM contents:

```bash
cosign verify-attestation \
  ghcr.io/grafana/provider-grafana:v2.10.0 \
  --certificate-identity-regexp 'https://github.com/grafana/crossplane-provider-grafana/.github/workflows/ci.*' \
  --certificate-oidc-issuer 'https://token.actions.githubusercontent.com' \
  --type spdxjson | jq -r '.payload' | base64 -d | jq '.predicate'
```

## Automatic verification in Crossplane

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

If signature verification is enabled, Crossplane sets a `SignatureVerificationComplete` status condition on the `ProviderRevision` resource once verification succeeds.
