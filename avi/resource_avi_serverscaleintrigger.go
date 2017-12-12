/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceServerScaleInTriggerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"servers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServerIdSchema()},
		},
	}
}
