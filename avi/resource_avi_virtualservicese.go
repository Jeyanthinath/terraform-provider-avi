/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVirtualServiceSeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apic_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"cluster_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"datapath_debug": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceDebugVirtualServiceSchema()},
			"first_se_assigned_time": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceTimeStampSchema()},
			"geo_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gs_refs_v2": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"gslb_clear_on_max_retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gslb_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gslb_send_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ipam_dns_records": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsRecordSchema()},
			"marked_for_delete": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"metrics_mgr_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"se_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeListSchema()},
			"tls_ticket_key": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceTLSTicketSchema()},
			"total_ses": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"total_vcpus": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"total_vips": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"virtual_service": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVirtualServiceSchema()},
		},
	}
}
