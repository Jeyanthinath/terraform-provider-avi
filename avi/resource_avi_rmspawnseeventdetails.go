/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceRmSpawnSeEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
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
			"memory": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"network_names": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
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
			"vcpus": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vs_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
