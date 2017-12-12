/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeAgentGraphDBNodeTxnDetailSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dp_deq_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dp_duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dp_enq_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"txn_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
