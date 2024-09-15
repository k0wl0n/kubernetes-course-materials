1. Install ElasticSearch
```
helm repo add elastic https://helm.elastic.co
helm repo update
helm show values elastic/elasticsearch > values.yaml
helm upgrade --install elastic elastic/elasticsearch -f values.yaml


1. Watch all cluster members come up.
  $ kubectl get pods --namespace=default -l app=elasticsearch-master -w
2. Retrieve elastic user's password.
  $ kubectl get secrets --namespace=default elasticsearch-master-credentials -ojsonpath='{.data.password}' | base64 -d
3. Test cluster health using Helm test.
  $ helm --namespace=default test elastic