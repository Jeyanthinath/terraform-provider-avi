/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVirtualServiceStateDBCacheSummarySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"ew": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_invalid": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"last_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "OPER_DISABLED"},
			"last_updated_time": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"num_events": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_old_events": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pool_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"se": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeVsRuntimeSummarySchema()},
			"summary": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVirtualServiceRuntimeSummarySchema()},
			"vip_last_states": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVirtualServiceVipMetaDataSchema()},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
