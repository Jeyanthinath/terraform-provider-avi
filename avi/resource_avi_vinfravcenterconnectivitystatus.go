/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVinfraVcenterConnectivityStatusSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"datacenter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"vcenter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
