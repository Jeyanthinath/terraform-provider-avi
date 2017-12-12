/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceServerResourcesScoreDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
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
			"avg_disk2_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_disk3_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_disk4_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_mem_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_uptime": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_open_conns": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pool_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
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
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
