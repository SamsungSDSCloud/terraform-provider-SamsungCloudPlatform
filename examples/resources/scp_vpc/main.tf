data "scp_region" "my_region" {
}

resource "scp_vpc" "vpc01" {
  name        = var.name
  description = "Vpc generated from Terraform"
  region      = data.scp_region.my_region.location
  tags = {
    tk01 = "tv01"
    tk02 = "tv02"
  }
}

