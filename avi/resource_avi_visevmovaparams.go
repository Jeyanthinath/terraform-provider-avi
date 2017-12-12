/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVISeVmOvaParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_cluster_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller_ip_addr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rm_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_auth_token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sevm_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"single_socket_affinity": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vcenter_cpu_reserv": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vcenter_ds_include": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vcenter_ds_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVcenterDatastoreSchema()},
			"vcenter_ds_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_internal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "UNIFIED ADMIN"},
			"vcenter_mem_reserv": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vcenter_num_mem": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2048,
			},
			"vcenter_num_se_cores": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"vcenter_ovf_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_se_disk_size_kb": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10485760,
			},
			"vcenter_se_mgmt_nw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_vm_folder": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
