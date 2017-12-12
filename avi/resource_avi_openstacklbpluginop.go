/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceOpenStackLbPluginOpSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"command": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"detail": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"elapsed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"prov": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"result": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
