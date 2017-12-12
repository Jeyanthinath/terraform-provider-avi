/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceEmailConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auth_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"from_email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "admin@avicontroller.net"},
			"mail_server_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "localhost"},
			"mail_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  25,
			},
			"smtp_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
