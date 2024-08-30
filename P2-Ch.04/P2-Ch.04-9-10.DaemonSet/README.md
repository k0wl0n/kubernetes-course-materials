### README for `P2-Ch.04-9-10.DaemonSet`

This guide provides steps to set up and manage a Kubernetes cluster using Minikube and configure DaemonSets, along with managing nodes and labels within the `P2-Ch.04-9-10.DaemonSet` folder.

### Steps

#### 1. Start Minikube Cluster

Start a Minikube cluster with a specified driver and container runtime:

```bash
minikube start --driver=docker --container-runtime=containerd --profile single
```

#### 2. Set Profile to Cluster

Set the active Minikube profile to the created cluster:

```bash
minikube profile single
```

#### 3. Check Cluster Info

Display information about the Kubernetes cluster:

```bash
kubectl cluster-info
```

#### 4. List All Resources

List all Kubernetes resources in the current namespace:

```bash
kubectl get all
```

#### 5. Delete All Resources

Delete all pods, services, and deployments in the current namespace:

```bash
kubectl delete pods,services,deployments --all
```

#### 6. List Nodes

List the nodes in the Minikube cluster:

```bash
minikube node list
```

or using Kubernetes:

```bash
kubectl get nodes
kubectl get nodes --show-labels
```

#### 7. Adding Node

Add a new node to the Minikube cluster:

```bash
minikube node add --profile=single
```

#### 8. Adding Labels

Label nodes to specify their roles or types:

```bash
kubectl label nodes single-m02 node_type=standard
kubectl label nodes single-m03 node_type=database
```

This setup is used for managing DaemonSets and other resources in a Kubernetes cluster, helping to ensure that specific pods are run on all or selected nodes as required by the configuration.