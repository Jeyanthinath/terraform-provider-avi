/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCreateInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mgmt_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mgmt_subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrPrefixSchema()},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"recoverable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"timed_out": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}
