/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceNetworkSubnetVMSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_addr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
