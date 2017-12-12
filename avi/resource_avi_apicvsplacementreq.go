/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceApicVSPlacementReqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"graph": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"lifs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceLifSchema()},
			"network_rel": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAPICNetworkRelSchema()},
			"tenant_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vdev": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
