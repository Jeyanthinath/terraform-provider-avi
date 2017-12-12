/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceReferencableSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bool_data": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"enum_data": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"int_data": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"string_data": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
		},
	}
}
