/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMetricsDataSchema() *schema.Resource {
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
			"is_null": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"num_samples": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"page_download_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"page_load_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"prediction_interval_high": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"prediction_interval_low": &schema.Schema{
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
			"timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true},
			"value_str": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"value_str_desc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"waiting_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}
