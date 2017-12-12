/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceServiceEngineRuntimeDetailSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"active_tags": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"at_curr_ver": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"counters": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceServiceEngineCountersSchema()},
			"gateway_up": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"hb_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeHbStatusSchema()},
			"inband_mgmt": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"migrate_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SE_MIGRATE_STATE_IDLE"},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"online_since": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"oper_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceOperationalStatusSchema()},
			"power_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resources": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeResourcesSchema()},
			"se_connected": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0"},
			"vinfra_discovered": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vnic": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVnicRuntimeSchema()},
		},
	}
}
