/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeResourceAddReqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"at_curr_ver": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"az": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"azure_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceAzureInfoSchema()},
			"container_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"container_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "CONTAINER_TYPE_HOST"},
			"disk_gb": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"enable_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"flavor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gateway_up": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"host_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hypervisor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_mac_addr": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"memory": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"online": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"se_group_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"subnet_ip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema()},
			"subnet_mask": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt}},
			"vcpus": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0"},
			"vinfra_discovered": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vnics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeResourceVnicSchema()},
		},
	}
}
