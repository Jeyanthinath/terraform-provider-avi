/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMemberInterfaceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"active": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"if_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"mac_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
