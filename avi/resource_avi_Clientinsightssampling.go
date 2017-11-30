/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceClientInsightsSamplingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrMatchSchema()},
			"sample_uris": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceStringMatchSchema()},
			"skip_uris": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceStringMatchSchema()},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
