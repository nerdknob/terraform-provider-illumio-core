---
layout: "illumio-core"
page_title: "illumio-core_workload_settings Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-workload-settings"
subcategory: ""
description: |-
  Manages Illumio Workload Settings
---

# illumio-core_workload_settings (Resource)

Manages Illumio Workload Settings

Example Usage
------------

```hcl
# INFO: cherry-picked attributes from terraform show after import
resource "illumio-core_workload_settings" "test" {
  workload_disconnected_timeout_seconds {
    value = -1
  }
  workload_goodbye_timeout_seconds {
    value = -1
  }
}
```

## Schema

### Optional

- **workload_disconnected_timeout_seconds** (Block Set) Workload Disconnected Timeout Seconds for Workload Settings (see [below for nested schema](#nestedblock--workload_disconnected_timeout_seconds))
- **workload_goodbye_timeout_seconds** (Block Set) Workload Goodbye Timeout Seconds for Workload Settings (see [below for nested schema](#nestedblock--workload_goodbye_timeout_seconds))

<a id="nestedatt--workload_disconnected_timeout_seconds"></a>
### Nested Schema for `workload_disconnected_timeout_seconds`

Optional:

- **scope** (Block Set) Assigned labels for Workload Disconnected Timeout Seconds (see [below for nested schema](#nestedblock--workload_disconnected_timeout_seconds--scope))
- **value** (Number) Property value associated with the scope

<a id="nestedobjatt--workload_disconnected_timeout_seconds--scope"></a>
### Nested Schema for `workload_disconnected_timeout_seconds.scope`

Optional:

- **href** (String) Label URI

<a id="nestedatt--workload_goodbye_timeout_seconds"></a>

### Nested Schema for `workload_goodbye_timeout_seconds`

Optional:

- **scope** (Block Set) Assigned labels for Workload Goodbye Timeout Seconds (see [below for nested schema](#nestedblock--workload_goodbye_timeout_seconds--scope))
- **value** (Number) Property value associated with the scope

<a id="nestedobjatt--workload_goodbye_timeout_seconds--scope"></a>
### Nested Schema for `workload_goodbye_timeout_seconds.scope`

Optional:

- **href** (String) Label URI


## Importing ##

This resource can only be imported and can not be created. Use below command to import resource. Since this resource auto determines ID based on provider config. You can pass anything as identifier. 

After importing, Cherry pick the configurable parameters from `terraform show` and paste it into .tf file.

Ref: https://www.terraform.io/docs/import/index.html


```
terraform import illumio-core_workload_settings.example <ANYTHING>
```