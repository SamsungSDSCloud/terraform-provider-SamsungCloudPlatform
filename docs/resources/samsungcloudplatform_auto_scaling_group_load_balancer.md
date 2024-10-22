---
page_title: "samsungcloudplatform_auto_scaling_group_load_balancer Resource - scp"
subcategory: ""
description: |-
  Provides LB Service resource connected to Auto-Scaling Group.
---

# Resource: samsungcloudplatform_auto_scaling_group_load_balancer

Provides LB Service resource connected to Auto-Scaling Group.


## Example Usage

```terraform
resource "samsungcloudplatform_auto_scaling_group_load_balancer" "my_auto_scaling_group_load_balancer" {
  asg_id = data.terraform_remote_state.auto_scaling_group.outputs.id
  lb_rule_ids = var.lb_rule_ids
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `asg_id` (String) Auto-Scaling Group ID
- `lb_rule_ids` (Set of String) LB rule ID list connected to Auto-Scaling Group

### Read-Only

- `id` (String) The ID of this resource.