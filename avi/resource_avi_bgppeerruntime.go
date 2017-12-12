/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceBgpPeerRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"active": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"advertise_snat_ip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"advertise_vip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"bgp_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"md5_secret": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nexthop_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"peer_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"peer_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_as": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"routes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"vs_names": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
		},
	}
}
