/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceApicAgentBridgeDomainVrfChangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bridge_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"new_vrf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"old_vrf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"vs_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
		},
	}
}
