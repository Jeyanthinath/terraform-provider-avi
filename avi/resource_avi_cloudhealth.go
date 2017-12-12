/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCloudHealthSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"first_fail": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_fail": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_ok": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_fails": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
