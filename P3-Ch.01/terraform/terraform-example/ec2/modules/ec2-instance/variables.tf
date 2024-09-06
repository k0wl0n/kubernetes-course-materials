variable "key_name" {
  description = "Name of the key pair"
  default     = "local-key"
}

variable "public_key_path" {
  description = "Path to the public SSH key"
  default     = "~/.ssh/id_rsa.pub"
  # on windows
  #  default     = "C:/Users/<YourUsername>/.ssh/id_rsa.pub"
  
  # if dont have ssh-key
  # ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
}

variable "ami_id" {
  description = "AMI ID to use"
  default     = "ami-0ac0f5ac9a9b402fa" 
  # os menggunakan amazon linux 2
  # setiap region berbeda
}

variable "instance_type" {
  description = "Type of the EC2 instance"
  default     = "t2.micro"
}

variable "ebs_volume_size" {
  description = "Size of the additional EBS volume in GB"
  default     = 10
}

variable "root_volume_size" {
  description = "Size of the root EBS volume in GB"
  default     = 8
}

variable "instance_name" {
  description = "Name for the instance"
  default     = "ec2-instance"
}
