/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMetricCollectionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avg_application_latency": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_client_latency": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_l4_client_latency": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_max_concurrent_connections": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_4xx_errors_excluded": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_5xx_errors_excluded": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_total_bandwidth": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_total_finished_connections": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"min_se_disk_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"node_obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rum_navigation_timing": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"rum_visits": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"server_total_frustrated_responses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"server_total_satisfactory_responses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"server_total_tolerated_responses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_total_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"virtualservice_total_frustrated_responses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"virtualservice_total_satisfactory_responses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"virtualservice_total_tolerated_responses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}
