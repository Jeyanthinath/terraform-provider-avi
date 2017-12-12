/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceDnsEdnsOptionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"addr_family": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"code": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"scope_prefix_len": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"source_prefix_len": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"subnet_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
