/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceServiceEngineCountersSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dp_hb_miss_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"hb_miss_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"hb_recv_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"hb_sent_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"last_down": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_dp_hb_miss": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_hb_miss": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_partitioned": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_up": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reg_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reg_fail_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_down_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_up_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
