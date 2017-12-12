/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMicroServiceMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"group_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
