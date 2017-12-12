/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAlertSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action_script_output": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alert_config_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"app_events": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceApplicationLogSchema()},
			"conn_events": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceConnectionLogSchema()},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"event_pages": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"events": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceEventLogSchema()},
			"last_throttle_timestamp": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"level": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"metric_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMetricLogSchema()},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"obj_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"obj_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"related_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"summary": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"throttle_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"timestamp": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
		},
	}
}
