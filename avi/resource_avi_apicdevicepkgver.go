/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceApicDevicePkgVerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"majorversion": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"minorversion": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"model": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"vendor": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
