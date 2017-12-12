/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMarathonServicePortConflictSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"marathon_url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
