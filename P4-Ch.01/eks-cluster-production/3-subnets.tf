resource "aws_subnet" "private_ap_southeast_1a" {
  vpc_id            = aws_vpc.vpc-prod.id
  cidr_block        = "10.0.0.0/19"  
  availability_zone = "ap-southeast-1a"

  # 3 AZ
  # 3 subnet public
  # 3 subnet private

  tags = {
    "Name"                            = "private-ap-southeast-1a"
    "kubernetes.io/role/internal-elb" = "1"
    "kubernetes.io/cluster/fastcampus" = "owned"
  }
}

resource "aws_subnet" "private_ap_southeast_1b" {
  vpc_id            = aws_vpc.vpc-prod.id
  cidr_block        = "10.0.32.0/19"  
  availability_zone = "ap-southeast-1b"

  tags = {
    "Name"                            = "private-ap-southeast-1b"
    "kubernetes.io/role/internal-elb" = "1"
    "kubernetes.io/cluster/fastcampus" = "owned"
  }
}

resource "aws_subnet" "private_ap_southeast_1c" {
  vpc_id            = aws_vpc.vpc-prod.id
  cidr_block        = "10.0.64.0/19"  
  availability_zone = "ap-southeast-1c"

  tags = {
    "Name"                            = "private-ap-southeast-1c"
    "kubernetes.io/role/internal-elb" = "1"
    "kubernetes.io/cluster/fastcampus" = "owned"
  }
}


resource "aws_subnet" "public_ap_southeast_1a" {
  vpc_id                  = aws_vpc.vpc-prod.id
  cidr_block              = "10.0.96.0/19"  
  availability_zone       = "ap-southeast-1a"
  map_public_ip_on_launch = true

  tags = {
    "Name"                       = "public-ap-southeast-1a"
    "kubernetes.io/role/elb"     = "1"
    "kubernetes.io/cluster/fastcampus" = "owned"
  }
}

resource "aws_subnet" "public_ap_southeast_1b" {
  vpc_id                  = aws_vpc.vpc-prod.id
  cidr_block              = "10.0.128.0/19"  
  availability_zone       = "ap-southeast-1b"
  map_public_ip_on_launch = true

  tags = {
    "Name"                       = "public-ap-southeast-1b"
    "kubernetes.io/role/elb"     = "1"
    "kubernetes.io/cluster/fastcampus" = "owned"
  }
}

resource "aws_subnet" "public_ap_southeast_1c" {
  vpc_id                  = aws_vpc.vpc-prod.id
  cidr_block              = "10.0.160.0/19"  
  availability_zone       = "ap-southeast-1c"
  map_public_ip_on_launch = true

  tags = {
    "Name"                       = "public-ap-southeast-1c"
    "kubernetes.io/role/elb"     = "1"
    "kubernetes.io/cluster/fastcampus" = "owned"
  }
}
