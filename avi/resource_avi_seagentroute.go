/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeAgentRouteSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"default_gw": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"delete_route": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dst_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"gateway": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"gwmon_disable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"if_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mask": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"metric": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"origin": &schema.Schema{
				Type:     schema.TypeString,
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
