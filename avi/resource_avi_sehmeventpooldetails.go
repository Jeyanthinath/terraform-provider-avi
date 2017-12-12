/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeHmEventPoolDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ha_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"percent_servers_up": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeHmEventServerDetailsSchema()},
			"src_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
