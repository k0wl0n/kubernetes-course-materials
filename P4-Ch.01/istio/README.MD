# Kubernetes Infrastructure Setup and Monitoring

This guide provides step-by-step instructions for setting up Kubernetes infrastructure, including Helm, Istio, certificate management, and monitoring using Prometheus and Grafana.

### 1. Initialize and Apply Terraform

go to folder P4-Ch.01/istio/istio-terraform:

```bash
terraform init
terraform apply
```

### 2. Check Istio Custom Resource Definitions (CRD)

Verify if the Istio CRDs have been installed in your cluster:
```bash
kubectl get crds | grep 'istio'
```

## 3. Install Certificate Manager & Gateway

go to folder cert-manager
```bash
kubectl apply -f cert-manager.yaml
```

# register cert manager to wildcard *.apps.kowlon.my.id
echo1.apps.kowlon.my.id
echo2.apps.kowlon.my.id
echo10.apps.kowlon.my.id

create token cloudflare api
reffer to this docs P4-Ch.01/istio/gateway/cloudflare-token.md
after you have api token
fill the token api to this file P4-Ch.01/istio/gateway/1-1-dns-secret.yaml
* NOTE that api-key must be encode to base64 first
* THE Domain name is different every user, you must be change 
* FOR example im use domain `kowlon.my.id` for this couse you should change the domain based on you domain

### Apply gateway
```bash
kubectl apply -f 2-gateway.yaml
```

### Apply willcard certificate
go to folder P4-Ch.01/istio/gateway

```bash
kubectl apply -f 1-1-dns-secret.yaml
kubectl apply -f 1-2-dns01.yaml
kubectl apply -f 1-3-certificate.yaml
```

### Add Gateway Load Balancer to Domain

Update your DNS records to point to the gateway's load balancer IP. 

#### Debugging Certificate Issues

If you face any issues with the certificates, you can debug using the following commands:
```bash
kubectl get certificate -n istio-ingress
kubectl describe certificate -n istio-ingress
kubectl describe certificaterequest -n istio-ingress
kubectl describe challenge -n istio-ingress
kubectl get challenges -A
kubectl get svc -n istio-ingress
kubectl get secret -n istio-ingress
```

### Enable TLS 

```
tls:
    credentialName: apps-kowlon-my-id-wc-crt  # Secret for wildcard TLS certificate
    mode: SIMPLE
```

and apply again
```bash
kubectl apply -f 2-gateway.yaml
```

# register istio ingress from LB to cloudflare
*.apps.kowlon.my.id

 kubectl get svc -n istio-ingress  
 get external-ip
 contoh: ad6c46c39632244459cd887585743a0d-954344150.ap-southeast-1.elb.amazonaws.com

 add to DNS in yuur domain panel on cloudflare

 add TYPE CNAME
 add name *.apps.kowlon.my.id
 add external-ip ad6c46c39632244459cd887585743a0d-954344150.ap-southeast-1.elb.amazonaws.com

 and save


### 4. Apply Kubernetes Manifests

#### Install Manifest 0

Apply the manifest for your application setup:
```bash
kubectl apply -f P4-Ch.01/istio/manifest/manifest-0/deployment.yaml
kubectl apply -f P4-Ch.01/istio/manifest/manifest-0/deployment-istio.yaml
```

kubectl get pods -n default

kubectl get pod echov2-d69b85f64-czv8h -o jsonpath='{.spec.containers[*].name}' -n default
kubectl get pod echov3-54966b7b57-hvk2f -o jsonpath='{.spec.containers[*].name}' -n default

and see deployment echov3 has no istio-proxy injected

#### Install Manifest 1

Apply the manifest for your application setup:
```bash
kubectl apply -f P4-Ch.01/istio/manifest/manifest-1
```
### Flow ingress istio

domain name -> Loadbalancer ( gateway) -> virtual service -> service -> deployment -> pod

#### Test Access from a Client Pod

You can test access to the application from a client pod:
```bash
kubectl exec -ti -n backend client -- sh
while true; do curl echo1.production:8080/ && echo "" && sleep 1; done
```

To check the pods in the Istio Ingress namespace:
```bash
kubectl get pods -n istio-ingress --show-labels
```