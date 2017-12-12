/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeVssPlacementEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"core": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_vs_placement_entry": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeVsPlacementEntrySchema()},
		},
	}
}
