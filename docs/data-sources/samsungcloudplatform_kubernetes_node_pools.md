---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "samsungcloudplatform_kubernetes_node_pools Data Source - scp"
subcategory: ""
description: |-
  Provides list of K8s node pools
---

# samsungcloudplatform_kubernetes_node_pools (Data Source)

Provides list of K8s node pools

## Example Usage

```terraform
# Find all nodepool for current project
data "samsungcloudplatform_kubernetes_node_pools" "my_scp_kubernetes_node_pools" {
  kubernetes_engine_id = "engine id"
}

output "result_scp_kubernetes" {
  value = data.samsungcloudplatform_kubernetes_node_pools.my_scp_kubernetes_node_pools
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `kubernetes_engine_id` (String) K8s engine id

### Optional

- `created_by` (String) The person who created the resource
- `node_pool_name` (String) K8s NodePool name
- `page` (Number) Page start number from which to get the list
- `size` (Number) Size to get list

### Read-Only

- `contents` (Block List) K8s Node Pool list (see [below for nested schema](#nestedblock--contents))
- `id` (String) The ID of this resource.
- `total_count` (Number) Total list count

<a id="nestedblock--contents"></a>
### Nested Schema for `contents`

Read-Only:

- `auto_recovery` (Boolean) Enable auto recovery
- `auto_scale` (Boolean) Enable auto scale
- `contract_id` (String) Contract id
- `created_by` (String) The person who created the resource
- `created_dt` (String) Creation Date
- `current_node_count` (Number) Current node count in the pool
- `desired_node_count` (Number) Desired node count in the pool
- `image_id` (String) Image id
- `in_progress` (Boolean) Check inProgress status
- `k8s_version` (String) K8s version
- `max_node_count` (Number) Maximum node count
- `min_node_count` (Number) Minimum node count
- `modified_by` (String) The person who modified the resource
- `modified_dt` (String) Modification Date
- `node_pool_id` (String) NodePool id
- `node_pool_name` (String) NodePool name
- `node_pool_state` (String) NodePool status
- `os_type` (String) Host OS type (Ubuntu, Window,..)
- `product_group_id` (String) Product group id
- `project_id` (String) Project id
- `region` (String) Modification Date
- `scale_id` (String) Scale id
- `service_level_id` (String) Service level id
- `storage_id` (String) Storage id
- `storage_size` (String) Storage size in GB
- `upgradable` (Boolean) Where to enable upgrade

