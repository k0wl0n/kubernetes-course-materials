### README for `P2-Ch.04-21.HPA`

This guide provides steps to enable and use the Metrics Server in Minikube to monitor resource usage and set up Horizontal Pod Autoscaling (HPA) based on these metrics.

### Overview

This project demonstrates how to enable the Metrics Server in Minikube and use it to gather metrics data, which can be utilized for setting up Horizontal Pod Autoscaling (HPA) to automatically scale the number of pod replicas based on CPU or memory usage.

### Steps

#### 1. Enable Metrics Server

First, check the list of available Minikube addons and enable the `metrics-server` addon:

```bash
minikube addons list
minikube addons enable metrics-server
```

#### 2. Check Node Metrics

Once the Metrics Server is enabled, you can check the metrics for nodes to see their resource usage:

```bash
kubectl top node
```

#### 3. Troubleshooting

If enabling the Metrics Server via the addon does not work, you can manually deploy the Metrics Server components by applying the following YAML file:

```bash
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
```

This command will deploy the Metrics Server components, allowing you to gather metrics for both nodes and pods.

#### 4. Set Up Horizontal Pod Autoscaler (HPA)

To set up HPA for a deployment based on CPU usage, use the following command:

```bash
kubectl autoscale deployment <deployment-name> --cpu-percent=50 --min=1 --max=10
```

Replace `<deployment-name>` with the name of your deployment. This command sets up an autoscaler that adjusts the number of replicas between 1 and 10 based on whether the average CPU usage across all pods exceeds 50%.

#### 5. Check HPA Status

Monitor the HPA status and see how it scales the deployment:

```bash
kubectl get hpa
```

### Summary

- **Enable Metrics Server**: Use Minikube addons or manually deploy to monitor resource usage.
- **Monitor Metrics**: Use `kubectl top node` and `kubectl top pod` to check real-time metrics.
- **Set Up HPA**: Automatically scale pods based on metrics data using `kubectl autoscale`.
- **Troubleshoot**: If Metrics Server isn't working, manually apply the latest components YAML.

This setup enables effective resource monitoring and autoscaling in Kubernetes, ensuring applications run efficiently by adjusting resources dynamically based on demand.