### README for `P2-Ch.04-7-8.Deployment`

This guide provides steps to manage Kubernetes Deployments, including adding nodes, changing container images, handling rollouts, and scaling deployments within the `P2-Ch.04-7-8.Deployment` folder.

### Steps

#### 1. Adding Nodes

Add nodes to the Minikube cluster and label them as worker nodes:

```bash
minikube node add --profile=single
kubectl label nodes single-m02 type=worker
kubectl label nodes single-m03 type=worker
```

#### 2. Manually Change Container Image

Update the container image for a deployment and view the deployment details:

```bash
kubectl set image deploy deployment-echo nginx-container=kowlon/go-echo:v1 --record
kubectl describe deploy deployment-echo
```

#### 3. Rollout Commands

Manage the rollout process for the deployment:

```bash
kubectl rollout status deployment/deployment-echo
kubectl rollout restart deployment/deployment-echo
kubectl rollout history deployment/deployment-echo
kubectl rollout undo deployment/deployment-echo
kubectl rollout pause deployment/deployment-echo
kubectl rollout resume deployment/deployment-echo
```

#### 4. Scaling Deployment Replicas

Scale the deployment to the desired number of replicas:

```bash
kubectl scale deployment deployment-echo --replicas=5
```

#### 5. Loop to Access Service

Run a loop to continuously access a service running on localhost to test the deployment:

```bash
while true; do
  curl -sSL http://localhost:8080/ && echo
  sleep 0.2
done
```

#### 6. Rollout Command Descriptions

- `kubectl rollout history deployment/deployment-echo`: View the revision history of the Deployment.
- `kubectl rollout pause deployment/deployment-echo`: Pause the ongoing rollout, preventing further updates.
- `kubectl rollout resume deployment/deployment-echo`: Resume a previously paused rollout.
- `kubectl rollout restart deployment/deployment-echo`: Restart the Deployment by triggering a rolling update.
- `kubectl rollout status deployment/deployment-echo`: Check the current status of the Deployment's rollout.
- `kubectl rollout undo deployment/deployment-echo`: Roll back the Deployment to the previous revision.

This setup demonstrates how to manage Kubernetes Deployments, including scaling, updating container images, and using rollout strategies to ensure smooth updates and recoveries.