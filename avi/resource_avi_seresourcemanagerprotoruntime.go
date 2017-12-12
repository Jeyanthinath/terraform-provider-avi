/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeResourceManagerProtoRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cold_start_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"pending_creates": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeCreatePendingProtoSchema()},
			"se_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcePlacementStatsSchema()},
			"ticks": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"upgrade_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
			"vrfs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRmVrfProtoSchema()},
			"vs_resync_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
