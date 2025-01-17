---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "samsungcloudplatform_iam_members Data Source - scp"
subcategory: ""
description: |-
  
---

# samsungcloudplatform_iam_members (Data Source)



## Example Usage

```terraform
data "samsungcloudplatform_iam_members" "my_members" {
}

output "result_my_members" {
  value = data.samsungcloudplatform_iam_members.my_members
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `company_name` (String) Company name
- `email` (String) Member's email
- `filter` (Block Set) (see [below for nested schema](#nestedblock--filter))
- `user_name` (String) Member's name

### Read-Only

- `contents` (List of Object) Contents list (see [below for nested schema](#nestedatt--contents))
- `id` (String) The ID of this resource.
- `total_count` (Number)

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

- `company_name` (String)
- `created_by` (String)
- `created_by_email` (String)
- `created_by_name` (String)
- `created_dt` (String)
- `email` (String)
- `last_access_dt` (String)
- `modified_by` (String)
- `modified_by_email` (String)
- `modified_by_name` (String)
- `modified_dt` (String)
- `organization_id` (String)
- `position_name` (String)
- `user_group_count` (Number)
- `user_id` (String)
- `user_name` (String)
- `user_password_reuse_limit_count` (String)


