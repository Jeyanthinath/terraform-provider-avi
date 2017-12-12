/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceConnpoolConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"life_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"lo_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"max_cache": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"max_reuse": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"strategy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"thresh_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
