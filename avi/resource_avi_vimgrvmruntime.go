/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVIMgrVMRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/api/cloud?name=Default-Cloud"},
			"connection_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller_cluster_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller_ip_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller_vm": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"cpu_reservation": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cpu_shares": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"creation_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"guest_nic": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIMgrGuestNicRuntimeSchema()},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"init_vnics": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"managed_object_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"mem_shares": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"memory": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"memory_reservation": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"num_cpu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ovf_avisetype_field": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"powerstate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_ver": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
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
			"vcenter_datacenter_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_rm_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_se_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_template_vm": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vcenter_vappname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_vappvendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_vm_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_vnic_discovered": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vm_lb_weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
