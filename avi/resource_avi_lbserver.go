/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceLbServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"affinity_core_open_conns": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"affinity_core_total_conns": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip_addr": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"non_affinity_core_open_conns": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"non_affinity_core_total_conns": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
