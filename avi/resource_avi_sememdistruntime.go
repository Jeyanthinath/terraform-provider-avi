/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeMemDistRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"clusters": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"conn_memory_mb": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"conn_memory_mb_per_core": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"huge_pages": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"hypervisor_type": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"num_queues": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"num_rxd": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"num_txd": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"os_reserved_memory_mb": &schema.Schema{
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
			"shm_conn_memory_mb": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"shm_memory_mb": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
