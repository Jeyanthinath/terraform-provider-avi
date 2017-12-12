/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeResourceFindReqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"active_standby_se_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ACTIVE_STANDBY_SE_1"},
			"apic_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"clear_placement_disabled_flags": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dns_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsInfoSchema()},
			"east_west_placement": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enable_rhi": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enable_rhi_snat": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"fip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"ign_pool_net_reach": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"lifs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"notify_on_avail": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"num_se": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"num_standby_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vnics": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"parent_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"res_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVsResProfileSchema()},
			"scaleout_ecmp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_group_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"servers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceConServerSchema()},
			"service_subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrPrefixSchema()},
			"services": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVsProtocolSchema()},
			"snat_ip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema()},
			"static_se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"static_se_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"svc_id": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceSeMgrSvcIdSchema()},
			"svc_notify_q_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_res_profile": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceConVipSchema()},
			"vrf_context_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
