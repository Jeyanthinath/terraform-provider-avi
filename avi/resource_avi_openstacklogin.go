/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceOpenstackLoginSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_tenant": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "admin"},
			"auth_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"keystone_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
