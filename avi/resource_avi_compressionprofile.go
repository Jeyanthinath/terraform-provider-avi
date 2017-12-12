/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCompressionProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"compressible_content_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"compression": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCompressionFilterSchema()},
			"remove_accept_encoding_header": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
