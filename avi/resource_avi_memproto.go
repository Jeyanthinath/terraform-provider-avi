/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMemProtoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"current_usage": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"max_mem_usage": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"memory_mapped": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
