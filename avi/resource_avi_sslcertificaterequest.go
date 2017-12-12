/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSSLCertificateRequestSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"certificate_management_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"challenge_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"common_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"days_until_expire": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  365,
			},
			"dynamic_params": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCustomParamsSchema()},
			"email_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hardwaresecuritymodulegroup_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"key_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSSLKeyParamsSchema()},
			"locality": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"organization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"organization_unit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"self_signed": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"subject_alt_names": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_CERTIFICATE_TYPE_VIRTUALSERVICE"},
		},
	}
}
