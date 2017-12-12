/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceEcnStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ecn_capable_transport_0": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ecn_capable_transport_1": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ecn_congestion_experienced": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ecn_successful_handshakes": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"times_ecn_reduced_cwnd": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
