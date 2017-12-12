/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceConnectionLogSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adf": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"average_turntime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"client_dest_port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"client_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"client_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_rtt": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"client_src_port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"connection_ended": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"dns_etype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_ips": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt}},
			"dns_qtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_request": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceDnsRequestSchema()},
			"dns_response": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceDnsResponseSchema()},
			"gslbpool_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gslbservice": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gslbservice_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"microservice": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"microservice_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mss": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"network_security_policy_rule_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_syn_retransmit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_transaction": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_window_shrink": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"out_of_orders": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"pool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxy_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"report_timestamp": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"retransmits": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"rx_bytes": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"rx_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"server_conn_src_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"server_dest_port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"server_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"server_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_num_window_shrink": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_out_of_orders": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"server_retransmits": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"server_rtt": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"server_rx_bytes": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"server_rx_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"server_src_port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"server_timeouts": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"server_total_bytes": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"server_total_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"server_tx_bytes": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"server_tx_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"server_zero_window_size_events": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"service_engine": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"significance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"significant": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"significant_log": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"ssl_cipher": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_session_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_timestamp": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"timeouts": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"total_bytes": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"total_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"total_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tx_bytes": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"tx_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udf": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"vcpu_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"virtualservice": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"vs_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"zero_window_size_events": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
