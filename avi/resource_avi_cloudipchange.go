/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCloudIpChangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"ip_mask": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  32,
			},
			"mac_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_vm_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
