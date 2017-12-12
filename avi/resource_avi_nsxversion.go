/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcensxVersionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"module": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcensxModuleSchema()},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
