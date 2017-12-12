/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceTechSupportFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"duration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
