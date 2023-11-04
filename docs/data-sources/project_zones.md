---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scp_project_zones Data Source - scp"
subcategory: ""
description: |-
  Provides list of service zones in project
---

# scp_project_zones (Data Source)

Provides list of service zones in project

## Example Usage

```terraform
data "scp_project_zones" "my_scp_project_zones" {
  project_id = "PROJECT-XXXXXXX"
}

output "output_my_scp_project" {
  value = data.scp_project_zones.my_scp_project_zones
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `project_id` (String) Project ID

### Optional

- `filter` (Block Set) (see [below for nested schema](#nestedblock--filter))

### Read-Only

- `contents` (List of Object) Zones in project (see [below for nested schema](#nestedatt--contents))
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

- `availability_zones` (List of Object) (see [below for nested schema](#nestedobjatt--contents--availability_zones))
- `block_id` (String)
- `is_multi_availability_zone` (Boolean)
- `service_zone_id` (String)
- `service_zone_location` (String)
- `service_zone_name` (String)

<a id="nestedobjatt--contents--availability_zones"></a>
### Nested Schema for `contents.availability_zones`

Read-Only:

- `availability_zone_name` (String)

