/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVsFsmEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"vip_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_rt": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVirtualServiceRuntimeSchema()},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
