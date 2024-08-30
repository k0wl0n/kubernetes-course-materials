### README for `P2-Ch.04-13-14.CronJob`

This guide provides steps to manage Kubernetes CronJobs, including listing, applying, and deleting CronJobs within the `P2-Ch.04-13-14.CronJob` folder.

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

#### 3. Apply Scheduled Job

Create and run a scheduled CronJob using a specified YAML configuration file:

```bash
kubectl apply -f schedule-job.yaml
```

#### 4. List CronJobs

Check the list of all CronJobs currently configured in the cluster:

```bash
kubectl get cronjobs
```

#### 5. Delete CronJob

Remove a specific CronJob from the cluster:

```bash
kubectl delete cronjobs [job-name]
```

Replace `[job-name]` with the actual name of the CronJob you want to delete.

This setup demonstrates how to create, manage, and clean up CronJobs in Kubernetes, allowing for the automation of tasks on a scheduled basis.