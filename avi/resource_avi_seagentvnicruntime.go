/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeAgentVnicRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avi_internal_network": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"can_se_dp_takeover": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"connected": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"consumed_by_dataplane": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dhcp_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"enabled_flag": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"if_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"is_complete": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_mgmt": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"linux_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"mac_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"network_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nw": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeAgentVnicNwRuntimeSchema()},
			"pci_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"pushed_to_controller": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"pushed_to_dataplane": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"running_flag": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vrf_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vrf_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
