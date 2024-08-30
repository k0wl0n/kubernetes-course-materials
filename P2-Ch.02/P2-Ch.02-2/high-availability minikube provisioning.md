## melihat cluster minikube list

```
minikube profile list
```

## membuat cluster high availability
```
minikube start --ha --driver=docker --container-runtime=containerd --profile ha-cluster
```

## List your HA cluster nodes:

```
kubectl get nodes -owide
minikube profile list
minikube status -p ha-cluster
minikube profile ha-cluster
```

## validate cluster

```
minikube ssh -p ha-cluster -- 'sudo /var/lib/minikube/binaries/v1.30.0/kubectl  --kubeconfig=/var/lib/minikube/kubeconfig logs -n kube-system pod/kube-vip-ha-cluster'
minikube ssh -p ha-cluster -- 'sudo /var/lib/minikube/binaries/v1.30.0/kubectl  --kubeconfig=/var/lib/minikube/kubeconfig logs -n kube-system pod/kube-vip-ha-cluster-m02'
minikube ssh -p ha-cluster -- 'sudo /var/lib/minikube/binaries/v1.30.0/kubectl  --kubeconfig=/var/lib/minikube/kubeconfig logs -n kube-system pod/kube-vip-ha-cluster-m03'
```

## get etcd status

```
minikube ssh -p ha-cluster -- 'sudo /var/lib/minikube/binaries/v1.30.0/kubectl --kubeconfig=/var/lib/minikube/kubeconfig exec -ti pod/etcd-ha-cluster -n kube-system -- /bin/sh -c "ETCDCTL_API=3 etcdctl member list --write-out=table --cacert=/var/lib/minikube/certs/etcd/ca.crt --cert=/var/lib/minikube/certs/etcd/server.crt --key=/var/lib/minikube/certs/etcd/server.key"'

```

## remove 1 control-plane
```
minikube node delete m02 -p ha-cluster
kubectl get nodes -owide
kubectl get componentstatuses
minikube profile list
```

## adding back control plane
```
minikube node add --control-plane -p ha-cluster
kubectl get nodes -owide
minikube profile list
```

## deploy some apps
```
kubectl apply -f deployment.yaml
for i in {1..10}; do curl -sSL service_url && echo; done
```



