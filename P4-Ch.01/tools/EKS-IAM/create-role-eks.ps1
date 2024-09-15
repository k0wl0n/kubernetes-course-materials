# Variables
$RoleName = "EKSAdminRole"
$PolicyName = "EKSCustomPolicy"
$AccountId = (aws sts get-caller-identity --query Account --output text)

# Step 1: Create IAM Role with EC2 and EKS Trust Policy
Write-Host "Creating IAM role: $RoleName..."
aws iam create-role `
    --role-name $RoleName `
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
    }'

# Step 2: Attach AWS Managed Policies
Write-Host "Attaching AWS managed policies..."
aws iam attach-role-policy `
    --role-name $RoleName `
    --policy-arn arn:aws:iam::aws:policy/AmazonEKSClusterPolicy

aws iam attach-role-policy `
    --role-name $RoleName `
    --policy-arn arn:aws:iam::aws:policy/AmazonEKSServicePolicy

aws iam attach-role-policy `
    --role-name $RoleName `
    --policy-arn arn:aws:iam::aws:policy/AmazonEC2FullAccess

aws iam attach-role-policy `
    --role-name $RoleName `
    --policy-arn arn:aws:iam::aws:policy/IAMFullAccess

aws iam attach-role-policy `
    --role-name $RoleName `
    --policy-arn arn:aws:iam::aws:policy/AmazonVPCFullAccess

aws iam attach-role-policy `
    --role-name $RoleName `
    --policy-arn arn:aws:iam::aws:policy/ElasticLoadBalancingFullAccess

# Step 3: Create a custom inline policy for EKS administration
Write-Host "Creating custom inline policy..."

$CustomPolicy = @{
    "Version" = "2012-10-17"
    "Statement" = @(
        @{
            "Effect" = "Allow"
            "Action" = @(
                "eks:*",
                "ec2:*",
                "iam:PassRole",
                "elasticloadbalancing:*",
                "cloudformation:*",
                "autoscaling:*",
                "logs:*",
                "vpc:*",
                "sts:AssumeRole"
            )
            "Resource" = "*"
        }
    )
}

$CustomPolicyJson = $CustomPolicy | ConvertTo-Json -Compress

# Save the policy to a file
$CustomPolicyFile = "eks_custom_policy.json"
$CustomPolicyJson | Out-File $CustomPolicyFile

aws iam put-role-policy `
    --role-name $RoleName `
    --policy-name $PolicyName `
    --policy-document file://$CustomPolicyFile

# Cleanup the policy file
Remove-Item $CustomPolicyFile

# Output the role information
Write-Host "IAM role $RoleName created and policies attached."
