1. Install node-exporter
```
helm repo add prometheus-community https://victoriametrics.github.io/helm-charts/
helm repo update
helm upgrade --install node-exporter prometheus-community/prometheus-node-exporter -f values.yaml
```