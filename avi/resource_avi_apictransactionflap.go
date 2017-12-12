/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAPICTransactionFlapSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ret_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ret_status_msg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
