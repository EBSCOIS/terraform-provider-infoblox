terraform {
  required_providers {
    infoblox = {
      source = "wstrydom-ebsco/infoblox"
    }
  }
}

provider "infoblox" {
  server   = "34.206.73.135"
  username = "admin"
}

// create containers from list of cidrs
resource "infoblox_ipv4_network_container" "this" {
  for_each = toset(split(", ", data.infoblox_ipv4_network_containers.this.network))
  network_view = "sheepstest"
  cidr = each.value
}

// Get container name to pass below
data "infoblox_network_view" "netview" {
  name = "vdiscovery"
}

output "netview" {
  value = data.infoblox_network_view.netview
}

// Get network view from ea search
data "infoblox_ipv4_network_containers" "this" {
  ea_search = {
    # "Tenant ID" = "514397636173" // for go
    "Tenant ID" = "454362366576" // for vdiscovery
    # "AWS_REGION" = "us-east-1" // for sheepstest
    # "Subnet Name" = "demo2-subnet-private1-us-east-1a" // for vdiscovery
  }
  network_view = data.infoblox_network_view.netview.name
}

output "this" {
  value = split(", ", data.infoblox_ipv4_network_containers.this.network)
}

// Create container
