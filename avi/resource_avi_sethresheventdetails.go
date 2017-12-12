/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeThreshEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"curr_value": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
