/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeHBEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hb_type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"se_ref1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_ref2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
