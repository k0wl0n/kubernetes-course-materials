# Terraform Demo Project

This project demonstrates how to set up and provision an AWS EC2 instance using Terraform. The project is structured with reusable modules for easy scalability and flexibility. The following guide will help you understand the project structure and how to use it.

## Project Structure

```bash
terraform-demo/
├── main.tf                # Main Terraform configuration file
├── variables.tf           # Variables used in the main configuration
├── outputs.tf             # Outputs for the main configuration
├── provider.tf            # AWS provider configuration
├── terraform.tfvars       # Variable values specific to the environment
├── README.md              # This README file
└── modules/
    └── ec2-instance/
        ├── main.tf        # EC2 module main configuration
        ├── variables.tf   # Variables for the EC2 module
        └── outputs.tf     # Outputs for the EC2 module
```

## Prerequisites

Before using this Terraform project, ensure you have the following:

- [Terraform](https://www.terraform.io/downloads.html) installed.
- An AWS account with appropriate credentials configured on your machine (via AWS CLI or environment variables).
- Basic knowledge of Terraform and AWS EC2.

## Getting Started

Follow these steps to initialize and use the project:

### 1. Initialize the project

Before applying any Terraform changes, initialize the project to download the required providers and modules:

```bash
cd terraform-demo
terraform init
```

### 2. Customize Variables

The project uses variables to make the configuration flexible. You can modify the `terraform.tfvars` file to set your specific values. For example, to set the instance type, you can add this to `terraform.tfvars`:

```hcl
instance_type = "t3.micro"
aws_region    = "us-east-1"
```

### 3. Plan the Infrastructure

To see what changes Terraform will make, you can run the `terraform plan` command:

```bash
terraform plan
```

This will output an execution plan of the infrastructure that will be created.

### 4. Apply the Infrastructure

To apply the infrastructure and provision the EC2 instance, run:

```bash
terraform apply
```

You will be asked to confirm before the resources are created. Type `yes` to proceed.

### 5. Destroy the Infrastructure

To tear down and remove all the resources that were created, run:

```bash
terraform destroy
```

Again, you will be prompted to confirm the destruction.

## Project Details

### main.tf

This file contains the main configuration for Terraform. It sets up the AWS provider and calls the EC2 instance module.

```hcl
provider "aws" {
  region = "us-east-1"
}

module "ec2_instance" {
  source = "./modules/ec2-instance"
  instance_type = var.instance_type
}
```

### variables.tf

This file defines the input variables for the project. It contains default values and descriptions for the variables:

```hcl
variable "instance_type" {
  description = "Type of EC2 instance"
  type        = string
  default     = "t2.micro"
}
```

### outputs.tf

This file defines the outputs that Terraform will return after the resources are created. In this case, it returns the EC2 instance ID:

```hcl
output "instance_id" {
  description = "EC2 Instance ID"
  value       = module.ec2_instance.instance_id
}
```

### provider.tf

This file configures the AWS provider. It uses the `aws_region` variable for flexibility:

```hcl
provider "aws" {
  region = var.aws_region
}

variable "aws_region" {
  description = "AWS region"
  default     = "us-east-1"
}
```

### Module: EC2 Instance

The `modules/ec2-instance/` directory contains the reusable module that provisions an EC2 instance. This module can be used in other projects by referencing its path.

- **main.tf**: The core EC2 instance configuration:

  ```hcl
  resource "aws_instance" "example" {
    ami           = "ami-0c55b159cbfafe1f0"  # Example AMI for the region
    instance_type = var.instance_type
  }

  output "instance_id" {
    value = aws_instance.example.id
  }
  ```

- **variables.tf**: The variables used by the EC2 module, such as `instance_type`:

  ```hcl
  variable "instance_type" {
    description = "Type of EC2 instance"
    type        = string
    default     = "t2.micro"
  }
  ```

- **outputs.tf**: The module output for the EC2 instance ID:

  ```hcl
  output "instance_id" {
    value = aws_instance.example.id
  }
  ```
