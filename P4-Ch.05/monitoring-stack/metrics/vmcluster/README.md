1. Install vmcluster
```
helm repo add vm https://victoriametrics.github.io/helm-charts/
helm repo update
helm show values vm/victoria-metrics-cluster > values.yaml
helm upgrade --install vmcluster vm/victoria-metrics-cluster -f values.yaml
```