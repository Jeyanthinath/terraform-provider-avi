/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVsSeAssignedSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"connected": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"oper_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceOperationalStatusSchema()},
			"primary": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"scalein_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"snat_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"standby": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}
