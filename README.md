# Terrajet Grafana Provider

`crossplane-provider-grafana` is a [Crossplane](https://crossplane.io/) provider that
is built using [Terrajet](https://github.com/crossplane/terrajet) code
generation tools and exposes XRM-conformant managed resources for the 
Grafana API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://github.com/grafana/crossplane-provider-grafana/releases):

```
kubectl crossplane install provider grafana/crossplane-provider-grafana:v0.1.0
```

You can see the API reference [here](https://doc.crds.dev/github.com/grafana/crossplane-provider-grafana).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go
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

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/grafana/crossplane-provider-grafana/issues).

## Contact

Please use the following to reach members of the community:

* Slack: Join our [slack channel](https://slack.crossplane.io)
* Forums:
  [crossplane-dev](https://groups.google.com/forum/#!forum/crossplane-dev)
* Twitter: [@crossplane_io](https://twitter.com/crossplane_io)
* Email: [info@crossplane.io](mailto:info@crossplane.io)

## Governance and Owners

crossplane-provider-grafana is run according to the same
[Governance](https://github.com/crossplane/crossplane/blob/master/GOVERNANCE.md)
and [Ownership](https://github.com/crossplane/crossplane/blob/master/OWNERS.md)
structure as the core Crossplane project.

## Code of Conduct

crossplane-provider-grafana adheres to the same [Code of
Conduct](https://github.com/crossplane/crossplane/blob/master/CODE_OF_CONDUCT.md)
as the core Crossplane project.

## Licensing

crossplane-provider-grafana is under the Apache 2.0 license.
