/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceFailActionHTTPRedirectSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HTTPS"},
			"query": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HTTP_REDIRECT_STATUS_CODE_302"},
		},
	}
}
