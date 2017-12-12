/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMallstatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bytes_allocated": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"bytes_mapped": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
