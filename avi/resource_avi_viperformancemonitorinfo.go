/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVIPerformanceMonitorInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceVIAdminCredentialsSchema()},
			"cloud_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dc_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"refresh_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
