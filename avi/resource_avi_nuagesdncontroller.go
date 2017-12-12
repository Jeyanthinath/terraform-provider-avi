/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceNuageSDNControllerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"nuage_organization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nuage_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nuage_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  8443,
			},
			"nuage_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nuage_vsd_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_enterprise": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_policy_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
