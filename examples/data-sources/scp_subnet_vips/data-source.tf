data "scp_subnets" "my_scp_subnets" {
}

data "scp_subnet_vips" "my_scp_subnet_vips1" {
  subnet_id = "SUBNET-XXXXXXXX"
}

output "contents" {
  value = data.scp_subnet_vips.my_scp_subnet_vips1.contents
}
