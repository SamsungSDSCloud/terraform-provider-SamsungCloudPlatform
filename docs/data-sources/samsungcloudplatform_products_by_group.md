---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "samsungcloudplatform_products_by_group Data Source - scp"
subcategory: ""
description: |-
  
---

# samsungcloudplatform_products_by_group (Data Source)



## Example Usage

```terraform
data "samsungcloudplatform_products_by_group" "my_products" {
  product_group_id = "PRODUCTGROUP-XXXXXXXXXX"
}

output "result_my_products" {
  value = data.samsungcloudplatform_products_by_group.my_products
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `product_group_id` (String) Product group id

### Read-Only

- `id` (String) The ID of this resource.
- `products` (List of Object) (see [below for nested schema](#nestedatt--products))

<a id="nestedatt--products"></a>
### Nested Schema for `products`

Read-Only:

- `item` (List of Object) (see [below for nested schema](#nestedobjatt--products--item))
- `product_description` (String)
- `product_id` (String)
- `product_name` (String)
- `product_state` (String)
- `product_type` (String)

<a id="nestedobjatt--products--item"></a>
### Nested Schema for `products.item`

Read-Only:

- `item_name` (String)
- `item_type` (String)
- `item_value` (String)


