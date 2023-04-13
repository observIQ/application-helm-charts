
```bash
k apply -f crds/
k apply -f operator.yaml
k apply -f custom-scrape-config.yaml
k apply -f secret.yaml
k apply -f agent.yaml
k apply -f metrics-instance.yaml
k apply -f podmonitor.yaml
```


Requires [operator](https://grafana.com/docs/agent/latest/operator/helm-getting-started/#install-the-agent-operator-helm-chart)
