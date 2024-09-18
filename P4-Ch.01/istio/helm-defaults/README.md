 helm repo add istio https://istio-release.storage.googleapis.com/charts
 helm repo update

 helm install my-istio-base-release -n istio-system --create-namespace istio/base --set global.istioNamespace=istio-system
 helm install my-istiod-release -n istio-system --create-namespace istio/istiod --set telemetry.enabled=true --set global.istioNamespace=istio-system
 helm install gateway -n istio-ingress --create-namespace istio/gateway

 a2ffd1611bec64a218657d7f586da69e-740223033.ap-southeast-1.elb.amazonaws.com

dig +short echo1.apps.kowlon.my.id
a2ffd1611bec64a218657d7f586da69e-740223033.ap-southeast-1.elb.amazonaws.com.
18.139.189.80
52.220.232.172
3.1.62.118