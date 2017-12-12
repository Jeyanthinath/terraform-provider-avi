/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceDatacenterMapSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dc_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"dc_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"disc_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"perf_monitor_list_ver": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"perf_monitor_ts": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"resource_monitor_list_ver": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"resource_monitor_ts": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"vm_monitor_list_ver": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"vm_monitor_ts": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
