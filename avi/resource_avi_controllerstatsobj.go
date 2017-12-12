/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceControllerStatsObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avg_cpu_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_disk_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_mem_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_num_active_vs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_num_backend_servers": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_num_se_cores": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_num_ses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_num_sockets": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_total_se_throughput": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_be_servers_lic_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_num_active_vs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_num_active_vs_lic_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_num_backend_servers": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_num_se_cores": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_num_ses": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_num_ses_lic_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_num_sockets": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_num_sockets_lic_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_se_cores_lic_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_se_throughput_lic_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_total_se_throughput": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"node_obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"sum_total_se_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_total_vs_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_total_vs_client_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_total_vs_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}
