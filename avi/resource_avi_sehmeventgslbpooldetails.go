/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeHmEventGslbPoolDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"gsgroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gslb_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gsmember": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeHmEventGslbPoolMemberDetailsSchema()},
			"ha_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
