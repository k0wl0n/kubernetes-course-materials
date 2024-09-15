# StackGres Installation with Helm

This guide provides step-by-step instructions for installing the StackGres operator in a Kubernetes cluster using Helm.

## Prerequisites

Before proceeding, ensure you have the following tools installed and configured:
- [Helm](https://helm.sh/docs/intro/install/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)

### 1. Add the StackGres Helm Repository

First, you need to add the StackGres Helm chart repository to Helm:

```bash
helm repo add stackgres-charts https://stackgres.io/downloads/stackgres-k8s/stackgres/helm/
```

This command registers the StackGres chart repository in Helm, allowing you to install the StackGres operator.

### 2. View Default Values for the StackGres Operator

To review the default configuration values for the StackGres operator, use the following command:

```bash
helm show values stackgres-charts/stackgres-operator
```

This command displays the default values used in the StackGres Helm chart, which you can customize if necessary.

### 3. Install the StackGres Operator

Install the StackGres operator in a Kubernetes cluster, creating a new namespace called `stackgres`:

```bash
helm install --create-namespace --namespace stackgres stackgres-operator stackgres-charts/stackgres-operator
```

This command will:
- Create the `stackgres` namespace (if it doesn't exist)
- Install the StackGres operator into that namespace

## Verifying the Installation

To verify that the StackGres operator was installed successfully, run the following command:

```bash
kubectl get pods --namespace stackgres
```

You should see the StackGres operator pod(s) running in the `stackgres` namespace.
