/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCpuUsageEntrySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cycles_idle": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"ethrx_cycles_idle": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"hwrx_cycles_idle": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"l7_cycles_idle": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"total_cycles": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"vnictx_cycles_idle": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
