## single node
```
minikube start --driver=docker --container-runtime=containerd --profile single
minikube profile single
```

## menambah node pertama
```
minikube node add --profile=single
```

## menambah node kedua
```
minikube node add --profile=single
```

## mengganti nama node pertama
```
kubectl label nodes single-m02 node-name=worker-node-1
```

## mengganti nama node kedua
```
kubectl label nodes single-m03 node-name=worker-node-2
```


## membuat deployment
```
kubectl create deployment hello-minikube --image=kicbase/echo-server:1.0   
kubectl scale deployment/hello-minikube --replicas=1 
kubectl scale deployment/hello-minikube --replicas=3
```