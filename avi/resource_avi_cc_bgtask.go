/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCC_BGTaskSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"abort": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"added": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"admin_state": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"bgworker": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"oper_state": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"poll_timer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"poll_timer_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prempt": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"task_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
