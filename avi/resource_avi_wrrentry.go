/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceWrrEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ip_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
