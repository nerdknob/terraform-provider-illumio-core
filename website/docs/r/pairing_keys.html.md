---
layout: "illumio-core"
page_title: "illumio-core_pairing_keys Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-pairing-keys"
subcategory: ""
description: |-
  Manages Illumio Pairing Keys
---


# illumio-core_pairing_keys (Resource)

Manages Illumio Pairing Keys


Example Usage
------------

This resource encrypts generated activation token using AES GCM algorithm before saving it to state. For this purpose, environment variable `ILLUMIO_AES_GCM_KEY` needs to be set with valid key.

**Note: On destroying the resource, all (including ones that are not generated by this resource) the pairing keys will deleted associated with pairing profile.**

```hcl
resource "illumio-core_pairing_keys" "example" {
    pairing_profile_href = "/orgs/1/pairing_profiles/1"
    token_count = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **pairing_profile_href** (String) Href of pairing profile
- **token_count** (Number) Count of token to generate/maintain. It can be accessed in `activation_tokens` On increasing the count, new activation tokens will be generated. On decreasing the count `activation_tokens` list will shrink to that size and extra keys will be discarded. Minimum: 1

### Read-Only

- **activation_tokens** (List of Object) List of activation tokens (see [below for nested schema](#nestedatt--activation_tokens))

<a id="nestedatt--activation_tokens"></a>
### Nested Schema for `activation_tokens`

Read-Only:

- **activation_token** (String) Encrypted activation token (encrypted by key set in env. variable `ILLUMIO_AES_GCM_KEY`)
- **nonce** (String) Nonce used in encrypting activation token

