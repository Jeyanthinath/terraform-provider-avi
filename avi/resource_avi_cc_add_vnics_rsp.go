/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Resourcecc_add_vnics_rspSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_vnics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCC_VnicInfoSchema()},
			"ret_status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"ret_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcevNICSchema()},
		},
	}
}
