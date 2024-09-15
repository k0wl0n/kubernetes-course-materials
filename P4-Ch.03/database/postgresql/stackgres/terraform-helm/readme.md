# StackGres Helm Installation with Terraform

This guide will help you provision the StackGres operator in a Kubernetes cluster using Terraform. The configuration files are set up to install the operator with custom values from a YAML file.

## Project Structure

```
.
├── 0-provider.tf        # Provider configuration for Helm and Kubernetes
├── 1-stackgres.tf       # Helm release configuration for StackGres
└── values
    └── values.yaml      # Custom values for the StackGres operator
```

## Prerequisites

Before running the Terraform commands, ensure you have the following installed:
- [Terraform](https://www.terraform.io/downloads)
- [Helm](https://helm.sh/docs/intro/install/)
- [AWS CLI](https://aws.amazon.com/cli/) (if working with AWS EKS)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)

Make sure your Kubernetes cluster is configured and `kubectl` is pointing to the right cluster context. 

## Steps to Run

1. **Initialize Terraform**

   Navigate to the directory containing the `.tf` files and run the following command to initialize the project and download any necessary providers:
   
   ```bash
   terraform init
   ```

2. **Review the Execution Plan**

   You can preview the changes Terraform will make by running:

   ```bash
   terraform plan
   ```

   This will display an execution plan without actually applying any changes.

3. **Apply the Terraform Configuration**

   To install the StackGres operator, run the following command:

   ```bash
   terraform apply
   ```

   Confirm the action when prompted. Terraform will install the StackGres operator using the Helm chart, and the custom values from `values/values.yaml` will be applied.

4. **Verify Installation**

   Once the StackGres operator is installed, you can check if the installation was successful by running:

   ```bash
   kubectl get pods -n stackgres
   ```

   You should see the StackGres operator pod running in the `stackgres` namespace.

5. **Destroy the Resources (Optional)**

   To remove all resources created by Terraform, you can run:

   ```bash
   terraform destroy
   ```

   This will tear down the StackGres operator and clean up any associated resources.

## Customizing Values

You can customize the StackGres installation by modifying the `values/values.yaml` file. This file allows you to override the default Helm chart values for the StackGres operator.