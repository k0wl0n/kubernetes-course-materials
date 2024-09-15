1. Install loki-distributed 
```
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
helm show values grafana/loki-distributed > values.yaml
helm upgrade --install loki-distributed grafana/loki-distributed -f values.yaml

enable cache if any: 
```
memcachedChunks line 1761
memcachedFrontend line 1830
memcachedIndexQueries line 1895
memcachedIndexWrites line 1960
```