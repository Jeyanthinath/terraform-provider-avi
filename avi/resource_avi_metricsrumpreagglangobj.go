/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMetricsRumPreaggLangObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"num_samples": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sum_application_response_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_blocking_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_browser_rendering_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_client_rtt": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_connection_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_dns_lookup_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_dom_content_load_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_page_download_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_page_load_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_redirection_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_rum_client_data_transfer_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_server_rtt": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_service_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_waiting_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}
