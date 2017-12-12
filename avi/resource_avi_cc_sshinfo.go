/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCC_SshInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key_file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_pwd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
