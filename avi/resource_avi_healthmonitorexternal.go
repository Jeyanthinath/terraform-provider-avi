/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHealthMonitorExternalSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"command_code": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"command_parameters": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"command_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"command_variables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
