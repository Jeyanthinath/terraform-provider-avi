/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceIfQStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ibytes": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ipackets": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"obytes": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"oerrors": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"opackets": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"rx_max_single_burst": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"rx_pkt_iterations": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"rx_queue_full": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
