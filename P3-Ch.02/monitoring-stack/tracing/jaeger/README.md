1. Install Jaeger
```
helm repo add jaegertracing https://jaegertracing.github.io/helm-charts
helm repo update
helm upgrade --install jaeger jaegertracing/jaeger -f values.yaml

option installation
1. jaeger all in one true using memory on storage, no need es/cassandra
2. jaeger all in one false must be define backend storage such as es/cassandra

edit 
1. set false jaeger agent line 377
2. if you want using jaeger all in one, set false for all provisionDataStore

apps
hit jaeger-collector port 4318

kubectl --namespace monitoring port-forward svc/jaeger-query 16686

http://localhost:16686

kubectl port-forward -n default service/example-hotrod 8888:frontend
http://localhost:8888