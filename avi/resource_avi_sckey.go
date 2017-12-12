/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSCKeySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SC_UPDATE"},
			"metadata": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
