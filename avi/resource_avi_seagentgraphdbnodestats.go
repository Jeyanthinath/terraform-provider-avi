/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeAgentGraphDBNodeStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"delete_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeAgentGraphDBNodeTxnStatsSchema()},
			"num_create": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_delete": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_read": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_update": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"read_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeAgentGraphDBNodeTxnStatsSchema()},
			"update_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeAgentGraphDBNodeTxnStatsSchema()},
		},
	}
}
