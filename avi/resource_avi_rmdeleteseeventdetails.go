/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceRmDeleteSeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_grp_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_grp_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status_code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
