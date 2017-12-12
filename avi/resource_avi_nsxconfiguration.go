/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceNsxConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avi_nsx_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsx_manager_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsx_manager_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsx_manager_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsx_poll_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
		},
	}
}
