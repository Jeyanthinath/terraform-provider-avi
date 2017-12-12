/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSyncacheStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bucket_overflow": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"cache_overflow": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"duplicate_syn_packet": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"entry_aborted": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"entry_added_to_syncache": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"entry_completed": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"entry_dropped_reply_failed": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"entry_removed_by_badack": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"entry_removed_by_reset": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"entry_retransmitted": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"entry_staled": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"icmp_unreachable_received": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"syn_cookie_received": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"syn_cookie_sent": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"zone_failures": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
