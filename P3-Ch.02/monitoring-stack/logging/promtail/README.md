1. Install promtail
```
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
helm upgrade --install promtail grafana/promtail -f values.yaml
```

line 421 add clients loki gateway
```
http://loki-distributed-gateway.{{namespace}}.svc.cluster.local/loki/api/v1/push
```