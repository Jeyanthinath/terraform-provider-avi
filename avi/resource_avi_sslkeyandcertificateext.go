/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSSLKeyAndCertificateExtSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"certificate_base64": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_PEM"},
			"key_base64": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"key_passphrase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
