/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceServerAutoScaleOutCompleteInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"launch_config_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nscaleout": &schema.Schema{
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
				Required: true},
			"scaled_out_servers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServerIdSchema()},
		},
	}
}
