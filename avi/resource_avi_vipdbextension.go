/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVipDbExtensionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"first_se_assigned_time": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceTimeStampSchema()},
			"num_additional_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"requested_resource": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVirtualServiceResourceSchema()},
			"se_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeListSchema()},
			"vip_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
