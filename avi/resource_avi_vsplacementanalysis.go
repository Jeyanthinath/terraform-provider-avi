/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVsPlacementAnalysisSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"east_west": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"notes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"num_hosts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se_in_segroup": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_placement_issues": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVsSePlacementAnalysisSchema()},
			"vip_analysis": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipPlacementAnalaysisSchema()},
		},
	}
}
