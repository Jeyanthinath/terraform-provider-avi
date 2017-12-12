/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCC_MgmtNwSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"az": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mgmt_nw_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
