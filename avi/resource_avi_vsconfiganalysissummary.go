/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVsConfigAnalysisSummarySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"eastwest_vs_config_inconsistent_ref": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"notes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"num_eastwest_vs_config_inconsistent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vs_config_inconsistent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vs_config_inconsistent_ref": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
		},
	}
}
