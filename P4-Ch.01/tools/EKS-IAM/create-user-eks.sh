#!/bin/bash

# Variables
USER_NAME="EKSAdminUser1"
ROLE_NAME="EKSAdminRole1"
POLICY_NAME="EKSCustomPolicy"
ACCOUNT_ID=$(aws sts get-caller-identity --query Account --output text)

# Step 1: Create IAM User
echo "Creating IAM user: $USER_NAME..."
aws iam create-user \
    --user-name $USER_NAME >/dev/null 2>&1

# Step 2: Create an inline policy to allow the user to assume the role
echo "Creating assume role policy for $USER_NAME to assume $ROLE_NAME..."
cat > assume_role_policy.json <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "sts:AssumeRole",
            "Resource": "arn:aws:iam::$ACCOUNT_ID:role/$ROLE_NAME"
        }
    ]
}
EOF

# Step 3: Attach the inline policy to the user
echo "Attaching assume role policy to $USER_NAME..."
aws iam put-user-policy \
    --user-name $USER_NAME \
    --policy-name "Assume${ROLE_NAME}Policy" \
    --policy-document file://assume_role_policy.json >/dev/null 2>&1

# Step 4: Create Access Keys for the User
echo "Creating access keys for $USER_NAME..."
aws iam create-access-key --user-name $USER_NAME > access_keys.json

# Extract and display the keys (ensure secure storage in production)
ACCESS_KEY=$(cat access_keys.json | jq -r '.AccessKey.AccessKeyId')
SECRET_KEY=$(cat access_keys.json | jq -r '.AccessKey.SecretAccessKey')

echo "Access Key: $ACCESS_KEY"
echo "Secret Key: $SECRET_KEY"

# Cleanup JSON file with access keys and policy
rm access_keys.json
rm assume_role_policy.json

echo "User $USER_NAME now has permission to assume the role $ROLE_NAME."
