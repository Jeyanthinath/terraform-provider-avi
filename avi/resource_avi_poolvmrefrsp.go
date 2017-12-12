/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcePoolVmRefRspSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"obj_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"vm_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePoolVmRefInfoSchema()},
			"vrf_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
