/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSystemUpgradeCliSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disruptive": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"force": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"image_path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"patch": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"patch_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_group_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"suspend_on_failure": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
