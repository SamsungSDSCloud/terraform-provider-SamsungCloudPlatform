---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "samsungcloudplatform_products_by_zone Data Source - scp"
subcategory: ""
description: |-
  
---

# samsungcloudplatform_products_by_zone (Data Source)



## Example Usage

```terraform
data "samsungcloudplatform_products_by_zone" "my_products" {
  service_zone_id = "ZONE-XXXXXXXXXXX"
}

output "result_my_products" {
  value = data.samsungcloudplatform_products_by_zone.my_products
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `service_zone_id` (String) Service zone id

### Optional

- `product_group_id` (String) Product group id
- `product_type` (String) Product type (SCALE, DISK, ...)

### Read-Only

- `contents` (List of Object) Contents (see [below for nested schema](#nestedatt--contents))
- `id` (String) The ID of this resource.
- `total_count` (Number)

<a id="nestedatt--contents"></a>
### Nested Schema for `contents`

Read-Only:

- `created_by` (String)
- `created_dt` (String)
- `modified_by` (String)
- `modified_dt` (String)
- `product_id` (String)
- `product_name` (String)
- `product_type` (String)


