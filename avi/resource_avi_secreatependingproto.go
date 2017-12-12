/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeCreatePendingProtoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"az": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"create_op_consumers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"flavor_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hypervisor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_ips_per_vnic": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_vnics": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"memory": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"mgmt_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"mgmt_net_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"op_start_ticks": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"oper_start_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"progress_percent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"se_group_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"vcpus": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"virtual_network_ids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"vrf_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
		},
	}
}
