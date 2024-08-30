## melihat cluster minikube list

```
minikube profile list
```

## membuat cluster minikube
```
minikube start --driver=docker --container-runtime=containerd --profile single
minikube profile single
minikube node add --profile=single
kubectl label nodes single-m01 node-name=worker-single-node-1
```

## membuat 2 cluster minikube
```
minikube start --driver=docker --container-runtime=containerd --profile second
minikube profile second
minikube node add --profile=second
kubectl label nodes second-m01 node-name=worker-second-node-1
```

## stop cluster minikube
```
minikube stop --profile single
```