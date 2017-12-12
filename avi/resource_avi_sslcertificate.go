/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSSLCertificateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate_signing_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"chain_verified": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"days_until_expire": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  365,
			},
			"expiry_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_CERTIFICATE_GOOD"},
			"fingerprint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"issuer": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSSLCertificateDescriptionSchema()},
			"key_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSSLKeyParamsSchema()},
			"not_after": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"not_before": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"self_signed": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"signature": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"signature_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subject": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSSLCertificateDescriptionSchema()},
			"subject_alt_names": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"text": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
