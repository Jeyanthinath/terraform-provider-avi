/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceTerminalSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"command_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
			"display_api_details": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"length": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"linux_cmd_output": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"output": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"session_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  15,
			},
			"timezone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"unhide": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}
