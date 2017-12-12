/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeMgrEventDetailsSchema() *schema.Resource {
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
			"enable_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gcp_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceGcpInfoSchema()},
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
			"migrate_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"reason": &schema.Schema{
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
			"vcpus": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vs_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
		},
	}
}
