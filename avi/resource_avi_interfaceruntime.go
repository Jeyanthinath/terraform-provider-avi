/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceInterfaceRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"proc_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"vnics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVnicInfoSchema()},
		},
	}
}
