terraform {
  required_providers {
    illumio-core = {
      version = "0.1.0"
      source  = "illumio/illumio-core"
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

data "illumio-core_security_rule" "example" {
  href = "/orgs/1/sec_policy/draft/rule_sets/6/sec_rules/7"
}

resource "illumio-core_security_rule" "example" {
  rule_set_href = "/orgs/1/sec_policy/draft/rule_sets/6"

  enabled = true

  resolve_labels_as {
    consumers = ["workloads"]
    providers = ["virtual_services"]
  }

  consumers {
    actors = "ams"
  }

  providers {
    label {
      href = "/orgs/1/labels/715"
    }
  }

  ingress_services {
    href = "/orgs/1/sec_policy/draft/services/19"
  }

  ingress_services {
    proto = 6
    port  = 1
  }

  ingress_services {
    proto   = 6
    port    = 1
    to_port = 12
  }
}
