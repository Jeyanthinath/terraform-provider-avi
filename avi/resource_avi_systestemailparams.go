/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSysTestEmailParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_emails": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subject": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"text": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"to_emails": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
