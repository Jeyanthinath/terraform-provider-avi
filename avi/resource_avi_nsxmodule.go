/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcensxModuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"baseuri": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
