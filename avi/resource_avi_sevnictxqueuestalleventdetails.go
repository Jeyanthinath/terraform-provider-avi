/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeVnicTxQueueStallEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"if_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"linux_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"queue": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
