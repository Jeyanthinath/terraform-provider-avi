/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceInterfaceStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"flow_probes_ignored_in_tw": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"flow_probes_ignored_same_vnic": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ibytes": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ierrors": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ifq_stats": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIfQStatsSchema()},
			"ip_checksum_drops": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ipackets": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"l4_checksum_drops": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"local_flow_probes_req_received": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"local_flow_probes_req_sent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"obytes": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"oerrors": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"opackets": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"rx_kni": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"rx_max_single_burst": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"rx_mim_etype_p2s": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"rx_mim_etype_s2p": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"rx_nombuf": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"rx_pkt_iterations": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"rx_queue_full": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"tx_frags_p2s": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tx_kni": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"tx_kni_errs": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"tx_mim_etype_p2s": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"tx_mim_etype_s2p": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"tx_mim_frags_etype_p2s": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"tx_mim_frags_etype_s2p": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"tx_queue_full_retries": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
