/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcePlacementGlobalsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cold_start_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"num_consumers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_create_pending": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_resources": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ticks": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"upgrade_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vs_resync_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
