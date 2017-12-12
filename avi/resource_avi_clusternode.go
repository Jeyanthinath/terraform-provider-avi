/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceClusterNodeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "node"},
			"public_ip_or_name": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"vm_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_mor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
