/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceServerStateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"server_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"server_state": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
