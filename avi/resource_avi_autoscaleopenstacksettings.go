/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAutoScaleOpenStackSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"heat_scale_down_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"heat_scale_up_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
