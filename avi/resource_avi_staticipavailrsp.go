/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceStaticIPAvailRspSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_avail": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceStaticIPAvailInfoSchema()},
			"status": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}
