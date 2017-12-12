/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceProcessStatsObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avg_ctx_switches_invol": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_ctx_switches_vol": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_fds": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_io_read_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_io_read_count": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_io_write_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_io_write_count": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_num_threads": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_pss": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rss": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_swap": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_vms": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_cpu_pct": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"node_obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
