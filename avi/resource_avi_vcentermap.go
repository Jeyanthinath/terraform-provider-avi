/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVcenterMapSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"datacenters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDatacenterMapSchema()},
			"vcenter_url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
