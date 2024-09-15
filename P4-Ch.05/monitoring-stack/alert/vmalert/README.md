1. Install vmalert
```
helm repo add vm https://victoriametrics.github.io/helm-charts/
helm repo update
helm show values vm/victoria-metrics-alert > values.yaml
helm upgrade --install vmalert vm/victoria-metrics-alert -f values.yaml
```

line 93 add url vmselect 
```
url: "http://vmcluster-victoria-metrics-cluster-vmselect.monitoring.svc.cluster.local:8481/select/0/prometheus/" 
```

line 135 add url alertmanager 
```
url: "http://alertmanager.monitoring.svc.cluster.local:9093"
```

line 281 add alerting rules

line 355 add alertmanager config 

line 365 change with your telegram bot config