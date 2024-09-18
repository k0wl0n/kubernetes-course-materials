# IAM Role for EKS Cluster
resource "aws_iam_role" "fastcampus" {
  name = "eks-cluster-fastcampus"

  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "eks.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
POLICY
}

# Attach EKS Cluster Policy to IAM Role
resource "aws_iam_role_policy_attachment" "fastcampus_amazon_eks_cluster_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"
  role       = aws_iam_role.fastcampus.name
}

# Create EKS Cluster
resource "aws_eks_cluster" "fastcampus" {
  name     = "fastcampus"
  version  = "1.30"
  role_arn = aws_iam_role.fastcampus.arn

  vpc_config {
    subnet_ids = [
      aws_subnet.private_ap_southeast_1a.id,
      aws_subnet.private_ap_southeast_1b.id,
      aws_subnet.private_ap_southeast_1c.id,
      aws_subnet.public_ap_southeast_1a.id,
      aws_subnet.public_ap_southeast_1b.id,
      aws_subnet.public_ap_southeast_1c.id
    ]
  }

  depends_on = [aws_iam_role_policy_attachment.fastcampus_amazon_eks_cluster_policy]
}

# IAM Role for EKS Node Group
resource "aws_iam_role" "nodes" {
  name = "eks-node-group-nodes"

  assume_role_policy = jsonencode({
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "ec2.amazonaws.com"
      }
    }]
    Version = "2012-10-17"
  })
}

# Attach Policies for EKS Worker Nodes
resource "aws_iam_role_policy_attachment" "nodes_amazon_eks_worker_node_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"
  role       = aws_iam_role.nodes.name
}

resource "aws_iam_role_policy_attachment" "nodes_amazon_eks_cni_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"
  role       = aws_iam_role.nodes.name
}

resource "aws_iam_role_policy_attachment" "nodes_amazon_ec2_container_registry_read_only" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
  role       = aws_iam_role.nodes.name
}

# Attach full access policy to the EKS cluster role
resource "aws_iam_role_policy_attachment" "eks_full_access_policy_attachment" {
  policy_arn = aws_iam_policy.eks_full_access_policy.arn
  role       = aws_iam_role.fastcampus.name
}

# Attach full access policy to the EKS worker node role
resource "aws_iam_role_policy_attachment" "nodes_full_access_policy_attachment" {
  policy_arn = aws_iam_policy.eks_full_access_policy.arn
  role       = aws_iam_role.nodes.name
}


# Optional: SSH Access via SSM
resource "aws_iam_role_policy_attachment" "amazon_ssm_managed_instance_core" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"
  role       = aws_iam_role.nodes.name
}

resource "aws_eks_node_group" "private_nodes" {
  cluster_name    = aws_eks_cluster.fastcampus.name
  node_group_name = "private-nodes"
  node_role_arn   = aws_iam_role.nodes.arn

  subnet_ids = [
    aws_subnet.private_ap_southeast_1a.id
  ]

  capacity_type  = "SPOT"
  instance_types = ["t3a.medium"]

  scaling_config {
    desired_size = 2
    max_size     = 6
    min_size     = 0
  }

  update_config {
    max_unavailable = 1
  }

  tags = {
    "Region"                                = "ap-southeast-1"
    "Version"                               = "v1.0"
    "CreatedBy"                             = "Terraform"
    "Environment"                           = "Production"
    "Application"                           = "Fastcampuse e-commerce"
    "ResourceType"                          = "EKSNodeGroup"
    "EnvironmentType"                       = "Prod"
    "kubernetes.io/cluster/fastcampus"      = "owned"
    "k8s.io/cluster-autoscaler/enabled"     = "true"
    "k8s.io/cluster-autoscaler/fastcampus"  = "owned"
  }

  depends_on = [
    aws_iam_role_policy_attachment.nodes_amazon_eks_worker_node_policy,
    aws_iam_role_policy_attachment.nodes_amazon_eks_cni_policy,
    aws_iam_role_policy_attachment.nodes_amazon_ec2_container_registry_read_only,
    aws_iam_role_policy_attachment.amazon_ssm_managed_instance_core,
  ]
}



