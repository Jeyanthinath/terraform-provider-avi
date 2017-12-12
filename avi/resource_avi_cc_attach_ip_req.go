/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Resourcecc_attach_ip_reqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"azure_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceAzureInfoSchema()},
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "cloud-0"},
			"cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_pfx": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceIpAddrPrefixSchema()},
			"is_primary": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mac_addr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"scaleout_ecmp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"se_group_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_intfs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeIntfSchema()},
			"se_vm_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"services": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVsProtocolSchema()},
			"tenant_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
