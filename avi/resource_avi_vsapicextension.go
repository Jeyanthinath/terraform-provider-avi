/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVsApicExtensionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"txn_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
			"vnic": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVsSeVnicSchema()},
		},
	}
}
