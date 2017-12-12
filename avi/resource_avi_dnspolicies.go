/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceDnsPoliciesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_policy_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
