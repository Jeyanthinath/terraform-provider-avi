/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcePageStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"madvise": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"page_active": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"page_dirty": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"purged": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"sweep": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
