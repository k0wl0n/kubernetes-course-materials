# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

provider "aws" {
  region = var.region
}

terraform {
  backend "s3" {
    bucket         = "terraform-eks-production-fascampus"    # Your S3 bucket name
    key            = "terraform.tfstate"                    # The path to the state file within the bucket
    region         = "ap-southeast-1"                       # The region where your bucket is located
    dynamodb_table = "terraform-eks-production-fascampus-lock"  # Your DynamoDB table name for state locking
    encrypt        = true                                   # Enable server-side encryption of the state file
  }
}

# Filter out local zones, which are not currently supported 
# with managed node groups
data "aws_availability_zones" "available" {
  filter {
    name   = "opt-in-status"
    values = ["opt-in-not-required"]
  }
}

locals {
  cluster_name = "eks-fastcampus-production"
}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "5.8.1"

  name = "production-vpc"

  cidr = "10.0.0.0/16"
  azs  = slice(data.aws_availability_zones.available.names, 0, 3)

  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets  = ["10.0.4.0/24", "10.0.5.0/24", "10.0.6.0/24"]

  enable_nat_gateway   = true
  single_nat_gateway   = true
  enable_dns_hostnames = true

  public_subnet_tags = {
    "kubernetes.io/role/elb" = 1
  }

  private_subnet_tags = {
    "kubernetes.io/role/internal-elb" = 1
  }
}

module "eks" {
  source  = "terraform-aws-modules/eks/aws"
  version = "20.8.5"

  cluster_name    = local.cluster_name
  cluster_version = "1.30"

  cluster_endpoint_public_access           = true
  enable_cluster_creator_admin_permissions = true

  cluster_addons = {
    aws-ebs-csi-driver = {
      service_account_role_arn = module.irsa-ebs-csi.iam_role_arn
    }
  }

  vpc_id     = module.vpc.vpc_id
  subnet_ids = module.vpc.private_subnets

  eks_managed_node_group_defaults = {
    ami_type = "AL2_x86_64"

  }

  eks_managed_node_groups = {
    spot = {
      name = "node-group-spot"

      instance_types = ["t3.medium"]

      min_size     = 2
      max_size     = 5
      desired_size = 3
      
      # Set the capacity type to use Spot instances
      capacity_type = "SPOT"
      # Spot Instances for cost optimization. but comes with the potential for interruptions.

      # capacity_type = "ON_DEMAND"
      # On-Demand to ensure stability

      additional_tags = {
        "Environment"      = "Production"
        "Application"      = "Fastcampuse e-commerce"
        "Service"          = "e-commerce"
        "Owner"            = "e-commerce"
        "OwnerEmail"       = "e-commerce@example.com"
        "BusinessUnit"     = "IT"
        "EnvironmentType"  = "Prod"
        "ResourceType"     = "EKSNodeGroup"
        "Version"          = "v1.0"
        "Region"           = "ap-southeast-1"
        "CreatedBy"        = "Terraform"
      }
    }
  }
}


# https://aws.amazon.com/blogs/containers/amazon-ebs-csi-driver-is-now-generally-available-in-amazon-eks-add-ons/ 
data "aws_iam_policy" "ebs_csi_policy" {
  arn = "arn:aws:iam::aws:policy/service-role/AmazonEBSCSIDriverPolicy"
}

module "irsa-ebs-csi" {
  source  = "terraform-aws-modules/iam/aws//modules/iam-assumable-role-with-oidc"
  version = "5.39.0"

  create_role                   = true
  role_name                     = "AmazonEKSTFEBSCSIRole-${module.eks.cluster_name}"
  provider_url                  = module.eks.oidc_provider
  role_policy_arns              = [data.aws_iam_policy.ebs_csi_policy.arn]
  oidc_fully_qualified_subjects = ["system:serviceaccount:kube-system:ebs-csi-controller-sa"]
}
