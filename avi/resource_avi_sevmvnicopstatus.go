/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeVmVnicOpStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_vnics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCC_VnicInfoSchema()},
			"cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"status_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"syserr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SYSERR_FAILURE"},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
			"vnics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcevNICSchema()},
		},
	}
}
