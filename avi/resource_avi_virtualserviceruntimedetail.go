/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVirtualServiceRuntimeDetailSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"config_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceConfigurationStatusSchema()},
			//"dns_stats": &schema.Schema{
			//	Type:     schema.TypeList,
			//	Optional: true,
			//	Elem:     ResourceVserverDNSStatsSchema()},
			"east_west": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"microservice_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"one_plus_one_ha": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"oper_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceOperationalStatusSchema()},
			"proc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "VS_TYPE_NORMAL"},
			"vh_child_vs_ref": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"vip_detail": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipRuntimeDetailSchema()},
			//"virtual_service_auth_stats": &schema.Schema{
			//	Type:     schema.TypeList,
			//	Optional: true,
			//	Elem:     ResourceVirtualServiceAuthStatsSchema()},
			"vs_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
