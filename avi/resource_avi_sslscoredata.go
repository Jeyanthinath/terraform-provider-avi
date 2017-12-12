/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSslScoreDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cert_chain_verified": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cert_chain_verified_score": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"disable_client_renegotiation": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"disable_client_renegotiation_score": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  "5.0"},
			"earliest_cert_expiry": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"earliest_cert_expiry_score": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"hsts_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"hsts_enabled_score": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"min_cipher_strength": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"min_cipher_strength_score": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"min_ssl_protocol_strength": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"min_ssl_protocol_strength_score": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pfs_support": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"pfs_support_score": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"reason_attr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"self_signed_cert": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"self_signed_cert_score": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"ssl_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"value": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  "5.0"},
			"weak_signature_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"weak_signature_algorithm_score": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"weakest_enc_algo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"weakest_enc_algo_score": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}
