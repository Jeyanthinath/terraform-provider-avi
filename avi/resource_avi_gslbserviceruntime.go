/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceGslbServiceRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"checksum": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"flr_state": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCfgStateSchema()},
			"groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbPoolRuntimeSchema()},
			"ldr_state": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCfgStateSchema()},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"oper_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceOperationalStatusSchema()},
			"send_event": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"send_status": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"services_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sp_oper_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceOperationalStatusSchema()},
			"tenant_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
		},
	}
}
