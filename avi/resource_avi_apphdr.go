/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAppHdrSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hdr_match_case": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"hdr_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"hdr_string_op": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
