provider "aws" {
  region = "ap-southeast-1"
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.56"
    }
  }
}


terraform {
  backend "s3" {
    bucket         = "terraform-state-production-fascampus"    # Your S3 bucket name
    key            = "terraform.tfstate"                    # The path to the state file within the bucket
    region         = "ap-southeast-1"                       # The region where your bucket is located
    dynamodb_table = "terraform-state-production-fascampus-lock"  # Your DynamoDB table name for state locking
    encrypt        = true                                   # Enable server-side encryption of the state file
  }
}