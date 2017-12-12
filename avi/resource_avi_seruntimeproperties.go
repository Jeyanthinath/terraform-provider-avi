/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeRuntimePropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_headers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAppHdrSchema()},
			"baremetal_dispatcher_handles_flows": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"connections_lossy_log_rate_limiter_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"connections_udfnf_log_rate_limiter_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"disable_flow_probes": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_gro": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_tso": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dos_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceDosThresholdProfileSchema()},
			"downstream_send_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3600000,
			},
			"dp_aggressive_hb_frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"dp_aggressive_hb_timeout_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"dp_hb_frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"dp_hb_timeout_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"dupip_frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dupip_timeout_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"enable_hsm_log": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"feproxy_vips_enable_proxy_arp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"flow_table_batch_push_frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"flow_table_new_syn_max_entries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  40000,
			},
			"global_mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"http_rum_console_log": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"http_rum_min_content_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  64,
			},
			"lbaction_num_requests_to_dispatch": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"lbaction_rq_per_request_max_retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  22,
			},
			"log_agent_compress_logs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"log_agent_conn_send_buffer_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  16384,
			},
			"log_agent_export_msg_buffer_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  524288,
			},
			"log_agent_export_wait_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"log_agent_file_sz_appl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"log_agent_file_sz_conn": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"log_agent_file_sz_debug": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"log_agent_file_sz_event": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"log_agent_log_storage_min_sz": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1024,
			},
			"log_agent_max_active_adf_files_per_vs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"log_agent_max_concurrent_rsync": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1024,
			},
			"log_agent_max_logmessage_proto_sz": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  65536,
			},
			"log_agent_max_storage_excess_percent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  110,
			},
			"log_agent_max_storage_ignore_percent": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  "20.0"},
			"log_agent_min_storage_per_vs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"log_agent_pause_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"log_agent_sleep_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"log_agent_unknown_vs_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1800,
			},
			"log_message_max_file_list_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  64,
			},
			"mcache_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mcache_fetch_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mcache_max_cache_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mcache_store_in_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mcache_store_in_max_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mcache_store_in_min_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mcache_store_out_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mcache_store_se_max_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ngx_free_connection_stack": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"persistence_mem_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scaleout_udp_per_pkt": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"se_auth_ldap_bind_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5000,
			},
			"se_auth_ldap_cache_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100000,
			},
			"se_auth_ldap_connect_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10000,
			},
			"se_auth_ldap_conns_per_server": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"se_auth_ldap_reconnect_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10000,
			},
			"se_auth_ldap_request_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10000,
			},
			"se_auth_ldap_servers_failover_only": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_dp_compression": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeRuntimeCompressionPropertiesSchema()},
			"se_dp_hm_drops": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_dp_if_state_poll_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"se_dp_log_nf_enqueue_percent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  70,
			},
			"se_dp_log_udf_enqueue_percent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  90,
			},
			"se_dp_vnic_queue_stall_event_sleep": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_dp_vnic_queue_stall_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2000,
			},
			"se_dp_vnic_queue_stall_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10000,
			},
			"se_handle_interface_routes": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_hb_persist_fudge_bits": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3,
			},
			"se_mac_error_threshold_to_disable_promiscious": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"se_memory_poison": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"se_metrics_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60000,
			},
			"se_metrics_rt_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"se_metrics_rt_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"se_packet_buffer_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_random_tcp_drops": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_rate_limiters": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeRateLimitersSchema()},
			"service_ip_subnets": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema()},
			"service_port_ranges": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePortRangeSchema()},
			"services_accessible_all_interfaces": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"spdy_fwd_proxy_parse_enable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"tcp_syn_cache_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  32768,
			},
			"tcp_syncache_max_retransmit_default": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"upstream_connect_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3600000,
			},
			"upstream_connpool_cache_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"upstream_connpool_conn_idle_thresh_tmo": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"upstream_connpool_conn_idle_tmo": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"upstream_connpool_conn_life_tmo": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"upstream_connpool_conn_max_reuse": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"upstream_connpool_core_max_cache": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"upstream_connpool_enable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"upstream_connpool_server_max_cache": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"upstream_connpool_strategy": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},
			"upstream_keepalive": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"upstream_read_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3600000,
			},
			"upstream_send_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3600000,
			},
			"user_defined_metric_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
		},
	}
}
