/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHTTPApplicationProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allow_dots_in_header_name": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"cache_config": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceHttpCacheConfigSchema()},
			"client_body_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30000,
			},
			"client_header_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10000,
			},
			"client_max_body_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"client_max_header_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  12,
			},
			"client_max_request_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  48,
			},
			"compression_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCompressionProfileSchema()},
			"connection_multiplexing_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"disable_keepalive_posts_msie6": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"enable_fire_and_forget": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enable_request_body_buffering": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"hsts_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"hsts_max_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  365,
			},
			"http_to_https": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"httponly_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"keepalive_header": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"keepalive_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30000,
			},
			"max_bad_rps_cip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_bad_rps_cip_uri": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_bad_rps_uri": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_rps_cip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_rps_cip_uri": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_rps_unknown_cip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_rps_unknown_uri": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_rps_uri": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pki_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"post_accept_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30000,
			},
			"secure_cookie_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"server_side_redirect_to_https": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"spdy_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"spdy_fwd_proxy_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ssl_client_certificate_action": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSSLClientCertificateActionSchema()},
			"ssl_client_certificate_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_CLIENT_CERTIFICATE_NONE"},
			"ssl_everywhere_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_app_keepalive_timeout": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"websockets_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"x_forwarded_proto_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"xff_alternate_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "X-Forwarded-For"},
			"xff_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}
