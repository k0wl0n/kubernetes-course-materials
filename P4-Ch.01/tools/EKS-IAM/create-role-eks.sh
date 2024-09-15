#!/bin/bash

# Variables
ROLE_NAME="EKSAdminRole"
POLICY_NAME="EKSCustomPolicy"
ACCOUNT_ID=$(aws sts get-caller-identity --query Account --output text)

# Step 1: Create IAM Role with EC2 and EKS Trust Policy (Silent)
echo "Creating IAM role: $ROLE_NAME..."
aws iam create-role \
    --role-name $ROLE_NAME \
    --assume-role-policy-document '{
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Principal": {
                    "Service": "eks.amazonaws.com"
                },
                "Action": "sts:AssumeRole"
            },
            {
                "Effect": "Allow",
                "Principal": {
                    "Service": "ec2.amazonaws.com"
                },
                "Action": "sts:AssumeRole"
            }
        ]
    }' --no-cli-pager >/dev/null 2>&1

# Step 2: Attach AWS Managed Policies (Silent)
echo "Attaching AWS managed policies..."
aws iam attach-role-policy \
    --role-name $ROLE_NAME \
    --policy-arn arn:aws:iam::aws:policy/AmazonEKSClusterPolicy \
    --no-cli-pager >/dev/null 2>&1

aws iam attach-role-policy \
    --role-name $ROLE_NAME \
    --policy-arn arn:aws:iam::aws:policy/AmazonEKSServicePolicy \
    --no-cli-pager >/dev/null 2>&1

aws iam attach-role-policy \
    --role-name $ROLE_NAME \
    --policy-arn arn:aws:iam::aws:policy/AmazonEC2FullAccess \
    --no-cli-pager >/dev/null 2>&1

aws iam attach-role-policy \
    --role-name $ROLE_NAME \
    --policy-arn arn:aws:iam::aws:policy/IAMFullAccess \
    --no-cli-pager >/dev/null 2>&1

aws iam attach-role-policy \
    --role-name $ROLE_NAME \
    --policy-arn arn:aws:iam::aws:policy/AmazonVPCFullAccess \
    --no-cli-pager >/dev/null 2>&1

aws iam attach-role-policy \
    --role-name $ROLE_NAME \
    --policy-arn arn:aws:iam::aws:policy/ElasticLoadBalancingFullAccess \
    --no-cli-pager >/dev/null 2>&1

# Step 3: Create a custom inline policy for EKS administration (Silent)
echo "Creating custom inline policy..."
cat > eks_custom_policy.json <<EOL
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "eks:*",
                "ec2:*",
                "iam:PassRole",
                "elasticloadbalancing:*",
                "cloudformation:*",
                "autoscaling:*",
                "logs:*",
                "vpc:*",
                "sts:AssumeRole"
            ],
            "Resource": "*"
        }
    ]
}
EOL

aws iam put-role-policy \
    --role-name $ROLE_NAME \
    --policy-name $POLICY_NAME \
    --policy-document file://eks_custom_policy.json \
    --no-cli-pager >/dev/null 2>&1

# Cleanup the policy file
rm eks_custom_policy.json

# Output the role creation status
echo "IAM role $ROLE_NAME created and policies attached silently."
