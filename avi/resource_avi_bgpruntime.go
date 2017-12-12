/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceBgpRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"active": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"debug_info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"local_as": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"peer_bmp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"peers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceBgpPeerRuntimeSchema()},
			"proc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"quagga_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrf": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
