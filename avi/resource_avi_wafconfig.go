/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceWafConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allowed_http_versions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"allowed_methods": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"allowed_request_content_types": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"argument_separator": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "&"},
			"buffer_response_body_for_inspection": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"client_file_upload_max_body_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1024,
			},
			"client_nonfile_upload_max_body_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  128,
			},
			"cookie_format_version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"request_body_default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "phase:2,deny,status:403,log,auditlog"},
			"request_hdr_default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "phase:1,deny,status:403,log,auditlog"},
			"response_body_default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "phase:4,deny,status:403,log,auditlog"},
			"response_hdr_default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "phase:3,deny,status:403,log,auditlog"},
			"restricted_extensions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"restricted_headers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"server_response_max_body_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  128,
			},
		},
	}
}
