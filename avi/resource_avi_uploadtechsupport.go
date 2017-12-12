/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceUploadTechSupportSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"case_number": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"portal_token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
