/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCloneServerStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"num_driver_failures": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_mbuf_alloc_failed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_succeeded": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_total_pkts_cloned": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
