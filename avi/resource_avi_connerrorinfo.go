/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceConnErrorInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_syn_retransmit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_window_shrink": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"out_of_orders": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"retransmits": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rx_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_num_window_shrink": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_out_of_orders": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_retransmits": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_rx_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_timeouts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_tx_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_zero_window_size_events": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"timeouts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tx_pkts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"zero_window_size_events": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
