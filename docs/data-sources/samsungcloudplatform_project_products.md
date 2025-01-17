---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "samsungcloudplatform_project_products Data Source - scp"
subcategory: ""
description: |-
  Provides list of products in given project
---

# samsungcloudplatform_project_products (Data Source)

Provides list of products in given project

## Example Usage

```terraform
data "samsungcloudplatform_project" "my_project" {

}

data "samsungcloudplatform_project_products" "my_scp_products" {
  project_id = data.samsungcloudplatform_project.my_project.id
}

output "samsungcloudplatform_project_products" {
  value = data.samsungcloudplatform_project_products.my_scp_products
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `project_id` (String) Project ID

### Optional

- `filter` (Block Set) (see [below for nested schema](#nestedblock--filter))
- `language_code` (String) Language code

### Read-Only

- `contents` (List of Object) List of products  in project (see [below for nested schema](#nestedatt--contents))
- `id` (String) The ID of this resource.
- `total_count` (Number) Total list size

<a id="nestedblock--filter"></a>
### Nested Schema for `filter`

Required:

- `name` (String) Filtering target name
- `values` (List of String) Filtering values. Each matching value is appended. (OR rule)

Optional:

- `use_regex` (Boolean) Enable regex match for values


<a id="nestedatt--contents"></a>
### Nested Schema for `contents`

Read-Only:

- `product_category_description` (String)
- `product_category_id` (String)
- `product_category_name` (String)
- `product_category_state` (String)
- `product_set` (String)
- `products` (List of Object) (see [below for nested schema](#nestedobjatt--contents--products))

<a id="nestedobjatt--contents--products"></a>
### Nested Schema for `contents.products`

Read-Only:

- `is_product_creatable` (String)
- `product_offering_description` (String)
- `product_offering_detail_info` (String)
- `product_offering_id` (String)
- `product_offering_name` (String)
- `product_offering_state` (String)


