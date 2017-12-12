/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceServiceEngineResourcesScoreDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avg_connection_mem_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_cpu_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_cpu_wait": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_disk1_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_dynamic_mem_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_mem_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_packet_buffer_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_persistent_table_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ssl_session_cache_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"oper_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pct_syn_cache_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"reason_attr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
