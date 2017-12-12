/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceDNSVsSyncInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"total_records": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
