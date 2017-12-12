/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcePerformanceLimitsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_concurrent_connections": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_throughput": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
