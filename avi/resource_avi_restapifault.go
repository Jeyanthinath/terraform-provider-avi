/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceRestApiFaultSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
