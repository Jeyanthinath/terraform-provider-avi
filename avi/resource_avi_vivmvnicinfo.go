/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVIVmVnicInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mac_addr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"vcenter_portgroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_vnic_nw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
