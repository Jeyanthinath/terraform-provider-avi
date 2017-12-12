/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceRouteEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"destination": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ifa_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rt_flags": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rt_refcnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vrf_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
