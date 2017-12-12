/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeResourceConsumedProtoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_down": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"attach_ip_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"attach_ip_fail_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"attach_ip_fail_str": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"attach_ip_fail_syserr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SYSERR_FAILURE"},
			"attach_ip_in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"attach_ip_success": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"consumer_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"consumer_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"failover_attempts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"floating_intf_ip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema()},
			"is_primary": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_stby": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"last_attach_ip_fail_ticks": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"last_attach_ip_start_ticks": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"last_failover_attempt_ticks": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"marked_for_delete": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"marked_for_scalein": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"memory": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"res_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"res_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"sec_idx": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"static_se_binding": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vcpus": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"vip_intf_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeVipInterfaceListSchema()},
			"vnics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeResourceVnicSchema()},
		},
	}
}
