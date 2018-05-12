output "jibjib_eip_public_ip" {
  value = "${aws_eip.jibjib_api.public_ip}"
}

output "jibjib_instance_id" {
  value = "${aws_instance.jibjib_api.id}"
}