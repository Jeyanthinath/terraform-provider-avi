/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVmStatsObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avg_cpu_maxlimited": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_cpu_ready": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_cpu_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_cpu_wait": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_cpus_system_time_secs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_cpus_user_time_secs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_disk1_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_disk2_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_disk3_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_disk4_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_disk_commands_aborted": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_disk_io": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_disk_max_total_latency": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_disk_number_read": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_disk_number_write": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_disk_read": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_disk_write": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_mem_swap_in": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_mem_swap_out": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_mem_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_mem_vmmemctl": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_net_dropped": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_net_dropped_rx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_net_dropped_tx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_net_errors_rx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_net_errors_tx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_net_received": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_net_transmitted": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_net_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_port_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_uptime": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_virtual_disk_commands_aborted": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_virtual_disk_read": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_virtual_disk_throughput_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_virtual_disk_write": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"node_obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
