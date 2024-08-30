### README for `P2-Ch.04-23.Pengelolaan Scheduling Pod`

This guide provides steps to manage Kubernetes pod scheduling through Minikube cluster setup, node addition, labeling, and tainting. It is focused on configuring nodes to control pod placement based on specific criteria.

### Steps

#### 1. Start Minikube Cluster

Initiate a Minikube cluster using Docker as the driver and Containerd as the container runtime:

```bash
minikube start --driver=docker --container-runtime=containerd --profile single
```

#### 2. Add and Label Nodes

Add new nodes to the Minikube cluster and assign specific labels to identify their disk types:

- **Add Node 1 and Label as SSD:**

  ```bash
  minikube node add --profile=single
  kubectl label nodes single-m02 disktype=ssd
  ```

- **Add Node 2 and Label as NVMe:**

  ```bash
  minikube node add --profile=single
  kubectl label nodes single-m03 disktype=nvme
  ```

- **Add Node 3 and Label as SSD:**

  ```bash
  minikube node add --profile=single
  kubectl label nodes single-m04 disktype=ssd
  ```

#### 3. Verify Node Labels

Ensure that all nodes have the correct labels by listing nodes with their assigned labels:

```bash
kubectl get nodes --show-labels
```

This command helps verify that each node has been correctly labeled for disk type identification.

#### 4. Apply Taints

Taint a node to restrict pod scheduling to nodes that match specific tolerations, ensuring that only pods with the correct toleration can be scheduled on tainted nodes:

```bash
kubectl taint nodes single key=master:NoSchedule
```

Applying taints helps in managing node utilization effectively by preventing certain pods from being scheduled on specific nodes unless they have the necessary tolerations.

### Summary

- **Start Minikube Cluster**: Initiates the Kubernetes environment using Minikube.
- **Add and Label Nodes**: Adds nodes to the cluster and assigns labels for resource management.
- **Verify Node Labels**: Confirms the correct labeling of nodes to control pod placement.
- **Apply Taints**: Uses taints to control where pods can be scheduled, providing an additional layer of scheduling control.

This configuration is useful for managing and optimizing resource allocation, ensuring that specific workloads are scheduled on nodes with appropriate capabilities or restrictions within your Kubernetes cluster.