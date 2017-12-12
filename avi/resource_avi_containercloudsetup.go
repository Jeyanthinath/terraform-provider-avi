/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceContainerCloudSetupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_access": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"failed_hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"fleet_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"master_nodes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"missing_hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"new_hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_deploy_method_access": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"se_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
