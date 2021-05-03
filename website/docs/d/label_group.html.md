---
layout: "illumio-core"
page_title: "illumio-core_label_group Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-label-group"
subcategory: ""
description: |-
  Represents Illumio Label Group
---

# illumio-core_label_group (Data Source)

Represents Illumio Label Group


Example Usage
------------

```hcl
data "illumio-core_label_group" "lg1"{
  label_group_id = "731fca37-7f57-4852-96ac-753ddfd359d3"
}
```



## Schema

### Required

- **label_group_id** (String) UUID of Label group

### Read-Only

- **created_at** (String) Timestamp when this Label Group was first created
- **created_by** (Map of String) User who originally created this Label Group
- **deleted_at** (String) Timestamp when this Label Group was last deleted
- **deleted_by** (Map of String) User who deleted this Label Group
- **description** (String) The long description of the Label Group
- **external_data_reference** (String) External Data reference identifier
- **external_data_set** (String) External Data set Identifier
- **href** (String) URI of Label Group
- **key** (String) Key in key-value pair of contained labels or Label Groups
- **labels** (Set of Object) Contained labels (see [below for nested schema](#nestedatt--labels))
- **name** (String) Name of the Label Group
- **sub_groups** (Set of Object) Contained Label Groups (see [below for nested schema](#nestedatt--sub_groups))
- **update_type** (String) Type of Update
- **updated_at** (String) Timestamp when this Label Group was last updated
- **updated_by** (Map of String) User who last updated this Label Group

<a id="nestedblock--labels"></a>
### Nested Schema for `labels`

Read-Only:

- **href** (String) URI of label
- **key** (String) Label Key same as label group key
- **value** (String) Label Value in key-value pair


<a id="nestedblock--sub_groups"></a>
### Nested Schema for `sub_groups`

Read-Only:

- **href** (String) URI of Label Group
- **name** (String) name of sub Label Group