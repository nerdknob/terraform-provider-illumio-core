---
layout: "illumio-core"
page_title: "illumio-core_container_clusters Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-container-clusters"
subcategory: ""
description: |-
  Represents Illumio Container Clusters
---

# illumio-core_container_clusters (Data Source)

Represents Illumio Container Clusters

Example Usage
------------

```hcl
data "illumio-core_container_clusters" "example" {
    max_results = "5"
}
```

## Schema

### Optional

- **max_results** (String) Maximum number of container clusters to return. The integer should be a non-zero positive integer
- **name** (String) Name of the container cluster(s) to return. Supports partial matches

### Read-Only

- **items** (List of Object) List of Container Clusters (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- **caps** (List of String) Permission types
- **container_runtime** (String) The Container Runtime used in this Cluster
- **description** (String) Description of the Cluster
- **errors** (List of Object) Errors for Cluster (see [below for nested schema](#nestedobjatt--items--errors))
- **href** (String) Href of Container Cluster
- **kubelink_version** (String) Kubelink software version string for Cluster
- **last_connected** (String) Time the Cluster last connected to
- **manager_type** (String) Manager for this Cluster (and version)
- **name** (String) Name of the Cluster
- **nodes** (List of Object) Nodes of the Cluster (see [below for nested schema](#nestedobjatt--items--nodes))
- **online** (Boolean) Whether the Cluster is online or not
- **pce_fqdn** (String) PCE FQDN for this container cluster. Used in Supercluster only

<a id="nestedobjatt--items--errors"></a>
### Nested Schema for `items.errors`

Read-Only:

- **audit_event** (Map of String) Audit Event of Error
- **duplicate_ids** (List of String) Duplicate IDs of Error
- **error_type** (String) Error Type of Error


<a id="nestedobjatt--items--nodes"></a>
### Nested Schema for `items.nodes`

Read-Only:

- **pod_subnet** (String) Pod Subnet of the node

