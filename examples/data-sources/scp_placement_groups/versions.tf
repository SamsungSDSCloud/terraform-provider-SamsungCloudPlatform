terraform {
  required_providers {
    samsungcloudplatform = {
      version = "3.5.0"
      source  = "SamsungSDSCloud/samsungcloudplatform"
    }
  }
  required_version = ">= 0.13"
}

# Provider setup
provider "samsungcloudplatform" {
}