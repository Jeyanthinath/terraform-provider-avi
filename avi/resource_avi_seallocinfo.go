/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeAllocInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_down": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"at_curr_ver": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"attach_ip_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Programming Network reachability to the Virtual Service IP in the Cloud"},
			"attach_ip_success": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"floating_intf_ip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema()},
			"is_connected": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"is_primary": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"is_standby": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"memory": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scaling_in": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"sec_idx": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"total_se_memory": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"total_se_vcpus": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"vcpus": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0"},
			"vip_intf_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeVipInterfaceListSchema()},
			"vip_subnet_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_subnet_mask": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vnics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeResourceVnicSchema()},
		},
	}
}
