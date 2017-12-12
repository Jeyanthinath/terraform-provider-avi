/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceApicAgentVsNetworkErrorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
