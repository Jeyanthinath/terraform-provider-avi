/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceLicenseExpiryDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backend_servers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cores": &schema.Schema{
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
			"license_tier": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"license_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_apps": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_ses": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sockets": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"throughput": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
