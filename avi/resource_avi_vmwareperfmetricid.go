/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVmwarePerfMetricIdSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"counter_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"instance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
