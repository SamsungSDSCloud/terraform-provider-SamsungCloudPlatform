---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "samsungcloudplatform_resource_groups_in_my_projects Data Source - scp"
subcategory: ""
description: |-
  
---

# samsungcloudplatform_resource_groups_in_my_projects (Data Source)



## Example Usage

```terraform
data "samsungcloudplatform_resource_groups_in_my_projects" "my_resource_groups_in_my_projects" {
}

output "result_my_resource_groups" {
  value = data.samsungcloudplatform_resource_groups_in_my_projects.my_resource_groups_in_my_projects
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `created_by_id` (String) The user id which created the resource group
- `modified_by_email` (String) The user email which modified the resource group
- `modified_by_id` (String) The user id which modified the resource group
- `project_ids` (List of String) Project id list
- `resource_group_name` (String) Resource group name

### Read-Only

- `contents` (List of Object) Resource group list (see [below for nested schema](#nestedatt--contents))
- `id` (String) The ID of this resource.
- `total_count` (Number) Total count

<a id="nestedatt--contents"></a>
### Nested Schema for `contents`

Read-Only:

- `account_name` (String)
- `created_by_email` (String)
- `created_by_id` (String)
- `created_by_name` (String)
- `created_dt` (String)
- `modified_by_email` (String)
- `modified_by_id` (String)
- `modified_by_name` (String)
- `modified_dt` (String)
- `project_id` (String)
- `project_name` (String)
- `resource_group_description` (String)
- `resource_group_id` (String)
- `resource_group_name` (String)


