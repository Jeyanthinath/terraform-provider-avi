/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceServiceEngineDosAttackLevelDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"avg_connections": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rx_bandwidth": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"avg_rx_pkts": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_connections_dropped": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_rx_bytes_dropped": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pct_rx_pkts_dropped": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"reason_attr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}
