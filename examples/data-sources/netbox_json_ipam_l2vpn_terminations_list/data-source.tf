data "netbox_json_ipam_l2vpn_terminations_list" "test" {
  limit = 0
}

output "example" {
  value = jsondecode(data.netbox_json_ipam_l2vpn_terminations_list.test.json)
}
