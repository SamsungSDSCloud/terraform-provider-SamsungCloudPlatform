---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scp_bm_block_storage Data Source - scp"
subcategory: ""
description: |-
  Provide Block Storage(BM) Detail
---

# scp_bm_block_storage (Data Source)

Provide Block Storage(BM) Detail

## Example Usage

```terraform
data "scp_bm_block_storage" "my_scp_bm_block_storage" {
  storage_id = "STORAGE-xxxxxxxxxxxxxxxxxxxxx"
}

output "output_my_scp_bm_block_storage_org"{
  value = data.scp_bm_block_storage.my_scp_bm_block_storage
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `storage_id` (String) baremetal_block_storage_id

### Read-Only

- `backup_bare_metal_block_storage_id` (String) backup_baremetal_block_storage_id
- `bare_metal_block_storage_name` (String) baremetal_block_storage_name
- `bare_metal_block_storage_purpose` (String) baremetal_block_storage_purpose
- `bare_metal_block_storage_size` (Number) baremetal_block_storage_size
- `bare_metal_block_storage_state` (String) baremetal_block_storage_state
- `block_id` (String) block_id
- `created_by` (String) created_by
- `created_dt` (String) created_dt
- `dr_bare_metal_block_storage_id` (String) dr_baremetal_block_storage_id
- `encryption_enabled` (Boolean) encryption_enabled
- `error_check` (Boolean) error_check
- `id` (String) The ID of this resource.
- `iscsi_target_ip` (List of String) iscsi_target_ip
- `modified_by` (String) modified_by
- `modified_dt` (String) modified_dt
- `origin_bare_metal_block_storage` (List of Object) origin_baremetal_block_storage (see [below for nested schema](#nestedatt--origin_bare_metal_block_storage))
- `product_id` (String) product_id
- `project_id` (String) project_id
- `servers` (List of Object) servers (see [below for nested schema](#nestedatt--servers))
- `service_zone_id` (String) service_zone_id

<a id="nestedatt--origin_bare_metal_block_storage"></a>
### Nested Schema for `origin_bare_metal_block_storage`

Read-Only:

- `bare_metal_block_storage_id` (String)
- `bare_metal_block_storage_name` (String)
- `block_id` (String)
- `service_zone_id` (String)


<a id="nestedatt--servers"></a>
### Nested Schema for `servers`

Read-Only:

- `server_id` (String)
- `server_type` (String)

