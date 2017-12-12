/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVCASetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"privilege": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
