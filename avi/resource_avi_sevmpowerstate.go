/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeVmPowerStateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"power_off": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "poweredOff"},
			"power_on": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "poweredOn"},
			"suspended": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "suspended"},
		},
	}
}
