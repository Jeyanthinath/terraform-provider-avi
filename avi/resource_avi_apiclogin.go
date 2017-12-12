/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceApicLoginSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apic_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"apic_password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"apic_username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
