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

resource "illumio-core_vens_unpair" "example" {
  vens {
    href = "/orgs/1/workloads/de7c1705-6e9d-46e7-b3b1-5ef5a638c0f8"
  }
  vens {
    href = "/orgs/1/workloads/32254e8b-eddc-428d-96fa-f8625416a0d6"
  }
  vens {
    href = "/orgs/1/workloads/11416eb7-43df-4acc-a4b0-c17c1e2b1b77"
  }
}
