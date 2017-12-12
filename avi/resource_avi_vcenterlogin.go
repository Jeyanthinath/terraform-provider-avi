/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVcenterLoginSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_ts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vcenter_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
