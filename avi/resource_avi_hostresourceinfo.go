/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHostResourceInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host_scale": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"managed_object_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"num_ses": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_fail_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_success_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
