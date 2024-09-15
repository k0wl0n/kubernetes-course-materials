#!/bin/bash

# Variables
ROLE_NAME="ECRAdminRole"
USER_NAME="ECRAdminUser"
TRUST_POLICY_FILE="trust-policy.json"
POLICY_ARN="arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryFullAccess"

# Step 1: Create a trust policy JSON file
cat > $TRUST_POLICY_FILE << EOL
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::$ACCOUNT_ID:user/$USER_NAME"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOL

# Step 2: Create the IAM Role
echo "Creating IAM role: $ROLE_NAME..."
aws iam create-role \
  --role-name $ROLE_NAME \
  --assume-role-policy-document file://$TRUST_POLICY_FILE \
  --description "Role for Admin access to ECR"

# Step 3: Attach the ECR Admin Policy to the Role
echo "Attaching ECR admin policy to $ROLE_NAME..."
aws iam attach-role-policy \
  --role-name $ROLE_NAME \
  --policy-arn $POLICY_ARN

# Step 4: Create an IAM User
echo "Creating IAM user: $USER_NAME..."
aws iam create-user --user-name $USER_NAME

# Step 5: Create access keys for the user
echo "Creating access keys for user: $USER_NAME..."
access_keys=$(aws iam create-access-key --user-name $USER_NAME)

# Step 6: Assign the user to the ECR Admin Role
echo "Attaching user to the ECR admin role..."
aws iam attach-user-policy \
  --user-name $USER_NAME \
  --policy-arn $POLICY_ARN

# Step 7: Display the Access Key ID and Secret Access Key
access_key_id=$(echo $access_keys | jq -r '.AccessKey.AccessKeyId')
secret_access_key=$(echo $access_keys | jq -r '.AccessKey.SecretAccessKey')

echo "Access Key ID: $access_key_id"
echo "Secret Access Key: $secret_access_key"

echo "IAM Role $ROLE_NAME and IAM user $USER_NAME created and assigned to the role successfully."
