/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSingleLicenseSchema() *schema.Resource {
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
			"created_on": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"customer_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"enforced_params": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"last_update": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"license_string": &schema.Schema{
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
			"sockets": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"start_on": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"throughput": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"valid_until": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
