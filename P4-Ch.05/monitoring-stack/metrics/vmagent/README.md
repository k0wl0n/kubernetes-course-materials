1. Install vmagent
```
helm repo add vm https://victoriametrics.github.io/helm-charts/
helm repo update
helm show values vm/victoria-metrics-agent > values.yaml
helm upgrade --install vmagent vm/victoria-metrics-agent -f values.yaml
```

line 76 enable remotewrite to vminsert
```
remoteWriteUrls:
  - http://vmcluster-victoria-metrics-cluster-vminsert.monitoring.svc.cluster.local:8480/insert/0/prometheus
```

line 288 create scrape config kube-state-metrics
```
- job_name: kube-state-metrics
  static_configs:
    - targets:
      - kube-state-metrics.monitoring.svc.cluster.local:8080
```

line 281
```
    external_labels: 
      cluster: staging
```