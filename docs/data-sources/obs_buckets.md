---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scp_obs_buckets Data Source - scp"
subcategory: ""
description: |-
  Provides list of Object Storage Buckets.
---

# scp_obs_buckets (Data Source)

Provides list of Object Storage Buckets.

## Example Usage

```terraform
data "scp_obs_buckets" "my_scp_obs_buckets" {
}

output "output_my_scp_obs_buckets" {
  value = data.scp_obs_buckets.my_scp_obs_buckets
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `created_by` (String) Created By
- `is_obs_bucket_sync` (Boolean) Perform Object Storage Bucket sync (true | false)
- `is_obs_system_bucket_enabled` (Boolean) Is Object Storage System Bucket enabled (true | false)
- `obs_bucket_id_list` (List of String) Object Storage Bucket ID List
- `obs_bucket_name` (String) Object Storage Bucket Name (Like)
- `obs_bucket_name_exact` (String) Object Storage Bucket Name (Equal)
- `obs_bucket_query_end_dt` (String) Object Storage Bucket Query End Date
- `obs_bucket_query_start_dt` (String) Object Storage Bucket Query Start Date
- `obs_bucket_state` (String) Object Storage Bucket State
- `obs_bucket_state_in` (List of String) Object Storage Bucket State List
- `obs_bucket_used_type` (List of String) Object Storage Bucket Used Type
- `obs_quota_id` (String) Object Storage Quota ID
- `obs_storage_id` (String) Object Storage Bucket Name
- `page` (Number) Page start number from which to get the list
- `pool_region` (String) Region
- `size` (Number) Size to get list
- `sort` (List of String) Sort

### Read-Only

- `contents` (List of Object) Object Storage Bucket List (see [below for nested schema](#nestedatt--contents))
- `id` (String) The ID of this resource.
- `total_count` (Number) Total List Size

<a id="nestedatt--contents"></a>
### Nested Schema for `contents`

Read-Only:

- `created_by` (String)
- `created_dt` (String)
- `is_obs_bucket_dr_enabled` (Boolean)
- `is_obs_bucket_ip_address_filter_enabled` (Boolean)
- `is_obs_object_creation_enabled` (Boolean)
- `is_obs_system_bucket_enabled` (Boolean)
- `modified_by` (String)
- `modified_dt` (String)
- `obs_bucket_dr_type` (String)
- `obs_bucket_id` (String)
- `obs_bucket_name` (String)
- `obs_bucket_state` (String)
- `obs_bucket_used_type` (String)
- `obs_id` (String)
- `obs_name` (String)
- `obs_quota_id` (String)
- `obs_quota_name` (String)
- `obs_tenant_name` (String)
- `pool_region` (String)
- `project_id` (String)
- `system_id` (String)
- `system_name` (String)
- `zone_name` (String)

