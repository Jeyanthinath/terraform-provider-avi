/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceConfigUserNotAuthrzByRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"roles": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "No Roles"},
			"tenants": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "No Tenants"},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
