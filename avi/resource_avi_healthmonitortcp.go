/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHealthMonitorTcpSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"maintenance_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tcp_half_open": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"tcp_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tcp_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
