/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeVssPlacementSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"core": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"core_nonaffinity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dedicated_dispatcher": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_flow_cores": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_subcores": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_vss_placement_entry": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeVssPlacementEntrySchema()},
			"stretch_factor": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sum_weights": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
