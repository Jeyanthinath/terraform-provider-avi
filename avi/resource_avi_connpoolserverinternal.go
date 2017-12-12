/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceConnpoolServerInternalSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"conn_bind": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceConnectionItemSchema()},
			"conn_free": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceConnectionItemSchema()},
			"ip_addr": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"max_connection": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_conn_bind": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_conn_free": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"oper_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceOperationalStatusSchema()},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceConnpoolStatsSchema()},
		},
	}
}
