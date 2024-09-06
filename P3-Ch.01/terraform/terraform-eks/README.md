# EKS Cluster Provisioning and Connection with Terraform

This guide will walk you through the steps to provision an AWS EKS (Elastic Kubernetes Service) cluster using Terraform and then connect to the cluster using `kubectl`.

## Prerequisites

Before you begin, ensure you have the following tools installed:
- [Terraform](https://www.terraform.io/downloads)
- [AWS CLI](https://aws.amazon.com/cli/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)

### 1. Provision the EKS Cluster

Use Terraform to plan and apply the infrastructure changes required to create an EKS cluster.

- **Plan the infrastructure**:
  ```bash
  terraform plan
  ```

  This command generates and displays an execution plan, detailing what changes Terraform will apply to provision the cluster.

- **Apply the infrastructure**:
  ```bash
  terraform apply
  ```

  This command provisions the EKS cluster as per the Terraform configuration.

### 2. Connect to the EKS Cluster

Once the cluster is provisioned, you can use the AWS CLI to configure your local `kubectl` to interact with the new EKS cluster.

- **Update the kubeconfig**:
  ```bash
  aws eks update-kubeconfig --region ap-southeast-1 --name fastcampus-cluster
  ```

  This command configures `kubectl` to communicate with your EKS cluster in the `ap-southeast-1` region.

## Verifying the Connection

To verify the connection, run the following command to list the nodes in the EKS cluster:

```bash
kubectl get nodes
```

If everything is set up correctly, you should see a list of nodes that are part of your EKS cluster.
