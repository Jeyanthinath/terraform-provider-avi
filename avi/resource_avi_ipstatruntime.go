/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceIpStatRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ips_badaddr": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_badhlen": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_badlen": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_badoptions": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_badsum": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_badtcpsum": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_badudpsum": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_badvers": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_cantforward": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_cantfrag": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_delivered": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_fastforward": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_forward": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_fragdropped": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_fragmented": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_fragments": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_fragtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_ipsum_large": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_localout": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_nogif": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_noproto": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_noroute": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_notmember": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_odropped": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_ofragments": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_rawout": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_reassembled": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_redirectsent": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_tcpsum_large": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_toolong": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_tooshort": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_toosmall": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_total": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ips_udpsum_large": &schema.Schema{
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
		},
	}
}
