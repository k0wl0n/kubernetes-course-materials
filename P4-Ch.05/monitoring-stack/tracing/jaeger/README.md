1. Install Jaeger
```
helm repo add jaegertracing https://jaegertracing.github.io/helm-charts
helm repo update
helm show values jaegertracing/jaeger > values.yaml
helm upgrade --install jaeger jaegertracing/jaeger -f values.yaml

option installation
1. jaeger all in one true using memory on storage, no need es/cassandra
2. jaeger all in one false must be define backend storage such as es/cassandra

edit 
1. set false jaeger agent line 377
2. if you want using jaeger all in one, set false for all provisionDataStore

apps
hit jaeger-collector port 4318
cd ../example-apps/kubernetes/base/hotrod/deployment.yaml
```
    spec:
      containers:
      - env:
          - name: OTEL_EXPORTER_OTLP_ENDPOINT
            value: "http://jaeger-collector.monitoring.svc.cluster.local:4318"
``` 