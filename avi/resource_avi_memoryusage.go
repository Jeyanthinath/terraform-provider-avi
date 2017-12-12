/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMemoryUsageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"free": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"total": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
