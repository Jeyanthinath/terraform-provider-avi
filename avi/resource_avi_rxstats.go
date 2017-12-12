/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceRxStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ack_byte_received": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ack_packets_received": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ack_too_much_packets_received": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"after_window_data_bytes_received": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"after_window_data_packets_received": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"bytes_received_in_sequence": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ignored_rst_packet_in_window": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"packets_received_in_sequence": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"total_packets_received": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"window_probe_packets_received": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"window_update_packet_received": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
