/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceContentRewriteProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"req_match_replace_pair": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMatchReplacePairSchema()},
			"request_rewrite_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"response_rewrite_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"rewritable_content_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rsp_match_replace_pair": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMatchReplacePairSchema()},
		},
	}
}
