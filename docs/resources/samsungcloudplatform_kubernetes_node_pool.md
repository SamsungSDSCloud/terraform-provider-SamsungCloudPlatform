---
page_title: "samsungcloudplatform_kubernetes_node_pool Resource - scp"
subcategory: ""
description: |-
  Provides a K8s Node Pool resource.
---

# Resource: samsungcloudplatform_kubernetes_node_pool

Provides a K8s Node Pool resource.


## Example Usage

```terraform
data "samsungcloudplatform_region" "region" {
}

data "samsungcloudplatform_standard_image" "ubuntu_image" {
  service_group = "CONTAINER"
  service       = "Kubernetes Engine VM"
  region        = data.samsungcloudplatform_region.region.location

  filter {
    name      = "image_name"
    values    = ["Ubuntu 18.04 (Kubernetes)-v1.24.8"]
    use_regex = false
  }
}

resource "samsungcloudplatform_kubernetes_node_pool" "pool" {
  name               = var.name
  engine_id          = data.terraform_remote_state.engine.outputs.id
  image_id           = data.samsungcloudplatform_standard_image.ubuntu_image.id
  desired_node_count = 2

  scale_name = var.scale_name
  storage_name = var.storage_name
  storage_size_gb = var.storage_size_gb

  // optional field
  availability_zone_name = null

  // update optional field
  auto_scale         = false
  min_node_count     = null
  max_node_count     = null
  auto_recovery      = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `engine_id` (String) ID of scp_kubernetes_engine resource
- `image_id` (String) Image ID (use scp_standard_image data source)
- `name` (String) Node pool name

### Optional

- `auto_recovery` (Boolean) Enable auto recovery
- `auto_scale` (Boolean) Enable auto scale
- `availability_zone_name` (String) Availability zone name.
- `desired_node_count` (Number) Desired node count in the pool (Desired node count is 0 when auto_scale is enabled)
- `encrypt_enabled` (Boolean) Encrypt enabled
- `max_node_count` (Number) Maximum node count
- `min_node_count` (Number) Minimum node count
- `scale_name` (String) Scale name
- `storage_name` (String) Storage name (Currently only SSD is supported)
- `storage_size_gb` (String) Storage size in GB (default 100)
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `update` (String)