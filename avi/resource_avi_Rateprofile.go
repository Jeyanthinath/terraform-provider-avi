/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceRateProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateLimiterActionSchema()},
			"burst_sz": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"explicit_tracking": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"fine_grain": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"http_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
