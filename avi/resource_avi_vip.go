/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVipSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auto_allocate_floating_ip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"auto_allocate_ip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"avi_allocated_fip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"avi_allocated_vip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"discovered_networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDiscoveredNetworkSchema()},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"floating_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"floating_ip6": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"floating_subnet6_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"floating_subnet_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6_address": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"ip_address": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"ipam_network_subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIPNetworkSubnetSchema()},
			"network_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrPrefixSchema()},
			"subnet6": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrPrefixSchema()},
			"subnet6_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnet_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
