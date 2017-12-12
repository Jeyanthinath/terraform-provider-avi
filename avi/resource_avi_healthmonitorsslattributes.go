/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHealthMonitorSSLAttributesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pki_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_key_and_certificate_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
