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


### 4. Apply Kubernetes Manifests

#### Install Manifest 0

Apply the manifest for your application setup:
```bash
kubectl apply -f P4-Ch.01/istio/manifest/manifest-0/deployment.yaml
kubectl apply -f P4-Ch.01/istio/manifest/manifest-0/deployment-istio.yaml
```

and see deployment echov2 has no istio-proxy injected

#### Install Manifest 1

Apply the manifest for your application setup:
```bash
kubectl apply -f P4-Ch.01/istio/manifest/manifest-1
```

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


#### Apply Certificate Manifest

To apply your HTTPS certificate:
```bash
kubectl apply -f manifest-2/7-certificate.yaml
```



### 7. Verify DNS Setup

To check if your domain is correctly pointing to the cluster IP:
```bash
dig +short k8s.kowlon.my.id
```

### 8. Apply Manifest 2

Apply the second set of manifests:
```bash
kubectl apply -f manifest/manifest-2
```
