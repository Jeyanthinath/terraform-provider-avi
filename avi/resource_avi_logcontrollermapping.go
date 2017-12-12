/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceLogControllerMappingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metrics_mgr_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRICS_MGR_PORT_0"},
			"node_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prev_controller_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prev_metrics_mgr_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRICS_MGR_PORT_0"},
			"static_mapping": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
