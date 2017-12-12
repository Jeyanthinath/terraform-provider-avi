/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeUpgradeErrorsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"from_se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se_group": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"se_group_ha_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_group_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"task": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"to_se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"traffic_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
