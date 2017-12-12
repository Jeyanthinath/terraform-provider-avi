/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCloudTenantCleanupSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_ports": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_secgrp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_svrgrp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
