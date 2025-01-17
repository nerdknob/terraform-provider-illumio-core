---
layout: "illumio-core"
page_title: "illumio-core_firewall_settings Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-firewall-settings"
subcategory: ""
description: |-
  Manages Illumio Firewall Settings
---

# illumio-core_firewall_settings (Resource)

Manages Illumio Firewall Settings


Example Usage
------------

```hcl
# INFO: cherry-picked attributes from terraform show after import
resource "illumio-core_firewall_settings" "example" {
    ike_authentication_type = "psk"

    blocked_connection_reject_scopes {
        label {
            href = "/orgs/1/labels/1"
        }
        label {
            href = "/orgs/1/labels/12"
        }
        label {
            href = "/orgs/1/labels/14"
        }
        label {
            href = "/orgs/1/labels/787"
        }
    }

    containers_inherit_host_policy_scopes {
        label {
            href = "/orgs/1/labels/1"
        }
        label {
            href = "/orgs/1/labels/11"
        }
        label {
            href = "/orgs/1/labels/14"
        }
        label {
            href = "/orgs/1/labels/787"
        }
    }

    firewall_coexistence {
        illumio_primary = true

        scope {
            href = "/orgs/1/labels/1"
        }
        scope {
            href = "/orgs/1/labels/787"
        }
        scope {
            href = "/orgs/1/labels/788"
        }
        scope {
            href = "/orgs/1/labels/11"
        }

        workload_mode = "enforced"
    }

    loopback_interfaces_in_policy_scopes {
        label {
            href = "/orgs/1/labels/1"
        }
        label {
            href = "/orgs/1/labels/12"
        }
        label {
            href = "/orgs/1/labels/787"
        }
        label {
            href = "/orgs/1/labels/788"
        }
    }

    static_policy_scopes {
        label {
            href = "/orgs/1/labels/1"
        }
        label {
            href = "/orgs/1/labels/12"
        }
        label {
            href = "/orgs/1/labels/14"
        }
        label_group {
            href = "/orgs/1/sec_policy/draft/label_groups/a715cd8f-04f3-4bc1-82bf-d650b01453a5"
        }
    }
}

```

## Schema

### Required

- **ike_authentication_type** (String) IKE authentication type to use for IPsec (SecureConnect and Machine Authentication). Allowed values are "psk" and "certificate"

### Optional

- **blocked_connection_reject_scopes** (Block List) scopes for reject connections. Either label or label_group can be specified (see [below for nested schema](#nestedblock--blocked_connection_reject_scopes))
- **containers_inherit_host_policy_scopes** (Block List) scopes for container inherit host policy. Either label or label_group can be specified (see [below for nested schema](#nestedblock--containers_inherit_host_policy_scopes))
- **firewall_coexistence** (Block List) Firewall coexistence configuration (see [below for nested schema](#nestedblock--firewall_coexistence))
- **loopback_interfaces_in_policy_scopes** (Block List) scopes for loopback interface. Either label or label_group can be specified (see [below for nested schema](#nestedblock--loopback_interfaces_in_policy_scopes))
- **static_policy_scopes** (Block List) scopes for static policy. Either label or label_group can be specified (see [below for nested schema](#nestedblock--static_policy_scopes))

### Read-Only

- **created_at** (String) Timestamp when these firewall settings were first created
- **created_by** (Map of String) User who created this resource
- **deleted_at** (String) Timestamp when these firewall settings were deleted
- **deleted_by** (Map of String) User who deleted this resource
- **href** (String) URI of Firewall Settings
- **update_type** (String) Type of Update
- **updated_at** (String) Timestamp when these firewall settings were last updated
- **updated_by** (Map of String) User who last updated this resource

<a id="nestedblock--blocked_connection_reject_scopes"></a>
### Nested Schema for `blocked_connection_reject_scopes`

Optional:

- **label** (Block Set) Href of Label (see [below for nested schema](#nestedblock--blocked_connection_reject_scopes--label))
- **label_group** (Block Set) Href of Label Group (see [below for nested schema](#nestedblock--blocked_connection_reject_scopes--label_group))

<a id="nestedblock--blocked_connection_reject_scopes--label"></a>
### Nested Schema for `blocked_connection_reject_scopes.label`

Required:

- **href** (String) URI of Label


<a id="nestedblock--blocked_connection_reject_scopes--label_group"></a>
### Nested Schema for `blocked_connection_reject_scopes.label_group`

Required:

- **href** (String) URI of Label Group



<a id="nestedblock--containers_inherit_host_policy_scopes"></a>
### Nested Schema for `containers_inherit_host_policy_scopes`

Optional:

- **label** (Block Set) Href of Label (see [below for nested schema](#nestedblock--containers_inherit_host_policy_scopes--label))
- **label_group** (Block Set) Href of Label Group (see [below for nested schema](#nestedblock--containers_inherit_host_policy_scopes--label_group))

<a id="nestedblock--containers_inherit_host_policy_scopes--label"></a>
### Nested Schema for `containers_inherit_host_policy_scopes.label`

Required:

- **href** (String) URI of Label


<a id="nestedblock--containers_inherit_host_policy_scopes--label_group"></a>
### Nested Schema for `containers_inherit_host_policy_scopes.label_group`

Required:

- **href** (String) URI of Label Group



<a id="nestedblock--firewall_coexistence"></a>
### Nested Schema for `firewall_coexistence`

Required:

- **illumio_primary** (Boolean) Whether Illumio is primary firewall or not
- **scope** (Block List, Min: 1) List of Href of label (see [below for nested schema](#nestedblock--firewall_coexistence--scope))

Optional:

- **workload_mode** (String) Match criteria to select workload(s). Allowed values are "enforced" and "illuminated"

<a id="nestedblock--firewall_coexistence--scope"></a>
### Nested Schema for `firewall_coexistence.scope`

Required:

- **href** (String) Href of Label



<a id="nestedblock--loopback_interfaces_in_policy_scopes"></a>
### Nested Schema for `loopback_interfaces_in_policy_scopes`

Optional:

- **label** (Block Set) Href of Label (see [below for nested schema](#nestedblock--loopback_interfaces_in_policy_scopes--label))
- **label_group** (Block Set) Href of Label Group (see [below for nested schema](#nestedblock--loopback_interfaces_in_policy_scopes--label_group))

<a id="nestedblock--loopback_interfaces_in_policy_scopes--label"></a>
### Nested Schema for `loopback_interfaces_in_policy_scopes.label`

Required:

- **href** (String) URI of Label


<a id="nestedblock--loopback_interfaces_in_policy_scopes--label_group"></a>
### Nested Schema for `loopback_interfaces_in_policy_scopes.label_group`

Required:

- **href** (String) URI of Label Group



<a id="nestedblock--static_policy_scopes"></a>
### Nested Schema for `static_policy_scopes`

Optional:

- **label** (Block Set) Href of Label (see [below for nested schema](#nestedblock--static_policy_scopes--label))
- **label_group** (Block Set) Href of Label Group (see [below for nested schema](#nestedblock--static_policy_scopes--label_group))

<a id="nestedblock--static_policy_scopes--label"></a>
### Nested Schema for `static_policy_scopes.label`

Required:

- **href** (String) URI of Label


<a id="nestedblock--static_policy_scopes--label_group"></a>
### Nested Schema for `static_policy_scopes.label_group`

Required:

- **href** (String) URI of Label Group


## Importing ##

This resource can only be imported and can not be created. Use the below command to import resource. This resource auto determines URI based on provider config. So no need of providing URI while importing.

After importing, Cherry-pick the configurable parameters from `terraform show` and paste it into .tf file.

Ref: https://www.terraform.io/docs/import/index.html


```
terraform import illumio-core_firewall_settings.example <ANYTHING>
```
