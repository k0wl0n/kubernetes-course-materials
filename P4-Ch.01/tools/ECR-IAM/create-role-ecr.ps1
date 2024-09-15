# Variables
$RoleName = "ECRAdminRole"
$UserName = "ECRAdminUser"
$PolicyArn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryFullAccess"
$TrustPolicyFile = "trust-policy.json"

# Step 1: Create the trust policy JSON file
$TrustPolicy = @{
    Version = "2012-10-17"
    Statement = @(
        @{
            Effect = "Allow"
            Principal = @{
                AWS = "*"
            }
            Action = "sts:AssumeRole"
        }
    )
} | ConvertTo-Json -Compress

Set-Content -Path $TrustPolicyFile -Value $TrustPolicy

# Step 2: Create the IAM Role
Write-Host "Creating IAM role: $RoleName..."
aws iam create-role `
    --role-name $RoleName `
    --assume-role-policy-document file://$TrustPolicyFile `
    --description "Role for Admin access to ECR"

# Step 3: Attach the ECR Admin Policy to the Role
Write-Host "Attaching ECR admin policy to $RoleName..."
aws iam attach-role-policy `
    --role-name $RoleName `
    --policy-arn $PolicyArn

# Step 4: Create an IAM User
Write-Host "Creating IAM user: $UserName..."
aws iam create-user --user-name $UserName

# Step 5: Create access keys for the user
Write-Host "Creating access keys for user: $UserName..."
$AccessKeys = aws iam create-access-key --user-name $UserName | ConvertFrom-Json

# Step 6: Assign the user to the ECR Admin Role by attaching the policy
Write-Host "Attaching user to the ECR admin role..."
aws iam attach-user-policy `
    --user-name $UserName `
    --policy-arn $PolicyArn

# Step 7: Display the Access Key ID and Secret Access Key
$AccessKeyId = $AccessKeys.AccessKey.AccessKeyId
$SecretAccessKey = $AccessKeys.AccessKey.SecretAccessKey

Write-Host "`nAccess Key ID: $AccessKeyId"
Write-Host "Secret Access Key: $SecretAccessKey"

Write-Host "`nIAM Role $RoleName and IAM user $UserName created and assigned to the role successfully."
