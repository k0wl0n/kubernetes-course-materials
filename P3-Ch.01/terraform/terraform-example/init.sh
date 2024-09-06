#!/bin/bash

# Function to create directory structure
create_structure() {
  echo "Creating Terraform folder structure..."

  # Create main project directories
  mkdir -p terraform-demo/modules/ec2-instance

  # Create main Terraform files
  touch terraform-demo/{main.tf,variables.tf,outputs.tf,provider.tf,terraform.tfvars,README.md}

  # Create module files
  touch terraform-demo/modules/ec2-instance/{main.tf,variables.tf,outputs.tf}

  echo "Writing basic Terraform configuration..."

  # Basic Terraform configuration in main.tf
  cat <<EOL > terraform-demo/main.tf
provider "aws" {
  region = "us-east-1"
}

module "ec2_instance" {
  source = "./modules/ec2-instance"
  instance_type = var.instance_type
}

EOL

  # Variable definition in variables.tf
  cat <<EOL > terraform-demo/variables.tf
variable "instance_type" {
  description = "Type of EC2 instance"
  type        = string
  default     = "t2.micro"
}

EOL

  # Outputs in outputs.tf
  cat <<EOL > terraform-demo/outputs.tf
output "instance_id" {
  description = "EC2 Instance ID"
  value       = module.ec2_instance.instance_id
}

EOL

  # Provider configuration in provider.tf
  cat <<EOL > terraform-demo/provider.tf
provider "aws" {
  region = var.aws_region
}

variable "aws_region" {
  description = "AWS region"
  default     = "us-east-1"
}
EOL

  # Module content in ec2-instance module
  cat <<EOL > terraform-demo/modules/ec2-instance/main.tf
resource "aws_instance" "example" {
  ami           = "ami-0c55b159cbfafe1f0"  # Example AMI for the region
  instance_type = var.instance_type
}

output "instance_id" {
  value = aws_instance.example.id
}

EOL

  # Variable file in module
  cat <<EOL > terraform-demo/modules/ec2-instance/variables.tf
variable "instance_type" {
  description = "Type of EC2 instance"
  type        = string
  default     = "t2.micro"
}
EOL

  echo "Bootstrap complete. Terraform folder structure and scripts created!"
}

# Call the function
create_structure
