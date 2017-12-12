/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceServiceEngineInternalSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"at_curr_ver": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"cloud_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller_created": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"controller_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_vnics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcevNICSchema()},
			"default_gws": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDefaultGatewaySchema()},
			"gateway_up": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"host_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"inband_mgmt": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mgmt_vnic": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcevNICSchema()},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resources": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeResourcesSchema()},
			"se_group_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0"},
		},
	}
}
