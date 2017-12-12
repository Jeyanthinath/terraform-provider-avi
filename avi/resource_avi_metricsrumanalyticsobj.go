/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMetricsRumAnalyticsObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"application_response_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"blocking_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"browser_rendering_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"client_rtt": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"connection_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"dns_lookup_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"dom_content_load_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"node_obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"page_download_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"page_load_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"redirection_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"rum_client_data_transfer_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"server_rtt": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"service_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"waiting_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}
