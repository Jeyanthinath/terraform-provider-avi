/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVIPGInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auto_expand": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dc_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dvs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"managed_object_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"num_ports": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pg_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"switch_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vlan_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vlan_id_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVlanRangeSchema()},
		},
	}
}
