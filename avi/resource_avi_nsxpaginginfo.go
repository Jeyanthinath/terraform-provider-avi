/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcensxPagingInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pagesize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sortby": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sortorderascending": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"startindex": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"totalcount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
