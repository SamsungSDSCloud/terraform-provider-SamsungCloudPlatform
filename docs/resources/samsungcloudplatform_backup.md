---
page_title: "samsungcloudplatform_backup Resource - scp"
subcategory: ""
description: |-
  Provides a Backup resource.
---

# Resource: samsungcloudplatform_backup

Provides a Backup resource.


## Example Usage

```terraform
resource "samsungcloudplatform_backup" "my_scp_backup" {
  backup_name = var.name
  backup_policy_type_category = "VM"
  backup_repository = "SD_STORAGE"
  is_backup_dr_enabled = "N"
  object_id = "INSTANCE-XXXXX"
  object_type = "INSTANCE"
  policy_type = "VMsnapshot"
  product_names = [
    "VM Image"
  ]
  retention_period = "4W"
  service_zone_id = "ZONE-XXXXX"

  dynamic "schedules" {
    for_each = var.schedules
    content {
      schedule_frequency = schedules.value["schedule_frequency"]
      schedule_frequency_detail = schedules.value["schedule_frequency_detail"]
      schedule_type = schedules.value["schedule_type"]
      start_time = schedules.value["start_time"]
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `backup_name` (String) Backup Name
- `backup_policy_type_category` (String) Backup Policy Type Category (VM, FILESYSTEM)
- `backup_repository` (String) Backup Repository (SD_STORAGE)
- `object_id` (String) Backup Object ID
- `object_type` (String) Backup Object Type
- `policy_type` (String) Backup Policy Type
- `product_names` (List of String) Product Names
- `schedules` (Block List, Min: 1) Backup Schedules (see [below for nested schema](#nestedblock--schedules))
- `service_zone_id` (String) Service Zone ID

### Optional

- `az_code` (String) Multi AZ Code
- `backup_dr_zone_id` (String) Backup(DR) Service Zone Id
- `dr_az_code` (String) Multi AZ(DR) Code
- `incremental_retention_period` (String) Incremental Backup Retention Period
- `is_backup_dr_destroy_enabled` (Boolean) IF 'Y', Destroy DR replica together.
- `is_backup_dr_enabled` (String) Backup(DR) Activation (If 'Y', Backup(DR) will be activated)
- `retention_period` (String) Full Backup Retention Period
- `tags` (Map of String)

### Read-Only

- `backup_dr_id` (String) Backup DR ID
- `id` (String) The ID of this resource.
- `is_backup_dr_deleted` (String) Is Backup DR Deleted.

<a id="nestedblock--schedules"></a>
### Nested Schema for `schedules`

Required:

- `schedule_frequency` (String) Backup Schedule Frequency (MONTHLY, WEEKLY, DAYS)
- `schedule_frequency_detail` (String) Backup Schedule Frequency details
- `schedule_type` (String) Backup Schedule Type (FULL, INCREMENTAL)
- `start_time` (String) Backup Start Time (format:HH:mm±hh:mm)

Read-Only:

- `schedule_id` (String) Backup Schedule ID
- `schedule_name` (String) Backup Schedule Name