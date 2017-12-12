/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVIMgrNWRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apic_vrf_context": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_expand": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/api/cloud?name=Default-Cloud"},
			"datacenter_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dvs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"host_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"interested_nw": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ip_subnet": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIMgrIPSubnetRuntimeSchema()},
			"managed_object_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"mgmtnw": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"num_ports": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"switch_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
			"vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vlan_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVlanRangeSchema()},
			"vm_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"vrf_context_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
