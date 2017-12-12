/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVsAwaitingSeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"awaitingse_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_assigned": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipSeAssignedSchema()},
			"se_requested": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVirtualServiceResourceSchema()},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
