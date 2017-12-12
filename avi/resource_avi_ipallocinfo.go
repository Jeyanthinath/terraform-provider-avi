/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceIpAllocInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}