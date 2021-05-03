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

data "illumio-core_traffic_collector_settings" "example" {
  traffic_collector_setting_id = "9c186bde-27aa-495b-89ac-8401f62ffbe8"
}

output "output" {
  value = data.illumio-core_traffic_collector_settings.name
}

resource "illumio-core_traffic_collector_settings" "example" {
  action       = "drop"
  transmission = "broadcast"
  target {
    dst_ip   = "1.1.1.2"
    dst_port = -1
    proto    = 6
  }
}