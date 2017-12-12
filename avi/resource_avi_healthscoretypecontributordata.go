/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHealthScoreTypeContributorDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"entities": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHealthScoreEntitySchema()},
			"hs_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"metric_ids": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
