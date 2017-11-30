/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceFailActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"backup_pool": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceFailActionBackupPoolSchema()},
			"local_rsp": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceFailActionHTTPLocalResponseSchema()},
			"redirect": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceFailActionHTTPRedirectSchema()},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
