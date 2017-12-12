/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_down_requested": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"at_curr_ver": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"attach_ip_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Programming Network reachability to the Virtual Service IP in the Cloud"},
			"attach_ip_success": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"delete_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"download_selist_only": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"floating_intf_ip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema()},
			"geo_download": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"geodb_download": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"gslb_download": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_connected": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"is_portchannel": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_primary": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"is_standby": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"memory": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2001,
			},
			"pending_download": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"scalein_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"sec_idx": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"snat_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"vcpus": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0"},
			"vip_intf_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"vip_intf_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeVipInterfaceListSchema()},
			"vip_intf_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_subnet_mask": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  32,
			},
			"vlan_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vnic": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVsSeVnicSchema()},
		},
	}
}
