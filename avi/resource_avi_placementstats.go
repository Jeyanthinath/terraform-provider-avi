/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcePlacementStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_se_create_call": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se_create_fail": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se_create_success": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se_create_timed_out": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se_freed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vnic_add_call": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vnic_add_fail": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vnic_add_success": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vnic_del_call": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vnic_del_fail": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vnic_del_success": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_name_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Avi-"},
			"se_name_suffix": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
