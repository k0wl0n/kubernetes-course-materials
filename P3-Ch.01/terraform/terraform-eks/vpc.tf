module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 4.0"

  name = local.name
  cidr = local.vpc_cidr

  azs             = local.azs
  private_subnets = local.private_subnets
  public_subnets  = local.public_subnets

  enable_nat_gateway = true

  private_subnet_tags = {
    "kubernetes.io/cluster/${local.name}" = "shared"
    "kubernetes.io/role/internal-elb" = true
  }
  public_subnet_tags = {
    "kubernetes.io/cluster/${local.name}" = "shared"
     "kubernetes.io/role/elb" = true
  }
  vpc_tags = {
    "kubernetes.io/cluster/${local.name}" = "shared"
  }
}