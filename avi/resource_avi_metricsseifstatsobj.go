/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMetricsSeIfStatsObjSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avg_bandwidth": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_connection_table_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rx_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rx_bytes_dropped_non_vs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rx_pkts": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rx_pkts_dropped_non_vs": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_syn_seen_entries_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_tx_bytes": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_tx_pkts": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_peak_bandwidth": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_rx_bytes_absolute": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_tx_bytes_absolute": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"node_obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"sum_connection_dropped_syn_seen_limit": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sum_connection_dropped_table_limit": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}
