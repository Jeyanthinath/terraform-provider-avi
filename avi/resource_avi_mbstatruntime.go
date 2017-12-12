/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMbStatRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"drv_compacts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"drv_compacts_failed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mbuf_allocation_failures": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"mbuf_allocations": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"mbuf_available": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"mbuf_cache_empty": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"mbuf_cache_max": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"mbuf_cached_allocations": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"mbuf_cached_frees": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"mbuf_frees": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"mbuf_total": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"mcopy_fail": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"mpullup_fail": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"mpullup_slow": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"pktbuf_allocation_failures": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"pktbuf_allocations": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"pktbuf_available": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"pktbuf_cache_empty": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"pktbuf_cache_max": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"pktbuf_cached_allocations": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"pktbuf_cached_frees": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"pktbuf_frees": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"pktbuf_total": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"proc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"small_pktbuf_allocation_failures": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"small_pktbuf_allocations": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"small_pktbuf_available": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"small_pktbuf_cache_empty": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"small_pktbuf_cache_max": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"small_pktbuf_cached_allocations": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"small_pktbuf_cached_frees": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"small_pktbuf_frees": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"small_pktbuf_total": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
