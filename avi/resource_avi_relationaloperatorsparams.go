/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceRelationalOperatorsParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"entity_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"right_arg": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
