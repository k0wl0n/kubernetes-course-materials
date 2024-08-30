### README for `P2-Ch.04-11-12.Job`

This guide provides steps to manage Kubernetes Jobs, including listing, applying, and deleting Jobs within the `P2-Ch.04-11-12.Job` folder.

### Steps

#### 1. List All Resources

Display all Kubernetes resources in the current namespace:

```bash
kubectl get all
```

#### 2. Delete All Resources

Remove all pods, services, deployments, daemonsets, and jobs in the current namespace:

```bash
kubectl delete pods,services,deployments,daemonsets,jobs --all
```

#### 3. Apply Job

Create and run a Job using a specified YAML configuration file:

```bash
kubectl apply -f simple-job.yaml
```

#### 4. List Jobs

Check the list of all Jobs currently running or completed in the cluster:

```bash
kubectl get jobs
```

#### 5. Delete Job

Remove a specific Job from the cluster:

```bash
kubectl delete jobs [job-name]
```

Replace `[job-name]` with the actual name of the Job you want to delete.

This setup demonstrates how to create, manage, and clean up Jobs in Kubernetes, allowing for the execution of short-lived or batch tasks efficiently.