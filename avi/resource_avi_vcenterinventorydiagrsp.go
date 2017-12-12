/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVcenterInventoryDiagRspSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"datacenter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"errored_hosts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVcenterHostDiagSchema()},
			"networks": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVcenterNetworkDiagSchema()},
			"total_hosts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vcenter_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
