/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceLogCommandSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"args": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"cmd": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
