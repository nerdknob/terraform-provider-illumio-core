---
layout: "illumio-core"
page_title: "illumio-core_ven Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-ven"
subcategory: ""
description: |-
  Represents Illumio VEN
---

# illumio-core_ven (Data Source)

Represents Illumio VEN

Example Usage
------------
```hcl
data "illumio-core_ven" "ven" {
  href = "/orgs/1/vens/80ec4e0a-e628-41c2-b79a-866f72a6b070"
}
```

## Schema

### Required

- **href** (String) URI of VEN

### Read-Only

- **activation_type** (String) The method by which the VEN was activated
- **active_pce_fqdn** (String) The FQDN of the PCE that the VEN last connected to
- **caps** (List of String) Permission types
- **conditions** (List of Object) (see [below for nested schema](#nestedatt--conditions))
- **container_cluster** (List of Object) container_cluster details for ven. Single element list (see [below for nested schema](#nestedatt--container_cluster))
- **created_at** (String) The time (rfc3339 timestamp) at which this VEN was created
- **created_by** (Map of String) The href of the user who created this VEN
- **description** (String) The description of the VEN
- **hostname** (String) The hostname of the host managed by the VEN
- **interfaces** (List of Object) Network interfaces of the host managed by the VEN (see [below for nested schema](#nestedatt--interfaces))
- **labels** (List of Object) Labels assigned to the host managed by the VEN (see [below for nested schema](#nestedatt--labels))
- **last_goodbye_at** (String) The time (rfc3339 timestamp) of the last goodbye from the VEN
- **last_heartbeat_at** (String) The last time (rfc3339 timestamp) a heartbeat was received from this VEN
- **name** (String) Friendly name for the VEN
- **os_detail** (String) Additional OS details from the host managed by the VEN
- **os_id** (String) OS identifier of the host managed by the VEN
- **os_platform** (String) OS platform of the host managed by the VEN
- **secure_connect** (List of Object) secure_connect details for vens (see [below for nested schema](#nestedatt--secure_connect))
- **status** (String) Status of the VEN
- **target_pce_fqdn** (String) The FQDN of the PCE that the VEN will use for future connections
- **uid** (String) The unique ID of the host managed by the VEN
- **unpair_allowed** (Boolean)
- **updated_at** (String) The time (rfc3339 timestamp) at which this VEN was last updated
- **updated_by** (Map of String) The href of the user who last updated this VEN
- **version** (String) Software version of the VEN
- **workloads** (List of Object) collection of Workloads (see [below for nested schema](#nestedatt--workloads))

<a id="nestedatt--conditions"></a>
### Nested Schema for `conditions`

Read-Only:

- **first_reported_timestamp** (String) The timestamp of the first event that reported this condition
- **latest_event** (List of Object) The latest notification event that was generated for the corresponding condition. Single element list (see [below for nested schema](#nestedobjatt--conditions--latest_event))

<a id="nestedobjatt--conditions--latest_event"></a>
### Nested Schema for `conditions.latest_event`

Read-Only:

- **href** (String) The href of the event
- **info** (List of Object) (see [below for nested schema](#nestedobjatt--conditions--latest_event--info))
- **notification_type** (String) The notification_type of the event
- **severity** (String) Severity of the condition, same as the event
- **timestamp** (String) RFC 3339 timestamp at which this event was created

<a id="nestedobjatt--conditions--latest_event--info"></a>
### Nested Schema for `conditions.latest_event.info`

Read-Only:

- **agent** (Map of String) Agent info




<a id="nestedatt--container_cluster"></a>
### Nested Schema for `container_cluster`

Read-Only:

- **href** (String) The URI of the container cluster managed by this VEN
- **name** (String) The name of the container cluster managed by this VEN, only present in expanded representations


<a id="nestedatt--interfaces"></a>
### Nested Schema for `interfaces`

Read-Only:

- **address** (String) The IP Address to assign to this interface
- **cidr_block** (Number) The number of bits in the subnet
- **default_gateway_address** (String) The IP Address of the default gateway
- **friendly_name** (String) User-friendly name for interface
- **href** (String) Interface URI
- **link_state** (String) Link State
- **loopback** (Boolean) loopback for interface
- **name** (String) Interface name
- **network** (Map of String) Network that the interface belongs to
- **network_detection_mode** (String) Network Detection Mode


<a id="nestedatt--labels"></a>
### Nested Schema for `labels`

Read-Only:

- **href** (String) Label URI
- **key** (String) Key of the label
- **value** (String) Value of the label


<a id="nestedatt--secure_connect"></a>
### Nested Schema for `secure_connect`

Read-Only:

- **matching_issuer_name** (String) Issuer name match criteria for certificate used during establishing secure connections


<a id="nestedatt--workloads"></a>
### Nested Schema for `workloads`

Read-Only:

- **enforcement_mode** (String) Policy enforcement mode
- **hostname** (String) The hostname of this workload
- **href** (String) URI of the Workload
- **interfaces** (List of Object) Network interfaces of the workload (see [below for nested schema](#nestedobjatt--workloads--interfaces))
- **labels** (List of Object) Labels assigned to the host managed by the VEN (see [below for nested schema](#nestedobjatt--workloads--labels))
- **mode** (String) Policy enforcement mode
- **name** (String) The short friendly name of the workload
- **online** (Boolean) If this workload is online and present in policy
- **os_detail** (String) Additional OS details
- **os_id** (String) OS identifier for the workload
- **public_ip** (String) The public IP address of the server
- **security_policy_applied_at** (String) Last reported time when policy was applied to the workload (UTC)
- **security_policy_received_at** (String) Last reported time when policy was received by the workload (UTC)
- **visibility_level** (String) Visibility level of the workload

<a id="nestedobjatt--workloads--interfaces"></a>
### Nested Schema for `workloads.interfaces`

Read-Only:


- **address** (String) The IP Address to assign to this interface
- **cidr_block** (Number) The number of bits in the subnet
- **default_gateway_address** (String) The IP Address of the default gateway
- **friendly_name** (String) User-friendly name for interface
- **link_state** (String) Link State
- **loopback** (Boolean) loopback for interface
- **name** (String) Interface name
- **network** (Map of String) Network that the interface belongs to
- **network_detection_mode** (String) Network Detection Mode


<a id="nestedobjatt--workloads--labels"></a>
### Nested Schema for `workloads.labels`

Read-Only:

- **href** (String) Label URI
- **key** (String) Key of the label
- **value** (String) Value of the label


