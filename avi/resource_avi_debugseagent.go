/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceDebugSeAgentSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"log_level": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"sub_module": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"trace_level": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
