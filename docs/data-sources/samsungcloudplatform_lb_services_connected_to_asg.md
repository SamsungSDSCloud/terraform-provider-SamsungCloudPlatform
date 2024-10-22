---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "samsungcloudplatform_lb_services_connected_to_asg Data Source - scp"
subcategory: ""
description: |-
  Provides list of load balancer service Connected to ASG
---

# samsungcloudplatform_lb_services_connected_to_asg (Data Source)

Provides list of load balancer service Connected to ASG

## Example Usage

```terraform
data "samsungcloudplatform_lb_services_connected_to_asg" "my_scp_lb_services_connected_to_asg" {
  auto_scaling_group_id = "AUTO_SCALING_GROUP-XXXXXXXXXXXXXXXXXXXXXX"
}

#Connected List
output "result_scp_lb_services_connected_to_asg" {
  value = data.samsungcloudplatform_lb_services_connected_to_asg.my_scp_lb_services_connected_to_asg
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `auto_scaling_group_id` (String) ASG ID

### Read-Only

- `contents` (List of Object) Load balancer service connected to asg list (see [below for nested schema](#nestedatt--contents))
- `id` (String) The ID of this resource.
- `total_count` (Number) Total list size

<a id="nestedatt--contents"></a>
### Nested Schema for `contents`

Read-Only:

- `auto_scale_group_ids` (List of String)
- `default_forwarding_ports` (String)
- `layer_type` (String)
- `lb_rules` (List of Object) (see [below for nested schema](#nestedobjatt--contents--lb_rules))
- `lb_service_id` (String)
- `lb_service_ip_address` (String)
- `lb_service_name` (String)
- `lb_service_state` (String)
- `load_balancer_id` (String)
- `load_balancer_name` (String)
- `nat_ip_address` (String)
- `persistence` (String)
- `protocol` (String)
- `service_ports` (String)

<a id="nestedobjatt--contents--lb_rules"></a>
### Nested Schema for `contents.lb_rules`

Read-Only:

- `auto_scale_group_ids` (List of String)
- `lb_rule_id` (String)
- `lb_server_group_id` (String)
- `pattern_url` (String)
- `seq` (Number)

