/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceQueryHostsCapabilityReqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceVIAdminCredentialsSchema()},
			"async": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"cloud_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"datacenter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHostInfoSchema()},
			"ignore_vcenter_query": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"memory": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"num_cpu": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"obj_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ovf_path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"se_mgmt_nw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"vcenter_se_disk_size_kb": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10485760,
			},
		},
	}
}
