/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeAgentPropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_echo_miss_aggressive_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"controller_echo_miss_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"controller_echo_rpc_aggressive_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2000,
			},
			"controller_echo_rpc_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2000,
			},
			"controller_heartbeat_miss_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  6,
			},
			"controller_heartbeat_timeout_sec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  12,
			},
			"controller_registration_timeout_sec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"controller_rpc_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"cpustats_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"ctrl_reg_pending_max_wait_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  150,
			},
			"debug_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dp_aggressive_deq_interval_msec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"dp_aggressive_enq_interval_msec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"dp_batch_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"dp_deq_interval_msec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  20,
			},
			"dp_enq_interval_msec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  20,
			},
			"dp_max_wait_rsp_time_sec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"dp_reg_pending_max_wait_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  75,
			},
			"headless_timeout_sec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ignore_docker_mac_change": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"sdb_flush_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"sdb_pipeline_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"sdb_scan_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000,
			},
			"vnic_dhcp_ip_check_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  6,
			},
			"vnic_dhcp_ip_max_retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"vnic_ip_delete_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"vnic_probe_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
		},
	}
}
