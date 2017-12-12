/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceFullClientLogsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_headers": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"throttle": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
		},
	}
}
