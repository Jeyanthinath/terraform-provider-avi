/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMetricsSourceInsightsObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avg_bandwidth": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_client_end2end_latency": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_complete_conns": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_complete_responses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dropped_conns": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_error_responses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_http_timeout": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_policy_drops": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_total_requests": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_open_conns": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"node_obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"sum_client_end2end_latency": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}
