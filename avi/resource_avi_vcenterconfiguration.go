/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcevCenterConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"datacenter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_ip_subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrPrefixSchema()},
			"management_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"privilege": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"vcenter_template_se_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
