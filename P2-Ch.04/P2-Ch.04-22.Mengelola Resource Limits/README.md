### README for `P2-Ch.04-22.Mengelola Resource Limits`

This guide provides steps and commands for monitoring and managing resource limits in Kubernetes, including how to watch pods, monitor resource usage, fetch logs, and debug resource quotas.

### Overview

Resource limits in Kubernetes help ensure that applications run efficiently and do not consume more resources than necessary. This guide includes commands for real-time monitoring, checking resource usage, and managing resource quotas.

### Commands

#### 1. Watch Pods

Continuously monitor the status of all pods in real-time:

```bash
watch -t kubectl get pods
```

This command provides a live view of the pod status, which is useful for monitoring deployments and checking for any issues.

#### 2. Monitor Resource Usage

Check the current CPU and memory usage of all pods within a specific namespace (e.g., `quota`):

```bash
kubectl top pods -n quota
```

This command helps identify resource-intensive pods and can be used to ensure that pods are running within their allocated resource limits.

#### 3. Get Logs

Retrieve detailed information, including configuration and logs, for a specific pod (e.g., `memory-demo`):

```bash
kubectl get pod memory-demo --output=yaml
```

This command outputs the entire YAML configuration of the pod, which can be useful for debugging and understanding the pod's setup.

#### 4. Debug Resource Quota

- List events related to resource quotas in a specific namespace (e.g., `quota`), sorted by the time they were created:

  ```bash
  kubectl get events --namespace=quota --sort-by=.metadata.creationTimestamp
  ```

  This command provides a history of events, which can help identify issues related to resource allocation and usage.

- Display the current resource quotas set for a specific namespace:

  ```bash
  kubectl get resourcequota --namespace=quota
  ```

  This command shows the defined resource quotas, helping you understand the constraints placed on resource usage within a namespace.

- Describe the details of a specific resource quota (e.g., `rs-quota`) to understand its configuration and current usage:

  ```bash
  kubectl describe resourcequota rs-quota --namespace=quota
  ```

  This command provides detailed information about a particular resource quota, including its limits and current usage.

### Summary

- **Watch Pods**: Use `watch -t kubectl get pods` to monitor pod status in real-time.
- **Monitor Resource Usage**: Use `kubectl top pods -n <namespace>` to check CPU and memory usage.
- **Get Logs**: Use `kubectl get pod <pod-name> --output=yaml` to fetch detailed pod configurations and logs.
- **Debug Resource Quota**: Use `kubectl get events`, `kubectl get resourcequota`, and `kubectl describe resourcequota` to troubleshoot and manage resource limits.

This setup enables effective monitoring and debugging of Kubernetes resources, helping maintain application stability and performance by managing resource limits and quotas efficiently.