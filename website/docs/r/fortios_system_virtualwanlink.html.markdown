---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_virtualwanlink"
description: |-
  Configure redundant internet connections using SD-WAN (formerly virtual WAN link).
---

# fortios_system_virtualwanlink
Configure redundant internet connections using SD-WAN (formerly virtual WAN link).

## Example Usage

```hcl
resource "fortios_system_virtualwanlink" "trname" {
  fail_detect       = "disable"
  load_balance_mode = "source-ip-based"
  status            = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable SD-WAN.
* `load_balance_mode` - Algorithm or mode to use for load balancing Internet traffic to SD-WAN members.
* `fail_detect` - Enable/disable SD-WAN Internet connection status checking (failure detection).
* `fail_alert_interfaces` - Physical interfaces that will be alerted.
* `members` - Physical FortiGate interfaces added to the virtual-wan-link.
* `health_check` - SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it.
* `service` - Create SD-WAN rules or priority rules (also called services) to control how sessions are distributed to physical interfaces in the SD-WAN.

The `fail_alert_interfaces` block supports:

* `name` - Physical interface name.

The `members` block supports:

* `seq_num` - Sequence number(1-255).
* `interface` - Interface name.
* `gateway` - The default gateway for this interface. Usually the default gateway of the Internet service provider that this interface is connected to.
* `source` - Source IP address used in the health-check packet to the server.
* `gateway6` - IPv6 gateway.
* `source6` - Source IPv6 address used in the health-check packet to the server.
* `cost` - Cost of this interface for services in SLA mode (0 - 4294967295, default = 0).
* `weight` - Weight of this interface for weighted load balancing. (0 - 255) More traffic is directed to interfaces with higher weights.
* `priority` - Priority of the interface (0 - 4294967295). Used for SD-WAN rules or priority rules.
* `spillover_threshold` - Egress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
* `ingress_spillover_threshold` - Ingress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.
* `volume_ratio` - Measured volume ratio (this value / sum of all values = percentage of link volume, 0 - 255).
* `status` - Enable/disable this interface in the SD-WAN.
* `comment` - Comments.

The `health_check` block supports:

* `name` - Status check or health check name.
* `probe_packets` - Enable/disable transmission of probe packets.
* `addr_mode` - Address mode (IPv4 or IPv6).
* `server` - IP address or FQDN name of the server.
* `protocol` - Protocol used to determine if the FortiGate can communicate with the server.
* `port` - Port number used to communicate with the server over the selected protocol.
* `security_mode` - Twamp controller security mode.
* `password` - Twamp controller password in authentication mode
* `packet_size` - Packet size of a twamp test session,
* `http_get` - URL used to communicate with the server if the protocol if the protocol is HTTP.
* `http_agent` - String in the http-agent field in the HTTP header.
* `http_match` - Response string expected from the server if the protocol is HTTP.
* `interval` - Status check interval, or the time between attempting to connect to the server (1 - 3600 sec, default = 5).
* `failtime` - Number of failures before server is considered lost (1 - 3600, default = 5).
* `recoverytime` - Number of successful responses received before server is considered recovered (1 - 3600, default = 5).
* `update_cascade_interface` - Enable/disable update cascade interface.
* `update_static_route` - Enable/disable updating the static route.
* `sla_fail_log_period` - Time interval in seconds that SLA fail log messages will be generated (0 - 3600, default = 0).
* `sla_pass_log_period` - Time interval in seconds that SLA pass log messages will be generated (0 - 3600, default = 0).
* `threshold_warning_packetloss` - Warning threshold for packet loss (percentage, default = 0).
* `threshold_alert_packetloss` - Alert threshold for packet loss (percentage, default = 0).
* `threshold_warning_latency` - Warning threshold for latency (ms, default = 0).
* `threshold_alert_latency` - Alert threshold for latency (ms, default = 0).
* `threshold_warning_jitter` - Warning threshold for jitter (ms, default = 0).
* `threshold_alert_jitter` - Alert threshold for jitter (ms, default = 0).
* `members` - Member sequence number list.
* `sla` - Service level agreement (SLA).

The `members` block supports:

* `seq_num` - Member sequence number.

The `sla` block supports:

* `id` - SLA ID.
* `link_cost_factor` - Criteria on which to base link selection.
* `latency_threshold` - Latency for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `jitter_threshold` - Jitter for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `packetloss_threshold` - Packet loss for SLA to make decision in percentage. (0 - 100, default = 0).

The `service` block supports:

* `id` - Priority rule ID (1 - 4000).
* `name` - Priority rule name.
* `addr_mode` - Address mode (IPv4 or IPv6).
* `input_device` - Source interface name.
* `mode` - Control how the priority rule sets the priority of interfaces in the SD-WAN.
* `quality_link` - Quality grade.
* `member` - Member sequence number.
* `tos` - Type of service bit pattern.
* `tos_mask` - Type of service evaluated bits.
* `protocol` - Protocol number.
* `start_port` - Start destination port number.
* `end_port` - End destination port number.
* `route_tag` - IPv4 route map route-tag.
* `dst` - Destination address name.
* `dst_negate` - Enable/disable negation of destination address match.
* `src` - Source address name.
* `dst6` - Destination address6 name.
* `src6` - Source address6 name.
* `src_negate` - Enable/disable negation of source address match.
* `users` - User name.
* `groups` - User groups.
* `internet_service` - Enable/disable use of Internet service for application-based load balancing.
* `internet_service_custom` - Custom Internet service name list.
* `internet_service_custom_group` - Custom Internet Service group list.
* `internet_service_id` - Internet service ID list.
* `internet_service_group` - Internet Service group list.
* `internet_service_ctrl` - Control-based Internet Service ID list.
* `internet_service_ctrl_group` - Control-based Internet Service group list.
* `health_check` - Health check.
* `link_cost_factor` - Link cost factor.
* `packet_loss_weight` - Coefficient of packet-loss in the formula of custom-profile-1.
* `latency_weight` - Coefficient of latency in the formula of custom-profile-1.
* `jitter_weight` - Coefficient of jitter in the formula of custom-profile-1.
* `bandwidth_weight` - Coefficient of reciprocal of available bidirectional bandwidth in the formula of custom-profile-1.
* `link_cost_threshold` - Percentage threshold change of link cost values that will result in policy route regeneration (0 - 10000000, default = 10).
* `hold_down_time` - Waiting period in seconds when switching from the back-up member to the primary member (0 - 10000000, default = 0).
* `dscp_forward` - Enable/disable forward traffic DSCP tag.
* `dscp_reverse` - Enable/disable reverse traffic DSCP tag.
* `dscp_forward_tag` - Forward traffic DSCP tag.
* `dscp_reverse_tag` - Reverse traffic DSCP tag.
* `sla` - Service level agreement (SLA).
* `priority_members` - Member sequence number list.
* `status` - Enable/disable SD-WAN service.
* `gateway` - Enable/disable SD-WAN service gateway.

The `input_device` block supports:

* `name` - Interface name.

The `dst` block supports:

* `name` - Address or address group name.

The `src` block supports:

* `name` - Address or address group name.

The `dst6` block supports:

* `name` - Address6 or address6 group name.

The `src6` block supports:

* `name` - Address6 or address6 group name.

The `users` block supports:

* `name` - User name.

The `groups` block supports:

* `name` - Group name.

The `internet_service_custom` block supports:

* `name` - Custom Internet service name.

The `internet_service_custom_group` block supports:

* `name` - Custom Internet Service group name.

The `internet_service_id` block supports:

* `id` - Internet service ID.

The `internet_service_group` block supports:

* `name` - Internet Service group name.

The `internet_service_ctrl` block supports:

* `id` - Control-based Internet Service ID.

The `internet_service_ctrl_group` block supports:

* `name` - Control-based Internet Service group name.

The `sla` block supports:

* `health_check` - Virtual WAN Link health-check.
* `id` - SLA ID.

The `priority_members` block supports:

* `seq_num` - Member sequence number.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System VirtualWanLink can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_virtualwanlink.labelname SystemVirtualWanLink
$ unset "FORTIOS_IMPORT_TABLE"
```
