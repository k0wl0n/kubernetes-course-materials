1. Install alertmanager
```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm show values prometheus-community/alertmanager > values.yaml
helm upgrade --install alertmanager prometheus-community/alertmanager -f values.yaml