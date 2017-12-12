/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcevNICSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adapter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"can_se_dp_takeover": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"connected": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"del_pending": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dhcp_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"if_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_asm": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
			"linux_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"members": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMemberInterfaceSchema()},
			"mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1500,
			},
			"network_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pci_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vlan_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vlan_interfaces": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVlanInterfaceSchema()},
			"vnic_networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcevNICNetworkSchema()},
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
