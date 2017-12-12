/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVIMgrVcenterRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"api_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"apic_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/api/cloud?name=Default-Cloud"},
			"datacenter_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"disc_end_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"disc_start_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"discovered_datacenter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"inventory_progress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"inventory_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"num_clusters": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_dcs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_hosts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_nws": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vcenter_req_pending": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vms": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"privilege": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"progress": &schema.Schema{
				Type:     schema.TypeInt,
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
			"vcenter_connected": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vcenter_fullname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_template_se_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
