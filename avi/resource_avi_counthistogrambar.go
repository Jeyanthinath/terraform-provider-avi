/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCountHistogramBarSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"x": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"y": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
