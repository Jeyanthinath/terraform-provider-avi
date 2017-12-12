/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMetricsUserMetricsObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avg_counter": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_counter": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"min_counter": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"node_obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"sum_counter": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_gauge": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}
