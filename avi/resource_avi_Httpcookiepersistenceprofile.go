/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHttpCookiePersistenceProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"always_send_cookie": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"cookie_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"encryption_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHttpCookiePersistenceKeySchema()},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
