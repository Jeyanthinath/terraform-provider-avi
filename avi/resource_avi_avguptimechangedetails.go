/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAvgUptimeChangeDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"current_value": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"metric_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_str": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"threshold": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}
