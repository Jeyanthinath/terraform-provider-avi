/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceGslbPoolDetailSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_ALGORITHM_ROUND_ROBIN"},
			"hmon": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceHealthMonitorStatRuntimeSchema()},
			"members": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbPoolMemberDetailSchema()},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_servers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_servers_enabled": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_servers_up": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"oper_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceOperationalStatusSchema()},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceGslbServiceGroupStatsSchema()},
		},
	}
}
