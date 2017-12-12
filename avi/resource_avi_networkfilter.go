/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceNetworkFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"server_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
