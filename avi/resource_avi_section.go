/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcesectionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"class": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"generationnumber": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"rule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceruleSchema()},
			"timestamp": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
