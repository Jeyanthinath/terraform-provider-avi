/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVsPlacementAnalysisSummarySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"eastwest_vs_fully_placed_ref": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"eastwest_vs_not_placed_ref": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"eastwest_vs_partially_placed_ref": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"notes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"num_eastwest_vs_fully_placed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_eastwest_vs_not_placed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_eastwest_vs_partially_placed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vs_fully_placed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vs_not_placed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vs_partially_placed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vs_fully_placed_ref": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"vs_not_placed_ref": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"vs_partially_placed_ref": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
		},
	}
}
