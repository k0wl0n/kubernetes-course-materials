1. Install grafana
```
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
helm upgrade --install grafana grafana/grafana -f values.yaml
```

line 367, takutnya ngerestart nanti ilang datanya balik lagi ke awal
```
persistence:
  type: pvc
  enabled: true
  # storageClassName: default
  accessModes:
    - ReadWriteOnce
  size: 2Gi
```

get userpass
```
kubectl get secret --namespace monitoring grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
```

kubens monitoring

victoriametrics query http path
http://vmcluster-victoria-metrics-cluster-vmselect.monitoring.svc.cluster.local:8481/select/0/prometheus/

loki query http path
http://loki-distributed-gateway.monitoring.svc.cluster.local

import kube-state-metrics dashboard
13332


# access to grafana

kubens monitoring

export POD_NAME=$(kubectl get pods --namespace monitoring -l "app.kubernetes.io/name=grafana,app.kubernetes.io/instance=grafana" -o jsonpath="{.items[0].metadata.name}")

# get user pass
kubectl get secret --namespace monitoring grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo

kubectl --namespace monitoring port-forward grafana-68fbd8b67-wd6c9 3000

username: admin
password: qcwtp6skgQAXq4SFnNdc9APq9G7W7eQbyNdw9XZO

# import dashboard
copy dashboard.json to grafana