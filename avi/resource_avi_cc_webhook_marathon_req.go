/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Resourcecc_webhook_marathon_reqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "cloud-0"},
			"event_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"task_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
