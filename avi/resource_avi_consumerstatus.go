/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceConsumerStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"num_se_assigned": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se_requested": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason_code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reason_code_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resources": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceResourceStatusSchema()},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
			"vip_placement_resolution_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVipPlacementResolutionInfoSchema()},
		},
	}
}
