/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCC_CronSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"average": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"last": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_fail": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_fail_msg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_ok": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"next": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
