/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceURIParamTokenSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"end_index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"start_index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"str_value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
