/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeGatewayHeartbeatFailedDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"gateway_ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"vrf_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrf_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
