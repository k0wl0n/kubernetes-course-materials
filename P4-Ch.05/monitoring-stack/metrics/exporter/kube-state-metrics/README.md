1. Install kube-state-metrics
```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm show values prometheus-community/kube-state-metrics > values.yaml
helm upgrade --install kube-state-metrics prometheus-community/kube-state-metrics -f values.yaml
```