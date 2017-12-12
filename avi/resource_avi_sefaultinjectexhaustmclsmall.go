/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeFaultInjectExhaustMclSmallSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_objects": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"proc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
