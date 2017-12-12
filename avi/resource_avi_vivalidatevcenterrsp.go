/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVIValidateVcenterRspSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"api_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ctrl_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ctrl_vnics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIControllerVnicInfoSchema()},
			"datacenters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVIDCInfoSchema()},
			"discovery_dc_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"response_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"vcenter_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
