/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceApicAgentInternalSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apic_build_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"apic_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"connected": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"connected_apic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"curr_monotonic_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"last_login": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"last_poll": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_ws_message": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_debug_notifications": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_ipg_notifications": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_pool_notifications": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se_notifications": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_subscriptions": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vcenter_notifications": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vs_notifications": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_ws_msgs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"poll_seconds": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"websocket_close_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"websocket_connected": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"websocket_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
