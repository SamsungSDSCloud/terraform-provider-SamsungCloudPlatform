---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "samsungcloudplatform_epas Data Source - scp"
subcategory: ""
description: |-
  Search single EPAS database.
---

# samsungcloudplatform_epas (Data Source)

Search single EPAS database.

## Example Usage

```terraform
data "samsungcloudplatform_epas" "my_scp_epas" {
  epas_cluster_id = "SERVICE-123456789"
}

output "output_my_scp_epas" {
  value = data.samsungcloudplatform_epas.my_scp_epas
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `epas_cluster_id` (String) EPAS Cluster Id

### Read-Only

- `audit_enabled` (Boolean) audit enabled
- `backup_config` (List of Object) backup config (see [below for nested schema](#nestedatt--backup_config))
- `block_id` (String) block id
- `contract` (List of Object) contract (see [below for nested schema](#nestedatt--contract))
- `created_by` (String) created by
- `created_dt` (String) created dt
- `database_version` (String) database version
- `epas_cluster_name` (String) EPAS Cluster Name
- `epas_cluster_state` (String) EPAS Cluster State
- `epas_initial_config` (List of Object) EPAS initial config (see [below for nested schema](#nestedatt--epas_initial_config))
- `epas_master_cluster_id` (String) EPAS master cluster id
- `epas_replica_cluster_ids` (List of String) EPAS replica cluster ids
- `epas_server_group` (List of Object) EPAS server group (see [below for nested schema](#nestedatt--epas_server_group))
- `id` (String) The ID of this resource.
- `image_id` (String) image Id
- `maintenance` (List of Object) maintenance (see [below for nested schema](#nestedatt--maintenance))
- `modified_by` (String) modified by
- `modified_dt` (String) modified dt
- `nat_ip_address` (String) nat ip address
- `project_id` (String) project id
- `security_group_ids` (List of String) security group ids
- `service_zone_id` (String) service zone id
- `subnet_id` (String) subnet Id
- `timezone` (String) timezone
- `vpc_id` (String) vPC Id

<a id="nestedatt--backup_config"></a>
### Nested Schema for `backup_config`

Read-Only:

- `full_backup_config` (List of Object) (see [below for nested schema](#nestedobjatt--backup_config--full_backup_config))
- `incremental_backup_config` (List of Object) (see [below for nested schema](#nestedobjatt--backup_config--incremental_backup_config))

<a id="nestedobjatt--backup_config--full_backup_config"></a>
### Nested Schema for `backup_config.full_backup_config`

Read-Only:

- `archive_backup_schedule_frequency` (String)
- `backup_retention_period` (String)
- `backup_start_hour` (Number)
- `object_storage_bucket_id` (String)


<a id="nestedobjatt--backup_config--incremental_backup_config"></a>
### Nested Schema for `backup_config.incremental_backup_config`

Read-Only:

- `archive_backup_schedule_frequency` (String)
- `backup_retention_period` (String)
- `backup_schedule_frequency` (String)
- `object_storage_bucket_id` (String)



<a id="nestedatt--contract"></a>
### Nested Schema for `contract`

Read-Only:

- `contract_end_date` (String)
- `contract_period` (String)
- `contract_start_date` (String)
- `next_contract_end_date` (String)
- `next_contract_period` (String)


<a id="nestedatt--epas_initial_config"></a>
### Nested Schema for `epas_initial_config`

Read-Only:

- `database_encoding` (String)
- `database_locale` (String)
- `database_name` (String)
- `database_port` (Number)
- `database_user_name` (String)


<a id="nestedatt--epas_server_group"></a>
### Nested Schema for `epas_server_group`

Read-Only:

- `block_storages` (List of Object) (see [below for nested schema](#nestedobjatt--epas_server_group--block_storages))
- `encryption_enabled` (Boolean)
- `epas_servers` (List of Object) (see [below for nested schema](#nestedobjatt--epas_server_group--epas_servers))
- `server_group_role_type` (String)
- `server_type` (String)
- `virtual_ip_address` (String)

<a id="nestedobjatt--epas_server_group--block_storages"></a>
### Nested Schema for `epas_server_group.block_storages`

Read-Only:

- `block_storage_group_id` (String)
- `block_storage_name` (String)
- `block_storage_role_type` (String)
- `block_storage_size` (Number)
- `block_storage_type` (String)


<a id="nestedobjatt--epas_server_group--epas_servers"></a>
### Nested Schema for `epas_server_group.epas_servers`

Read-Only:

- `availability_zone_name` (String)
- `created_by` (String)
- `created_dt` (String)
- `epas_server_id` (String)
- `epas_server_name` (String)
- `epas_server_state` (String)
- `modified_by` (String)
- `modified_dt` (String)
- `server_role_type` (String)
- `subnet_ip_address` (String)



<a id="nestedatt--maintenance"></a>
### Nested Schema for `maintenance`

Read-Only:

- `maintenance_period` (Number)
- `maintenance_start_day_of_week` (String)
- `maintenance_start_time` (String)


