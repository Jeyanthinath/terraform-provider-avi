/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Resourcecc_update_vip_rspSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alloced_fip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"alloced_vip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
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
			"fip_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_dns": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"public_dns": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ret_status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"ret_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_subnet_prefix": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrPrefixSchema()},
		},
	}
}
