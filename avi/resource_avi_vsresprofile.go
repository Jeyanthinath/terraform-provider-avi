/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVsResProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_cpu_per_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_mem_per_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_cpu_per_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_mem_per_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"res": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceVirtualServiceResourceSchema()},
		},
	}
}
