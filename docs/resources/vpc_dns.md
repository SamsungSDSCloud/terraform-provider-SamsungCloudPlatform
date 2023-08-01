---
page_title: "scp_vpc_dns Resource - scp"
subcategory: ""
description: |-
  Provides a VPC Dns resource.
---

# Resource: scp_vpc_dns

Provides a VPC Dns resource.


## Example Usage

```terraform
data "scp_vpcs" "vpcs" {
}

data "scp_subnets" "subnets"{
  vpc_id = data.scp_vpcs.vpcs.contents[0].vpc_id
}

resource "scp_vpc_dns" "vpcdns01" {
  vpc_id = data.scp_vpcs.vpcs.contents[0].vpc_id
  subnet_id = data.scp_subnets.subnets.contents[0].subnet_id
  domain = "hello.com"
  dns_ip = "10.254.10.254"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `dns_ip` (String) DNS ip address
- `domain` (String) domain (ex: abc.com)
- `subnet_id` (String) Subnet Id for Source IP
- `vpc_id` (String) VPC id

### Optional

- `source_ip` (String) Source Ip address

### Read-Only

- `id` (String) The ID of this resource.