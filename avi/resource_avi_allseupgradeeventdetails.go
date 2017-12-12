/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAllSeUpgradeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"notes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"num_se": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"num_vs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"request": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeUpgradeParamsSchema()},
		},
	}
}
