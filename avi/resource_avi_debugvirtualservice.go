/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceDebugVirtualServiceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"capture": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"capture_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceDebugVirtualServiceCaptureSchema()},
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/api/cloud?name=Default-Cloud"},
			"debug_hm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DEBUG_VS_HM_NONE"},
			"debug_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceDebugIpAddrSchema()},
			"flags": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDebugVsDataplaneSchema()},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"se_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceDebugVirtualServiceSeParamsSchema()},
			"tenant_ref": &schema.Schema{
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
