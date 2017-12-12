/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCC_VnicInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mac_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"port_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SYSERR_SUCCESS"},
			"status_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_uuid": &schema.Schema{
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
