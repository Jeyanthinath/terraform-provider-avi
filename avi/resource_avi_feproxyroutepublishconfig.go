/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceFeProxyRoutePublishConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "FE_PROXY_ROUTE_PUBLISH_NONE"},
			"publisher_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  80,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  32,
			},
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
