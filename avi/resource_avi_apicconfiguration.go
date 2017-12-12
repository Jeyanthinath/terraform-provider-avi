/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAPICConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apic_admin_tenant": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "common"},
			"apic_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"apic_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"apic_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"apic_product": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ASP"},
			"apic_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"apic_vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Avi"},
			"avi_controller_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"avi_controller_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "admin"},
			"context_aware": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SINGLE_CONTEXT"},
			"deployment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"minor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "2"},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "1.0"},
		},
	}
}
