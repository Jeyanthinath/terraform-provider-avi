/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceServiceEngineDbExtensionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"active_tags": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"azure_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceAzureInfoSchema()},
			"consumers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceConInfoSchema()},
			"container_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gcp_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceGcpInfoSchema()},
			"inband_mgmt": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"migrate_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SE_MIGRATE_STATE_IDLE"},
			"networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeNetworkSubnetSchema()},
			"online_since": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_connected": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0"},
			"vinfra_discovered": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vs_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
		},
	}
}
