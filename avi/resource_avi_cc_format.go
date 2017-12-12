/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCC_FormatSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"img_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"img_pfx": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name_pfx": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secgrp_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secgrp_pfx": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
