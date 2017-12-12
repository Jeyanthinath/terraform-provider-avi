/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcePoolGroupMemberSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"deployment_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"priority_label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ratio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
		},
	}
}
