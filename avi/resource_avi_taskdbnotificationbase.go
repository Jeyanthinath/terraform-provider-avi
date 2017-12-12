/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceTaskDbNotificationBaseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"method": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"method_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"pb_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pb_runtime_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
			"ver": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
