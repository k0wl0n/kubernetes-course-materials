# Full access policy for the EKS cluster
resource "aws_iam_policy" "eks_full_access_policy" {
  name        = "eks-full-access-policy"
  description = "This policy grants full access to all AWS resources for EKS"
  policy      = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "*",
            "Resource": "*"
        }
    ]
}
EOF
}
