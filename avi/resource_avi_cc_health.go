/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCC_HealthSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"first_fail": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_fail": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_fails": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"last_ok": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_fails": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"services": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCC_Service_HealthSchema()},
		},
	}
}
