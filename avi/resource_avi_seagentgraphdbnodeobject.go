/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeAgentGraphDBNodeObjectSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"config": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeAgentGraphDBNodeConfigSchema()},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeAgentGraphDBNodeStatsSchema()},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
