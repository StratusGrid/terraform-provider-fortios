---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_centralmanagement"
description: |-
  Configure central management.
---

# fortios_system_centralmanagement
Configure central management.

## Example Usage

```hcl
resource "fortios_system_centralmanagement" "trname" {
  allow_monitor                 = "enable"
  allow_push_configuration      = "enable"
  allow_push_firmware           = "enable"
  allow_remote_firmware_upgrade = "enable"
  enc_algorithm                 = "high"
  fmg_source_ip                 = "0.0.0.0"
  fmg_source_ip6                = "::"
  include_default_servers       = "enable"
  mode                          = "normal"
  schedule_config_restore       = "enable"
  schedule_script_restore       = "enable"
  type                          = "fortimanager"
  vdom                          = "root"
}
```

## Argument Reference

The following arguments are supported:

* `mode` - Central management mode.
* `type` - Central management type.
* `schedule_config_restore` - Enable/disable allowing the central management server to restore the configuration of this FortiGate.
* `schedule_script_restore` - Enable/disable allowing the central management server to restore the scripts stored on this FortiGate.
* `allow_push_configuration` - Enable/disable allowing the central management server to push configuration changes to this FortiGate.
* `allow_push_firmware` - Enable/disable allowing the central management server to push firmware updates to this FortiGate.
* `allow_remote_firmware_upgrade` - Enable/disable remotely upgrading the firmware on this FortiGate from the central management server.
* `allow_monitor` - Enable/disable allowing the central management server to remotely monitor this FortiGate
* `serial_number` - Serial number.
* `fmg` - IP address or FQDN of the FortiManager.
* `fmg_source_ip` - IPv4 source address that this FortiGate uses when communicating with FortiManager.
* `fmg_source_ip6` - IPv6 source address that this FortiGate uses when communicating with FortiManager.
* `vdom` - Virtual domain (VDOM) name to use when communicating with FortiManager.
* `server_list` - Additional severs that the FortiGate can use for updates (for AV, IPS, updates) and ratings (for web filter and antispam ratings) servers.
* `include_default_servers` - Enable/disable inclusion of public FortiGuard servers in the override server list.
* `enc_algorithm` - Encryption strength for communications between the FortiGate and central management.

The `server_list` block supports:

* `id` - ID.
* `server_type` - FortiGuard service type.
* `addr_type` - Indicate whether the FortiGate communicates with the override server using an IPv4 address, an IPv6 address or a FQDN.
* `server_address` - IPv4 address of override server.
* `server_address6` - IPv6 address of override server.
* `fqdn` - FQDN address of override server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System CentralManagement can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_centralmanagement.labelname SystemCentralManagement
$ unset "FORTIOS_IMPORT_TABLE"
```
