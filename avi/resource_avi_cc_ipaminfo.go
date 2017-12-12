/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCC_IpamInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aws_htoken_expiry": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"aws_stoken_expiry": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"err_msg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipam_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_subnet_cidr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_subnet_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"park_intfs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCC_ParkIntfSchema()},
			"ready": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vpc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
