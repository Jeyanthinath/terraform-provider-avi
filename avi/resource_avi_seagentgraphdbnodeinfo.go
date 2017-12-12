/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeAgentGraphDBNodeInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_obj": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_obj_active": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_obj_awaiting_dp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_obj_error": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_obj_ew_subnet_error": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"obj": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeObjectSchema()},
		},
	}
}
