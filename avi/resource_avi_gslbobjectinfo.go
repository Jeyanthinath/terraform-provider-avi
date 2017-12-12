/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceGslbObjectInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"object_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pb_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
