/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeTaskSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"from_se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"notes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
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
			"se_group_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"thread": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"to_se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"worker": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
