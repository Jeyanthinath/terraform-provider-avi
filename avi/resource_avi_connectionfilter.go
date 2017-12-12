/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceConnectionFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"conn_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
