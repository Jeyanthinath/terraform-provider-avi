/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVipRuntimeDetailSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"config_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceConfigurationStatusSchema()},
			//"fel4stats": &schema.Schema{
			//	Type:     schema.TypeList,
			//	Optional: true,
			//	Elem:     ResourceVserverL4StatsSchema()},
			//"fel7stats": &schema.Schema{
			//	Type:     schema.TypeList,
			//	Optional: true,
			//	Elem:     ResourceVserverL7StatsSchema()},
			"first_se_assigned_time": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceTimeStampSchema()},
			//"l7_virtual_service_stats_runtime": &schema.Schema{
			//	Type:     schema.TypeList,
			//	Optional: true,
			//	Elem:     ResourceL7VirtualServiceStatsRuntimeSchema()},
			"last_scale_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceScaleStatusSchema()},
			"microservice_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"migrate_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"migrate_scalein_pending": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"migrate_scaleout_pending": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"num_se_assigned": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se_requested": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"oper_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceOperationalStatusSchema()},
			"proc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"progress_percent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scale_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceScaleStatusSchema()},
			"scalein_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"scaleout_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_engine": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipSeAssignedSchema()},
			"user_scaleout_pending": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vip_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_placement_resolution_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVipPlacementResolutionInfoSchema()},
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
