# charts

## Usage

[Helm](https://helm.sh) must be installed to use the charts.  Please refer to
Helm's [documentation](https://helm.sh/docs) to get started.


### Add Repo

Add the observIQ help repository.

```bash
helm repo add observiq https://observiq.github.io/charts
```

### Update Repo

If you had already added this repo earlier, run `helm repo update` to retrieve
the latest versions of the packages.  You can then run `helm search repo
observiq` to see the charts.

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
