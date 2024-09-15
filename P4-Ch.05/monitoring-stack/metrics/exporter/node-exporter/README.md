1. Install node-exporter
```
helm repo add prometheus-community https://victoriametrics.github.io/helm-charts/
helm repo update
helm show values prometheus-community/prometheus-node-exporter > values.yaml
helm upgrade --install node-exporter prometheus-community/prometheus-node-exporter -f values.yaml
```