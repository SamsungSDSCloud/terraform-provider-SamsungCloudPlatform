---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scp_key_pairs Data Source - scp"
subcategory: ""
description: |-
  Provides list of key pairs
---

# scp_key_pairs (Data Source)

Provides list of key pairs



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `created_by` (String) The person who created the resource
- `key_pair_id` (String) Key Pair Id
- `key_pair_name` (String) Key Pair Name
- `page` (Number) Page start number from which to get the list
- `size` (Number) Size to get list
- `sort` (String) Sorting

### Read-Only

- `contents` (Block List) Key Pair list (see [below for nested schema](#nestedblock--contents))
- `id` (String) The ID of this resource.
- `total_count` (Number) Total list size

<a id="nestedblock--contents"></a>
### Nested Schema for `contents`

Read-Only:

- `created_by` (String) Person who created the resource
- `created_dt` (String) Creation time
- `key_pair_id` (String) Key Pair Id
- `key_pair_name` (String) Key Pair Name
- `key_pair_state` (String) Key Pair State
- `modified_by` (String) Person who modified the resource
- `modified_dt` (String) Modification time
- `project_id` (String) Project id
- `virtual_server_id_list` (List of String) Virtual Server Id List

