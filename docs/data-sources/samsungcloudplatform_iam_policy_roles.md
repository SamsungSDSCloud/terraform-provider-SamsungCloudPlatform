---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "samsungcloudplatform_iam_policy_roles Data Source - scp"
subcategory: ""
description: |-
  Provide a list of roles that contain this policy
---

# samsungcloudplatform_iam_policy_roles (Data Source)

Provide a list of roles that contain this policy

## Example Usage

```terraform
data "samsungcloudplatform_iam_policy_roles" "my_policy_roles" {
  policy_id = "policy-XXXX"
}

output "result_my_policy_roles" {
  value = data.samsungcloudplatform_iam_policy_roles.my_policy_roles
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `policy_id` (String) Policy ID

### Optional

- `filter` (Block Set) (see [below for nested schema](#nestedblock--filter))
- `role_name` (String) Role name

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

- `created_by` (String)
- `created_by_email` (String)
- `created_by_name` (String)
- `created_dt` (String)
- `description` (String)
- `modified_by` (String)
- `modified_by_email` (String)
- `modified_by_name` (String)
- `modified_dt` (String)
- `principal_policy_id` (String)
- `role_id` (String)
- `role_name` (String)


