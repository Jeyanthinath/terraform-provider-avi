/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceScaleStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time_str": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_se_assigned": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se_requested": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"reason_code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reason_code_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"scale_se": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time_str": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_placement_resolution_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVipPlacementResolutionInfoSchema()},
		},
	}
}
