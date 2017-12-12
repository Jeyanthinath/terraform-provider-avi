/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSePlacementAnalysisSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_vs_limit_on_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"notes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"num_eastwest_vs_on_controller": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_eastwest_vs_on_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_expected_vs_on_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vs_issues": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vs_on_controller": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vs_on_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vs_on_controller_not_on_se": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeVsPlacementAnalysisSchema()},
			"vs_on_se_not_on_controller": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeVsPlacementAnalysisSchema()},
		},
	}
}
