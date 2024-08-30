### README for `P2-Ch.04-15-16.Statefulset`

This guide provides steps to manage a Redis cluster using Kubernetes StatefulSets, including creating, scaling, and getting information about the cluster within the `P2-Ch.04-15-16.Statefulset` folder.

### Steps

#### 1. Get All Redis IPs

Retrieve the IP addresses of all Redis pods to set up the cluster:

```bash
IPs=$(kubectl get pods -l app=redis-cluster -o jsonpath='{range.items[*]}{.status.podIP}:6379 {end}')
```

#### 2. Create Redis Cluster

Initialize the Redis cluster using the IPs of the Redis pods:

```bash
kubectl exec -ti redis-cluster-0 -c redis -- /bin/sh -c "redis-cli -h 127.0.0.1 -p 6379 --cluster create ${IPs}"
```

#### 3. Get Redis Cluster Info

Check the status and details of the Redis cluster:

```bash
kubectl exec -it redis-cluster-0 -- /bin/sh -c "redis-cli -h 127.0.0.1 -p 6379 cluster info"
```

#### 4. Scale Application

Scale the deployment of the application (e.g., hit-counter-app) to the desired number of replicas:

```bash
kubectl scale deployment/hit-counter-app --replicas=5
```

#### 5. Scale Redis Cluster

Adjust the number of replicas in the Redis StatefulSet to scale up or down:

```bash
kubectl scale statefulset/redis-cluster --replicas=5
kubectl scale statefulset/redis-cluster --replicas=2
```

This setup demonstrates how to manage and scale a Redis cluster using StatefulSets in Kubernetes, ensuring stateful applications are handled properly and can be scaled dynamically.