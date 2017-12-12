/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceIpstkQStatsRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_ctrl_rx": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"num_ctrl_tx": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"num_data_rx": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"num_data_tx": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"num_db_rx": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"num_db_tx": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"num_drv_rx": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"num_drv_tx": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"proc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
