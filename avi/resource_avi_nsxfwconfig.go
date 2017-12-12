/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcensxFwConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"contextid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"generatenumber": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"layer2sections": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcesectionsSchema()},
			"layer3redirectsections": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcesectionsSchema()},
			"layer3sections": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcesectionsSchema()},
			"timestamp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
