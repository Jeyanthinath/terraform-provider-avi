/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceLicenseDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backend_servers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"expiry_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
