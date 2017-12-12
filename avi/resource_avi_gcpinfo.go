/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceGcpInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"project": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"zone": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
