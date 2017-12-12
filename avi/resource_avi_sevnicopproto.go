/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeVnicOpProtoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"macs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"oper_start_ticks": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"oper_start_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"se_vm_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_network_ids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"vnic_op_infra": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnic_op_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"vrf_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
		},
	}
}
