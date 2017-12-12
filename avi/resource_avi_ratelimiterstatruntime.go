/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceRatelimiterStatRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"arp_rx_rl_cfg_pps": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"arp_rx_rl_confirming": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"arp_rx_rl_drops": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"flow_probe_tx_rl_cfg_pps": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"flow_probe_tx_rl_confirming": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"flow_probe_tx_rl_drops": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_rsp_rx_rl_cfg_pps": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_rsp_rx_rl_confirming": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_rsp_rx_rl_drops": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_rx_rl_cfg_pps": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_rx_rl_confirming": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_rx_rl_drops": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"proc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tcp_rst_tx_rl_cfg_pps": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"tcp_rst_tx_rl_confirming": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"tcp_rst_tx_rl_drops": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
