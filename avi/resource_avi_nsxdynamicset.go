/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcensxDynamicSetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dynamiccriteria": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcensxDynamicCriteriaSchema()},
			"operator": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
