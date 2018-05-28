provider "aws" {
  region = "${var.region}"
  access_key = "${var.access_key}"
  secret_key = "${var.secret_key}"
}

# AWS API instance
resource "aws_instance" "jibjib_api" {
  ami = "${var.ami}"
  instance_type = "${var.instance_type}"
  key_name = "${var.key_name}"
  security_groups = ["${aws_security_group.jibjib_api.name}"]
  tags {
    Name = "${var.project_name}_api"
    Project = "${var.project_name}"
  }

  provisioner "remote-exec" {
    script = "${path.root}/bootstrap/ubuntu.sh"
    connection {
      type = "ssh"
      user = "${var.ssh_user}"
      private_key = "${file("${var.private_key}")}"
    }
  }
}

# AWS DB instance
resource "aws_instance" "jibjib_db" {
  ami = "${var.ami}"
  instance_type = "${var.instance_type}"
  key_name = "${var.key_name}"
  security_groups = ["${aws_security_group.jibjib_api.name}"]
  tags {
    Name = "${var.project_name}_db"
    Project = "${var.project_name}"
  }

  provisioner "remote-exec" {
    script = "${path.root}/bootstrap/ubuntu.sh"
    connection {
      type = "ssh"
      user = "${var.ssh_user}"
      private_key = "${file("${var.private_key}")}"
    }
  }  
}

# Use a specific key par for deployment
resource "aws_key_pair" "deployer" {
  key_name = "${var.key_name}"
  public_key = "${file("${var.public_key}")}"
}

# API IP
resource "aws_eip" "jibjib_api" {
  instance = "${aws_instance.jibjib_api.id}"
}

# DB IP
resource "aws_eip" "jibjib_db" {
  instance = "${aws_instance.jibjib_db.id}"
}

# Create a Security Group
resource "aws_security_group" "jibjib_api" {
  name = "${var.project_name}_sg"
  description = "${var.sg_desc}"
  vpc_id = "${var.vpc}"

  # Allow all TCP inside SG with itself
  ingress {
    from_port = 0
    to_port = 65535
    protocol = "tcp"
    self = true
  }

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

  # MongoDB
  # ingress {
  #   from_port = 27017
  #   to_port = 27017
  #   protocol = "tcp"
  #   cidr_blocks = ["0.0.0.0/0"]
  #   ipv6_cidr_blocks = ["::/0"]
  # }

  # Allow all to exit
  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
    self = true
  }
}