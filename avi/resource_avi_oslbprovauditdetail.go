/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceOsLbProvAuditDetailSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"diff": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_ts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"job_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"job_ts": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"passed": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"start_ts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
