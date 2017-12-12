/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeUpgradeStatusSummarySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"duration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"notes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"se_already_upgraded_at_start": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"se_disconnected_at_start": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"se_group_status": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeGroupStatusSchema()},
			"se_ip_missing_at_start": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"se_poweredoff_at_start": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"se_upgrade_completed": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"se_upgrade_errors": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeUpgradeErrorsSchema()},
			"se_upgrade_failed": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"se_upgrade_in_progress": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"se_upgrade_not_started": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"se_upgrade_retry_completed": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"se_upgrade_retry_failed": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"se_upgrade_retry_in_progress": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"start_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_errors": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVsErrorSchema()},
		},
	}
}
