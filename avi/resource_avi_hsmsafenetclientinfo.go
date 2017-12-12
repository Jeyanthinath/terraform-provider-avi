/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHSMSafenetClientInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"chrystoki_conf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"client_priv_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"session_major_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"session_minor_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
