provider "aws" {
  region = "${var.region}"
  access_key = "${var.access_key}"
  secret_key = "${var.secret_key}"
}

# The AWS Instance to deploy the API to
resource "aws_instance" "jibjib_api" {
  ami = "${var.ami}"
  instance_type = "${var.instance_type}"
  key_name = "${var.key_name}"
  security_groups = ["${aws_security_group.jibjib_api.name}"]
  tags {
    Name = "${var.instance_name}"
    Project = "${var.project_name}"
  }
}

# Assign an EIP
resource "aws_eip" "jibjib_api" {
  instance = "${aws_instance.jibjib_api.id}"
}

# Create a Security Group
resource "aws_security_group" "jibjib_api" {
  name = "${var.instance_name}"
  description = "${var.sg_desc}"
  vpc_id = "${var.vpc}"

  # SSH
  ingress {
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  # JibJib API
  ingress {
    from_port = 8080
    to_port = 8080
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  # Allow all to exit
  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }
}