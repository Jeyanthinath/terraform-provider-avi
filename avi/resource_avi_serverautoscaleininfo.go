/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceServerAutoScaleInInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alertconfig_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alertconfig_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"available_capacity": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"load": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"num_scalein_servers": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"num_servers_up": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SYSERR_SUCCESS"},
			"scalein_server_candidates": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServerIdSchema()},
		},
	}
}
