---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scp_resource_group_service_types Data Source - scp"
subcategory: ""
description: |-
  
---

# scp_resource_group_service_types (Data Source)



## Example Usage

```terraform
data "scp_resource_group_service_types" "my_service_types" {
}

output "result_my_service_types" {
  value = data.scp_resource_group_service_types.my_service_types
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `service_type` (String) Service type

### Read-Only

- `id` (String) The ID of this resource.
- `service_types` (List of String) Service types

