terraform {
  required_providers {
    powerstore = {
      version = "0.0.1"
      source = "terraform.example.com/local/powerstore"
    }
  }
} 

provider "powerstore" {
  username = "Username"
  password = "Password"
  host = "https://<host_url>"
}

variable "volume" {}

data "powerstore_host" "map_hosts" {
  for_each = var.volume
  name = each.value["hosts"]
}

/* output "test" {
  value = data.powerstore_host.map_hosts.id
} */


resource "powerstore_volume" "test" {
  # (resource arguments)
  for_each = var.volume
  name = each.key
  size = each.value["size"]
  host_id = data.powerstore_host.map_hosts[each.key].id
      #appliance_id = "A1"
      #creation_timestamp = "2020-11-17T01:29:56.356792+00:00"
      #is_replication_destination = false
      #node_affinity = "System_Selected_Node_A"
      #node_affinity_l10n = "System Selected Node A"
      #nsid = 0
      #state  = "Ready"
      #state_l10n = "Ready"
      #type = "Primary"
      #type_l10n   = "Primary"
      #wwn = "naa.68ccf09800f9ecd28abfa5b729f25466"
      #is_app_consistent = false
      #creator_type = "User"
      #creator_type_l10n = "User"
      #family_id = "97036821-b960-452b-9655-dbcb2761c818"

}
