/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeBootupPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"distribute_queues": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"distribute_vnics": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"docker_backend_portend": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30720,
			},
			"docker_backend_portstart": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  20480,
			},
			"fair_queueing_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"geo_db_granularity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"l7_conns_per_core": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  16384,
			},
			"l7_resvd_listen_conns_per_core": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  256,
			},
			"log_agent_debug_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"log_agent_trace_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"se_dp_compression": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeBootupCompressionPropertiesSchema()},
			"se_dpdk_pmd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_emulated_cores": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_ip_encap_ipc": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_l3_encap_ipc": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_log_buffer_app_blocking_dequeue": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_log_buffer_applog_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4096,
			},
			"se_log_buffer_chunk_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1024,
			},
			"se_log_buffer_conn_blocking_dequeue": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_log_buffer_connlog_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1024,
			},
			"se_log_buffer_events_blocking_dequeue": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"se_log_buffer_events_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  512,
			},
			"se_lro": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"se_pcap_pkt_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  256,
			},
			"se_pcap_pkt_sz": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  65536,
			},
			"se_rum_sampling_nav_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"se_rum_sampling_nav_percent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"se_rum_sampling_res_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"se_rum_sampling_res_percent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"se_tx_batch_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  64,
			},
			"se_use_dpdk": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssl_sess_cache_per_vs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4096,
			},
			"ssl_sess_cache_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  86400,
			},
			"tcp_syncache_hashsize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  8192,
			},
		},
	}
}
