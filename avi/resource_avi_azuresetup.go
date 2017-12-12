/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAzureSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alb_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nic_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vips": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema()},
			"vnet_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
		},
	}
}
