### README for `P2-Ch.04-2-5.Pods`

This guide provides steps to create, describe, and access Kubernetes Pods within the `P2-Ch.04-2-5.Pods` folder.

### Steps

#### 1. Create Pod

Deploy a Pod using a YAML configuration file:

```bash
kubectl create -f filepod.yaml
```

#### 2. Describe Pod

Get detailed information about the running Pod:

```bash
kubectl get pod
kubectl get pod -o wide
kubectl describe pod <pod-name>
```

Replace `<pod-name>` with the actual name of the Pod you want to describe.

#### 3. Access Pod

Use port forwarding to access the Pod on a specific port:

```bash
kubectl port-forward <pod-name> <local-port>:<pod-port>
kubectl port-forward <pod-name> 8888:8080
```

Replace `<pod-name>` with the actual name of the Pod, `<local-port>` with the desired local port, and `<pod-port>` with the Pod's port.

This setup provides basic Pod operations, including creation, inspection, and accessing Pods via port forwarding.