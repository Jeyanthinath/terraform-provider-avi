/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVnicRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"if_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_avi_internal_network": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_hsm": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_mgmt": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_portchannel": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"mac_addr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1500,
			},
			"present_in_se": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"present_in_vinfra": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"virtual_network_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcevNICNetworkSchema()},
		},
	}
}
