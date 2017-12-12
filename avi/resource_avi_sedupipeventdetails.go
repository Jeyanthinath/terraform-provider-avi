/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeDupipEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"local_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
