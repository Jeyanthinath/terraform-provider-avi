/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceJestatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bytes_active": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"bytes_allocated": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"bytes_allocated_huge": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"bytes_allocated_large": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"bytes_allocated_small": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"bytes_mapped": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"large": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceAllocInfoSchema()},
			"page_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcePageStatsSchema()},
			"small": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceAllocInfoSchema()},
			"total": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceAllocInfoSchema()},
		},
	}
}
