/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVirtualServiceRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apic_extension": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVsApicExtensionSchema()},
			"controller_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"datapath_debug": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceDebugVirtualServiceSchema()},
			"east_west": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"gslb_dns_geo_update": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceGslbDnsGeoUpdateSchema()},
			"gslb_dns_update": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceGslbDnsUpdateSchema()},
			"ipam_dns_records": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsRecordSchema()},
			"is_dns_vs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"last_changed_time": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceTimeStampSchema()},
			"lif": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"manual_placement": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"marked_for_delete": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"metrics_mgr_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_additional_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"one_plus_one_ha": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"prev_controller_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prev_metrics_mgr_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"redis_db": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"redis_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"redis_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rules_configured": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"servers_configured": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"tls_ticket_key": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceTLSTicketSchema()},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "VS_TYPE_NORMAL"},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vh_child_vs_ref": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"vip_runtime": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipRuntimeSchema()},
		},
	}
}
