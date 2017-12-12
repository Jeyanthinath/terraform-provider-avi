/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSSLKeyRSAParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"exponent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  65537,
			},
			"key_size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_KEY_2048_BITS"},
		},
	}
}
