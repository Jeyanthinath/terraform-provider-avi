/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSwitchoverInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"new_primary_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"new_primary_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"old_primary_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"old_primary_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
