# Terraform Infrastructure Setup

This guide outlines basic Terraform commands for initializing, planning, applying, and destroying infrastructure, as well as generating SSH keys.

## Prerequisites

Make sure you have [Terraform](https://www.terraform.io/) installed on your system before proceeding.

## Terraform Commands

### 1. Initialize Terraform
This command initializes your Terraform project, downloading provider plugins and setting up the backend:
```bash
terraform init
```

### 2. Plan Infrastructure Changes
Generate an execution plan to review the changes that Terraform will apply:
```bash
terraform plan
```

### 3. Apply Infrastructure Changes
Apply the changes specified in the plan to provision or modify the infrastructure:
```bash
terraform apply
```

### 4. Destroy Infrastructure
When you're done with your infrastructure, you can destroy it using the following command:
```bash
terraform destroy
```

## SSH Key Generation

For secure access to resources like virtual machines, you can generate an SSH key pair. Use the following command to generate a 4096-bit RSA key:

```bash
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
```

Follow the prompts to save the key to the desired location. The public key can be added to your cloud provider, and the private key will remain on your local machine for authentication.
