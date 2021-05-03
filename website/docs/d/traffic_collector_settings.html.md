---
layout: "illumio-core"
page_title: "illumio-core_traffic_collector_settings Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-traffic-collector-settings"
subcategory: ""
description: |-
  Represents Illumio Traffic Collector Settings
---

# illumio-core_traffic_collector_settings (Data Source)

Represents Illumio Traffic Collector Settings

Example Usage
------------

```hcl
data "illumio-core_traffic_collector_settings" "example" {
  traffic_collector_setting_id = "9c186bde-27aa-495b-89ac-8401f62ffbe8"
}
```

## Schema

### Required

- **traffic_collector_setting_id** (String) Traffic Collector Settings ID

### Read-Only

- **action** (String) action for target traffic
- **href** (String) URI of traffic collecter settings
- **target** (List of Object) target for traffic collector settings (see [below for nested schema](#nestedatt--target))
- **transmission** (String) transmission type

<a id="nestedatt--target"></a>
### Nested Schema for `target`

Read-Only:

- **dst_ip** (String) single ip address or CIDR
- **dst_port** (Number) destination port for target
- **proto** (Number) protocol for target

