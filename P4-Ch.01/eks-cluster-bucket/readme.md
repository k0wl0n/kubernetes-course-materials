# AWS Infrastructure Setup with Terraform

This guide will help you set up an IAM user with the necessary policies to manage AWS resources using Terraform and AWS CLI.

## Prerequisites

### Install Terraform CLI
- **Windows**: 
  ```bash
  choco install terraform --pre
  ```
- **Mac/Linux**: 
  ```bash
  brew install terraform
  ```

### Install AWS CLI
- **Windows**: 
  ```bash
  choco install awscli
  ```
- **Mac/Linux**: 
  ```bash
  brew install awscli
  ```

## Setting Up an Admin IAM User

### 1. Create an IAM User
Create a new IAM user for administrative access:
```bash
aws iam create-user --user-name admin-user
```

### 2. Attach Policies to the IAM User
Attach the following AWS-managed policies to the IAM user to grant necessary permissions:

```bash
aws iam attach-user-policy --user-name admin-user --policy-arn arn:aws:iam::aws:policy/AdministratorAccess
aws iam attach-user-policy --user-name admin-user --policy-arn arn:aws:iam::aws:policy/IAMFullAccess
aws iam attach-user-policy --user-name admin-user --policy-arn arn:aws:iam::aws:policy/AmazonEC2FullAccess
aws iam attach-user-policy --user-name admin-user --policy-arn arn:aws:iam::aws:policy/AmazonS3FullAccess
aws iam attach-user-policy --user-name admin-user --policy-arn arn:aws:iam::aws:policy/AmazonDynamoDBFullAccess
```

### 3. Create Access Key for the IAM User
Generate an access key for programmatic access to AWS resources:
```bash
aws iam create-access-key --user-name admin-user > admin_user_credentials.json
```

The access key and secret key will be saved in `admin_user_credentials.json`.

### 4. Verify Policies Attached to the IAM User
To confirm the policies are attached correctly, list them with the following command:
```bash
aws iam list-attached-user-policies --user-name admin-user
```

### Output Access Key Details
View the access key details:
```bash
cat admin_user_credentials.json
```

Use this access key and secret key to configure your AWS CLI or in Terraform configuration.
