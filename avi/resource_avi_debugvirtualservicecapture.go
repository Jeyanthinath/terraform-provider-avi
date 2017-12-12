/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceDebugVirtualServiceCaptureSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pkt_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  128,
			},
		},
	}
}
