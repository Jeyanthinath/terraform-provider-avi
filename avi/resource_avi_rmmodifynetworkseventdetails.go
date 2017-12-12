/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceRmModifyNetworksEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRmModifyVnicSchema()},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
		},
	}
}
