/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceUdpStatRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"proc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"udpps_pcbcachemiss": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udpps_pcbhashmiss": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_badlen": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_badpkts": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_badsum": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_dns_domain_drops": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_gs_down": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_invalid_edns_option": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_invalid_qd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_local_responses": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_no_record": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_no_valid_gs_member": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_pass_through": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_pass_through_errors": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_policy_drops": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_query_a": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_query_aaaa": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_query_cname": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_query_mx": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_query_ns": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_query_others": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_query_srv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_response_a": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_response_aaaa": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_response_cname": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_response_mx": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_response_ns": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_response_nxdomain": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_response_others": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_response_srv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_rx_responses": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_dns_unsupported_queries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_errored_conns": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_fastout": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_filtermcast": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_finished_conns": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_fullsock": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_hdrops": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_ipackets": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_noport": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_noportbcast": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_noportmcast": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_nosum": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_opackets": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_pktsfrag": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_port_unreach_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_rxbytes": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_rxpkts": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_started_conns": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udps_timedout_conns": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_txbytes": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"udps_txpkts": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
