/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceIndexingTaskResponseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"error_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"filepos": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_timestamp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_timestamp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"success": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
		},
	}
}
