## membuat deployment
```
kubectl create deployment hello-minikube --image=kicbase/echo-server:1.0
```

## expose deployment ke service
```
kubectl expose deployment hello-minikube --type=NodePort --port=8080
minikube service hello-minikube --url=true
```

## akses service url
```
curl service_url
kubectl scale deployment/hello-minikube --replicas=3
```

## loop akses 10x
```
for i in {1..10}; do curl -sSL service_url && echo; done
```