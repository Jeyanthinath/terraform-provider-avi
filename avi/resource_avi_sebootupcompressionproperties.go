/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeBootupCompressionPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"buf_num": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  128,
			},
			"buf_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4096,
			},
			"hash_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  16384,
			},
			"level_aggressive": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"level_normal": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"window_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4096,
			},
		},
	}
}
