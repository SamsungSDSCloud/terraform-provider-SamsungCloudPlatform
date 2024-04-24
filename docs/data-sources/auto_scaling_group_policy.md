---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scp_auto_scaling_group_policy Data Source - scp"
subcategory: ""
description: |-
  Provides details of Auto-Scaling Group policy
---

# scp_auto_scaling_group_policy (Data Source)

Provides details of Auto-Scaling Group policy

## Example Usage

```terraform
# Find details of Auto-Scaling Group policy
data "scp_auto_scaling_group_policy" "my_auto_scaling_group_policy" {
  asg_id = "AUTO_SCALING_GROUP-XXXXX"
  policy_id = "ASG_POLICY-XXXXX"
}

output "output_scp_auto_scaling_group_policy1" {
  value = data.scp_auto_scaling_group_policy.my_auto_scaling_group_policy
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `asg_id` (String) Auto-Scaling Group ID
- `policy_id` (String) Policy ID

### Read-Only

- `block_id` (String) Block ID
- `comparison_operator` (String) Comparison operator
- `cooldown_seconds` (Number) Cooldown seconds
- `created_by` (String) The person who created the resource
- `created_dt` (String) Creation date
- `evaluation_minutes` (Number) Evaluation minutes
- `id` (String) The ID of this resource.
- `metric_method` (String) Metric method
- `metric_type` (String) Metric type
- `modified_by` (String) The person who modified the resource
- `modified_dt` (String) Modification date
- `policy_name` (String) Policy name
- `policy_state` (String) Policy state
- `project_id` (String) Project ID
- `scale_method` (String) Scale method
- `scale_type` (String) Scale type
- `scale_value` (Number) Scale value
- `service_zone_id` (String) Service zone ID
- `threshold` (String) Threshold
- `threshold_unit` (String) Threshold unit

