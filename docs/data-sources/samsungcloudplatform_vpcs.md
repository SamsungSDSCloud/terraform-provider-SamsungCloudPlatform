---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "samsungcloudplatform_vpcs Data Source - scp"
subcategory: ""
description: |-
  Provides list of vpcs.
---

# samsungcloudplatform_vpcs (Data Source)

Provides list of vpcs.

## Example Usage

```terraform
data "samsungcloudplatform_vpcs" "vpcs" {
}

output "contents" {
  value = data.samsungcloudplatform_vpcs.vpcs.contents
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `created_by` (String) Person who created the resource
- `page` (Number) Page start number from which to get the list
- `service_zone_id` (String) Service zone id
- `size` (Number) Size to get list
- `vpc_id` (String) VPC id
- `vpc_name` (String) VPC name
- `vpc_states` (String) VPC status

### Read-Only

- `contents` (Block List) VPC list (see [below for nested schema](#nestedblock--contents))
- `id` (String) The ID of this resource.
- `total_count` (Number) Total list size

<a id="nestedblock--contents"></a>
### Nested Schema for `contents`

Read-Only:

- `block_id` (String) Block id
- `created_by` (String) The person who created the resource
- `created_dt` (String) Creation date
- `modified_by` (String) The person who modified the resource
- `modified_dt` (String) Modification date
- `project_id` (String) Project id
- `service_zone_id` (String) Service zone id
- `vpc_id` (String) VPC id
- `vpc_name` (String) VPC name
- `vpc_state` (String) VPC status

