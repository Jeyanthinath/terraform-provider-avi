/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Resourcecc_vip_infoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alloc_fip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"alloc_vip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"avi_alloced_fip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"avi_alloced_vip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dns_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsInfoSchema()},
			"fip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"fip_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"vip_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
