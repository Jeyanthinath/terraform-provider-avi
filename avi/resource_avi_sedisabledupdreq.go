/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeDisabledUpdReqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_state": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
