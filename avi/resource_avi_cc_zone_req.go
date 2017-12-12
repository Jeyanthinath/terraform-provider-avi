/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Resourcecc_zone_reqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "cloud-0"},
			"detail": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"search": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"unavail": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
