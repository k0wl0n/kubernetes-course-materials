provider "aws" {
  region = "ap-southeast-1"
}

module "ec2_instance" {
  source = "./modules/ec2-instance"
  instance_type = var.instance_type
}

