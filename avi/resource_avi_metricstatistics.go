/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMetricStatisticsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"last_sample": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_ts": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mean": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"min": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"min_ts": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_samples": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sum": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"trend": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}
