/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSSLKeyECParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"curve": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_KEY_EC_CURVE_SECP256R1"},
		},
	}
}
