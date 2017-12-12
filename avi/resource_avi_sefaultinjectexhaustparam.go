/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSEFaultInjectExhaustParamSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"leak": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"num": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
