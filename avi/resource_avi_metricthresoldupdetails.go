/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMetricThresoldUpDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"current_value": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"entity_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"pool_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server": &schema.Schema{
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
