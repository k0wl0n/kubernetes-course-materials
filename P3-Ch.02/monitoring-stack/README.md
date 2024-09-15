# Step by step deployment 

* install logging/loki
* install logging/promtail
* install metrics/exporter/node_exporter
* install metrics/exporter/kube-state-metrics
* install metrics/vmcluster
* install metrics/vmagent
* Install visualize/grafana
  * setup grafana datasource loki and vmselect
  * add example dashboard metrics grafana
  * check logging on grafana explore menu
* setup jaeger tracing/jeager
* install example apps with otel SDK
  * tracing/jaeger/apps/hot_rod
* check jaeger UI

* install alert/vmalert