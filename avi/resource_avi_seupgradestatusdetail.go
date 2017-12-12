/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeUpgradeStatusDetailSchema() *schema.Resource {
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
			"estimated_time_secs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"notes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"progress_percent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"summary": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeUpgradeStatusSummarySchema()},
			"tasks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeTaskSchema()},
		},
	}
}
