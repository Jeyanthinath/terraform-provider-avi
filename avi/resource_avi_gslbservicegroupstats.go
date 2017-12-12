/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceGslbServiceGroupStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bad_connections": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"current_connections": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"lb_fail_add_pending": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"lb_fail_persistent_server_down": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"lb_fail_persistent_server_invalid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"lb_fail_server_down": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pool_load": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"total_connections": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
