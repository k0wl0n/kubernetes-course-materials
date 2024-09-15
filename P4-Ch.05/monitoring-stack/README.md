# Step by step deployment 

* install logging/loki
* install logging/promtail
* install metrics/node_exporter
* install metrics/ksm
* install metrics/vmagent
* install metrics/vmcluster
* Install visualize/grafana
  * setup grafana datasource loki and vmselect
  * add example dashboard metrics grafana
  * check logging on grafana explore menu
* setup jaeger
* install example apps with otel SDK
* check jaeger UI
* install alert/vmalert
  * deploy example oom service for alert oom 