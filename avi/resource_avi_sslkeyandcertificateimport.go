/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSSLKeyAndCertificateImportSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate_base64": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enckey_base64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"enckey_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_PEM"},
			"hardwaresecuritymodulegroup_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"key_base64": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"key_passphrase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_CERTIFICATE_TYPE_VIRTUALSERVICE"},
		},
	}
}
