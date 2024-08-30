### README for `P2-Ch.04-17.RolloutRollback`

This guide provides steps to manage and revert changes to Kubernetes Deployments using rollout and rollback strategies within the `P2-Ch.04-17.RolloutRollback` folder.

### Steps

#### 1. Apply Rollout and Rollback Configuration

Deploy the configuration that includes rollout and rollback strategies:

```bash
kubectl apply -f rollout-rollback.yaml
```

#### 2. Check Rollout History

View the revision history of the specified deployment to see previous versions:

```bash
kubectl rollout history deployment/rolling-update-demo
```

#### 3. Rollback Deployment

Revert the deployment to a previous revision if necessary:

```bash
kubectl rollout undo deployment/rolling-update-demo
```

This setup demonstrates how to manage deployment updates and handle rollbacks effectively, providing a way to ensure application stability and recover from problematic updates.