/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceNetworkSubnetInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"free": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"network_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrPrefixSchema()},
			"total": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"used": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
