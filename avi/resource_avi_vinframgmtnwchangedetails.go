/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVinfraMgmtNwChangeDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"existing_nw": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"new_nw": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"vcenter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
