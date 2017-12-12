/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHTTPHdrValueSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"val": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"var": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
