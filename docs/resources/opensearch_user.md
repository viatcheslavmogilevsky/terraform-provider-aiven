---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "aiven_opensearch_user Resource - terraform-provider-aiven"
subcategory: ""
description: |-
  The OpenSearch User resource allows the creation and management of Aiven OpenSearch Users.
---

# aiven_opensearch_user (Resource)

The OpenSearch User resource allows the creation and management of Aiven OpenSearch Users.

## Example Usage

```terraform
resource "aiven_opensearch_user" "foo" {
  service_name = aiven_opensearch.bar.service_name
  project      = "my-project"
  username     = "user-1"
  password     = "Test$1234"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `project` (String) Identifies the project this resource belongs to. To set up proper dependencies please refer to this variable as a reference. Changing this property forces recreation of the resource.
- `service_name` (String) Specifies the name of the service that this resource belongs to. To set up proper dependencies please refer to this variable as a reference. Changing this property forces recreation of the resource.
- `username` (String) The actual name of the OpenSearch User. To set up proper dependencies please refer to this variable as a reference. Changing this property forces recreation of the resource.

### Optional

- `password` (String, Sensitive) The password of the OpenSearch User.
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The ID of this resource.
- `type` (String) Type of the user account. Tells whether the user is the primary account or a regular account.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `default` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import aiven_opensearch_user.foo project/service_name/username
```
