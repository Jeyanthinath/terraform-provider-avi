/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceOsLbProvAuditFormSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"passwd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prov_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"tenant": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
