/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAppInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_hdr_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"app_hdr_value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
