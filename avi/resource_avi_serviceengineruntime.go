/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceServiceEngineRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"active_tags": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"at_curr_ver": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
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
			"counters": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceServiceEngineCountersSchema()},
			"creation_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"gateway_up": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"gcp_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceGcpInfoSchema()},
			"hb_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeHbStatusSchema()},
			"inband_mgmt": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"migrate_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SE_MIGRATE_STATE_IDLE"},
			"migrate_timestamp": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceTimeStampSchema()},
			"networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeNetworkSubnetSchema()},
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
				Default:  "SE_POWER_ON"},
			"prev_se_connected": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"se_connected": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"state_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"state_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SeMgrFsmMap::Inactive"},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
			"vcenter_rm_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"vnic": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVnicRuntimeSchema()},
			"vs_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
		},
	}
}
