/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMetricsServerDNSObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avg_complete_queries": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_errored_queries": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_req_type_a": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_req_type_aaaa": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_type_a": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_type_aaaa": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_type_cname": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_type_mx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_type_ns": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_type_other": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_resp_type_srv": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_tcp_queries": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_udp_queries": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"node_obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"pct_errored_queries": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}
