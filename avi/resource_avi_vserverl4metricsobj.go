/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVserverL4MetricsObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apdexc": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"apdexrtt": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_application_dos_attacks": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_bandwidth": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_bytes_policy_drops": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_complete_conns": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_connections_dropped": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_app_error": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_attacks": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_bad_rst_flood": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_bandwidth": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_conn": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_conn_ip_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_conn_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_fake_session": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_http_abort": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_http_error": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_http_timeout": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_malformed_flood": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_non_syn_flood": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_cookie_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_hdr_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_ip_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_ip_rl_drop_bad": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_ip_scan_bad_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_ip_scan_unknown_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_ip_uri_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_ip_uri_rl_drop_bad": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_uri_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_uri_rl_drop_bad": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_uri_scan_bad_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_req_uri_scan_unknown_rl_drop": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_rx_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_slow_uri": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_small_window_stress": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_ssl_error": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_syn_flood": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_total_req": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_tx_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dos_zero_window_stress": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_errored_connections": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_l4_client_latency": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_lossy_connections": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_lossy_req": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_network_dos_attacks": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_new_established_conns": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_pkts_policy_drops": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_policy_drops": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rx_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rx_bytes_dropped": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rx_pkts": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rx_pkts_dropped": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_syns": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_total_connections": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_total_rtt": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_tx_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_tx_pkts": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_num_active_se": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_open_conns": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_rx_bytes_absolute": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_rx_pkts_absolute": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_tx_bytes_absolute": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_tx_pkts_absolute": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"node_obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"pct_application_dos_attacks": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_connection_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_connections_dos_attacks": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_dos_bandwidth": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_dos_rx_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_network_dos_attacks": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_pkts_dos_attacks": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_policy_drops": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_conn_duration": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_connection_dropped_user_limit": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_connection_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_connections_dropped": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_dup_ack_retransmits": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_end_to_end_rtt": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_end_to_end_rtt_bucket1": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_end_to_end_rtt_bucket2": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_finished_conns": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_lossy_connections": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_lossy_req": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_out_of_orders": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_packet_dropped_user_bandwidth_limit": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_rtt_valid_connections": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_sack_retransmits": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_server_flow_control": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_timeout_retransmits": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_zero_window_size_events": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}
