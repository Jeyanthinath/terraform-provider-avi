/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVISetMgmtIpSEReqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVIAdminCredentialsSchema()},
			"all_vnic_connected": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"cloud_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dc_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_params": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceVISeVmIpConfParamsSchema()},
			"power_on": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"rm_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sevm_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
