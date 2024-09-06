locals {
  region = "ap-southeast-1"
  name   = "fastcampus-cluster"
  vpc_cidr = "10.123.0.0/16"
  azs      = ["ap-southeast-1a", "ap-southeast-1b", "ap-southeast-1c"]
  public_subnets  = ["10.123.1.0/24", "10.123.2.0/24", "10.123.3.0/24"]
  private_subnets = ["10.123.4.0/24", "10.123.5.0/24", "10.123.6.0/24"]
  tags = {
    cluster = local.name
  }
}

provider "aws" {
  region = "ap-southeast-1"
}

terraform {
  backend "s3" {
    bucket         = "terraform-state-production-252525"    # Your S3 bucket name
    key            = "terraform.tfstate"                    # The path to the state file within the bucket
    region         = "ap-southeast-1"                       # The region where your bucket is located
    dynamodb_table = "terraform-state-production-252525-lock"  # Your DynamoDB table name for state locking
    encrypt        = true                                   # Enable server-side encryption of the state file
  }
}