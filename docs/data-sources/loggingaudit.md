---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scp_loggingaudit Data Source - scp"
subcategory: ""
description: |-
  Provides detailed logging information for a given logging id
---

# scp_loggingaudit (Data Source)

Provides detailed logging information for a given logging id

## Example Usage

```terraform
data "scp_loggingaudits" "logs" {
}

data "scp_loggingaudit" "mylog" {
  logging_id = data.scp_loggingaudits.logs.contents[0].id
}

output "result_scp_loggingaudits" {
  value = data.scp_loggingaudit.mylog
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `logging_id` (String) Logging ID

### Read-Only

- `audit_content` (String) Audit content
- `audit_detail_content` (String) Audit details
- `cluster_id` (String) Cluster ID
- `cluster_namespace_id` (String) Cluster namespace ID
- `event_topic_name` (String) Event topic name
- `id` (String) Logging ID
- `log_error_message` (String) Log error message
- `object_id` (String) Object ID
- `object_name` (String) Object name
- `product_name` (String) Product name
- `project_id` (String) Project ID
- `project_name` (String) Project name
- `region` (String) Region
- `request_client_type` (String) requesting client type
- `request_dt` (String) Request date
- `requested_by` (String) Requester
- `resource_type` (String) Resource type
- `state` (String) Job's result state
- `user_email` (String) User email
- `user_name` (String) User name

