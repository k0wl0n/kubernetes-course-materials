# Key Pair
resource "aws_key_pair" "ssh_key" {
  key_name   = var.key_name
  public_key = file(var.public_key_path)
}

# Security Group allowing SSH access
resource "aws_security_group" "ssh_access" {
  name        = "ssh_access"
  description = "Allow SSH access from everywhere"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # Open to all IPs, restrict as needed
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# EC2 instance
resource "aws_instance" "vm-ec2" {
  ami           = var.ami_id
  instance_type = var.instance_type
  key_name      = aws_key_pair.ssh_key.key_name
  security_groups = [aws_security_group.ssh_access.name]

  # Attach additional EBS Volume
  ebs_block_device {
    device_name           = "/dev/sdh"
    volume_size           = var.ebs_volume_size
    delete_on_termination = true
  }

  # Root block device
  root_block_device {
    volume_size = var.root_volume_size
  }

  tags = {
    Name = var.instance_name
  }
}

# Output for instance id
output "instance_id" {
  value = aws_instance.vm-ec2.id
}
