# Specify the AWS provider
provider "aws" {
  region = "ap-southeast-1"  # Change to your preferred region
  access_key = "<access_key>"
  secret_key = "<secret_key>"

  # https://registry.terraform.io/providers/hashicorp/aws/2.70.1/docs#static-credentials
}

# Define the S3 bucket
resource "aws_s3_bucket" "terraform_state_bucket" {
  bucket = "terraform-state-production-252525"  # Replace with your bucket name
  acl    = "private"

  versioning {
    enabled = true
  }

  tags = {
    Name        = "Terraform State Bucket"
    Environment = "Production"
  }
}

# Block public access to the bucket
resource "aws_s3_bucket_public_access_block" "terraform_state_public_access_block" {
  bucket = aws_s3_bucket.terraform_state_bucket.id

  block_public_acls   = true
  ignore_public_acls  = true
  block_public_policy = false  # Set to false to allow bucket policy to be applied
  restrict_public_buckets = true
}

# Define a bucket policy allowing only specific IAM users or roles
resource "aws_s3_bucket_policy" "terraform_state_policy" {
  bucket = aws_s3_bucket.terraform_state_bucket.id

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::490004610217:user/admin-user" 
      },
      "Action": [
        "s3:PutObject",
        "s3:GetObject",
        "s3:DeleteObject",
        "s3:PutBucketPolicy",
        "s3:ListBucket"
      ],
      "Resource": [
        "arn:aws:s3:::terraform-state-production-252525",
        "arn:aws:s3:::terraform-state-production-252525/*"
      ]
    }
  ]
}
POLICY
}

# Define the DynamoDB table for state locking
resource "aws_dynamodb_table" "terraform_state_lock" {
  name         = "terraform-state-production-252525-lock"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }

  tags = {
    Name        = "Terraform State Lock Table"
    Environment = "Production"
  }
}
