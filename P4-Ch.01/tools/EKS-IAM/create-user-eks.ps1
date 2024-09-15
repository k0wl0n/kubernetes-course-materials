# Variables
$UserName = "EKSAdminUser"
$RoleName = "EKSAdminRole"
$AccountId = (aws sts get-caller-identity --query Account --output text)

# Step 1: Create IAM User
Write-Host "Creating IAM user: $UserName..."
aws iam create-user `
    --user-name $UserName | Out-Null

# Step 2: Create an inline policy to allow the user to assume the role
Write-Host "Creating assume role policy for $UserName to assume $RoleName..."
$AssumeRolePolicy = @"
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "sts:AssumeRole",
            "Resource": "arn:aws:iam::$AccountId:role/$RoleName"
        }
    ]
}
"@

# Save the assume-role policy to a file
$AssumeRolePolicyFile = "assume_role_policy.json"
$AssumeRolePolicy | Set-Content -Path $AssumeRolePolicyFile

# Step 3: Attach the inline policy to the user
Write-Host "Attaching assume role policy to $UserName..."
aws iam put-user-policy `
    --user-name $UserName `
    --policy-name "Assume${RoleName}Policy" `
    --policy-document "file://$AssumeRolePolicyFile" | Out-Null

# Step 4: Create Access Keys for the User
Write-Host "Creating access keys for $UserName..."
$AccessKeys = aws iam create-access-key --user-name $UserName | ConvertFrom-Json

# Extract Access Key and Secret Key
$AccessKey = $AccessKeys.AccessKey.AccessKeyId
$SecretKey = $AccessKeys.AccessKey.SecretAccessKey

# Display the keys (ensure secure storage in production)
Write-Host "Access Key: $AccessKey"
Write-Host "Secret Key: $SecretKey"

# Cleanup the JSON policy file
Remove-Item $AssumeRolePolicyFile

Write-Host "User $UserName now has permission to assume the role $RoleName."
