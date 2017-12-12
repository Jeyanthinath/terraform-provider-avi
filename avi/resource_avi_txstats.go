/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceTxStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ack_only_packet": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"data_bytes_retransmitted": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"data_bytes_sent": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"data_packets_retransmitted": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"data_packets_sent": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"delayed_acks_sent": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"keepalive_probes_sent": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"total_packets_sent": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"unnecessary_packet_retransmit": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"urg_only_packets_sent": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"window_probes_sent": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"window_update_only_packets_sent": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
