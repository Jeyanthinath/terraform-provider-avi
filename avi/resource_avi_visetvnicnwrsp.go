/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVISetvNicNwRspSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"sevm_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"status_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_sevm_mor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
