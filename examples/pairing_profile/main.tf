terraform {
  required_providers {
    illumio-core = {
      version = "0.1"
      source  = "illumio.com/labs/illumio-core"
    }
  }
}

provider "illumio-core" {
  //  pce_host              = "https://pce.my-company.com:8443"
  //  api_username          = "api_xxxxxx"
  //  api_secret            = "big-secret"
  request_timeout = 30
  org_id          = 1
}

data "illumio-core_pairing_profile" "example" {
  pairing_profile_id = 8
}

resource "illumio-core_pairing_profile" "example" {
  name    = "example1234"
  enabled = false
  label {
    href = "/orgs/1/labels/1"
  }
  label {
    href = "/orgs/1/labels/7"
  }
  allowed_uses_per_key  = "50"
  key_lifespan          = "50"
  env_label_lock        = false
  loc_label_lock        = true
  role_label_lock       = true
  app_label_lock        = true
  log_traffic           = false
  log_traffic_lock      = true
  visibility_level      = "flow_off"
  visibility_level_lock = false
}

resource "illumio-core_pairing_profile" "example1" {
  name    = "example123"
  enabled = true
  label {
    href = "/orgs/1/labels/4"
  }
  env_label_lock        = true
  loc_label_lock        = true
  role_label_lock       = true
  app_label_lock        = false
  visibility_level      = "flow_summary"
  visibility_level_lock = true
}