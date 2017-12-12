/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMetricsVserverDNSObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avg_avi_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_complete_queries": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_domain_lookup_failures": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_errored_queries": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_gslbpool_member_not_available": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_invalid_queries": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_local_nxdomains": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_local_responses": &schema.Schema{
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
			"avg_req_type_mx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_req_type_ns": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_req_type_other": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_req_type_srv": &schema.Schema{
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
			"avg_tcp_passthrough_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_tcp_passthrough_queries": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_tcp_queries": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_udp_passthrough_errors": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_udp_passthrough_queries": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_udp_passthrough_resp_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_udp_queries": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_unsupported_queries": &schema.Schema{
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
			"sum_udp_passthrough_resp_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}
