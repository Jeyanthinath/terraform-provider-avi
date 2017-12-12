/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVIMgrSEVMRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"azure_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceAzureInfoSchema()},
			"cloud_name": &schema.Schema{
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
			"creation_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"deletion_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"discovery_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"discovery_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"flavor": &schema.Schema{
				Type:     schema.TypeString,
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
			"host_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hypervisor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"init_vnics": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"last_discovery": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"managed_object_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"powerstate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_group_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"segroup_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_group_uuid": &schema.Schema{
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
		},
	}
}
