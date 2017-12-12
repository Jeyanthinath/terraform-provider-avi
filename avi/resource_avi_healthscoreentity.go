/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHealthScoreEntitySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"entity_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"metrics_entity": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
