[![CI](https://github.com/observIQ/charts/actions/workflows/ci.yaml/badge.svg)](https://github.com/observIQ/charts/actions/workflows/ci.yaml)
[![Release Container Images](https://github.com/observIQ/charts/actions/workflows/container.yaml/badge.svg)](https://github.com/observIQ/charts/actions/workflows/container.yaml)

# Charts

This repository contains Helm charts suitable for developing integrations. This repository
is not intended for production use.

## Usage

[Helm](https://helm.sh) must be installed to use the charts.  Please refer to
Helm's [documentation](https://helm.sh/docs) to get started.


### Add Repo

Add the observIQ Helm repository.

```bash
helm repo add observiq https://observiq.github.io/charts
```

### Update Repo

If you had already added this repo earlier, run `helm repo update` to retrieve
the latest versions of the packages.  You can then run `helm search repo
observiq` to see the charts.

### Updating container images

1. Open PR for changes to container image
2. Request a review. Once approved, merge PR
3. Watch main branch CI. It will build and push the image. The CI output will usually have the image tag.
4. Open another PR changing the image used by the helm chart. Be sure to update the chart version.
5. Request a review. Once approved, merge PR
6. CI will release the helm chart
7. Update your local helm repo to get the new chart version, which will have your new image.

### Applications

#### Clickhouse

To install the clickhouse chart:

```bash
helm install clickhouse observiq/clickhouse
```

To uninstall the chart:

```bash
helm delete clickhouse
```

#### Discourse(R)

To install the Discourse(R) chart:

```bash
helm install discourse observiq/discourse
```

To uninstall the chart:

```bash
helm delete discourse
```

#### Wildfly

To install the Wildfly chart:

```bash
helm install wildfly observiq/wildfly
```

To uninstall the chart:

```bash
helm delete wildfly
```

#### MSSQL

To install the Wildfly chart:

```bash
helm install mssql observiq/mssql
```

To uninstall the chart:

```bash
helm delete mssql
```

#### OracleDB Express Edition 21c (21.3.0)

To install the OracleDB XE chart:

```bash
helm install oracledbxe observiq/oracledbxe
```

To uninstall the chart:

```bash
helm delete oracledbxe
```
