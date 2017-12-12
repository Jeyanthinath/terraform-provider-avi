/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeRateLimitersSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"arp_rl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"default_rl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"flow_probe_rl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  250,
			},
			"icmp_rl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"icmp_rsp_rl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  500,
			},
			"rst_rl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
		},
	}
}
