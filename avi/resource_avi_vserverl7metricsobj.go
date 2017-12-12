/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVserverL7MetricsObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apdexr": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_application_response_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_blocking_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_browser_rendering_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_cache_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_cache_hits": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_cacheable_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_cacheable_hits": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_client_data_transfer_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_client_rtt": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_client_txn_latency": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_complete_responses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_connection_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dns_lookup_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dom_content_load_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_error_responses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_errors_excluded": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_frustrated_responses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_page_download_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_page_load_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_post_compression_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_pre_compression_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_redirection_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_reqs_per_session": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_1xx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_2xx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_3xx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_4xx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_4xx_avi_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_5xx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_5xx_avi_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rum_client_data_transfer_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_satisfactory_responses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_server_rtt": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_service_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_auth_dsa": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_auth_ecdsa": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_auth_rsa": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_connections": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_ecdsa_non_pfs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_ecdsa_pfs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_failed_connections": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_handshake_network_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_handshake_protocol_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_handshakes_new": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_handshakes_non_pfs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_handshakes_pfs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_handshakes_reused": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_handshakes_timedout": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_kx_dh": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_kx_ecdh": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_kx_rsa": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_rsa_non_pfs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_rsa_pfs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_ver_ssl30": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_ver_tls10": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_ver_tls11": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_ver_tls12": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_tolerated_responses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_total_requests": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_attacks": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_evaluated": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_evaluated_request_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_evaluated_request_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_evaluated_response_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_evaluated_response_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_flagged": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_flagged_request_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_flagged_request_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_flagged_response_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_flagged_response_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_latency_request_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_latency_request_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_latency_response_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_latency_response_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_matched": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_matched_request_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_matched_request_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_matched_response_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_matched_response_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_rejected": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_rejected_request_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_rejected_request_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_rejected_response_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waf_rejected_response_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_waiting_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_concurrent_sessions": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_ssl_open_sessions": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"node_obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"pct_cache_hits": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_cacheable_hits": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_response_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_ssl_failed_connections": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_waf_attacks": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_waf_evaluated": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_waf_flagged": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_waf_matched": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_waf_rejected": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"rum_apdexr": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"ssl_protocol_strength": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_application_response_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_blocking_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_browser_rendering_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_client_data_transfer_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_client_rtt": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_connection_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_dns_lookup_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_dom_content_load_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_finished_sessions": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_get_client_txn_latency": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_get_client_txn_latency_bucket1": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_get_client_txn_latency_bucket2": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_get_reqs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_num_page_load_time_bucket1": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_num_page_load_time_bucket2": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_num_rum_samples": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_other_client_txn_latency": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_other_client_txn_latency_bucket1": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_other_client_txn_latency_bucket2": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_other_reqs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_page_download_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_page_load_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_post_client_txn_latency": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_post_client_txn_latency_bucket1": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_post_client_txn_latency_bucket2": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_post_reqs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_redirection_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_reqs_finished_sessions": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_resp_1xx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_resp_2xx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_resp_3xx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_resp_4xx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_resp_5xx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_rum_client_data_transfer_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_server_rtt": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_service_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_total_responses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_attacks": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_evaluated_request_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_evaluated_request_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_evaluated_response_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_evaluated_response_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_flagged": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_flagged_request_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_flagged_request_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_flagged_response_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_flagged_response_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_latency_request_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_latency_request_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_latency_response_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_latency_response_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_matched_request_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_matched_request_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_matched_response_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_matched_response_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_rejected": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_rejected_request_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_rejected_request_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_rejected_response_body_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waf_rejected_response_header_phase": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waiting_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}
